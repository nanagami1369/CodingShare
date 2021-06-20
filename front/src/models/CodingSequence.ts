import CodeMirror from 'codemirror'

export type CodingSequence = {
  timestamp: number
  changeData: CodeMirror.EditorChangeLinkedList | undefined
  cursor: CodeMirror.Position | undefined
}
