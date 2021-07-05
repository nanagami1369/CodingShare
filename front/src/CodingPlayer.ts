import { CodingStream } from './CodingStream'
import { Video } from './models/Video'
import CodeMirror from 'codemirror'
import { VideoInfo } from './models/VideoInfo'
import { PlayerInfo } from '@/models/PlayerInfo'
import { CodingSequence } from './models/CodingSequence'
import { Snapshot } from './models/Snapshot'

export class CodingPlayer {
  private _stream?: CodingStream
  private _info: PlayerInfo = {
    elapsedTime: 0,
    totalTime: 0,
    speed: 100,
  }
  private _snapshot: Snapshot[] = []

  private _currentTimeoutId = -1
  public set currentTimeoutId(value: number) {
    this._currentTimeoutId = value
  }
  public get currentTimeoutId(): number {
    return this._currentTimeoutId
  }

  private _isPlay = false
  public get isPlay(): boolean {
    return this._isPlay
  }

  public set isPlay(value: boolean) {
    if (!value) {
      clearTimeout(this.currentTimeoutId)
    }
    this._isPlay = value
  }

  private setElapsedTime(stream: CodingStream): void {
    this._info.elapsedTime = stream.current.timestamp
  }

  public load(
    video: Video,
    editor?: CodeMirror.Editor,
    backgroundEditor?: CodeMirror.Editor
  ): void {
    // スナップショット作成
    if (backgroundEditor == null) {
      throw new Error('backgroundEditor is undefined')
    }
    setTimeout(() => {
      this._snapshot = createSnapshot(backgroundEditor, video, 30000)
    }, 0)
    // エディタ準備
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    this._stream = new CodingStream(video)
    const { language } = video.header
    if (language == null) {
      throw new Error('video is not language data')
    }
    editor.setOption('mode', language.tag)
    editor.setValue('')
    const normalizationVideo = NormalizationForVideo(500, video)
    this._stream = new CodingStream(normalizationVideo)
    this._info.totalTime = this._stream.videoInfo.recordingTime
    // 最初の要素を描画
    readAndExecCodingSequence(editor, this._stream.current)
    this.setElapsedTime(this._stream)
    this._stream.next()
  }

  public start(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == null) {
      throw new Error('video is not Load')
    }
    this._isPlay = true
    doSomethingLoop(this, (): { isNext: boolean; nextSpan: number } => {
      if (this._stream == null) {
        throw new Error('video is not Load')
      }
      if (!this.isPlay) {
        return { isNext: false, nextSpan: 0 }
      }
      readAndExecCodingSequence(editor, this._stream.current)
      this.setElapsedTime(this._stream)
      this._stream.next()
      const isNext = this._stream.isNext()
      if (this._stream.to == null) {
        // 次の要素が無いので最後の要素を表示して終了
        console.log('終了')
        readAndExecCodingSequence(editor, this._stream.current)
        this.setElapsedTime(this._stream)
        this.isPlay = false
        return { isNext: isNext, nextSpan: 1 }
      }
      this.setElapsedTime(this._stream)
      return {
        isNext: isNext,
        nextSpan:
          // prettier-ignore
          (this._stream.to.timestamp - this._stream.current.timestamp) / (this.info.speed / 100),
      }
    })
  }

  public pause(): void {
    this.isPlay = false
  }

  public backToTheBeginning(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == null) {
      throw new Error('video is not Load')
    }
    this.pause()
    editor.setValue('')
    this._stream.reset()
    // 最初の要素を描画
    readAndExecCodingSequence(editor, this._stream.current)
    this.setElapsedTime(this._stream)
    this._stream.next()
  }

  public stepForward(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == null) {
      throw new Error('video is not Load')
    }
    this.pause()
    while (
      this._stream.to != null &&
      this._stream.current.changeData == null &&
      this._stream.current.cursor == null
    ) {
      console.log(this._stream.current)
      this._stream.next()
    }
    console.log(`do!:${JSON.stringify(this._stream.current)}`)
    readAndExecCodingSequence(editor, this._stream.current)
    this.setElapsedTime(this._stream)
    this._stream.next()
  }

  public fastForward(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == null) {
      throw new Error('video is not Load')
    }
    this.pause()
    editor.setValue(this._snapshot.slice(-1)[0].value)
    this._info.elapsedTime = this._stream.videoInfo.recordingTime
    this._stream.seek(this._stream.length - 1)
  }

  public get videoInfo(): VideoInfo | undefined {
    return this._stream?.videoInfo
  }

  public get info(): PlayerInfo {
    return this._info
  }

  public get isLoaded(): boolean {
    return this._stream != null && this._snapshot != null
  }
}

function doSomethingLoop(
  player: CodingPlayer,
  doSomething: () => { nextSpan: number; isNext: boolean }
): void {
  const { isNext, nextSpan } = doSomething()
  if (isNext) {
    player.currentTimeoutId = setTimeout(function () {
      doSomethingLoop(player, doSomething)
    }, nextSpan)
  }
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
    snapshots.push(new Snapshot(timestamp,editor.getValue()))
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

const readAndExecCodingSequence = (
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

const getPreviousTimeSpan = (timestamp: number, timeSpan: number): number => {
  timestamp
  return Math.ceil(timestamp / timeSpan) * timeSpan - timeSpan
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
const NormalizationForVideo = (timeSpan: number, video: Video): Video => {
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
