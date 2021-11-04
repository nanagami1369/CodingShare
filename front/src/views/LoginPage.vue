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
    <div class="control-panel">
      <p class="error-message" v-show="errorMessage != ''">
        {{ errorMessage }}
      </p>
      <p>ユーザID</p>
      <input
        type="text"
        name="name"
        autocomplete="username"
        autocorrect="off"
        autocapitalize="off"
        v-model="id"
      />
      <p>パスワード</p>
      <input
        type="password"
        name="password"
        autocomplete="current-password"
        v-model="password"
      />
      <button @click="login" type="button">ログイン</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

type DataType = {
  id: string
  password: string
  errorMessage: string
}

export default Vue.extend({
  name: 'HomePage',
  data(): DataType {
    return {
      id: '',
      password: '',
      errorMessage: '',
    }
  },
  methods: {
    login: async function (): Promise<void> {
      if (this.id == null || this.id.length == 0) {
        this.errorMessage = 'ユーザーIDが空です'
        return
      }
      if (this.password == null || this.password.length == 0) {
        this.errorMessage = 'パスワードが空です'
        return
      }
      const isStudentNumber = /^[0-9]{7}$/
      if (isStudentNumber.test(this.id)) {
        this.errorMessage = '学生はKBookからログインしてください'
        return
      }
      if (this.id)
        try {
          const response = await fetch('/api/login', {
            method: 'POST',
            mode: 'cors',
            credentials: 'include',
            body: JSON.stringify({
              id: this.id,
              password: this.password,
            }),
          })
          switch (response.status) {
            case 200:
              // ログイン後の処理を書く
              this.$router.push('mypage')
              break
            case 401:
              this.errorMessage = 'ログインに失敗しました'
              break
            default:
              // それ以外はHTTP STATUSを返却するようにする
              this.errorMessage = `Error:${response.status} ${response.statusText}`
              break
          }
        } catch (error: unknown) {
          // fetchの例外はエラーとして処理
          this.errorMessage = (error as Error).message
        }
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

.control-panel {
  display: flex;
  flex-direction: column;
  width: 60vh;
}

.control-panel input {
  font-size: 1.5em;
  border-radius: 10px;
  border: solid 2px black;
}

.control-panel p {
  font-size: 1em;
  text-align: left;
}

.control-panel .error-message {
  color: #a94442;
  text-decoration: underline;
  background-color: #f2dede;
  border: solid 2px #ebccd1;
  border-radius: 10px;
  padding: 5px;
}

.control-panel button {
  padding: 5px 10px;
  margin: 5px 0px;
  background-color: #3dcf64;
  border: solid green 2px;
  border-radius: 25px;
  width: 150px;
}
.control-panel button:active {
  background-color: #77eb96;
}
</style>
