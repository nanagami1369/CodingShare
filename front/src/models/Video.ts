import { CodingSequence } from '@/models/CodingSequence'
import { ViewVideo } from './ViewVideo'

export type Video = {
  header: ViewVideo
  value: CodingSequence[]
}
