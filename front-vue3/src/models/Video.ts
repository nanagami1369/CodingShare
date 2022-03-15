import type { CodingSequence } from '@/models/CodingSequence'
import type { VideoInfo } from '@/models/VideoInfo'

export type Video = {
  header: VideoInfo
  value: CodingSequence[]
}
