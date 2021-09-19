import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import Home from '../views/HomePage.vue'
import NotFoundPage from '../views/NotFoundPage.vue'

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
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/LoginPage.vue'),
  },
  {
    path: '/mypage',
    name: 'MyPage',
    component: () => import('../views/MyPage.vue'),
    meta: { authRequire: true },
  },
  {
    path: '*',
    name: 'NotFound',
    component: NotFoundPage,
  },
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes,
})

router.beforeEach(async (to, from, next) => {
  const authRequire = to.matched.some((record) => record.meta.authRequire)
  if (!authRequire) {
    // ログイン不要なら進む
    next()
    return
  }
  const response = await fetch(
    `${process.env.VUE_APP_CODING_SHARE_API_URL}/api/islogin`,
    {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
    }
  )
  if (response.status == 200) {
    // 認証成功なら進む
    next()
    return
  }
  // 認証エラー
  console.log(response.status)
  if (response.status == 403) {
    // 認証エラーならログインページへ
    next('login')
    return
  }
  // それ以外の状態なら直接結果を表示
  alert(response.body)
})

export default router
