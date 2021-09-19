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
      const response = await fetch(
        `${process.env.VUE_APP_CODING_SHARE_API_URL}/api/logout`,
        {
          method: 'POST',
          mode: 'cors',
          credentials: 'include',
        }
      )
      let message =
        `message:${await response.text()}\n` +
        `http status:${response.status} ${response.statusText}`
      alert(message)
      this.$router.push('/login')
    },
  },
  async created(): Promise<void> {
    const response = await fetch(
      `${process.env.VUE_APP_CODING_SHARE_API_URL}/api/auth`,
      {
        method: 'POST',
        mode: 'cors',
        credentials: 'include',
      }
    )
    this.userName = (await response.json()).userId
  },
})
</script>
