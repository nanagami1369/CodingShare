import { CodingStream } from './CodingStream'
import { Video } from './models/Video'
import CodeMirror from 'codemirror'
import { VideoInfo } from './models/VideoInfo'
import { PlayerInfo } from '@/models/PlayerInfo'
export class CodingPlayer {
  private _stream: CodingStream | undefined
  private _info: PlayerInfo = {
    elapsedTime: 0,
    totalTime: 0,
    isPlay: false,
  }

  private setElapsedTime(stream: CodingStream): void {
    this._info.elapsedTime = stream.current.timestamp
  }

  public load(video: Video, editor: CodeMirror.Editor | undefined): void {
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
    editor.focus()
    this._stream = this._stream.toNormalization(500)
    this._info.totalTime = this._stream.videoInfo.recordingTime
    // 最初の要素を描画
    if (this._stream.current.changeData != undefined) {
      const { text, from, to, origin } = this._stream.current.changeData
      editor.replaceRange(text, from, to, origin)
    }
    if (this._stream.current.cursor != undefined) {
      const cursor = this._stream.current.cursor
      editor.setCursor(cursor)
    }
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
    this._info.isPlay = true
    doSomethingLoop((): { isNext: boolean; nextSpan: number } => {
      if (this._stream == undefined) {
        throw new Error('video is not Load')
      }
      if (!this.info.isPlay) {
        return { isNext: false, nextSpan: 0 }
      }
      if (this._stream.current.changeData != undefined) {
        const { text, from, to, origin } = this._stream.current.changeData
        editor.replaceRange(text, from, to, origin)
      }
      if (this._stream.current.cursor != undefined) {
        const cursor = this._stream.current.cursor
        editor.setCursor(cursor)
      }
      this.setElapsedTime(this._stream)
      this._stream.next()
      const isNext = this._stream.isNext()
      if (this._stream.to === undefined) {
        // 次の要素が無いので最後の要素を表示して終了
        console.log('終了')
        if (this._stream.current.changeData != undefined) {
          const { text, from, to, origin } = this._stream.current.changeData
          editor.replaceRange(text, from, to, origin)
        }
        if (this._stream.current.cursor != undefined) {
          const cursor = this._stream.current.cursor
          editor.setCursor(cursor)
        }
        this.setElapsedTime(this._stream)
        this._info.isPlay = false
        return { isNext: isNext, nextSpan: 1 }
      }
      this.setElapsedTime(this._stream)
      return {
        isNext: isNext,
        nextSpan: this._stream.to.timestamp - this._stream.current.timestamp,
      }
    })
  }

  public pause(): void {
    this.info.isPlay = false
  }
  public get videoInfo(): VideoInfo | undefined {
    return this._stream?.videoInfo
  }

  public get info(): PlayerInfo {
    return this._info
  }

  public get isLoaded(): boolean {
    return this._stream !== undefined
  }
}

function doSomethingLoop(
  doSomething: () => { nextSpan: number; isNext: boolean }
): void {
  const { isNext, nextSpan } = doSomething()
  if (isNext) {
    setTimeout(function () {
      doSomethingLoop(doSomething)
    }, nextSpan)
  }
}
