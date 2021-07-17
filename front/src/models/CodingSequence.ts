import CodeMirror from 'codemirror'

export type CodingSequence = {
  timestamp: number
  changeData: CodeMirror.EditorChangeLinkedList | null
  cursor: CodeMirror.Position | null
}
