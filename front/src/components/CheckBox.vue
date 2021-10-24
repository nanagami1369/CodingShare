<template>
  <div class="checkbox">
    <input
      type="checkbox"
      :id="id"
      :value="value"
      v-model="internalChecked"
      @change.stop="change"
      :name="name"
      class="checkbox-origin"
    />
    <label :for="id" class="checkbox-appearance">
      <FontAwesomeIcon v-if="!internalChecked" icon="square" />
      <FontAwesomeIcon v-else icon="check-square" />
    </label>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faCheckSquare, faSquare } from '@fortawesome/free-solid-svg-icons'

library.add(faCheckSquare, faSquare)

type DataType = {
  internalChecked: boolean
}

export default Vue.extend({
  name: 'CheckBox',
  components: {
    FontAwesomeIcon,
  },
  data(): DataType {
    return {
      internalChecked: false,
    }
  },
  props: {
    id: {
      type: String,
      required: true,
    },
    value: {
      type: String,
      required: true,
    },
    checked: {
      type: Boolean,
      required: false,
      default: () => false,
    },
    name: {
      type: String,
      default: () => '',
    },
    resetFlag: {
      type: Boolean,
      default: () => false,
    },
  },
  methods: {
    change: function (event: Event) {
      const checkbox = event.target as HTMLInputElement
      this.$emit('change', checkbox.checked)
    },
  },
  watch: {
    checked: function (newChecked: boolean): void {
      this.internalChecked = newChecked
    },
    resetFlag: function (): void {
      this.internalChecked = false
      this.$emit('change', false)
    },
  },
  model: {
    prop: 'checked',
    event: 'change',
  },
})
</script>

<style scoped>
.checkbox-origin {
  display: none;
}

.checkbox-appearance {
  background-color: #ee9d28;
  border-radius: 45px;
  height: 45px;
  width: 45px;
  line-height: 45px;
  margin: auto;
  font-size: 1.7em;
  color: white;
  border: none;
  display: inline-block;
}
</style>
