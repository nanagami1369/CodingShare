<template>
  <div class="slider">
    <div class="slider-value" :style="sliderValueStyle"></div>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'SliderBar',
  props: {
    value: {
      type: Number,
      default: 0,
    },
  },
  computed: {
    sliderValueStyle: function (): { width: string } {
      return {
        width: `${this.value}%`,
      }
    },
  },
})
</script>

<style scoped>
.slider {
  height: 15px;
  background-color: #858585;
}

.slider-value {
  height: 15px;
  background-color: #28e270;
}
</style>
