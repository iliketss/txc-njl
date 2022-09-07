import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    redirect:"/index",
    component: () => import(/* webpackChunkName: "about" */ '../views/HomeView.vue'),
    children:[
      {
        path: '/index',
        name: 'index',
        meta:{
          title:"首页",
          isShow:true
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/Index.vue')

      },
      {
        path: '/useroder',
        name: 'userOrder',
        meta:{
          title:"我的订单",
          isShow:true
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/userOrder.vue')

      }
    ]
  },
  {
    path: '/about',
    name: 'about',

    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "login" */ '../views/LoginView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
