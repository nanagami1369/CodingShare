<template>
  <modal
    name="save-video-modal"
    :adaptive="true"
    :clickToClose="false"
    :scrollable="true"
    :height="350"
  >
    <ValidationObserver :immediate="true" v-slot="{ invalid }">
      <div class="save-video-modal">
        <h4>録画データの保存</h4>
        <label v-if="!isLogin">
          <span class="label">作者</span>
          <ValidationProvider
            :immediate="true"
            rules="required"
            v-slot="{ errors }"
          >
            <input type="text" name="作者" v-model="data.name" />
            <span class="error-message">{{ errors[0] }}</span>
          </ValidationProvider>
        </label>
        <div v-else>
          <span class="label">作者</span>
          <span>{{ userId }}</span>
        </div>
        <div v-if="$route.query.title">
          <span class="label">タイトル</span>
          <span>{{ $route.query.title }}</span>
        </div>
        <div v-else>
          <label>
            <span class="label">タイトル</span>
            <ValidationProvider
              :immediate="true"
              rules="required"
              v-slot="{ errors }"
            >
              <input type="text" name="タイトル" v-model="data.title" />
              <span class="error-message">{{ errors[0] }}</span>
            </ValidationProvider>
          </label>
        </div>
        <label>コメント</label>
        <textarea rows="5" resize="none" v-model="data.comment"></textarea>
        <button
          class="submit-button"
          type="button"
          @click="submit"
          :disabled="invalid"
        >
          録画データを保存する
        </button>
        <button class="submit-button" type="button" @click="cancel">
          録画データを保存せずに消す
        </button>
        <p v-show="invalid" class="error-message">入力ミスがあります</p>
      </div>
    </ValidationObserver>
  </modal>
</template>

<script lang="ts">
import { SaveVideoUserEditData } from '@/models/SaveVideoUserEditData'

import { ValidationProvider, ValidationObserver, extend } from 'vee-validate'
import { required } from 'vee-validate/dist/rules'
import Vue from 'vue'

extend('required', {
  ...required,
  message: '{_field_}は必須です',
})

type DataType = {
  data: SaveVideoUserEditData
}

export default Vue.extend({
  name: 'SaveVideoModal',
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  data(): DataType {
    return {
      data: {
        name: '',
        title: '',
        comment: '',
      },
    }
  },
  computed: {
    isLogin: function (): boolean {
      return this.$store.getters.isLogin
    },
    userId: function (): string {
      return this.$store.getters.userId
    },
  },
  methods: {
    submit: function (): void {
      if (this.isLogin) {
        // ログイン中なら名前にユーザIDを挿入
        this.data.name = this.userId
      }
      if (this.$route.query.title) {
        // クエリにtitleがあるならtitleを挿入
        this.data.title = this.$route.query.title as string
      }
      this.$emit('submit', this.data)
      this.$modal.hide('save-video-modal')
      this.data = {
        name: '',
        title: '',
        comment: '',
      }
    },
    cancel: function (): void {
      this.$emit('cancel')
      this.$modal.hide('save-video-modal')
      this.data = {
        name: '',
        title: '',
        comment: '',
      }
    },
  },
})
</script>

<style scoped>
.save-video-modal {
  display: flex;
  flex-direction: column;
  margin: 10px;
}

.label {
  display: inline-block;
  width: 5em;
}

.submit-button {
  font-size: 1em;
  margin: 0.5rem 5rem;
}
#submit-button:disabled {
  background-color: gray;
}
.error-message {
  text-decoration: underline;
}
.error-message:not(:empty)::before {
  content: '※ ';
  color: red;
}
</style>
