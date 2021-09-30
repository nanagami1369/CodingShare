import { store } from '@/store'
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
    redirect: '/player/file',
  },
  {
    path: '/player/:id',
    name: 'Player',
    component: () => import('../views/PlayerPage.vue'),
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/LoginPage.vue'),
    meta: { notLogged: true },
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

const isLogin = (): boolean => {
  return store.getters.isLogin
}

router.beforeEach(async (to, from, next) => {
  const authRequire = to.matched.some((record) => record.meta.authRequire)
  const notLogged = to.matched.some((record) => record.meta.notLogged)
  // ログイン判定
  let response: Response
  try {
    response = await fetch('/api/auth', {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
    })
  } catch (error: unknown) {
    // 通信エラーの場合はアラートで表示
    alert((error as Error).message)
    return
  }
  // 認証成功ならユーザ情報をstoreに保存
  if (response.ok) {
    store.dispatch('setUserIdAction', (await response.json()).userId)
  } else if (!authRequire && !notLogged && response.status == 404) {
    // サーバー無しでフロントエンドのデバッグをするため
    // 認証情報が不要なページかつ404の場合はエラーを出しつつ先に進める
    console.error('通信エラー\n', await response.text())
    next()
    return
  } else {
    // それ以外の状態なら直接結果を表示して進めない
    alert(await response.text())
    return
  }

  if (authRequire) {
    if (isLogin()) {
      // 認証成功なら進む
      next()
      return
    } else {
      // 認証エラーならログインページへ
      next('login')
      return
    }
  }
  if (notLogged) {
    if (!isLogin()) {
      // // 認証エラーなら進む
      next()
      return
    } else {
      // 認証成功ならマイページへ
      next('mypage')
      return
    }
  }
  // authLogin notLogged 両方とも無い場合は先に進む
  next()
})

export default router
