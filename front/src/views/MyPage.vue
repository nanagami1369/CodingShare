<template>
  <div class="my-page">
    <div class="my-page-header">
      <h1 class="welcome-title">ようこそ {{ userId }} 様</h1>
      <div class="my-page-control">
        <label>
          <span>jsonファイルからimport</span>
          <input type="file" @change="importFromJson" />
        </label>
      </div>
    </div>
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

function readTextFile(file: File): Promise<string | ArrayBuffer | null> {
  return new Promise((resolve, reject) => {
    var fr = new FileReader()
    fr.onload = () => {
      resolve(fr.result)
    }
    fr.onerror = reject
    fr.readAsText(file)
  })
}

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
    loadMyVideo: async function (): Promise<void> {
      try {
        var response = await fetch(`/api/private/myvideo`)
        if (response.ok) {
          this.videos = (await response.json()) as Video[]
        } else {
          const message =
            'サーバーへのimportに失敗しました\n' +
            `message:${await response.text()}\n` +
            `http status:${response.status} ${response.statusText}`
          alert(message)
        }
      } catch (error: unknown) {
        // 通信エラーの場合はアラートで表示
        alert((error as Error).message)
        return
      }
    },
    importFromJson: async function (event: Event) {
      const target = event.target as HTMLInputElement
      const file = target.files?.item(0)
      if (file == null) {
        // ファイルがなければ何もしない
        return
      }

      try {
        const videoJson = (await readTextFile(file)) as string
        const video: Video = JSON.parse(videoJson)
        const response = await fetch('/api/private/savevideo', {
          method: 'POST',
          mode: 'cors',
          credentials: 'include',
          body: JSON.stringify(video),
        })
        if (response.ok) {
          alert('サーバーにimportしました')
          await this.loadMyVideo()
        } else {
          const message =
            'サーバーへのimportに失敗しました\n' +
            `message:${await response.text()}\n` +
            `http status:${response.status} ${response.statusText}`
          alert(message)
        }
      } catch (error: unknown) {
        alert('サーバーへのimportに失敗しました\n' + (error as Error).message)
        return
      }
      return
    },
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
    await this.loadMyVideo()
  },
})
</script>

<style scoped>
.welcome-title {
  text-align: start;
}

.my-page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 30px;
}

.my-page-control {
  flex: 1 1 auto;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.my-page-control input[type='file'] {
  display: none;
}

.my-page-control label {
  border: solid 2px #202020;
  background: #ffffff;
  padding: 10px 20px;
}

.my-page-control label:hover {
  background: #cccccc;
}

.my-page-control label:active {
  background: #aaaaaa;
}

.my-video-list {
  display: flex;
  flex-wrap: wrap;
  margin: 10px;
}
</style>
