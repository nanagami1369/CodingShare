import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

export const store = new Vuex.Store({
  state: {
    userId: 'ログインしていません',
  },
  getters: {
    userId: (state) => state.userId, // 2
  },
  mutations: {
    setUserId(state, newUserId) {
      state.userId = newUserId
    },
    resetUserId(state) {
      state.userId = 'ログインしていません'
    },
  },
  actions: {
    setUserIdAction({ commit }, newUserId) {
      commit('setUserId', newUserId)
    },
    resetUserIdAction({ commit }) {
      commit('resetUserId')
    },
  },
  modules: {},
})
