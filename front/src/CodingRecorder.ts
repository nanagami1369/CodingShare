import CodeMirror from 'codemirror'
import { CodingSequence } from '@/models/CodingSequence'
export class CodingRecorder {
  private _data: CodingSequence[] = []
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
    const scrollInfo = editor.getScrollInfo()
    const codingSequence = new CodingSequence(changeObj, scrollInfo)
    this._data.push(codingSequence)
    console.log(JSON.stringify(codingSequence))
  }

  public getData(): CodingSequence[] {
    return this._data.concat()
  }
}