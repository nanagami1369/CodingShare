import CodeMirror from 'codemirror'
import { CodingSequence } from '@/models/CodingSequence'

export class CodingRecorder {
  private _data: CodingSequence[] = []
  private _timer: number = new Date().getTime()
  private _isRecording = false
  
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

  public stop(): void {
    if (!this._isRecording) {
      throw new Error('Recorder is not start')
    }
    this._isRecording = false
    console.log(JSON.stringify(this.getData(), null, 4))
  }

  public getData(): CodingSequence[] {
    return this._data.concat()
  }

  public get isRecording(): boolean {
    return this._isRecording
  }
}
