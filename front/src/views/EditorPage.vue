<template>
  <div id="editor-page">
    <div id="sidePanel">
    <h1>Editor</h1>
      <h2 v-if="isRecoding"><span style="color: red">●</span>録画中</h2>
      <h2 v-else>■停止中</h2>
    <select v-model="selectedLanguage">
      <option v-for="lang in languages" :key="lang.tag" :value="lang">
        {{ lang.name }}
      </option>
    </select>
      <button @click="onSaveEditorImage">保存</button>
    </div>
    <div id="editorPanel">
      <textarea id="editorAria">#TEST</textarea>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import html2canvas from 'html2canvas'
import '../codemirror-global.js'
import CodeMirror, { EditorConfiguration } from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/monokai.css'
import 'codemirror/mode/clike/clike.js'
import 'codemirror/addon/hint/show-hint.js'
import 'codemirror/addon/hint/show-hint.css'
import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/addon/hint/javascript-hint.js'
import { Language } from '@/models/language'
type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  defualtConfig: EditorConfiguration
  languages: Language[]
  selectedLanguage: Language
  editorRecodeOptions: editorRecodeOptions
  isRecoding: boolean
}

type editorRecodeOptions = {
  imageConnt: number
  interval: number
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
        showHint: true,
        extraKeys: { 'Ctrl-Space': 'autocomplete' },
      },
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
      editorRecodeOptions: {
        imageConnt: 20,
        interval: 500,
      },
      isRecoding: false,
    }
  },
  watch: {
    selectedLanguage: function (newLang: Language, _: Language) {
      this.editor?.setOption('mode', newLang.tag)
    },
  },
  methods: {
    onSaveEditorImage: async function (): Promise<void> {
      const editorHistoryContainer: HTMLElement[] = []
      let counter = 0
      let timer = -1
      this.isRecoding = true
      timer = setInterval(async () => {
        if (counter >= this.editorRecodeOptions.imageConnt) {
          clearInterval(timer)
          this.isRecoding = false
          alert('終了')
          const body: HTMLElement | null = document.querySelector('body')
          if (body == null) {
            throw new Error('can not find body tag')
          }
          editorHistoryContainer
            .map((html) => {
              // HTMLElementをキャンバスへ変換
              html.id = 'data-for-canvas-nokohunjyata'
              body.appendChild(html)
              const dataForCanvas: HTMLElement | null = document.querySelector(
                '#data-for-canvas-nokohunjyata'
              )
              if (dataForCanvas == null) {
                throw new Error('can not find data for canvas')
              }
              try {
                const canvas = html2canvas(dataForCanvas)
                return canvas
              } finally {
                body.removeChild(dataForCanvas)
              }
            })
            .forEach(async (canvas, index) => {
              // キャンバスのデータを開く
              window.open((await canvas).toDataURL())
            })
        }
      const editor: HTMLElement | null = document.querySelector('.CodeMirror')
      if (editor == null) {
        throw new Error('can not find CodeMirror dom')
      }
        editorHistoryContainer.push(editor.cloneNode(true) as HTMLElement)
        console.log(counter)
        counter++
      }, this.editorRecodeOptions.interval)
    },
  },
  mounted() {
    const editorAria: HTMLTextAreaElement | null = document.querySelector(
      '#editorAria'
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

#editor-page {
  grid-row: 2;
  display: grid;
  grid-template-columns: 1fr 1fr;
}

#sidePanel {
  grid-column: 1;
}
#editorPanel {
  grid-column: 2;
}
</style>
