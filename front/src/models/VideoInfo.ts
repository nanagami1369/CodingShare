import { Language } from '@/models/language'

export type VideoInfo = {
  userId: string
  title: string
  name: string
  language?: Language
  uploadTime: number
  recordingTime: number
  comment: string
}
