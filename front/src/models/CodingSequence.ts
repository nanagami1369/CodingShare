import CodeMirror from 'codemirror'

export class CodingSequence {
  private _changeData: CodeMirror.EditorChangeLinkedList
  private _scrollInfo: CodeMirror.ScrollInfo

  constructor(
    changeData: CodeMirror.EditorChangeLinkedList,
    scrollInfo: CodeMirror.ScrollInfo
  ) {
    this._changeData = changeData
    this._scrollInfo = scrollInfo
  }

  public get changeData(): CodeMirror.EditorChangeLinkedList {
    return this._changeData
  }

  public get scrollInfo(): CodeMirror.ScrollInfo {
    return this._scrollInfo
  }
}
