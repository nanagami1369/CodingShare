<template>
  <header class="header">
    <button class="header-button" @click="changeHomePage">
      <FontAwesomeIcon icon="home" />
    </button>
    <span class="header-free-space"></span>
    <button
      v-if="isLogin"
      class="header-button"
      @click="toggleHeaderContextMenu"
    >
      <FontAwesomeIcon icon="user-circle" />
    </button>
    <button v-else class="header-login-button" @click="changeLogin">
      ログイン
    </button>
    <div
      v-show="isHeaderContextMenu"
      class="header-context-menu-wrapper"
      @click="toggleHeaderContextMenu"
    >
      <div class="header-context-menu">
        <div class="header-context-menu-item" @click="changeMyPage">
          マイページ
        </div>
      </div>
    </div>
  </header>
</template>
<script lang="ts">
import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faHome, faUserCircle } from '@fortawesome/free-solid-svg-icons'

library.add(faHome, faUserCircle)

type DataType = {
  isHeaderContextMenu: boolean
}

export default Vue.extend({
  name: 'HomePage',
  components: {
    FontAwesomeIcon,
  },
  data(): DataType {
    return {
      isHeaderContextMenu: false,
    }
  },
  computed: {
    isLogin() {
      return this.$store.getters.isLogin
    },
  },
  methods: {
    changeHomePage: function (): void {
      this.$router.push({ path: '/' })
    },
    changeMyPage: function (): void {
      this.$router.push({ path: '/mypage' })
    },
    toggleHeaderContextMenu: function (): void {
      this.isHeaderContextMenu = !this.isHeaderContextMenu
    },
    changeLogin: function (): void {
      this.$router.push({ path: '/login' })
    },
  },
})
</script>

<style scoped>
.header {
  background: #202020;
  color: #eeeeee;
  display: flex;
  padding: 0px 15px;
}
.header-button {
  cursor: pointer;
  color: #eeeeee;
  height: 100%;
  font-size: 1.5em;
  flex: 0 0 auto;
  margin: 0px 5px;
  background-color: transparent;
  border: none;
}
.header-free-space {
  flex: 1 1 auto;
}

.header-context-menu-wrapper {
  position: absolute;
  height: 100%;
  width: 100%;
  right: 0px;
  top: 0px;
  z-index: 100;
}
.header-context-menu {
  position: absolute;
  right: 30px;
  top: 25px;
  font-size: 1.2em;
  padding: 5px 7px;
  background: white;
  box-shadow: 1px 0 5px 1px #999;
}

.header-context-menu {
  color: #202020;
}

.header-login-button {
  background-color: #111111;
  border: solid 2px #eeeeee;
  border-radius: 25px;
  padding: 0px 30px;
  margin: 5px;
  color: #eeeeee;
}
.header-login-button:hover {
  background-color: #404040;
}
.header-login-button:active {
  background-color: #6d6d6d;
}
</style>
