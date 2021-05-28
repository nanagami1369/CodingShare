import { CodingSequence } from './models/CodingSequence'
import { Video } from './models/Video'

export class CodingStream {
  private _video: Video
  private _index = 0

  constructor(video: Video) {
    this._video = JSON.parse(JSON.stringify(video)) as Video
  }

  public toNormalization(): CodingStream {
    const video = JSON.parse(JSON.stringify(this._video)) as Video
    const codingSequence: CodingSequence[] = []
    const firstCodingSequence = video.value[0]
    codingSequence.push(firstCodingSequence)
    for (let index = 1; index < video.value.length; index++) {
      const startTimestamp = video.value[index - 1].timestamp
      const elapsedTime =
        video.value[index].timestamp - video.value[index - 1].timestamp
      codingSequence.push(
        ...this.createNonCodingSequence(startTimestamp, elapsedTime)
      )
      codingSequence.push(video.value[index])
    }
    console.log(codingSequence)
    console.log(video)
    video.value = codingSequence
    console.log(video)
    return new CodingStream(video)
  }

  private createNonCodingSequence(
    startTimestamp: number,
    elapsedTime: number
  ): CodingSequence[] {
    const oneSecond = 1000
    const codingSequence: CodingSequence[] = []
    const howManyTimes1SecondPassed = Math.floor(elapsedTime / oneSecond)
    let currentTimestamp = startTimestamp
    for (let count = 0; count < howManyTimes1SecondPassed; count++) {
      currentTimestamp += oneSecond
      codingSequence.push(
        new CodingSequence(currentTimestamp, undefined, undefined)
      )
    }
    return codingSequence
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
