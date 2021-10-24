<template>
  <div class="my-page">
    <ConfirmSelectedVideoModel @submit="removeVideo" />
    <div class="my-page-header">
      <h1 class="welcome-title">ようこそ {{ userId }} 様</h1>
      <div class="my-page-control">
        <button @click="confirmremoveVideo">
          <FontAwesomeIcon icon="eraser" />
          <span> 選択した録画データを削除する</span>
        </button>
        <button @click="resetSelecter">選択をリセットする</button>
        <label>
          <span>jsonファイルからimport</span>
          <input type="file" @change="importFromJson" />
        </label>
      </div>
    </div>
    <h2>あなたの動画一覧</h2>
    <article class="my-video-list">
      <div
        v-for="video in videos"
        :key="video.header.videoId"
        class="video-container"
      >
        <div class="video-controller">
          <CheckBox
            :id="'my_video_' + video.header.videoId"
            :value="video.header.videoId"
            :name="selectCheckboxName"
            :style="{
              position: 'relative',
              left: '284px',
              top: '31px',
            }"
            :resetFlag="resetSelecterFlag"
          />
        </div>
        <Thumbnail :video="video" />
      </div>
    </article>
    <button @click="logout">ログアウト</button>
  </div>
</template>

<script lang="ts">
import Thumbnail from '@/components/Thumbnail.vue'
import { Video } from '@/models/Video'
import CheckBox from '@/components/CheckBox.vue'
import ConfirmSelectedVideoModel from '@/components/ConfirmSelectedVideoModel.vue'

import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faEraser } from '@fortawesome/free-solid-svg-icons'
import { ConfirmSelectedVideoModalProperty } from '@/models/ConfirmSelectedVideoModalProperty'

library.add(faEraser)

const removeOneVideo = async (video: Video): Promise<Video> => {
  const response = await fetch(
    `/api/private/removevideo/${video.header.videoId}`,
    {
      method: 'DELETE',
      mode: 'cors',
      credentials: 'include',
    }
  )
  if (response.ok) {
    return video
  } else {
    const message =
      `id:${video.header.videoId}, title:${video.header.title}の削除に失敗しました\n` +
      `message:${await response.text()}\n` +
      `http status:${response.status} ${response.statusText}`
    throw new Error(message)
  }
}

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
  selectCheckboxName: string
  resetSelecterFlag: boolean
}

export default Vue.extend({
  name: 'MyPage',
  components: {
    Thumbnail,
    CheckBox,
    FontAwesomeIcon,
    ConfirmSelectedVideoModel,
  },
  data(): DataType {
    return {
      videos: [],
      selectCheckboxName: 'my-page-video-selecter',
      resetSelecterFlag: false,
    }
  },
  computed: {
    userId() {
      // return 'tanaka'
      return this.$store.getters.userId
    },
  },
  methods: {
    resetSelecter: function (): void {
      this.resetSelecterFlag = !this.resetSelecterFlag
    },
    loadMyVideo: async function (): Promise<void> {
      this.resetSelecter()
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
    getSelectedIdsFromCheckBox: function (): string[] {
      return [
        ...document.querySelectorAll(
          `input[name="${this.selectCheckboxName}"]`
        ),
      ]
        .map((c) => c as HTMLInputElement)
        .filter((c) => c.checked)
        .map((c) => c.value)
    },
    confirmremoveVideo: function (): void {
      const selectVideoIds = this.getSelectedIdsFromCheckBox()
      if (selectVideoIds.length == 0) {
        alert('削除するデータが選択されていません')
        return
      }
      console.log(selectVideoIds)
      const selectedVideos = selectVideoIds
        .map((id) => parseInt(id, 10))
        .map((id) => this.videos.find((v) => v.header.videoId == id))
        // filter処理で消してるのでundefindは絶対に入ってこない
        .filter((v) => v != null) as Video[]

      const eventProperty: ConfirmSelectedVideoModalProperty = {
        videos: selectedVideos,
        title: '録画データを削除する',
        message: '選択された録画データを削除しますか？',
        OKMessage: '削除する',
      }
      this.$modal.show('confirm-selected-video-modal', eventProperty)
    },
    removeVideo: async function (videoList: Video[]): Promise<void> {
      const resurt = await Promise.allSettled(
        videoList.map(async (v) => await removeOneVideo(v))
      )
      const succsessVideos: Video[] = []
      const regectMessages: string[] = []
      resurt.forEach((resurt) => {
        if (resurt.status == 'rejected') {
          regectMessages.push(resurt.reason)
        } else {
          succsessVideos.push(resurt.value)
        }
      })
      let message = ''
      if (regectMessages.length == 0) {
        message += '## すべての録画データの削除に成功しました ##\n'
        message += '成功したデータ\n'
        succsessVideos.forEach(
          (v) =>
            (message += `id:${v.header.videoId}, title:${v.header.title}\n`)
        )
      } else {
        message += '## 一部の録画データの削除に失敗しました ##\n'
        regectMessages.forEach((m) => (message += `${m}\n`))
        message += '## 成功したデータ ##\n'
        succsessVideos.forEach(
          (v) =>
            (message += `id:${v.header.videoId}, title:${v.header.title}\n`)
        )
      }
      alert(message)
      this.loadMyVideo()
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
  flex: 1 0 auto;
}

.my-page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 30px;
  flex-wrap: wrap;
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

.my-page-control :is(label, button) {
  border: solid 2px #202020;
  background: #ffffff;
  padding: 10px 20px;
  font-size: 1em;
  margin: 0px 5px;
}

.my-page-control :is(label, button):hover {
  background: #cccccc;
}

.my-page-control :is(label, button):active {
  background: #aaaaaa;
}

.my-video-list {
  display: flex;
  flex-wrap: wrap;
  margin: 10px;
}
</style>
