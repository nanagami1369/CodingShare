<template>
  <div class="my-page">
    <p>ようこそ {{ userId }} 様</p>
    <button @click="logout">ログアウト</button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'MyPage',
  computed: {
    userId() {
      return this.$store.getters.userId
    },
  },
  methods: {
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
