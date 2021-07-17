import CodeMirror from 'codemirror'

export type CodingSequence = {
  timestamp: number
  changeData: CodeMirror.EditorChange | null
  cursor: CodeMirror.Position | null
}
