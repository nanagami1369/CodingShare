<template>
  <div class="search-page">
    <div class="search-box">
      <input type="search" v-model="searchWord" @keydown.enter="search" />
      <button @click="search">
        <span>検索</span><span>&nbsp;</span><FontAwesomeIcon icon="search" />
      </button>
    </div>
    <div v-if="searchedVideos.length == 0" class="search-resurlt">
      {{ this.$route.query.q }}に一致する検索結果はありません
    </div>
    <div v-else>
      <div class="search-ward-title">{{ this.$route.query.q }}の検索結果</div>
      <article class="search-resurlt-list">
        <Thumbnail
          v-for="video in searchedVideos"
          :key="video.header.videoId"
          :video="video"
        />
      </article>
    </div>
  </div>
</template>

<script lang="ts">
import { Video } from '@/models/Video'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faSearch } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import Thumbnail from '@/components/Thumbnail.vue'
import Vue from 'vue'

library.add(faSearch)

type DataType = {
  searchWord: string | (string | null)[]
  searchedVideos: Video[]
}

export default Vue.extend({
  name: 'HomePage',
  components: {
    FontAwesomeIcon,
    Thumbnail,
  },
  data(): DataType {
    return {
      searchWord: '',
      searchedVideos: [],
    }
  },
  computed: {
    query: function (): string | (string | null)[] {
      return this.$route.query.q
    },
  },
  methods: {
    search: function (): void {
      if (!this.searchWord) {
        return
      }
      if (this.searchWord == this.query) {
        return
      }
      this.$router.push({ path: '/search', query: { q: this.searchWord } })
    },
    observerUrlDo: async function (): Promise<void> {
      this.searchWord = this.query
      if (!this.searchWord) {
        return
      }
      try {
        var response = await fetch(`/api/searchvideo?q=${this.searchWord}`)
        if (response.ok) {
          this.searchedVideos = (await response.json()) as Video[]
        } else {
          alert(`通信エラー\n${await response.text()}`)
        }
      } catch (error: unknown) {
        // 通信エラーの場合はアラートで表示
        alert((error as Error).message)
        return
      }
    },
  },
  created(): void {
    this.observerUrlDo()
  },
  watch: {
    async $route(): Promise<void> {
      this.observerUrlDo()
    },
  },
})
</script>

<style scoped>
.search-page {
  display: flex;
  flex-direction: column;
}
.search-box {
  flex: 0 0 40px;
  display: flex;
  flex-direction: row;
  margin: 0px 25px;
  margin-top: 25px;
}
.search-box input[type='search'] {
  flex: 0 1 600px;
  border: solid 2px #cccccc;
  outline: none;
}
.search-box input[type='search']:focus {
  border: solid 2px #66afe9;
}

.search-box button {
  border: solid 2px #202020;
  background: #ffffff;
  flex: 0 0 80px;
}
.search-box button:hover {
  background: #cccccc;
}
.search-box button:active {
  background: #aaaaaa;
}
.search-ward-title {
  flex: 0 0 auto;
  text-align: left;
  font-size: 1.5em;
  margin: 10px 25px;
}
.search-resurlt-list {
  flex: 1 1 auto;
  display: flex;
  flex-wrap: wrap;
  margin: 10px;
}
</style>
