<template>
  <div>
    <NotFoundPage v-show="isNotFound" :style="{ height: '100%' }" />
    <div
      v-show="!isNotFound"
      id="player-page"
      :class="{ 'is-play': player.isPlay }"
    >
      <div id="side-panel">
        <h1>Player</h1>
        <input
          v-show="isFileMode"
          class="video_label"
          type="file"
          @change="loadData"
          value="読み込み"
          :disabled="player.isLoading || player.isPlay"
        />
        <div v-show="!isFileMode" class="video_label"></div>
        <VideoInfoViewer :videoInfo="player.videoInfo" />
        <textarea id="background-editor-aria"></textarea>
        <LoadingViwer v-show="player.isLoading" />
      </div>
      <div id="player-panel">
        <textarea id="editor-aria"></textarea>
        <VideoSliderBar
          :elapsedTime="player.info.elapsedTime"
          :totalTime="player.info.totalTime"
          :disabled="!player.isLoaded || player.isPlay"
          @change="move"
        />
        <div class="player-control">
          <button
            @click="backToTheBeginning"
            class="player-control-button"
            :disabled="!player.isLoaded"
          >
            <FontAwesomeIcon icon="undo" />
          </button>
          <button
            v-if="player.isPlay"
            @click="pouse"
            class="player-control-button"
            :disabled="!player.isLoaded"
          >
            <FontAwesomeIcon icon="pause" />
          </button>
          <button
            v-else
            @click="start"
            class="player-control-button"
            :disabled="!player.isLoaded"
          >
            <FontAwesomeIcon icon="play" />
          </button>
          <button
            class="player-control-button"
            @click="stepForward"
            :disabled="!player.isLoaded"
          >
            <FontAwesomeIcon icon="step-forward" />
          </button>
          <button
            class="player-control-button"
            @click="fastForward"
            :disabled="!player.isLoaded"
          >
            <FontAwesomeIcon icon="fast-forward" />
          </button>
          <span class="elapsed-time">{{ playbackPosition }}</span>
          <div class="speed-control">
            <button
              class="player-control-button speed-control-button"
              @click="toggleSpeedMenu"
            >
              速度 {{ speed }}%
            </button>
            <div
              v-show="isSpeedMenuOpen"
              class="spped-control-wrapper"
              @click="toggleSpeedMenu"
            ></div>
            <div class="speed-context-menu">
              <div
                v-show="isSpeedMenuOpen"
                v-for="speedIndex in speedSliderIndex"
                :key="speedIndex"
                class="speed-context-menu-item"
                @click="
                  speed = speedIndex
                  toggleSpeedMenu()
                "
              >
                <span class="speed-context-menu-check-space">{{
                  speed == speedIndex ? '✓' : ''
                }}</span>
                <span>{{ speedIndex }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import '../codemirror-global.js'
import CodeMirror, { EditorConfiguration } from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/monokai.css'
import 'codemirror/mode/clike/clike.js'
import 'codemirror/addon/hint/show-hint.js'
import 'codemirror/addon/hint/show-hint.css'
import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/addon/hint/javascript-hint.js'
import 'codemirror/addon/edit/closebrackets.js'
import { Video } from '@/models/Video'
import NotFoundPage from '@/views/NotFoundPage.vue'
import VideoInfoViewer from '@/components/VideoInfoViewer.vue'
import VideoSliderBar from '@/components/VideoSliderBar.vue'
import LoadingViwer from '@/components/LoadingViewer.vue'
import 'vue-slider-component/theme/default.css'
import { CodingPlayer } from '@/CodingPlayer'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faPlay,
  faPause,
  faUndo,
  faStepForward,
  faFastForward,
} from '@fortawesome/free-solid-svg-icons'
import { formatRecordingTime } from '@/util'

library.add(faPlay, faPause, faUndo, faStepForward, faFastForward)

type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  backgroundEditor?: CodeMirror.EditorFromTextArea
  defualtConfig: EditorConfiguration
  player: CodingPlayer
  isNotFound: boolean
  speedSliderIndex: number[]
  isSpeedMenuOpen: boolean
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

