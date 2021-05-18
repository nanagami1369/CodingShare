<template>
  <div id="player-page">
    <div id="side-panel">
      <h1>Player</h1>
      <input type="file" @change="loadData" value="読み込み" />
    </div>
    <div id="player-panel">
      <textarea id="editor-aria">#TEST</textarea>
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
import { Language } from '@/models/language'
import { CodingRecorder } from '@/CodingRecorder'
import { Video } from '@/models/Video.js'
type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  recorder: CodingRecorder
  defualtConfig: EditorConfiguration
  selectedLanguage?: Language
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

function doSomethingLoop(
  maxCount: number,
  count: number,
  doSomething: (count: number) => number
) {
  if (count <= maxCount) {
    const nextSpan = doSomething(count)
    setTimeout(function () {
      doSomethingLoop(maxCount, ++count, doSomething)
    }, nextSpan)
  }
}

export default Vue.extend({
  name: 'PlayerPage',
  data(): DataType {
    return {
      defualtConfig: {
        mode: 'javascript',
        lineNumbers: true,
        indentUnit: 4,
        theme: 'monokai',
        showHint: true,
        extraKeys: { 'Ctrl-Space': 'autocomplete' },
      },
      recorder: new CodingRecorder(),
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
      console.log(video)
      this.editor?.setValue('')
      doSomethingLoop(video.value.length - 1, 0, (index: number) => {
        const text = video.value[index].changeData.text
        const from = video.value[index].changeData.from
        const to = video.value[index].changeData.to
        const origin = video.value[index].changeData.origin
        this.editor?.replaceRange(text, from, to, origin)
        return 250
      })
    },
  },
  mounted() {
    const editorAria: HTMLTextAreaElement | null = document.querySelector(
      '#editor-aria'
    )
    if (editorAria == null) {
      throw new Error('textarea not found for CodeMirror')
    }
    const config = this.defualtConfig
    this.editor = CodeMirror.fromTextArea(editorAria, config)
    this.editor?.setSize(1280, 720)
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
  grid-template-columns: 1fr 1fr;
}

#side-panel {
  grid-column: 1;
}
#player-panel {
  grid-column: 2;
}
</style>
