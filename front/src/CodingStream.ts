import { CodingSequence } from './models/CodingSequence'
import { Video } from './models/Video'

export class CodingStream {
  private _video: Video
  private _index = 0

  constructor(video: Video) {
    this._video = JSON.parse(JSON.stringify(video)) as Video
  }

  public isNext(): boolean {
    const index = this._index + 1
    return index !== this._video.value.length
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
    return this._video.value[index]
  }

  public get current(): CodingSequence {
    const index = this._index
    return this._video.value[index]
  }

  public get to(): CodingSequence | undefined {
    const index = this._index + 1
    if (index === this._video.value.length) {
      return undefined
    }
    return this._video.value[index]
  }
}
