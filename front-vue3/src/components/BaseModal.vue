<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'BaseModal',
  props: {
    isShow: {
      type: Boolean,
      default: false,
    },
    clickToClose: {
      type: Boolean,
      default: true,
    },
  },
  emits: ['update:isShow'],

  methods: {
    closeModal: function (): void {
      if (this.clickToClose) {
        this.$emit('update:isShow', false)
      } else {
        // 何もしない
      }
    },
  },
})
</script>
<template>
  <Teleport to="#modal-target">
    <div class="modal-background" v-show="isShow" @click.self="closeModal">
      <div class="modal-content">
        <slot>ここに要素を入れる（予定）</slot>
      </div>
    </div>
  </Teleport>
</template>

<style scoped>
.modal-background {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  top: 0px;
  left: 0px;
  width: 100%;
  height: 100%;
  background-color: #202020e6;
  animation: show 0.1s ease-in-out 0s;
}

.modal-content {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  background-color: white;
  flex: 0 1 500px;
  margin: 2px;
  box-sizing: border-box;
}

@keyframes show {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
