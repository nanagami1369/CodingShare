<template>
  <div class="my-page">
    <h1 class="welcome-title">ようこそ {{ userId }} 様</h1>
    <h2>あなたの動画一覧</h2>
    <article class="my-video-list">
      <Thumbnail
        v-for="video in videos"
        :key="video.header.videoId"
        :video="video"
      />
    </article>
    <button @click="logout">ログアウト</button>
  </div>
</template>

<script lang="ts">
import Thumbnail from '@/components/Thumbnail.vue'
import { Video } from '@/models/Video'
import Vue from 'vue'

type DataType = {
  videos: Video[]
}

export default Vue.extend({
  name: 'MyPage',
  components: {
    Thumbnail,
  },
  data(): DataType {
    return {
      videos: [],
    }
  },
  computed: {
    userId() {
      // return 'tanaka'
      return this.$store.getters.userId
    },
  },
  methods: {
    logout: async function () {
      try {
        const response = await fetch('/api/private/logout', {
          method: 'POST',
          mode: 'cors',
          credentials: 'include',
        })
        if (response.status != 200) {
          // 失敗時はエラーを表示
          let message =
            `message:${await response.text()}\n` +
            `http status:${response.status} ${response.statusText}`
          alert(message)
          return
        }
        this.$router.push('/login')
      } catch (error: unknown) {
        // 通信エラーの場合はアラートで表示
        alert((error as Error).message)
      }
    },
  },
  async created(): Promise<void> {
    try {
      var response = await fetch(`/api/private/myvideo`)
      if (response.ok) {
        this.videos = (await response.json()) as Video[]
      } else {
        alert(`通信エラー\n${await response.text()}`)
      }
    } catch (error: unknown) {
      // 通信エラーの場合はアラートで表示
      alert((error as Error).message)
      return
    }
  },
})
</script>

<style scoped>
.welcome-title {
  margin: 30px;
  text-align: start;
}

.my-video-list {
  display: flex;
  flex-wrap: wrap;
  margin: 10px;
}
</style>
