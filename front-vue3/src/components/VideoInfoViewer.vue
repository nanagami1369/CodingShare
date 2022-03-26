<script setup lang="ts">
import type { VideoInfo } from '@/models/VideoInfo'
import { format } from 'date-fns'
import { formatRecordingTime } from '@/util'
import { computed, type PropType } from 'vue'

const props = defineProps({
  videoInfo: {
    type: Object as PropType<VideoInfo>,
    default: (): VideoInfo => ({
      userId: -1,
      title: '',
      name: '',
      language: undefined,
      comment: '',
      uploadTime: -1,
      recordingTime: -1,
    }),
    required: false,
  },
})
const language = computed((): string => props.videoInfo.language?.name ?? '')
const uploadTime = computed((): string => {
  if (props.videoInfo.uploadTime < 0) {
    return ''
  }
  return format(props.videoInfo.uploadTime, 'yyyy/MM/dd hh:mm:ss')
})
const recordingTime = computed((): string => {
  if (props.videoInfo.recordingTime < 0) {
    return ''
  }
  return formatRecordingTime(props.videoInfo.recordingTime)
})
</script>

<template>
  <div id="video-info-viewer">
    <p>
      タイトル:<cite>{{ videoInfo.title }}</cite>
    </p>
    <p>
      作者:<cite>{{ videoInfo.name }}</cite>
    </p>
    <p>言語:{{ language }}</p>
    <p>投稿日:{{ uploadTime }}</p>
    <p>動画時間:{{ recordingTime }}</p>
    <p>コメント</p>
    <p>{{ videoInfo.comment }}</p>
  </div>
</template>

<style scoped>
#video-info-viewer {
  margin: 1em;
}
#video-info-viewer p {
  text-align: left;
}
</style>
