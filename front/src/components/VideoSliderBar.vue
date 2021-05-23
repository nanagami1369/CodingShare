<template>
  <div class="slider-component">
    <div class="slider">
      <div class="slider-value" :style="sliderValueStyle"></div>
    </div>
    <span class="elapsedTime">{{ playbackPosition }}</span>
  </div>
</template>
<script lang="ts">
import { format } from 'date-fns'
import Vue from 'vue'
export default Vue.extend({
  name: 'VideoSliderBar',
  props: {
    elapsedTime: {
      type: Number,
      default: 0,
    },
    totalTime: {
      type: Number,
      default: 0,
    },
  },
  computed: {
    sliderValueStyle: function (): { width: string } {
      if (this.totalTime == 0) {
        return { width: '0%' }
      }
      return {
        width: `${(this.elapsedTime / this.totalTime) * 100}%`,
      }
    },
    playbackPosition: function (): string {
      // prettier-ignore
      return `${format(this.elapsedTime, 'mm:ss')}/${format(this.totalTime,'mm:ss')}`
    },
  },
})
</script>

<style scoped>
.slider-component {
  display: flex;
  justify-content: space-between;
}
.slider {
  height: 15px;
  width: calc(100% - 6em);
  background-color: #858585;
}

.slider-value {
  height: 15px;
  background-color: #28e270;
}
</style>
