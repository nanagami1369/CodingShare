import type { Video } from '@/models/Video'
import { NormalizationForVideo } from '@/util'

onmessage = (message: MessageEvent<{ timeSpan: number; video: Video }>) => {
  const { timeSpan, video } = message.data
  const normalizationForVideo = NormalizationForVideo(timeSpan, video)
  postMessage(normalizationForVideo)
}