export default Vue.extend({
  name: 'PlayerPage',
  components: {
    VideoInfoViewer,
    VideoSliderBar,
    FontAwesomeIcon,
    LoadingViwer,
    NotFoundPage,
  },
  data(): DataType {
    const snapShotTimeSpan = 30000
    return {
      defualtConfig: {
        mode: 'javascript',
        lineNumbers: true,
        indentUnit: 4,
        theme: 'monokai',
        showHint: true,
        readOnly: true,
      },
      player: new CodingPlayer(snapShotTimeSpan, 100),
      isNotFound: false,
      speedSliderIndex: [50, 100, 200],
      isSpeedMenuOpen: false,
    }
  },
  computed: {
    speed: {
      get: function (): number {
        return this.player.info.speed
      },
      set: function (value: number) {
        this.player.setSpeed(value)
      },
    },
    playbackPosition: function (): string {
      // prettier-ignore
      return `${formatRecordingTime(this.player.info.elapsedTime)} / ${formatRecordingTime(this.player.info.totalTime)}`
    },
    isFileMode: function (): boolean {
      return this.$route.params.id == 'file'
    },
  },
  methods: {
    setEditorSize: function () {
      const sidePanelWidth = 300
      const bottomPanelHeight = 200
      let width = document.body.clientWidth - sidePanelWidth
      let height = document.body.clientHeight - bottomPanelHeight
      width = width >= 0 ? width : 0
      height = height >= 0 ? height : 0
      console.log('resize w :' + width + 'h:' + height)
      this.editor?.setSize(width, height)
    },
    toggleSpeedMenu: function (): void {
      this.isSpeedMenuOpen = !this.isSpeedMenuOpen
    },
    loadData: async function (event: Event): Promise<void> {
      this.player.pause()
      this.player.clear(this.editor)
      const target = event.target as HTMLInputElement
      const file = target.files?.item(0)
      if (file == null) {
        // ファイルがなければ何もしない
        return
      }
      const videoJson = (await readTextFile(file)) as string
      const video: Video = JSON.parse(videoJson)
      this.player.load(video, this.editor, this.backgroundEditor)
    },
    observerUrlDo: async function (): Promise<void> {
      this.player.pause()
      this.player.clear(this.editor)
      if (this.isFileMode) {
        this.isNotFound = false
        return
      }
      // サーバーからデータを取得
      const id = this.$route.params.id
      try {
        const response = await fetch('/api/loadvideo/' + id, {
          method: 'GET',
          mode: 'cors',
          credentials: 'include',
        })
        if (response.ok) {
          this.isNotFound = false
          const video = (await response.json()) as Video
          this.player.load(video, this.editor, this.backgroundEditor)
          return
        } else if (response.status == 404) {
          this.isNotFound = true
          return
        } else {
          // それ以外の場合はエラーを表示
          const message =
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
    start: function (): void {
      this.player.start(this.editor)
    },
    pouse: function (): void {
      this.player.pause()
    },
    backToTheBeginning: function (): void {
      this.player.backToTheBeginning(this.editor)
    },
    stepForward: function (): void {
      this.player.stepForward(this.editor)
    },
    fastForward: function (): void {
      this.player.fastForward(this.editor)
    },
    move: function (time: number): void {
      this.player.move(time, this.editor)
    },
  },
  watch: {
    async $route(): Promise<void> {
      // 2回目呼び出し
      await this.observerUrlDo()
    },
  },
  async mounted(): Promise<void> {
    // 裏で動かす用のエディタ
    const backgroundEditorArea: HTMLTextAreaElement | null =
      document.querySelector('#background-editor-aria')
    if (backgroundEditorArea == null) {
      throw new Error('textarea not found for Background CodeMirror')
    }
    this.backgroundEditor = CodeMirror.fromTextArea(backgroundEditorArea)
    this.backgroundEditor.setSize('0px', '0px')
    // 画面を再生する用のエディタ
    const editorAria: HTMLTextAreaElement | null =
      document.querySelector('#editor-aria')
    if (editorAria == null) {
      throw new Error('textarea not found for CodeMirror')
    }
    const config = this.defualtConfig
    this.editor = CodeMirror.fromTextArea(editorAria, config)
    this.setEditorSize()
    window.onresize = this.setEditorSize
    // 1回目呼び出し
    await this.observerUrlDo()
  },
  beforeDestroy(): void {
    window.onresize = null
  },
})
</script>

<style scoped>
h1 {
  color: #28e270;
}

#player-page {
  grid-row: 2;
  display: grid;
  grid-template-columns: 300px 3fr;
}

#side-panel {
  grid-column: 1;
}

#player-panel {
  grid-column: 2;
}

.player-control {
  height: 35px;
  line-height: 35px;
  display: flex;
  background-color: #222222;
}

.player-control-button {
  font-size: 1.2em;
  padding: 0px 10px;
  border: none;
  background-color: transparent;
  color: #eeeeee;
  cursor: pointer;
}

.player-control-button:disabled {
  color: #555;
  cursor: no-drop;
}

.player-control-button:hover {
  background-color: #777777;
}

.elapsed-time {
  padding: 0px 15px;
  vertical-align: middle;
  display: inline-block;
  color: white;
}

.spped-control-wrapper {
  position: absolute;
  height: 100%;
  width: 100%;
  top: 0px;
  right: 0px;
  z-index: 10;
}

.speed-control-button {
  width: 6em;
  background-color: #333333;
}

.speed-context-menu {
  background-color: #111111;
  position: relative;
  color: #eeeeee;
  top: -160px;
  z-index: 11;
}

.speed-context-menu-item {
  padding: 0px 10px;
  display: flex;
  justify-content: space-between;
}

.speed-context-menu-item:hover {
  background-color: #777777;
}

.speed-context-menu-check-space {
  width: 0.8em;
  display: inline-block;
}
.video_label {
  height: 70px;
  line-height: 70px;
}
</style>
