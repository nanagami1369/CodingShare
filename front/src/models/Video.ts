import { CodingSequence } from '@/models/CodingSequence'
import { VideoInfo } from '@/models/VideoInfo'

export type Video = {
  header: VideoInfo
  value: CodingSequence[]
}
