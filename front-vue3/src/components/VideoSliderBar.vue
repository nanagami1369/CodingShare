<script setup lang="ts">
import { computed, ref, toRefs, watch } from 'vue'

const _position = ref(0)
const position = computed({
  get: (): number => _position.value,
  set: (value: number): void => {
    _position.value = value
  },
})

const props = defineProps({
  elapsedTime: {
    type: Number,
    default: 1,
  },
  totalTime: {
    type: Number,
    default: 1,
  },
  disabled: {
    type: Boolean,
    default: true,
  },
})

const { totalTime, elapsedTime } = toRefs(props)

watch(
  totalTime,
  (newTotalTime: number) => {
    position.value = (elapsedTime.value / newTotalTime) * 100
  },
  { immediate: true }
)

watch(
  elapsedTime,
  (newElapsedTime: number) => {
    position.value = (newElapsedTime / totalTime.value) * 100
  },
  { immediate: true }
)

const isDisabled = computed((): true | undefined =>
  props.disabled ? true : undefined
)

interface Emits {
  (e: 'seek', id: number): void
}

const emit = defineEmits<Emits>()

const positionStyle = computed(() => {
  return { width: position.value + '%' }
})

const bar = ref<HTMLDivElement>()
const mouseMove = (e: MouseEvent) => {
  if (bar?.value == null) {
    throw new Error('video slider bar Error:bar not found')
  }
  const { left, right } = bar.value.getBoundingClientRect()

  if (left > e.clientX) {
    position.value = 0
    emit('seek', Math.trunc((position.value * props.totalTime) / 100))
    return
  } else if (e.clientX >= right) {
    position.value = 100
    emit('seek', Math.trunc((position.value * props.totalTime) / 100))
    return
  } else {
    const width = bar.value.clientWidth
    const currentPosition = e.clientX - left
    position.value = (currentPosition / width) * 100
    emit('seek', Math.trunc((position.value * props.totalTime) / 100))
  }
}

const isHold = ref(false)
const mouseDown = (e: MouseEvent): void => {
  if (props.disabled) {
    return
  }
  isHold.value = true
  mouseMove(e)
  document.addEventListener('mouseup', mouseUp)
  document.addEventListener('mousemove', mouseMove)
}

const mouseUp = (): void => {
  if (props.disabled) {
    return
  }

  isHold.value = false
  document.removeEventListener('mouseup', mouseUp)
  document.removeEventListener('mousemove', mouseMove)
}
</script>

<template>
  <div :aria-disabled="isDisabled" :class="[isHold]" class="video-slider">
    <div ref="bar" class="bar" @mousedown="mouseDown">
      <div class="slider" :style="positionStyle"></div>
      <div class="pointer"></div>
    </div>
  </div>
</template>

<style scoped>
.video-slider {
  display: flex;
  background-color: #222;
  align-items: center;
  height: 20px;
  user-select: none;
}
.bar {
  cursor: pointer;
  background-color: #f0f0f0;
  height: 5px;
  display: flex;
  width: calc(100% - 20px);
  margin: 0px 5px;
  user-select: none;
  align-items: center;
}

.slider {
  height: 100%;
  width: 0px;
  background-color: aquamarine;
}

.video-slider[aria-disabled] .slider {
  background-color: rgb(192, 255, 234);
}

.video-slider:hover[aria-disabled] .bar {
  cursor: not-allowed;
}

.video-slider:hover:not([aria-disabled]) .bar {
  height: 10px;
  transition: 0.1s;
}

.isHold .bar {
  height: 10px;
  transition: 0.1s;
}

.pointer {
  width: 0px;
  height: 0px;
  background-color: #7fffd4;
  display: block;
}

.video-slider:hover[aria-disabled] .pointer {
  cursor: not-allowed;
}

.video-slider:hover:not([aria-disabled]) .pointer {
  width: 1px;
  height: 1px;
  transform: scale(20, 20);
  transition: 0.1s;
  border-radius: 10px;
}

.isHold .pointer {
  width: 1px;
  height: 1px;
  transform: scale(20, 20);
  transition: 0.1s;
  border-radius: 10px;
}
</style>
