<template>
  <div id="editor">
    <h1>Editor</h1>
    <select v-model="selectedLanguage">
      <option v-for="lang in languages" :key="lang.tag" :value="lang">
        {{ lang.name }}
      </option>
    </select>
    <div class="editorAria">
      <textarea id="editorAria">#TEST</textarea>
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
import { Language } from '@/models/language'
type DataType = {
  editor?: CodeMirror.EditorFromTextArea
  defualtConfig: EditorConfiguration
  languages: Language[]
  selectedLanguage: Language
}

export default Vue.extend({
  name: 'Editor',
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
    }
  },
  watch: {
    selectedLanguage: function (newLang: Language, _: Language) {
      this.editor?.setOption('mode', newLang.tag)
    },
  },
  mounted() {
    const editorAria: HTMLTextAreaElement | null = document.querySelector(
      '#editorAria'
    )
    if (editorAria != null) {
      const config = this.defualtConfig
      this.editor = CodeMirror.fromTextArea(editorAria, config)
    }
  },
})
</script>

<style scoped>
h1 {
  color: #28e270;
}

.editorAria {
  text-align: start;
  font-size: 2em;
}
</style>
