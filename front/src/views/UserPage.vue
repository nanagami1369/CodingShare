<template>
  <div class="user-page">
    <h1 class="user-page-title">{{ userId }}のページ</h1>
    <h2>{{ userId }}の動画一覧</h2>
    <article class="user-video-list">
      <Thumbnail
        v-for="video in videos"
        :key="video.header.videoId"
        :video="video"
      />
    </article>
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
  name: 'UserPage',
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
      return this.$route.params.id
    },
  },
  methods: {
    observerUrlDo: async function (): Promise<void> {
      const id = this.$route.params.id
      try {
        var response = await fetch(`/api/uservideo/${id}`)
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
  },
  watch: {
    async $route(): Promise<void> {
      await this.observerUrlDo()
    },
  },
  async created(): Promise<void> {
    await this.observerUrlDo()
  },
})
</script>

<style scoped>
.user-page {
  overflow-y: auto;
  overflow-x: hidden;
}

.user-page-title {
  margin: 30px;
  text-align: start;
}

.user-video-list {
  display: flex;
  flex-wrap: wrap;
  margin: 10px;
}
</style>
