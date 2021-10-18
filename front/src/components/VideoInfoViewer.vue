<template>
  <div id="video-info-viewer">
    <p>
      タイトル:<cite>{{ videoInfo.title }}</cite>
    </p>
    <p>
      <span>作者:</span>
      <span v-if="videoInfo.userId != 'file'" @click.stop>
        <router-link :to="{ name: 'UserPage', params: { id: videoInfo.name } }">
          <cite>{{ videoInfo.name }}</cite>
        </router-link>
      </span>
      <span v-else>
        <cite>{{ videoInfo.name }}</cite>
      </span>
    </p>
    <p>言語:{{ language }}</p>
    <p>投稿日:{{ uploadTime }}</p>
    <p>動画時間:{{ recordingTime }}</p>
    <p>コメント</p>
    <p>{{ videoInfo.comment }}</p>
  </div>
</template>

<script lang="ts">
import { VideoInfo } from '@/models/VideoInfo'
import { format } from 'date-fns'
import { formatRecordingTime } from '@/util'
import Vue, { PropType } from 'vue'

export default Vue.extend({
  name: 'VideoInfoViewer',
  props: {
    videoInfo: {
      type: Object as PropType<VideoInfo>,
      default: (): VideoInfo => ({
        videoId: -1,
        userId: '',
        title: '',
        name: '',
        language: undefined,
        comment: '',
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
      return formatRecordingTime(this.videoInfo.recordingTime)
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
