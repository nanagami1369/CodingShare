import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import Home from '../views/HomePage.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/editor',
    name: 'Editor',
    component: () => import('../views/EditorPage.vue'),
  },
  {
    path: '/player',
    name: 'Player',
    component: () => import('../views/PlayerPage.vue'),
  },
]

const router = new VueRouter({
  mode: 'history',
  routes,
})

export default router
