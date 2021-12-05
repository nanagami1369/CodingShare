<template>
  <div id="editor-page">
    <div id="side-panel">
      <SaveVideoModal @submit="recordSave" @cancel="recordCancel" />
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
import 'codemirror/addon/comment/comment.js'
import { Language } from '@/models/language'
import { CodingRecorder } from '@/CodingRecorder'
import SaveVideoModal from '@/components/SaveVideoModal.vue'
import { SaveVideoUserEditData } from '@/models/SaveVideoUserEditData.js'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faCircle, faStop } from '@fortawesome/free-solid-svg-icons'

library.add(faCircle, faStop)

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
  components: {
    SaveVideoModal,
    FontAwesomeIcon,
  },
  data(): DataType {
    return {
      defualtConfig: {
        lineNumbers: true,
        indentUnit: 4,
        theme: 'monokai',
        autoCloseBrackets: true,
        showHint: true,
        extraKeys: {
          'Ctrl-Space': 'autocomplete',
          'Ctrl-/': (e: CodeMirror.Editor) => e.toggleComment(),
        },
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
        tag: 'text/x-csrc',
        name: 'C言語',
      },
    }
  },
  computed: {
    isLogin: function (): boolean {
      return this.$store.getters.isLogin
    },
    userId: function (): string {
      return this.$store.getters.userId
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
    recordStart: function () {
      this.recorder.start(this.editor)
    },
    recordStop: function () {
      this.recorder.stop()
      this.$modal.show('save-video-modal')
    },
    recordSave: async function (data: SaveVideoUserEditData): Promise<void> {
      const video = this.recorder.outputVideo(
        this.isLogin ? this.userId : 'file',
        this.isLogin ? this.userId : data.name,
        data.title,
        this.selectedLanguage,
        data.comment
      )
      if (this.isLogin) {
        try {
          const response = await fetch('/api/private/savevideo', {
            method: 'POST',
            mode: 'cors',
            credentials: 'include',
            body: JSON.stringify(video),
          })
          if (response.ok) {
            const videoId = (
              (await response.json()) as { message: string; videoId: number }
            ).videoId
            this.$router.push('/savedvideo/' + videoId)
          } else {
            alert('サーバーへの保存に失敗しました')
          }
        } catch (error: unknown) {
          // 通信エラーの場合はアラートで表示
          alert('サーバーへの保存に失敗しました:' + (error as Error).message)
          return
        }
        return
      } else {
        var url = (window.URL || window.webkitURL).createObjectURL(
          new Blob([JSON.stringify(video)], { type: 'application/json' })
        )
        const a = document.createElement('a')
        a.href = url
        a.download = `${data.name}_${data.title}_video.json`
        a.click()
        return
      }
    },
    recordCancel: function (): void {
      this.recorder.clearVideo()
    },
    observerUrlDo: async function (): Promise<void> {
      this.editor?.setValue('')
      const template = this.$route.query.template
      if (!template) {
        return
      }
      try {
        const response = await fetch('/templates/' + template)
        if (response.ok) {
          const text = await response.text()
          const headerText = text.split('\n')[0].substring(3)
          const header = JSON.parse(headerText) as {
            language: { tag: string; name: string }
          }
          const body = text
            .split('\n')
            .slice(1, text.split('\n').length)
            .join('\n')
          this.editor?.setOption('mode', header.language.tag)
          this.editor?.setValue(body)
          this.selectedLanguage = header.language
          return
        }
        if (response.status == 404) {
          return
        }
        const message =
          `録画データの取得に失敗しました\n` +
          `message:${await response.text()}\n` +
          `http status:${response.status} ${response.statusText}`
        alert(message)
        return
      } catch (error: unknown) {
        // 通信エラーの場合はアラートで表示
        alert((error as Error).message)
      }
    },
  },
  watch: {
    selectedLanguage: function (newLang: Language) {
      this.editor?.setOption('mode', newLang.tag)
    },
    async $route(): Promise<void> {
      // 2回目呼び出し
      await this.observerUrlDo()
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
    this.setEditorSize()
    this.editor.setOption('mode', this.selectedLanguage.tag)
    this.recorder.register(this.editor)
    window.onresize = this.setEditorSize
    // 1回目呼び出し
    this.observerUrlDo()
  },
  beforeDestroy() {
    this.recorder.unregister(this.editor)
    window.onresize = null
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
