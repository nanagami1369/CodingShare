<template>
  <div id="editor-page">
    <div id="side-panel">
      <h1>Editor</h1>
      <select v-model="selectedLanguage">
        <option v-for="lang in languages" :key="lang.tag" :value="lang">
          {{ lang.name }}
        </option>
      </select>
      <button @click="recordStart">開始</button>
      <button @click="recordStop">停止</button>
    </div>
    <div id="editor-panel">
      <textarea id="editor-aria"></textarea>
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
type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  recorder: CodingRecorder
  defualtConfig: EditorConfig
  languages: Language[]
  selectedLanguage: Language
}
interface EditorConfig extends EditorConfiguration {
  autoCloseBrackets: boolean
}
export default Vue.extend({
  name: 'EditorPage',
  data(): DataType {
    return {
      defualtConfig: {
        mode: 'javascript',
        lineNumbers: true,
        indentUnit: 4,
        theme: 'monokai',
        autoCloseBrackets: true,
        showHint: true,
        extraKeys: { 'Ctrl-Space': 'autocomplete' },
      },
      recorder: new CodingRecorder(),
      languages: [
        {
          tag: 'javascript',
          name: 'JavaScript',
        },
        {
          tag: 'text/x-csrc',
          name: 'C言語',
        },
        {
          tag: 'text/x-c++src',
          name: 'C++',
        },
        {
          tag: 'text/x-java',
          name: 'Java',
        },
      ],
      selectedLanguage: {
        tag: 'javascript',
        name: 'JavaScript',
      },
    }
  },
  methods: {
    recordStart: function () {
      this.recorder.start()
    },
    recordStop: function () {
      this.recorder.stop()
      const video = this.recorder.outputVideo(
        -1,
        'test',
        'test',
        this.selectedLanguage
      )
      var url = (window.URL || window.webkitURL).createObjectURL(
        new Blob([JSON.stringify(video)], { type: 'application/json' })
      )
      const a = document.createElement('a')
      a.href = url
      a.download = 'video.json'
      a.click()
    },
  },
  watch: {
    selectedLanguage: function (newLang: Language) {
      this.editor?.setOption('mode', newLang.tag)
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
    this.recorder.register(this.editor)
  },
  beforeDestroy() {
    this.recorder.unregister(this.editor)
  },
})
</script>

<style scoped>
h1 {
  color: #28e270;
}

#editor-page {
  grid-row: 2;
  display: grid;
  grid-template-columns: 1fr 1fr;
}

#side-panel {
  grid-column: 1;
}
#editor-panel {
  grid-column: 2;
}
</style>
