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
    <div class="controll-panel">
      <p class="error-message" v-show="errorMessage != ''">
        {{ errorMessage }}
      </p>
      <p>ユーザID</p>
      <input
        type="text"
        name="name"
        autocomplete="name"
        autocorrect="off"
        autocapitalize="off"
        v-model="id"
      />
      <p>パスワード</p>
      <input
        type="password"
        name="password"
        autocomplete="password"
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
      try {
        const response = await fetch('http://localhost:8081/login', {
          method: 'POST',
          body: JSON.stringify({
            id: this.id,
            password: this.password,
          }),
        })
        const message = await response.text()
        switch (response.status) {
          case 200:
            // ログイン後の処理を書く
            alert(JSON.parse(message).message)
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

.controll-panel {
  display: flex;
  flex-direction: column;
  width: 60vh;
}

.controll-panel input {
  font-size: 1.5em;
  border-radius: 10px;
  border: solid 2px black;
}

.controll-panel p {
  font-size: 1em;
  text-align: left;
}

.controll-panel .error-message {
  color: #a94442;
  text-decoration: underline;
  background-color: #f2dede;
  border: solid 2px #ebccd1;
  border-radius: 10px;
  padding: 5px;
}

.controll-panel button {
  padding: 5px 10px;
  margin: 5px 0px;
  background-color: #3dcf64;
  border: solid green 2px;
  border-radius: 25px;
  width: 150px;
}
.controll-panel button:active {
  background-color: #77eb96;
}
</style>
