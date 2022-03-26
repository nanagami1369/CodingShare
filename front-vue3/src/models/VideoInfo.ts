import type { Language } from '@/models/language'

export type VideoInfo = {
  userId: number
  title: string
  name: string
  language?: Language
  uploadTime: number
  recordingTime: number
  comment: string
}
