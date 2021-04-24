<template>
  <div class="editor">
    <h1>Editor</h1>
    <div id="editor" style="height: 500px"></div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import * as monaco from 'monaco-editor'

export default Vue.extend({
  name: 'Editor',
  mounted() {
    const monacoEditorContainer: HTMLElement | null = document.querySelector(
      '#editor'
    )
    if (monacoEditorContainer) {
      monaco.editor.setTheme('vs-dark')
      monaco.editor.create(monacoEditorContainer, {
        value: 'console.log("Hello, world")',
        language: 'javascript',
        fontSize: 25,
      })
    }
  },
})
</script>

<style scoped>
h1 {
  color: #28e270;
}
#editor {
  height: 600px;
  text-align: start;
  font-size: 2em;
}
</style>
