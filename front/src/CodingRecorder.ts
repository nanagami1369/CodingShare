import CodeMirror from 'codemirror'
import { CodingSequence } from '@/models/CodingSequence'
import { CordingSequenceData } from './models/CordingSequenceData'
import { Language } from './models/language'

export class CodingRecorder {
  private _data: CodingSequence[] = []
  private _timer: number = new Date().getTime()
  private _isRecording = false
  private _userId?: number
  private _name?: string
  private _title?: string
  private _language?: Language

  public register(editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    editor.on(
      'change',
      (
        editor: CodeMirror.Editor,
        changeObj: CodeMirror.EditorChangeLinkedList
      ) => {
        this.record(editor, changeObj)
      }
    )
  }

  public unregister(editor: CodeMirror.Editor | undefined): void {
    if (editor == null) {
      throw new Error('editor is undefined')
    }
    editor.on(
      'change',
      (
        editor: CodeMirror.Editor,
        changeObj: CodeMirror.EditorChangeLinkedList
      ) => {
        this.record(editor, changeObj)
      }
    )
  }

  private record(
    this: CodingRecorder,
    editor: CodeMirror.Editor,
    changeObj: CodeMirror.EditorChangeLinkedList
  ): void {
    if (this._isRecording) {
      const scrollInfo = editor.getScrollInfo()
      const time = new Date().getTime() - this._timer
      const codingSequence = new CodingSequence(time, changeObj, scrollInfo)
      this._data.push(codingSequence)
      console.log(JSON.stringify(codingSequence))
    }
  }

  public start(): void {
    if (this._isRecording) {
      throw new Error('Recorder has already started')
    }
    this._isRecording = true
    this._timer = new Date().getTime()
  }

  public stop(
    userId: number,
    name: string,
    title: string,
    language: Language
  ): CordingSequenceData {
    if (!this._isRecording) {
      throw new Error('Recorder is not start')
    }
    this._isRecording = false
    const time = new Date().getTime()
    return {
      header: {
        userId: userId,
        name: name,
        title: title,
        language: language,
        uploadTime: time,
        recordingTime: time - this._timer,
      },
      value: this.getData(),
    }
  }

  public getData(): CodingSequence[] {
    return this._data.concat()
  }

  public get isRecording(): boolean {
    return this._isRecording
  }
}
