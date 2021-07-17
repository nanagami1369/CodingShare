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
        <label
          >作者
          <ValidationProvider
            :immediate="true"
            rules="required"
            v-slot="{ errors }"
          >
            <input type="text" name="作者" v-model="data.name" />
            <span class="error-message">{{ errors[0] }}</span>
          </ValidationProvider>
        </label>
        <label
          >タイトル
          <ValidationProvider
            :immediate="true"
            rules="required"
            v-slot="{ errors }"
          >
            <input type="text" name="タイトル" v-model="data.title" />
            <span class="error-message">{{ errors[0] }}</span>
          </ValidationProvider>
        </label>
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
  methods: {
    submit: function (): void {
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
