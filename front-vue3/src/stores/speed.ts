import { defineStore } from 'pinia'

export const useSpeedStore = defineStore({
  id: 'speed',
  state: () => {
    return { speed: 200 }
  },
  actions: {
    setSpeed: function (newSpeed: number): void {
      this.speed = newSpeed
    },
  },
  persist: {
    enabled: true,
    strategies: [{ storage: localStorage, paths: ['speed'] }],
  },
})
