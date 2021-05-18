import CodeMirror from 'codemirror'

export class CodingSequence {
  constructor(
    public timestamp: number,
    public changeData: CodeMirror.EditorChangeLinkedList,
    public scrollInfo: CodeMirror.ScrollInfo
  ) {}
}
