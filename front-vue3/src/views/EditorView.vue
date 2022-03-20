<script setup lang="ts">
import CodeMirror from 'codemirror'
import type { EditorConfiguration } from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/monokai.css'
import 'codemirror/mode/clike/clike.js'
import 'codemirror/addon/hint/show-hint.js'
import 'codemirror/addon/hint/show-hint.css'
import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/addon/hint/javascript-hint.js'
import 'codemirror/addon/edit/closebrackets.js'
import 'codemirror/addon/comment/comment.js'
import type { Language } from '@/models/language'
import { CodingRecorder } from '@/CodingRecorder'
import SaveVideoModal from '@/components/SaveVideoModal.vue'
import type { SaveVideoModalData } from '@/models/SaveVideoModalData'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faCircle, faStop } from '@fortawesome/free-solid-svg-icons'
import { onMounted, onUnmounted, reactive, ref, watch } from 'vue'

library.add(faCircle, faStop)

interface EditorConfig extends EditorConfiguration {
  autoCloseBrackets: boolean
}

const defaultConfig = reactive<EditorConfig>({
  lineNumbers: true,
  indentUnit: 4,
  theme: 'monokai',
  autoCloseBrackets: true,
  showHint: true,
  extraKeys: {
    'Ctrl-Space': 'autocomplete',
    'Ctrl-/': (e: CodeMirror.Editor) => e.toggleComment(),
  },
})
const recorder = reactive(new CodingRecorder())

const languages: Language[] = [
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
]
const selectedLanguage = reactive<Language>({
  tag: 'text/x-csrc',
  name: 'C言語',
})
let editor: CodeMirror.Editor | undefined
const isModalShow = ref(false)
const editorAria = ref<HTMLTextAreaElement>()
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
const recordStart = (): void => {
  recorder.start(editor)
}
const recordStop = (): void => {
  recorder.stop()
  isModalShow.value = true
}
const recordSave = (data: SaveVideoModalData): void => {
  const video = recorder.outputVideo(
    -1,
    data.author,
    data.title,
    selectedLanguage,
    data.message
  )
  var url = (window.URL || window.webkitURL).createObjectURL(
    new Blob([JSON.stringify(video)], { type: 'application/json' })
  )
  const a = document.createElement('a')
  a.href = url
  a.download = `${data.author}_${data.title}_video.json`
  a.click()
}
const recordCancel = (): void => {
  recorder.clearVideo()
}

watch(selectedLanguage, (newLang: Language) => {
  editor?.setOption('mode', newLang.tag)
})
onMounted(() => {
  console.log('マウント')
  console.log(editorAria?.value)
  const config = defaultConfig
  if (editorAria.value == null) {
    throw new Error('textarea not found for CodeMirror')
  }
  editor = CodeMirror.fromTextArea(editorAria.value, config)

  console.log(editor)
  setEditorSize()
  editor.setOption('mode', selectedLanguage.tag)
  recorder.register(editor)
  window.onresize = setEditorSize
})
onUnmounted(() => {
  recorder.unregister(editor)
  window.onresize = null
})
</script>
<template>
  <div id="editor-page">
    <div id="side-panel">
      <SaveVideoModal
        v-model:isShow="isModalShow"
        @submit="recordSave"
        @cancel="recordCancel"
      />
      <h1>Editor</h1>
      <select v-model="selectedLanguage">
        <option v-for="lang in languages" :key="lang.tag" :value="lang">
          {{ lang.name }}
        </option>
      </select>
      <button
        v-if="!recorder.isRecording"
        @click="recordStart"
        class="recorder-control-button recoding-button"
      >
        <FontAwesomeIcon icon="circle" style="color: red" />
      </button>
      <button v-else @click="recordStop" class="recorder-control-button">
        <FontAwesomeIcon icon="stop" />
      </button>
      <div v-if="recorder.isRecording" class="recoding-status">
        <span class="recoding-icon">●</span>録画中
      </div>
    </div>
    <div id="editor-panel">
      <textarea ref="editorAria"></textarea>
    </div>
  </div>
</template>

<style scoped>
h1 {
  color: #28e270;
}

#editor-page {
  grid-row: 2;
  display: grid;
  grid-template-columns: 300px 3fr;
}

#side-panel {
  grid-column: 1;
}
#editor-panel {
  grid-column: 2;
}

.recorder-control-button {
  font-size: 1.2em;
  padding: 5px 20px;
  border-radius: 25px;
  border: solid 2px #9f9f9f;
}

.recorder-control-button:disabled {
  border: solid 2px #dddddd;
}

.recorder-control-button:active {
  border: solid 2px #222222;
}

select {
  margin: 0px 10px;
}

.recoding-icon {
  color: red;
  font-size: 2em;
  vertical-align: sub;
}
.recoding-status {
  font-size: 1.5em;
  text-align: center;
}
</style>
