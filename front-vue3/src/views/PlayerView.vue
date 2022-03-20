<script setup lang="ts">
import { computed, onMounted, onUnmounted, reactive, ref } from 'vue'
import CodeMirror, { type EditorConfiguration } from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/monokai.css'
import 'codemirror/mode/clike/clike.js'
import 'codemirror/addon/hint/show-hint.js'
import 'codemirror/addon/hint/show-hint.css'
import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/addon/hint/javascript-hint.js'
import 'codemirror/addon/edit/closebrackets.js'
import type { Video } from '@/models/Video'
import VideoInfoViewer from '@/components/VideoInfoViewer.vue'
import VideoSliderBar from '@/components/VideoSliderBar.vue'
import LoadingViewer from '@/components/LoadingViewer.vue'
// import 'vue-slider-component/theme/default.css'
import { CodingPlayer } from '@/CodingPlayer'
import { useSpeedStore } from '@/stores/speed'
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
import { useRoute } from 'vue-router'

library.add(faPlay, faPause, faUndo, faStepForward, faFastForward)

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

const snapShotTimeSpan = 30000

const defualtConfig = reactive<EditorConfiguration>({
  mode: 'javascript',
  lineNumbers: true,
  indentUnit: 4,
  theme: 'monokai',
  showHint: true,
  readOnly: true,
})

const speedStore = useSpeedStore()
const player = reactive(new CodingPlayer(snapShotTimeSpan, speedStore.speed))
const speedSliderIndex = [50, 100, 200]
const isSpeedMenuOpen = ref(false)
const editorAria = ref<HTMLTextAreaElement>()
const backgroundEditorArea = ref<HTMLTextAreaElement>()

let editor: CodeMirror.Editor | undefined
let backgroundEditor: CodeMirror.Editor | undefined

const speed = computed<number>({
  get: (): number => player.info.speed,
  set: (value: number): void => {
    speedStore.$patch({ speed: value })
    player.setSpeed(value)
  },
})
const playbackPosition = computed(
  (): string =>
    `${formatRecordingTime(player.info.elapsedTime)} / ${formatRecordingTime(
      player.info.totalTime
    )}`
)
const setEditorSize = (): void => {
  const sidePanelWidth = 300
  const bottomPanelHeight = 200
  let width = document.body.clientWidth - sidePanelWidth
  let height = document.body.clientHeight - bottomPanelHeight
  width = width >= 0 ? width : 0
  height = height >= 0 ? height : 0
  console.log('resize w :' + width + 'h:' + height)
  editor?.setSize(width, height)
}
const toggleSpeedMenu = (): void => {
  isSpeedMenuOpen.value = !isSpeedMenuOpen.value
}
const changeSpeed = (newSpeed: number): void => {
  speed.value = newSpeed
  toggleSpeedMenu()
}
const loadData = async (event: Event): Promise<void> => {
  player.pause()
  player.clear(editor)
  const target = event.target as HTMLInputElement
  const file = target.files?.item(0)
  if (file == null) {
    // ファイルがなければ何もしない
    return
  }
  const videoJson = (await readTextFile(file)) as string
  const video: Video = JSON.parse(videoJson)
  player.load(video, editor, backgroundEditor)
}
const start = (): void => {
  player.start(editor)
}
const pouse = (): void => {
  player.pause()
}
const backToTheBeginning = (): void => {
  player.backToTheBeginning(editor)
}
const stepForward = (): void => {
  player.stepForward(editor)
}
const fastForward = (): void => {
  player.fastForward(editor)
}
const move = (time: number): void => {
  player.move(time, editor)
}
const observerUrlDo = async (): Promise<void> => {
  speed.value = speedStore.speed
  player.pause()
  player.clear(editor)
  const route = useRoute()
  if (route.query.src == null) {
    return
  }
  try {
    const response = await fetch('/video/' + route.query.src.toString())
    if (response.ok) {
      const video = (await response.json()) as Video
      player.load(video, editor, backgroundEditor)
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
}
// watch(selectedLanguage, (newLang: Language) => {
//   editor?.setOption('mode', newLang.tag)
// })

// watch: {
//     async $route(): Promise<void> {
//       // 2回目呼び出し
//       await this.observerUrlDo()
//     },
//   },
onMounted(async (): Promise<void> => {
  // 裏で動かす用のエディタ
  if (backgroundEditorArea.value == null) {
    throw new Error('textarea not found for Background CodeMirror')
  }
  backgroundEditor = CodeMirror.fromTextArea(backgroundEditorArea.value)
  backgroundEditor.setSize('0px', '0px')
  // 画面を再生する用のエディタ
  if (editorAria.value == null) {
    throw new Error('textarea not found for CodeMirror')
  }
  const config = defualtConfig
  editor = CodeMirror.fromTextArea(editorAria.value, config)
  setEditorSize()
  window.onresize = setEditorSize
  await observerUrlDo()
})
onUnmounted((): void => {
  window.onresize = null
})
</script>
<template>
  <div id="player-page" :class="{ 'is-play': player.isPlay }">
    <div id="side-panel">
      <h1>Player</h1>
      <input
        type="file"
        @change="loadData"
        :disabled="player.isLoading || player.isPlay"
      />
      <VideoInfoViewer :videoInfo="player.videoInfo" />
      <textarea ref="backgroundEditorArea"></textarea>
      <LoadingViewer v-show="player.isLoading" />
    </div>

    <div id="player-panel">
      <textarea ref="editorAria"></textarea>

      <VideoSliderBar
        :elapsedTime="player.info.elapsedTime"
        :totalTime="player.info.totalTime"
        :disabled="!player.isLoaded || player.isPlay"
        @seek="move"
      />
      <!-- <VideoSliderBar
        :elapsedTime="player.info.elapsedTime"
        :totalTime="player.info.totalTime"
        :disabled="!player.isLoaded || player.isPlay"
        @change="move"
      /> -->
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
            class="speed-control-wrapper"
            @click="toggleSpeedMenu"
          ></div>
          <div class="speed-context-menu">
            <div
              v-show="isSpeedMenuOpen"
              v-for="speedIndex in speedSliderIndex"
              :key="speedIndex"
              class="speed-context-menu-item"
              @click="changeSpeed(speedIndex)"
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
</template>

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
  width: 1.8em;
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

.speed-control-wrapper {
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

input[type='file'] {
  padding: 20px 0px;
}
</style>
