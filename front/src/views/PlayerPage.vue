<template>
  <div id="player-page">
    <div id="side-panel">
      <h1>Player</h1>
      <input type="file" @change="loadData" value="読み込み" />
      <VideoInfoViewer :videoInfo="player.videoInfo" />
    </div>
    <div id="player-panel">
      <textarea id="editor-aria"></textarea>
      <VideoSliderBar
        :elapsedTime="player.info.elapsedTime"
        :totalTime="player.info.totalTime"
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
import { CodingPlayer } from '@/CodingPlayer'
type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  defualtConfig: EditorConfiguration
  player: CodingPlayer
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
  },
  data(): DataType {
    return {
      defualtConfig: {
        mode: 'javascript',
        lineNumbers: true,
        indentUnit: 4,
        theme: 'monokai',
        showHint: true,
        readOnly: true,
      },
      player: new CodingPlayer(),
    }
  },
  methods: {
    loadData: async function (event: Event): Promise<void> {
      const target = event.target as HTMLInputElement
      const file = target.files?.item(0)
      if (file == undefined || file == null) {
        // ファイルがなければ何もしない
        return
      }
      const videoJson = (await readTextFile(file)) as string
      const video: Video = JSON.parse(videoJson)
      this.player.load(video, this.editor)
      this.player.start(this.editor)
    },
  },
  mounted() {
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
</style>
