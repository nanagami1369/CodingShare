import CodeMirror from 'codemirror'

export class CodingSequence {
  public timestamp: number
  public changeData: CodeMirror.EditorChangeLinkedList
  public scrollInfo: CodeMirror.ScrollInfo

  constructor(
    timestamp: number,
    changeData: CodeMirror.EditorChangeLinkedList,
    scrollInfo: CodeMirror.ScrollInfo
  ) {
    this.timestamp = timestamp
    this.changeData = changeData
    this.scrollInfo = scrollInfo
  }
}
