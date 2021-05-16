import CodeMirror from 'codemirror'

export class CodingSequence {
  private _time: number
  private _changeData: CodeMirror.EditorChangeLinkedList
  private _scrollInfo: CodeMirror.ScrollInfo

  constructor(
    time: number,
    changeData: CodeMirror.EditorChangeLinkedList,
    scrollInfo: CodeMirror.ScrollInfo
  ) {
    this._time = time
    this._changeData = changeData
    this._scrollInfo = scrollInfo
  }

  public get time(): number {
    return this._time
  }

  public get changeData(): CodeMirror.EditorChangeLinkedList {
    return this._changeData
  }

  public get scrollInfo(): CodeMirror.ScrollInfo {
    return this._scrollInfo
  }
}
