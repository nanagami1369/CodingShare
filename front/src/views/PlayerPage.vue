<template>
  <div id="player-page" :class="{ 'is-play': player.isPlay }">
    <div id="side-panel">
      <h1>Player</h1>
      <input
        type="file"
        @change="loadData"
        value="読み込み"
        :disabled="player.isLoading || player.isPlay"
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
        <p>速度</p>
        <VueSlider
          :style="{ padding: '1em 2em' }"
          :process-style="{ backgroundColor: '#28e270' }"
          :tooltip-style="{
            backgroundColor: '#116230',
            borderColor: '#116230',
          }"
          :data="speedSliderIndex"
          v-model="speed"
          :marks="true"
          :adsorb="true"
          :lazy="true"
          :contained="true"
        />
      </div>
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
import VideoInfoViewer from '@/components/VideoInfoViewer.vue'
import VideoSliderBar from '@/components/VideoSliderBar.vue'
import LoadingViwer from '@/components/LoadingViewer.vue'
import VueSlider from 'vue-slider-component'
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

library.add(faPlay, faPause, faUndo, faStepForward, faFastForward)

type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  backgroundEditor?: CodeMirror.EditorFromTextArea
  defualtConfig: EditorConfiguration
  player: CodingPlayer
  speedSliderIndex: string[]
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
    VueSlider,
    LoadingViwer,
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
      player: new CodingPlayer(snapShotTimeSpan),
      speedSliderIndex: ['50%', '100%', '200%'],
    }
  },
  computed: {
    speed: {
      get: function (): string {
        return `${this.player.info.speed}%`
      },
      set: function (value: string) {
        // valueから「％」を取る
        const stringNumber = value.slice(0, -1)
        this.player.info.speed = parseInt(stringNumber, 10)
      },
    },
  },
  methods: {
    loadData: async function (event: Event): Promise<void> {
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
  mounted() {
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
    this.editor?.setSize('100%', '70vh')
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

.player-control-button {
  font-size: 1.2em;
  padding: 5px 20px;
  border-radius: 25px;
  border: solid 2px #9f9f9f;
}

.player-control-button:disabled {
  border: solid 2px #dddddd;
}

.player-control-button:active {
  border: solid 2px #222222;
}

input[type='file'] {
  padding: 20px 0px;
}
</style>
