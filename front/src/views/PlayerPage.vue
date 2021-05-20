<template>
  <div id="player-page">
    <div id="side-panel">
      <h1>Player</h1>
      <input type="file" @change="loadData" value="読み込み" />
    </div>
    <div id="player-panel">
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
import { Video } from '@/models/Video.js'
import { CodingStream } from '@/CodingStream'
type DataType = {
  editor?: CodeMirror.EditorFromTextArea
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
  doSomething: () => { nextSpan: number; isNext: boolean }
): void {
  const { isNext, nextSpan } = doSomething()
  if (isNext) {
    setTimeout(function () {
      doSomethingLoop(doSomething)
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
        readOnly: true,
      },
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
      const stream = new CodingStream(video)
      console.log(video)
      const language = video.header.language
      this.editor?.setOption('mode', language.tag)
      this.editor?.setValue('')
      this.editor?.focus()
      doSomethingLoop((): { isNext: boolean; nextSpan: number } => {
        const { text, from, to, origin } = stream.current.changeData
        const cursor = stream.current.cursor
        this.editor?.replaceRange(text, from, to, origin)
        this.editor?.setCursor(cursor)
        stream.next()
        const isNext = stream.isNext()
        if (stream.from === undefined) {
          return { isNext: isNext, nextSpan: stream.current.timestamp }
        }
        if (stream.to === undefined) {
          // 次の要素が無いので最後の要素を表示して終了
          console.log('終了')
          const { text, from, to, origin } = stream.current.changeData
          const cursor = stream.current.cursor
          this.editor?.replaceRange(text, from, to, origin)
          this.editor?.setCursor(cursor)
          return { isNext: isNext, nextSpan: 1 }
        }
        return {
          isNext: isNext,
          nextSpan: stream.to.timestamp - stream.current.timestamp,
        }
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
