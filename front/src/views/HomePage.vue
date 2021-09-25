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
      <p>{{ userId }}</p>
    </div>
    <div class="controll-panel">
      <!-- prettier-ignore -->
      <router-link to="/editor"><FontAwesomeIcon icon="video" />録画画面へ</router-link>
      <!-- prettier-ignore -->
      <router-link to="/player"><FontAwesomeIcon icon="file-video" />再生画面へ</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { library } from '@fortawesome/fontawesome-svg-core'
import { faFileVideo, faVideo } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import Vue from 'vue'

library.add(faFileVideo, faVideo)

export default Vue.extend({
  name: 'HomePage',
  components: {
    FontAwesomeIcon,
  },
  computed: {
    userId() {
      return this.$store.getters.userId != ''
        ? this.$store.getters.userId
        : 'ログインしていません'
    },
  },
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
