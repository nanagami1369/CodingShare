import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueJsModal from 'vue-js-modal'

Vue.config.productionTip = false

Vue.use(VueJsModal)

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app')
