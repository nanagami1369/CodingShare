import { Language } from '@/models/language'

export type ViewVideo = {
  userId: number
  title: string
  name: string
  language: Language | undefined
  uploadTime: number
  recordingTime: number
}
