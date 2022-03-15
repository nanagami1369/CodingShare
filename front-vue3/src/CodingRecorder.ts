import type CodeMirror from 'codemirror'
import type { CodingSequence } from '@/models/CodingSequence'
import type { Video } from '@/models/Video'
import type { Language } from '@/models/language'

export class CodingRecorder {
  private _video: CodingSequence[] = []
  private _timer: number = new Date().getTime()
  private _isRecording = false
  private _uploadTime = -1

  public register(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    editor.on(
      'change',
      (editor: CodeMirror.Editor, changeObj: CodeMirror.EditorChange) => {
        this.record(editor, changeObj)
      }
    )
  }

  public unregister(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    editor.on(
      'change',
      (editor: CodeMirror.Editor, changeObj: CodeMirror.EditorChange) => {
        this.record(editor, changeObj)
      }
    )
  }

  private record(
    this: CodingRecorder,
    editor: CodeMirror.Editor,
    changeObj: CodeMirror.EditorChange
  ): void {
    if (this._isRecording) {
      const cursor = editor.getCursor()
      const time = new Date().getTime() - this._timer
      const codingSequence: CodingSequence = {
        timestamp: time,
        changeData: changeObj,
        cursor: cursor,
      }
      this._video.push(codingSequence)
      console.log(JSON.stringify(codingSequence))
    }
  }

  public start(editor?: CodeMirror.Editor): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    if (this._isRecording) {
      throw new Error('Recorder has already started')
    }
    if (this._video.length != 0) {
      throw new Error('Video exists !! Clear or output video')
    }
    const startData = editor.getValue().split('\n')
    const startCursor = editor.getCursor()
    const startTimestamp = 0
    const startChangeData: CodeMirror.EditorChange = {
      from: { line: 0, ch: 0, sticky: undefined },
      to: { line: 0, ch: 0, sticky: undefined },
      text: startData,
      removed: [''],
      origin: '+input',
    }
    const startCodingSequence: CodingSequence = {
      timestamp: startTimestamp,
      changeData: startChangeData,
      cursor: startCursor,
    }
    this._video.push(startCodingSequence)
    this._isRecording = true
    this._timer = new Date().getTime()
  }
  public stop(): void {
    if (!this._isRecording) {
      throw new Error('Recorder is not start')
    }
    this._isRecording = false
    const time = new Date().getTime()
    this._uploadTime = time
    const recordingTime = time - this._timer
    const lastCodingSequence: CodingSequence = {
      timestamp: recordingTime,
      changeData: null,
      cursor: null,
    }
    this._video.push(lastCodingSequence)
  }

  public outputVideo(
    userId: number,
    name: string,
    title: string,
    language: Language,
    comment: string
  ): Video {
    if (this._video.length == 0) {
      throw new Error('video is not found')
    }
    if (this._isRecording) {
      throw new Error('is recording')
    }
    const video = this._video.concat()
    this._video = []
    return {
      header: {
        userId: userId,
        name: name,
        title: title,
        language: language,
        uploadTime: this._uploadTime,
        comment: comment,
        recordingTime: video.slice(-1)[0].timestamp,
      },
      value: video,
    }
  }

  public clearVideo(): void {
    this._video = []
  }

  public get isRecording(): boolean {
    return this._isRecording
  }
}
