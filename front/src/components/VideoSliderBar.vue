<template>
  <div class="slider-component">
    <VueSlider
      :lazy="true"
      v-model="elapsedTimePositon"
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
    />
    <span class="elapsedTime">{{ playbackPosition }}</span>
  </div>
</template>
<script lang="ts">
import { format } from 'date-fns'
import VueSlider from 'vue-slider-component'
import Vue from 'vue'
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
  computed: {
    playbackPosition: function (): string {
      // prettier-ignore
      return `${format(this.elapsedTime, 'mm:ss')}/${format(this.totalTime,'mm:ss')}`
    },
    elapsedTimePositon: {
      get: function (): number {
        return this.elapsedTime
      },
      set: function (value: number): void {
        // 入れられても困るので何もしない
      },
    },
  },
  methods: {
    tooltipFormatter: function (value: number): string {
      return format(value, 'mm:ss')
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
