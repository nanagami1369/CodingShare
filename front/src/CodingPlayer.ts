import { CodingStream } from './CodingStream'
import { Video } from './models/Video'
import CodeMirror from 'codemirror'
import { VideoInfo } from './models/VideoInfo'
import { PlayerInfo } from '@/models/PlayerInfo'
import { Snapshot } from './models/Snapshot'
import { AsyncLoop } from './AsyncLoop'
import NormalizationWorker from 'worker-loader!@/worker/Normalization.worker.ts'
import {
  createSnapshot,
  getClosestTimeSpanIndex,
  identificationBetweenSequences,
  readAndExecCodingSequence,
} from './util'

export class CodingPlayer {
  private _stream?: CodingStream
  private _snapShotTimeSpan: number
  private _info: PlayerInfo = {
    elapsedTime: 0,
    totalTime: 0,
    speed: 100,
  }
  private _isLoading = false

  public get isLoading(): boolean {
    return this._isLoading
  }

  private _snapshot: Snapshot[] = []

  private _isPlay = false
  public get isPlay(): boolean {
    return this._isPlay
  }

  public set isPlay(value: boolean) {
    if (!value) {
      this._asyncLoop.stop()
    }
    this._isPlay = value
  }

  private _asyncLoop = new AsyncLoop()

  public constructor(snapShotTimeSpan: number, speed: number) {
    this._snapShotTimeSpan = snapShotTimeSpan
    this.setSpeed(speed)
  }

  public setSpeed(speed: number): void {
    this._info.speed = speed
  }

  private setElapsedTime(stream: CodingStream): void {
    this._info.elapsedTime = stream.current.timestamp
  }

  public clear(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    editor.setValue('')
    this._info.elapsedTime = 0
    this._info.totalTime = 0
    this._stream = undefined
    this._snapshot = []
  }

  public load(
    video: Video,
    editor?: CodeMirror.Editor,
    backgroundEditor?: CodeMirror.Editor
  ): void {
    this._isLoading = true
    console.time('total time')
    const worker = new NormalizationWorker()
    worker.onmessage = (message: MessageEvent<Video>) => {
      const normalizationVideo = message.data
      this._stream = new CodingStream(normalizationVideo)
      this._info.totalTime = this._stream.videoInfo.recordingTime
      this.setElapsedTime(this._stream)
      this._stream.next()
      console.timeEnd('nor time')
      console.log('nor end')
      console.timeEnd('total time')
      this._isLoading = false
      worker.terminate()
    }
    console.log('nor start')
    console.time('nor time')
    worker.postMessage({ timeSpan: 500, video: video })

    // スナップショット作成
    if (backgroundEditor == null) {
      throw new Error('backgroundEditor is undefined')
    }
    setTimeout(() => {
      console.log('snapshot start')
      console.time('snapshot time')
      this._snapshot = createSnapshot(
        backgroundEditor,
        video,
        this._snapShotTimeSpan
      )
      console.timeEnd('snapshot time')
      console.log('snapshot end')
    }, 0)
    // エディタ準備
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    const { language } = video.header
    if (language == null) {
      throw new Error('video is not language data')
    }
    editor.setOption('mode', language.tag)
    editor.setValue('')
    // 最初の要素を描画
    readAndExecCodingSequence(editor, video.value[0])
  }

  public start(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == null) {
      throw new Error('video is not Load')
    }
    this._isPlay = true
    this._asyncLoop.loop((): { isNext: boolean; nextSpan: number } => {
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

  public move(time: number, editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == null) {
      throw new Error('stream is undefined')
    }
    let snapshot: Snapshot | null = null
    // スナップショットを展開
    if (time === this.videoInfo?.recordingTime) {
      snapshot = this._snapshot.slice(-1)[0]
      editor.setValue(snapshot.value)
      this._stream.seek(this._stream.length - 1)
      return
    }

    if (time === 0) {
      snapshot = this._snapshot[0]
      editor.setValue(snapshot.value)
      this._stream.reset()
      this._stream.next()
      return
    }

    const snapshotIndex = getClosestTimeSpanIndex(time, this._snapShotTimeSpan)
    snapshot = this._snapshot[snapshotIndex]
    editor.setValue(snapshot.value)
    const { index, indexStatus } = identificationBetweenSequences(
      snapshot.timestamp,
      this._stream
    )
    switch (indexStatus) {
      case 'None':
        // 基本的にこのエラーは出ない
        throw new Error('データが見つかりませんでした')
      case 'Equals':
        this._stream?.seek(index)
        break
      case 'Greater':
        this._stream.seek(index + 1)
        break
      case 'Smaller':
        this._stream.seek(index)
        break
    }
    while (this._stream.to != null && this._stream.current.timestamp <= time) {
      readAndExecCodingSequence(editor, this._stream.current)
      this.setElapsedTime(this._stream)
      this._stream.next()
    }
  }

  public get videoInfo(): VideoInfo | undefined {
    return this._stream?.videoInfo
  }

  public get info(): Readonly<PlayerInfo> {
    return this._info
  }

  public get isLoaded(): boolean {
    return this._stream != null && this._snapshot != null
  }
}
