import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersist from 'vuex-persist'

type Store = {
  userId: string
  speed: number
}

const vuexLocal = new VuexPersist<Store>({
  storage: window.localStorage,
  key: 'coding_share_vuex_key',
  reducer: (state) => ({ speed: state.speed }),
  filter: (mutation) => mutation.type === 'setSpeed',
})

Vue.use(Vuex)

export const store = new Vuex.Store<Store>({
  state: {
    userId: '',
    speed: 200,
  },
  getters: {
    userId: (state) => state.userId,
    isLogin: (state) => state.userId != '',
    speed: (state) => state.speed,
  },
  mutations: {
    setUserId(state, newUserId) {
      state.userId = newUserId
    },
    setSpeed(state, newSpeed) {
      state.speed = newSpeed
    },
  },
  actions: {
    setUserIdAction({ commit }, newUserId) {
      commit('setUserId', newUserId)
    },
    setSpeedAction({ commit }, newSpeed) {
      commit('setSpeed', newSpeed)
    },
  },
  plugins: [vuexLocal.plugin],
  modules: {},
})
