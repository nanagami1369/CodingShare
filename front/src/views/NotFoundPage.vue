<template>
  <div class="not-found-page">
    <div>
      <h1>404 Not found</h1>
      <p>そのURLは存在しません</p>
      <div class="page-change-controller">
        <router-link
          v-if="this.$store.getters.isLogin"
          :style="{ 'text-decoration': 'none' }"
          to="/mypage"
        >
          <div class="change-page-button">
            <div>マイページ</div>
            <FontAwesomeIcon icon="user" />
          </div>
        </router-link>
        <router-link v-else :style="{ 'text-decoration': 'none' }" to="/login">
          <div class="change-page-button">
            <div>ログインページ</div>
            <FontAwesomeIcon icon="sign-in-alt" />
          </div>
        </router-link>
        <router-link :style="{ 'text-decoration': 'none' }" to="/editor">
          <div class="change-page-button">
            <div>録画ページ</div>
            <FontAwesomeIcon icon="video" />
          </div>
        </router-link>
        <router-link :style="{ 'text-decoration': 'none' }" to="/player">
          <div class="change-page-button">
            <div>再生ページ</div>
            <FontAwesomeIcon icon="file-video" />
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faUser,
  faSignInAlt,
  faVideo,
  faFileVideo,
} from '@fortawesome/free-solid-svg-icons'

library.add(faUser, faSignInAlt, faVideo, faFileVideo)

export default Vue.extend({
  name: 'NotFoundPage',
  components: {
    FontAwesomeIcon,
  },
  computed: {
    topPageUrl: function (): string {
      return `${location.host}${process.env.BASE_URL}#/`
    },
  },
})
</script>
<style scoped>
.not-found-page {
  font-size: 3em;
  display: flex;
  justify-content: center;
  align-items: center;
}
.not-found-page pre {
  text-align: left;
}
.page-change-controller {
  display: flex;
}
.change-page-button {
  border: solid 2px #202020;
  background: #ffffff;
  padding: 10px 20px;
  font-size: 1em;
  margin: 0px 5px;
  text-decoration: none;
  color: black;
}

.change-page-button:hover {
  color: #28e270;
}

.change-page-button:active {
  background: #cccccc;
}
</style>
