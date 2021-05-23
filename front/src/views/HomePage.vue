<template>
  <div class="home-page">
    <div class="logo-panel">
      <h1 class="app-logo">CodingShare{}</h1>
      <p>本サイトはベータ版です。</p>
      <p>
        本ソフトを使用したことによるいかなる損害について、開発者は何も保証しません。
      </p>
      <p>
        これに同意できない場合、あなたは本ソフトを使用することができません。
      </p>
    </div>
    <div class="controll-panel">
      <a href="/#/editor"><i class="fas fa-video"></i>録画画面へ</a>
      <a href="/#/player"><i class="fas fa-file-video"></i>再生画面へ</a>
    </div>
  </div>
</template>

<script lang="ts">
import '@fortawesome/fontawesome-free/js/fontawesome'
import '@fortawesome/fontawesome-free/js/solid'
import Vue from 'vue'

export default Vue.extend({
  name: 'HomePage',
})
</script>

<style scoped>
.app-logo {
  font-size: 5em;
}

.home-page {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.controll-panel a {
  text-decoration: none;
  border-bottom: solid 2px black;
  color: black;
}
.controll-panel a:hover {
  color: #28e270;
  border-bottom: solid 2px #28e270;
}

.controll-panel a:active {
  color: #ff4500;
  border-bottom: solid 2px #ff4500;
}

.controll-panel {
  display: flex;
  flex-direction: column;
  font-size: 2em;
}
</style>
