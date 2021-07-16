import { getPreviousTimeSpan } from '@/getPreviousTimeSpan'
import { CodingSequence } from '@/models/CodingSequence'
import { Video } from '@/models/Video'

onmessage = (message: MessageEvent<{ timeSpan: number; video: Video }>) => {
  const { timeSpan, video } = message.data
  const normalizationForVideo = NormalizationForVideo(timeSpan, video)
  postMessage(normalizationForVideo)
}

/**
 * @function NormalizationForVideo ビデオを正規化する
 * timeSpanに合わせて空のCodingSequenceを入れるtimeSpanにCodingSequenceがある場合は何もしない
 * @param timeSpan CodingSequenceを入れる間隔
 * @param video 正規化処理を行う素となるVideo
 * @returns 正規化処理が終わったVideo
 */
const NormalizationForVideo = (timeSpan: number, video: Video): Video => {
  const processingVideo = JSON.parse(JSON.stringify(video)) as Video

  // 正規化処理ができる範囲を調べる
  const divisibleRecodingTime = getPreviousTimeSpan(
    processingVideo.value.slice(-1)[0].timestamp,
    timeSpan
  )
  // 正規化処理
  const codingSequence: CodingSequence[] = []
  let index = 0
  // prettier-ignore
  for (let timestamp = timeSpan; timestamp <= divisibleRecodingTime; timestamp += timeSpan) {

    // TimeSpanと次のTimeSpanの間に入るCodingSequenceを追加
    while (timestamp > processingVideo.value[index].timestamp) {
      codingSequence.push(processingVideo.value[index])
      index++
    }

    // TimeSpanにCodingSequenceがあれば追加，なければ空のCodingSequenceを追加する
    if (timestamp === processingVideo.value[index].timestamp) {
      codingSequence.push(processingVideo.value[index])
      index++
    } else {
      codingSequence.push(createNonCodingSequence(timestamp))
    }
  }
  // 残ったCodingSequenceがあれば追加
  while (index < processingVideo.value.length) {
    codingSequence.push(processingVideo.value[index])
    index++
  }
  return {
    header: processingVideo.header,
    value: codingSequence,
  }
}

const createNonCodingSequence = (timestamp: number): CodingSequence => {
  return { timestamp: timestamp, changeData: null, cursor: null }
}
