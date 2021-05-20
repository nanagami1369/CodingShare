import { CodingSequence } from '@/models/CodingSequence'
import { Language } from './language'

export type Video = {
  header: {
    userId: number
    name: string
    title: string
    language: Language
    uploadTime: number
    recordingTime: number
  }
  value: CodingSequence[]
}
