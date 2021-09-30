<template>
  <div class="search-page">
    <div class="search-box">
      <input type="search" v-model="searchWord" @keydown.enter="search" />
      <button @click="search">
        <span>検索</span><span>&nbsp;</span><FontAwesomeIcon icon="search" />
      </button>
    </div>
    <article class="search-resurlt-list">
      {{ this.$route.query.q }}に一致する検索結果はありません
    </article>
  </div>
</template>

<script lang="ts">
import { library } from '@fortawesome/fontawesome-svg-core'
import { faSearch } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import Vue from 'vue'

library.add(faSearch)

type DataType = {
  searchWord: string | (string | null)[]
}

export default Vue.extend({
  name: 'HomePage',
  components: {
    FontAwesomeIcon,
  },
  data(): DataType {
    return {
      searchWord: '',
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
    observerUrlDo: function (): void {
      this.searchWord = this.query
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
  margin: 25px 25px;
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

.search-resurlt-list {
  flex: 1 1 auto;
  margin-top: 20px;
}
</style>
