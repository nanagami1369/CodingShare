<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/editor">Editor</router-link> |
      <router-link to="/player">Player</router-link> |
      <router-link to="/list">list</router-link>
    </div>
    <router-view />
  </div>
</template>

<style>
* {
  margin: 0px;
  padding: 0px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
  display: grid;
  grid-template-rows: 50px 1fr;
  /* 画面がスクロールしないようにする */
  overflow-x: hidden;
  overflow-y: hidden;
}

#nav {
  padding: 15px;
  grid-row: 1;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

html {
  height: 100%;
}

body {
  height: 100%;
}

.CodeMirror {
  text-align: start;
  font-size: 2em;
}

.CodeMirror-hint {
  font-size: 2em !important;
}

/* PlayerPageで再生中のみカーソルを表示させるCSS */
#player-page.is-play .CodeMirror:not(.CodeMirror-focused) .CodeMirror-cursors {
  visibility: visible !important;
}

#player-page.is-play .CodeMirror:not(.CodeMirror-focused) .CodeMirror-cursor {
  animation: flash 1s linear infinite;
}

@keyframes flash {
  0%,
  50% {
    opacity: 1;
  }

  51%,
  100% {
    opacity: 0;
  }
}
</style>
