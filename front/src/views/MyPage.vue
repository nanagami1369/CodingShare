<template>
  <div class="my-page">
    <p>ようこそ {{ userName }} 様</p>
    <button @click="logout">ログアウト</button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

type DataType = {
  userName: string
}

export default Vue.extend({
  name: 'MyPage',
  data(): DataType {
    return {
      userName: '',
    }
  },
  methods: {
    logout: async function () {
      try {
        const response = await fetch(
          `${process.env.VUE_APP_CODING_SHARE_API_URL}/api/logout`,
          {
            method: 'POST',
            mode: 'cors',
            credentials: 'include',
          }
        )
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
  async created(): Promise<void> {
    try {
      const response = await fetch(
        `${process.env.VUE_APP_CODING_SHARE_API_URL}/api/auth`,
        {
          method: 'POST',
          mode: 'cors',
          credentials: 'include',
        }
      )
      if (response.status != 200) {
        // 失敗時はエラーを表示
        let message =
          `message:${await response.text()}\n` +
          `http status:${response.status} ${response.statusText}`
        alert(message)
        return
      }
      this.userName = (await response.json()).userId
    } catch (error: unknown) {
      // fetchの例外はエラーとして処理
      alert((error as Error).message)
    }
  },
})
</script>
