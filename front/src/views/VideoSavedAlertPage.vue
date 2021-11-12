<template>
  <div>
    <NotFoundPage v-if="isNotFound" :style="{ height: '100%' }" />
    <div v-else class="video-saved-alert-page">
      <h2>サーバーへの保存が完了しました</h2>
      <h3>保存された動画</h3>
      <div class="saved-video">
        <Thumbnail v-if="savedVideo.header.videoId != -1" :video="savedVideo" />
      </div>
      <router-link :style="{ 'text-decoration': 'none' }" to="/editor">
        <div class="change-page-button">
          <FontAwesomeIcon icon="file-video" />
          録画ページへ戻る
        </div>
      </router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Video } from '@/models/Video'
import Thumbnail from '@/components/Thumbnail.vue'
import Vue from 'vue'
import NotFoundPage from '@/views/NotFoundPage.vue'
import { faVideo } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'

library.add(faVideo)

type DataType = {
  isNotFound: boolean
  savedVideo: Video
}

export default Vue.extend({
  name: 'VideoSavedAlertPage',
  components: {
    Thumbnail,
    NotFoundPage,
    FontAwesomeIcon,
  },
  data(): DataType {
    return {
      isNotFound: false,
      savedVideo: {
        header: {
          videoId: -1,
          userId: 'null',
          title: 'null',
          name: 'null',
          uploadTime: 0,
          recordingTime: 0,
          comment: '',
          language: {
            tag: 'text/x-csrc',
            name: 'C言語',
          },
        },
        value: [],
      },
    }
  },
  methods: {
    observerUrlDo: async function (): Promise<void> {
      const videoId = this.$route.params.id
      try {
        const response = await fetch('/api/loadvideo/' + videoId, {
          method: 'GET',
          mode: 'cors',
          credentials: 'include',
        })
        if (response.ok) {
          this.isNotFound = false
          this.savedVideo = (await response.json()) as Video
          return
        } else if (response.status == 404) {
          this.isNotFound = true
        } else {
          // それ以外の場合はエラーを表示
          const message =
            `録画データの取得に失敗しました\n` +
            `message:${await response.text()}\n` +
            `http status:${response.status} ${response.statusText}`
          alert(message)
          return
        }
      } catch (error: unknown) {
        // 通信エラーの場合はアラートで表示
        alert((error as Error).message)
        return
      }
    },
  },
  watch: {
    async $route(): Promise<void> {
      // 2回目呼び出し
      await this.observerUrlDo()
    },
  },
  async created(): Promise<void> {
    this.observerUrlDo()
  },
})
</script>

<style scoped>
.video-saved-alert-page h2 {
  margin-top: 40px;
  margin-bottom: 20px;
}
.saved-video {
  display: flex;
  justify-content: center;
  margin-bottom: 30px;
}

.change-page-button {
  border: solid 2px #202020;
  background: #ffffff;
  padding: 10px 20px;
  font-size: 1em;
  margin: 0px 5px;
  text-decoration: none;
  color: black;
}

.change-page-button:hover {
  color: #28e270;
}

.change-page-button:active {
  background: #cccccc;
}
</style>
