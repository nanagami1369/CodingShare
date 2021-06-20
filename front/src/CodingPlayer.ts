import { CodingStream } from './CodingStream'
import { Video } from './models/Video'
import CodeMirror from 'codemirror'
import { VideoInfo } from './models/VideoInfo'
import { PlayerInfo } from '@/models/PlayerInfo'
import { CodingSequence } from './models/CodingSequence'
import { Snapshot } from './models/Snapshot'
export class CodingPlayer {
  private _stream: CodingStream | undefined
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
    editor: CodeMirror.Editor | undefined,
    backgroundEditor: CodeMirror.Editor | undefined
  ): void {
    // スナップショット作成
    if (backgroundEditor == null) {
      throw new Error('backgroundEditor is undefined')
    }
    setTimeout(() => {
      this._snapshot = createSnapshot(video, backgroundEditor)
      console.log(this._snapshot)
    }, 0)
    // エディタ準備
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    this._stream = new CodingStream(video)
    const { language } = video.header
    if (language == undefined) {
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

  public start(editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == undefined) {
      throw new Error('video is not Load')
    }
    this._isPlay = true
    doSomethingLoop(this, (): { isNext: boolean; nextSpan: number } => {
      if (this._stream == undefined) {
        throw new Error('video is not Load')
      }
      if (!this.isPlay) {
        return { isNext: false, nextSpan: 0 }
      }
      readAndExecCodingSequence(editor, this._stream.current)
      this.setElapsedTime(this._stream)
      this._stream.next()
      const isNext = this._stream.isNext()
      if (this._stream.to === undefined) {
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

  public backToTheBeginning(editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == undefined) {
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

  public stepForward(editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == undefined) {
      throw new Error('video is not Load')
    }
    this.pause()
    while (
      this._stream.to != undefined &&
      this._stream.current.changeData == undefined &&
      this._stream.current.cursor == undefined
    ) {
      console.log(this._stream.current)
      this._stream.next()
    }
    console.log(`do!:${JSON.stringify(this._stream.current)}`)
    readAndExecCodingSequence(editor, this._stream.current)
    this.setElapsedTime(this._stream)
    this._stream.next()
  }

  public fastForward(editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == undefined) {
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

const readAndExecCodingSequence = (
  editor: CodeMirror.Editor,
  codingSequence: CodingSequence
) => {
  if (codingSequence.changeData != undefined) {
    const { text, from, to, origin } = codingSequence.changeData
    editor.replaceRange(text, from, to, origin)
  }
  if (codingSequence.cursor != undefined) {
    editor.setCursor(codingSequence.cursor)
  }
}

const createSnapshot = (
  video: Video,
  editor: CodeMirror.Editor
): Snapshot[] => {
  const stream = new CodingStream(video)
  const snapshots: Snapshot[] = []
  readAndExecCodingSequence(editor, stream.current)
  const fastData = editor.getValue()
  snapshots.push(new Snapshot(fastData))
  while (stream.to != undefined) {
    readAndExecCodingSequence(editor, stream.current)
    stream.next()
  }
  const lastData = editor.getValue()
  editor.setValue('')
  snapshots.push(new Snapshot(lastData))
  return snapshots
}

const getNextTimeSpan = (timestamp: number, timeSpan: number): number => {
  timestamp
  return Math.floor(timestamp / timeSpan) * timeSpan + timeSpan
}

const createNonCodingSequence = (timestamp: number): CodingSequence => {
  return { timestamp: timestamp, changeData: undefined, cursor: undefined }
}

const NormalizationForVideo = (timeSpan: number, video: Video): Video => {
  const processingVideo = JSON.parse(JSON.stringify(video)) as Video

  // 最後の要素のTimestampがTimeSpanで割り切れる数字になるように調整
  const recordingTime = getNextTimeSpan(
    processingVideo.value.slice(-1)[0].timestamp,
    timeSpan
  )
  if (processingVideo.value.slice(-1)[0].timestamp !== recordingTime) {
    processingVideo.value.push(createNonCodingSequence(recordingTime))
  }

  //最後の要素のTimestamp を書き換えたものに調整
  const header = processingVideo.header
  header.recordingTime = recordingTime

  // 正規化
  const codingSequence: CodingSequence[] = []
  let index = 0
  // prettier-ignore
  for (let timestamp = timeSpan; timestamp <= recordingTime; timestamp += timeSpan) {

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

  return {
    header: processingVideo.header,
    value: codingSequence,
  }
}
