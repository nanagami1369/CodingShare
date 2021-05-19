import CodeMirror from 'codemirror'

export class CodingSequence {
  constructor(
    public timestamp: number,
    public changeData: CodeMirror.EditorChangeLinkedList,
    public cursor: CodeMirror.Position
  ) {}
}
