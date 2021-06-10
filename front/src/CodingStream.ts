import { CodingSequence } from './models/CodingSequence'
import { Video } from './models/Video'
import { VideoInfo } from './models/VideoInfo'

export class CodingStream {
  private _video: Video
  private _index = 0

  constructor(video: Video) {
    this._video = JSON.parse(JSON.stringify(video)) as Video
  }

  public toNormalization(timeSpan: number): CodingStream {
    const video = JSON.parse(JSON.stringify(this._video)) as Video
    // 最後の要素がキリの良い数字になるように調整
    const recordingTime = this.getNextTimeSpan(
      video.value.slice(-1)[0].timestamp,
      timeSpan
    )
    if (video.value.slice(-1)[0].timestamp !== recordingTime) {
      video.value.push(new CodingSequence(recordingTime, undefined, undefined))
    }
    // 正規化
    const codingSequence: CodingSequence[] = []
    let index = 0
    // prettier-ignore
    for (let timestamp = timeSpan; timestamp <= recordingTime; timestamp += timeSpan) {
      while (timestamp > video.value[index].timestamp) {
        codingSequence.push(video.value[index])
        index++
      }
      if (timestamp === video.value[index].timestamp) {
        codingSequence.push(video.value[index])
        index++
      } else {
        codingSequence.push(this.createNonCodingSequence(timestamp))
      }
    }
    video.value = codingSequence
    video.header.recordingTime = recordingTime
    console.log(video)
    return new CodingStream(video)
  }

  public reset(): void {
    this._index = 0
  }

  private getNextTimeSpan(timestamp: number, timeSpan: number): number {
    timestamp
    return Math.floor(timestamp / timeSpan) * timeSpan + timeSpan
  }

  private createNonCodingSequence(timestamp: number): CodingSequence {
    return new CodingSequence(timestamp, undefined, undefined)
  }

  public get videoInfo(): VideoInfo {
    return this._video.header
  }

  public get index(): number {
    return this._index
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
