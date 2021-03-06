import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersist from 'vuex-persist'

type Store = {
  speed: number
}

const vuexLocal = new VuexPersist<Store>({
  storage: window.localStorage,
  key: 'coding_share_vuex_key',
  reducer: (state) => ({ speed: state.speed }),
  filter: (mutation) => mutation.type === 'setSpeed',
})

Vue.use(Vuex)

export default new Vuex.Store<Store>({
  state: {
    speed: 100,
  },
  getters: {
    speed: (state) => state.speed,
  },
  mutations: {
    setSpeed(state, newSpeed) {
      state.speed = newSpeed
    },
  },
  actions: {
    setSpeedAction({ commit }, newSpeed) {
      commit('setSpeed', newSpeed)
    },
  },
  modules: {},
  plugins: [vuexLocal.plugin],
})
