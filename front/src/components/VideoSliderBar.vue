<template>
  <div class="slider-component">
    <VueSlider
      :lazy="true"
      v-model="value"
      :max="totalTime"
      :disabled="disabled"
      :tooltip-formatter="tooltipFormatter"
      :style="{
        width: 'calc(100% - 6em)',
        padding: '0px',
      }"
      :process-style="{
        backgroundColor: '#28e270',
        height: '1em',
        'border-radius': '0px',
      }"
      :tooltip-style="{
        backgroundColor: '#116230',
        borderColor: '#116230',
      }"
      :railStyle="{
        height: '1em',
        'border-radius': '0px',
      }"
      @change="change"
    />
    <span class="elapsedTime">{{ playbackPosition }}</span>
  </div>
</template>
<script lang="ts">
import { format } from 'date-fns'
import VueSlider from 'vue-slider-component'
import Vue from 'vue'

type DataType = { value: number }

export default Vue.extend({
  name: 'VideoSliderBar',
  components: {
    VueSlider,
  },
  props: {
    elapsedTime: {
      type: Number,
      default: 0,
    },
    totalTime: {
      type: Number,
      default: 0,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  },
  data(): DataType {
    return { value: 0 }
  },
  computed: {
    playbackPosition: function (): string {
      // prettier-ignore
      return `${format(this.elapsedTime, 'mm:ss')}/${format(this.totalTime,'mm:ss')}`
    },
  },
  methods: {
    tooltipFormatter: function (value: number): string {
      return format(value, 'mm:ss')
    },
    change: function (time: number): void {
      this.$emit('change', time)
    },
  },
  watch: {
    elapsedTime: function (newElapsedTime: number): void {
      this.value = newElapsedTime
    },
  },
})
</script>

<style scoped>
.slider-component {
  display: flex;
  justify-content: space-between;
}
</style>
