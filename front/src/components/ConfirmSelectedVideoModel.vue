<template>
  <modal name="confirm-selected-video-modal" @before-open="setVideos">
    <div class="confirm-selected-video-modal">
      <h1>{{ title }}</h1>
      <p>{{ message }}</p>
      <table>
        <tr>
          <th>タイトル</th>
          <th>言語</th>
          <th>投稿時間</th>
          <th>動画時間</th>
        </tr>
        <tr v-for="video in videoList" :key="video.header.videoId">
          <td>
            <cite>{{ video.header.title }}</cite>
          </td>
          <td class="lang-viewer">
            {{ video.header.language.name }}
          </td>
          <td class="upload-time-viewer">
            {{ video | uploadDate }}
          </td>
          <td class="video-time-viewer">
            {{ video | recordingTime }}
          </td>
        </tr>
      </table>
      <div>
        <button @click="submit">{{ OKMessage }}</button>
        <button @click="cancel">いいえ</button>
      </div>
    </div>
  </modal>
</template>

<script lang="ts">
import { Video } from '@/models/Video'
import { formatRecordingTime } from '@/util'
import { format } from 'date-fns'
import { ConfirmSelectedVideoModalProperty } from '@/models/ConfirmSelectedVideoModalProperty'
import Vue from 'vue'

type DataType = {
  videoList: Video[]
  title: string
  message: string
  OKMessage: string
}

interface ModalEventParam {
  params: ConfirmSelectedVideoModalProperty
}

export default Vue.extend({
  name: 'ConfirmSelectedVideoModel',
  data(): DataType {
    return {
      videoList: [],
      title: '',
      message: '',
      OKMessage: '',
    }
  },
  filters: {
    uploadDate: function (video: Video): string {
      return format(video.header.uploadTime, 'yyyy/MM/dd')
    },
    recordingTime: function (video: Video): string {
      return formatRecordingTime(video.header.recordingTime)
    },
  },
  methods: {
    setVideos: function (event: ModalEventParam): void {
      this.videoList = event.params.videos
      this.title = event.params.title
      this.message = event.params.message
      this.OKMessage = event.params.OKMessage
    },
    submit: function (): void {
      this.$emit('submit', this.videoList)
      this.$modal.hide('confirm-selected-video-modal')
    },
    cancel: function (): void {
      this.$modal.hide('confirm-selected-video-modal')
    },
  },
})
</script>

<style scoped>
.confirm-selected-video-modal {
  padding: 5px;
}

.confirm-selected-video-modal table {
  width: calc(100% - 20px);
  margin: 10px;
  border-collapse: collapse;
}
.confirm-selected-video-modal table th {
  text-align: center;
}
.confirm-selected-video-modal table td {
  overflow-wrap: break-word;
}
.confirm-selected-video-modal table :where(td, th) {
  padding: 2px 5px;
  border: solid 1px black;
}

.confirm-selected-video-modal :is(label, button) {
  border: solid 2px #202020;
  background: #ffffff;
  padding: 10px 20px;
  font-size: 1em;
  height: 3em;
  margin: 0px 5px;
}

.confirm-selected-video-modal button:hover {
  background: #cccccc;
}

.confirm-selected-video-modal button:active {
  background: #aaaaaa;
}

.video-time-viewer {
  width: 4em;
  text-align: center;
}
.upload-time-viewer {
  width: 9.2em;
  text-align: center;
}
</style>
