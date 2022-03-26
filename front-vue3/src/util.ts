import { CodingStream } from '@/CodingStream'
import type { CodingSequence } from '@/models/CodingSequence'
import { Snapshot } from '@/models/Snapshot'
import type { Video } from '@/models/Video'

export const formatRecordingTime = (value: number): string => {
  const secondTime = value / 1000
  const hours = ('00' + Math.floor(secondTime / 3600)).slice(-2)
  const minutes = ('00' + Math.floor((secondTime / 60) % 60)).slice(-2)
  const seconds = ('00' + Math.floor(secondTime % 60)).slice(-2)
  return `${hours}:${minutes}:${seconds}`
}

export const getPreviousTimeSpan = (
  timestamp: number,
  timeSpan: number
): number => {
  timestamp
  return Math.ceil(timestamp / timeSpan) * timeSpan - timeSpan
}

export type IndexStatus = 'None' | 'Equals' | 'Smaller' | 'Greater'

export const identificationBetweenSequences = (
  searchValue: number,
  stream: CodingStream
): { index: number; indexStatus: IndexStatus } => {
  const currentIndex = stream.index
  let range = stream.length / 2
  let rowIndex = range
  let index = Math.ceil(rowIndex)
  stream.seek(index)
  let value = stream.current.timestamp
  let previousValue = -1
  let indexStatus: IndexStatus = 'None'
  while (previousValue !== value) {
    previousValue = value
    range /= 2
    if (value == searchValue) {
      indexStatus = 'Equals'
      break
    }
    if (value > searchValue) {
      rowIndex -= range
      index = Math.ceil(rowIndex)
      stream.seek(index)
      value = stream.current.timestamp
      indexStatus = 'Smaller'
    } else {
      rowIndex += range
      index = Math.ceil(rowIndex)
      stream.seek(index)
      value = stream.current.timestamp
      indexStatus = 'Greater'
    }
  }
  stream.seek(currentIndex)
  return { index, indexStatus }
}

/**
 * @function createSnapshot CodingSequenceからスナップショットを作成する
 * @description
 * 差分は初期状態と最終状態の２つとTimeSpanの時間毎に作成される。
 * CodingSequenceと同じタイミングでSnapshotが作成された場合CodingSequenceは実行されていないものとする
 * @param editor CodingSequenceをテキスト形式に変換するのに使う
 * @param video CodingSequenceを見るため
 * @param timeSpan 何ミリ秒ごとにスナップショットを作成するか？
 *
 * @returns スナップショット
 *
 */

export const createSnapshot = (
  editor: CodeMirror.Editor,
  video: Video,
  timeSpan: number
): Snapshot[] => {
  // 途中のスナップショットが作れる作成範囲を調べる
  const CanBeCreatedRecodingTime = getPreviousTimeSpan(
    video.value.slice(-1)[0].timestamp,
    timeSpan
  )
  // 正規化処理
  const snapshots: Snapshot[] = []
  const stream = new CodingStream(video)
  // 開始地点
  readAndExecCodingSequence(editor, stream.current)
  stream.next()
  const fastData = editor.getValue()
  const fastTimestamp = 0
  snapshots.push(new Snapshot(fastTimestamp, fastData))
  // prettier-ignore
  for (let timestamp = timeSpan; timestamp <= CanBeCreatedRecodingTime; timestamp += timeSpan) {
    // TimeSpanと次のTimeSpanの間に入るCodingSequenceを実行
    while (timestamp > stream.current.timestamp) {
      readAndExecCodingSequence(editor, stream.current)
      stream.next()
    }
    // Snapshotを作成し次の処理へ移動する
    snapshots.push(new Snapshot(timestamp, editor.getValue()))
  }
  // 残ったCodingSequenceがあれば実行
  while (stream.isNext()) {
    readAndExecCodingSequence(editor, stream.current)
    stream.next()
  }
  // 最終地点
  const lastData = editor.getValue()
  snapshots.push(new Snapshot(video.header.recordingTime, lastData))
  // エディタの初期化処理
  editor.setValue('')
  return snapshots
}

export const readAndExecCodingSequence = (
  editor: CodeMirror.Editor,
  codingSequence: CodingSequence
): void => {
  if (codingSequence.changeData != null) {
    const { text, from, to, origin } = codingSequence.changeData
    editor.replaceRange(text, from, to, origin)
  }
  if (codingSequence.cursor != null) {
    editor.setCursor(codingSequence.cursor)
  }
}

export const getClosestTimeSpanIndex = (
  timestamp: number,
  timeSpan: number
): number => {
  const divisibleTimeSpan = Math.ceil(timestamp / timeSpan) * timeSpan
  if (divisibleTimeSpan == timestamp) {
    return Math.ceil(timestamp / timeSpan)
  }
  return Math.ceil(timestamp / timeSpan) - 1
}

const createNonCodingSequence = (timestamp: number): CodingSequence => {
  return { timestamp: timestamp, changeData: null, cursor: null }
}

/**
 * @function NormalizationForVideo ビデオを正規化する
 * timeSpanに合わせて空のCodingSequenceを入れるtimeSpanにCodingSequenceがある場合は何もしない
 * @param timeSpan CodingSequenceを入れる間隔
 * @param video 正規化処理を行う素となるVideo
 * @returns 正規化処理が終わったVideo
 */
export const NormalizationForVideo = (
  timeSpan: number,
  video: Video
): Video => {
  const processingVideo = JSON.parse(JSON.stringify(video)) as Video

  // 正規化処理ができる範囲を調べる
  const divisibleRecodingTime = getPreviousTimeSpan(
    processingVideo.value.slice(-1)[0].timestamp,
    timeSpan
  )
  // 正規化処理
  const codingSequence: CodingSequence[] = []
  let index = 0
  // prettier-ignore
  for (let timestamp = timeSpan; timestamp <= divisibleRecodingTime; timestamp += timeSpan) {

    // TimeSpanと次のTimeSpanの間に入るCodingSequenceを追加
    while (timestamp > processingVideo.value[index].timestamp) {
      codingSequence.push(processingVideo.value[index])
      index++
    }

    // TimeSpanにCodingSequenceがあれば追加，なければ空のCodingSequenceを追加する
    if (timestamp === processingVideo.value[index].timestamp) {
      codingSequence.push(processingVideo.value[index])
      index++
    } else {
      codingSequence.push(createNonCodingSequence(timestamp))
    }
  }
  // 残ったCodingSequenceがあれば追加
  while (index < processingVideo.value.length) {
    codingSequence.push(processingVideo.value[index])
    index++
  }
  return {
    header: processingVideo.header,
    value: codingSequence,
  }
}
