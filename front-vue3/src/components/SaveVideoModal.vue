<script setup lang="ts">
import BaseModal from '@/components/BaseModal.vue'
import RequireTag from '@/components/RequireTag.vue'
import type { SaveVideoModalData } from '@/models/SaveVideoModalData'
import { reactive, computed } from 'vue'

type FormValue = SaveVideoModalData

const formValue = reactive<FormValue>({
  author: '',
  title: '',
  message: '',
})

interface Props {
  isShow: boolean
}

const props = defineProps<Props>()

const isFormAuthorError = computed(() =>
  formValue.author.length === 0 ? '作者は必須です' : ''
)
const isFormTitleError = computed(() =>
  formValue.title.length === 0 ? 'タイトルは必須です' : ''
)
const isFormError = computed(
  () => isFormAuthorError.value !== '' || isFormTitleError.value !== ''
)

const isShow = computed({
  get: (): boolean => props.isShow,
  set: (value: boolean): void => emit('update:isShow', value),
})

interface Emits {
  (e: 'submit', value: FormValue): void
  (e: 'cancel'): void
  (e: 'update:isShow', value: boolean): void
}
const emit = defineEmits<Emits>()
const submit = (): void => {
  emit('submit', formValue)
  emit('update:isShow', false)

  // 初期化
  formValue.author = ''
  formValue.title = ''
  formValue.message = ''
}
const cancel = (): void => {
  emit('cancel')
  emit('update:isShow', false)

  // 初期化
  formValue.author = ''
  formValue.title = ''
  formValue.message = ''
}
</script>

<!-- <script lang="ts">
import type { SaveVideoUserEditData } from '@/models/SaveVideoUserEditData'

type DataType = {
  data: SaveVideoUserEditData
}

</script> -->

<template>
  <BaseModal :clickToClose="false" v-model:isShow="isShow">
    <div class="save-video-modal">
      <h1>録画データの保存</h1>
      <label for="save-modal-author">作者<RequireTag /></label>
      <input
        id="save-modal-author"
        type="text"
        name="作者"
        v-model="formValue.author"
      />
      <span class="error-message">{{ isFormAuthorError }}</span>
      <label for="save-modal-title">タイトル <RequireTag /></label>
      <input
        id="save-modal-title"
        type="text"
        name="タイトル"
        v-model="formValue.title"
      />
      <span class="error-message">{{ isFormTitleError }}</span>
      <label for="save-modal-message">コメント</label>
      <textarea
        id="save-modal-message"
        rows="5"
        resize="none"
        v-model="formValue.message"
      ></textarea>
      <div class="submit-button-area">
        <button
          type="button"
          class="submit-button"
          :disabled="isFormError"
          @click="submit"
        >
          録画データを保存する
        </button>
        <button type="button" @click="cancel">
          録画データを<strong>保存せずに</strong>消す
        </button>
      </div>
    </div>
  </BaseModal>
</template>

<style scoped>
.save-video-modal {
  display: flex;
  flex-direction: column;
  margin: 10px;
  width: 600;
}
.save-video-modal h1 {
  font-weight: bold;
  font-size: 1em;
}

.save-video-modal input[type='text']:not(:focus) {
  border: solid #111 1.5px;
}
.save-video-modal textarea:not(:focus) {
  border: solid #111 1.5px;
}

button {
  font-size: 1em;
  margin: 5px auto;
  padding: 4px 1em;
  border: solid #111 2px;
  background-color: solid #fff 2px;
}

button:disabled {
  background-color: #ccc;
  color: #777;
}

button:hover:not(:disabled) {
  border: solid #111 2px;
  background-color: #ccc;
}

button:active:not(:disabled) {
  border: solid #111 2px;
  background-color: #aaa;
}

.submit-button:disabled {
  background-color: #9cffe1;
}

.submit-button {
  background-color: #43ff92;
}

.submit-button:hover:not(:disabled) {
  background-color: #a4ffca;
}

.submit-button:active:not(:disabled) {
  background-color: #d6ffe7;
}

.submit-button-area {
  display: flex;
  justify-content: space-between;
}

.save-video-modal textarea {
  resize: none;
}

.error-message {
  text-decoration: underline;
  height: 1.4em;
  color: #da5c58;
}
.error-message:not(:empty)::before {
  content: '※ ';
}
</style>
