import { CodingSequence } from './models/CodingSequence'
import { Video } from './models/Video'
import { VideoInfo } from './models/VideoInfo'

export class CodingStream {
  private _video: Video
  private _index = 0

  constructor(video: Video) {
    this._video = JSON.parse(JSON.stringify(video)) as Video
  }

  public reset(): void {
    this._index = 0
  }

  public seek(index: number): boolean {
    const isSeek = 0 <= index && index < this._video.value.length
    if (isSeek) {
      this._index = index
    }
    return isSeek
  }

  public get videoInfo(): VideoInfo {
    return this._video.header
  }

  public get index(): number {
    return this._index
  }

  public get length(): number {
    return this._video.value.length
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
