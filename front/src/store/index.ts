import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

export const store = new Vuex.Store({
  state: {
    userId: '',
  },
  getters: {
    userId: (state) => state.userId, // 2
  },
  mutations: {
    setUserId(state, newUserId) {
      state.userId = newUserId
    },
  },
  actions: {
    setUserIdAction({ commit }, newUserId) {
      commit('setUserId', newUserId)
    },
  },
  modules: {},
})
