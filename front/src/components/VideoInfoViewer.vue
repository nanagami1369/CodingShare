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
  </div>
</template>

<script lang="ts">
import { VideoInfo } from '@/models/VideoInfo'
import { format } from 'date-fns'
import Vue, { PropType } from 'vue'

export default Vue.extend({
  name: 'VideoInfoViewer',
  props: {
    videoInfo: {
      type: Object as PropType<VideoInfo>,
      default: (): VideoInfo => ({
        userId: -1,
        title: '',
        name: '',
        language: undefined,
        uploadTime: -1,
        recordingTime: -1,
      }),
      required: false,
    },
  },
  computed: {
    language: function (): string {
      return this.videoInfo.language?.name ?? ''
    },
    uploadTime: function (): string {
      if (this.videoInfo.uploadTime < 0) {
        return ''
      }
      return format(this.videoInfo.uploadTime, 'yyyy/MM/dd hh:mm:ss')
    },
    recordingTime: function (): string {
      if (this.videoInfo.recordingTime < 0) {
        return ''
      }
      return format(this.videoInfo.recordingTime, 'mm分ss秒')
    },
  },
})
</script>
<style scoped>
#video-info-viewer {
  margin: 1em;
}
#video-info-viewer p {
  text-align: left;
}
</style>
