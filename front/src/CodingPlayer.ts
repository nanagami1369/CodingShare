import { CodingStream } from './CodingStream'
import { Video } from './models/Video'
import CodeMirror from 'codemirror'
import { VideoInfo } from './models/VideoInfo'

export class CodingPlayer {
  private _stream: CodingStream | undefined
  private _videoInfo: VideoInfo | undefined
  public load(video: Video, editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    this._videoInfo = video.header
    this._stream = new CodingStream(video)
    const { language } = this._videoInfo
    if (language == undefined) {
      throw new Error('video is not language data')
    }
    editor.setOption('mode', language.tag)
    editor.setValue('')
    editor.focus()
  }

  public start(
    editor: CodeMirror.Editor | undefined,
    appendDo: (stream: CodingStream) => void
  ): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._stream == undefined) {
      throw new Error('video is not Load')
    }

    doSomethingLoop((): { isNext: boolean; nextSpan: number } => {
      if (this._stream == undefined) {
        throw new Error('video is not Load')
      }
      if (this._stream.current.changeData != undefined) {
        const { text, from, to, origin } = this._stream.current.changeData
        editor.replaceRange(text, from, to, origin)
      }
      if (this._stream.current.cursor != undefined) {
        const cursor = this._stream.current.cursor
        editor.setCursor(cursor)
      }
      appendDo(this._stream)
      this._stream.next()
      const isNext = this._stream.isNext()
      if (this._stream.from === undefined) {
        appendDo(this._stream)
        return { isNext: isNext, nextSpan: this._stream.current.timestamp }
      }
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
        appendDo(this._stream)
        return { isNext: isNext, nextSpan: 1 }
      }
      appendDo(this._stream)
      return {
        isNext: isNext,
        nextSpan: this._stream.to.timestamp - this._stream.current.timestamp,
      }
    })
  }

  public get videoInfo(): VideoInfo {
    if (this._videoInfo == undefined) {
      throw new Error('video is not Load')
    }
    return this._videoInfo
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
