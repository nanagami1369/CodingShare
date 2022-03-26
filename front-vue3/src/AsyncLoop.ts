/**
 * @class AsyncLoop 非同期ループ
 */
export class AsyncLoop {
  private _timeOutId = -1
  public get timeOutId(): number {
    return this._timeOutId
  }

  /**
   * @function loop Do AsyncLoop
   * @param this AsyncLoop in instance
   * @param func function in process for loop
   */
  public loop(
    this: AsyncLoop,
    func: () => { nextSpan: number; isNext: boolean }
  ): void {
    const { isNext, nextSpan } = func()
    if (isNext) {
      this._timeOutId = setTimeout((): void => {
        this.loop(func)
      }, nextSpan)
    }
  }

  /**
   * @function stop Stop AsyncLoop
   */
  public stop(): void {
    clearTimeout(this.timeOutId)
    this._timeOutId = -1
  }
}
