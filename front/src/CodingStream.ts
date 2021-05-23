import { CodingSequence } from './models/CodingSequence'
import { Video } from './models/Video'

export class CodingStream {
  private _codingSequence: CodingSequence[]
  private _index = 0

  constructor(video: Video) {
    this._codingSequence = video.value.concat()
  }

  public isNext(): boolean {
    const index = this._index + 1
    return index !== this._codingSequence.length
  }

  public next(): boolean {
    if (this.isNext()) {
      this._index++
      return true
    }
    return false
  }
  public get from(): CodingSequence | undefined {
    const index = this._index - 1
    if (index === -1) {
      return undefined
    }
    return this._codingSequence[index]
  }

  public get current(): CodingSequence {
    const index = this._index
    return this._codingSequence[index]
  }

  public get to(): CodingSequence | undefined {
    const index = this._index + 1
    if (index === this._codingSequence.length) {
      return undefined
    }
    return this._codingSequence[index]
  }
}
