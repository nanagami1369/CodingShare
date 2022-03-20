import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/editor',
      name: 'Editor',
      component: () => import('../views/EditorView.vue'),
    },
    {
      path: '/player',
      name: 'Player',
      component: () => import('../views/PlayerView.vue'),
    },
  ],
})

export default router
