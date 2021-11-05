<template>
  <header class="header">
    <button class="header-button" @click="changeHomePage">
      <FontAwesomeIcon icon="home" />
    </button>
    <span class="header-free-space"></span>
    <div class="header-search-box">
      <input type="search" v-model="searchWord" @keydown.enter="search" />
      <button class="search-button" @click="search">
        <div class="search-button-style">
          <FontAwesomeIcon icon="search" />
        </div>
      </button>
    </div>
    <span class="header-free-space"></span>
    <button class="header-button" @click="changeEditorPage">
      <FontAwesomeIcon icon="edit" />
    </button>
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
        <div class="header-context-menu-item" @click="logout">ログアウト</div>
      </div>
    </div>
  </header>
</template>
<script lang="ts">
import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faHome,
  faUserCircle,
  faEdit,
  faSearch,
} from '@fortawesome/free-solid-svg-icons'

library.add(faHome, faUserCircle, faEdit, faSearch)

type DataType = {
  isHeaderContextMenu: boolean
  searchWord: string
}

export default Vue.extend({
  name: 'HomePage',
  components: {
    FontAwesomeIcon,
  },
  data(): DataType {
    return {
      isHeaderContextMenu: false,
      searchWord: '',
    }
  },
  computed: {
    isLogin() {
      return this.$store.getters.isLogin
    },
  },
  methods: {
    changeHomePage: function (): void {
      if (this.$route.path == '/') {
        return
      }
      this.$router.push({ path: '/' })
    },
    changeMyPage: function (): void {
      if (this.$route.path == '/mypage') {
        return
      }
      this.$router.push({ path: '/mypage' })
    },
    toggleHeaderContextMenu: function (): void {
      this.isHeaderContextMenu = !this.isHeaderContextMenu
    },
    changeLogin: function (): void {
      if (this.$route.path == '/login') {
        return
      }
      this.$router.push({ path: '/login' })
    },
    changeEditorPage: function (): void {
      if (this.$route.path == '/editor') {
        return
      }
      this.$router.push({ path: '/editor' })
    },
    search: function (): void {
      if (!this.searchWord) {
        return
      }
      if (this.searchWord != this.$route.query.q) {
        this.$router.push({ path: '/search', query: { q: this.searchWord } })
      }
      this.searchWord = ''
    },
    logout: async function () {
      try {
        const response = await fetch('/api/private/logout', {
          method: 'POST',
          mode: 'cors',
          credentials: 'include',
        })
        if (response.status != 200) {
          // 失敗時はエラーを表示
          let message =
            `message:${await response.text()}\n` +
            `http status:${response.status} ${response.statusText}`
          alert(message)
          return
        }
        this.$router.push('/login')
      } catch (error: unknown) {
        // 通信エラーの場合はアラートで表示
        alert((error as Error).message)
      }
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
.header-search-box {
  background-color: white;
  border-radius: 10px;
  border: solid 2px transparent;
  outline: solid 2px transparent;
  outline-offset: 1px;
  display: flex;
  flex: 3 3 auto;
  margin: 7px 5px;
}
.header-search-box:focus-within {
  border: solid 2px #0060df;
  outline: solid 2px #b2ddf9;
}
.header-search-box input[type='search'] {
  flex: 1 1 auto;
  padding: 0px 10px;
  font-size: 1em;
  border: none;
  border-top-left-radius: 11px;
  border-bottom-left-radius: 11px;
}
.header-search-box input[type='search']:focus {
  outline: none;
}

.search-button {
  cursor: pointer;
  background-color: transparent;
  border: solid 1px transparent;
}
.search-button-style {
  padding: 0px 10px;
  border-radius: 45px;
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
  background: white;
  box-shadow: 1px 0 5px 1px #999;
}

.header-context-menu {
  color: #202020;
}

.header-context-menu-item {
  border: 1px solid #aaaaaa;
  padding: 5px 7px;
}

.header-context-menu-item:hover {
  cursor: pointer;
  background-color: #ddd;
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
