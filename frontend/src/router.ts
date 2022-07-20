import { createRouter, createWebHashHistory } from 'vue-router'
import Layout from '@/layout/Layout.vue'
import NotFound from '@/layout/NotFound.vue'
import { HomeFilled } from '@element-plus/icons-vue'
import { IconWrap } from './components'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Layout,
    redirect: '/dashboard',
    meta: {
      icon: IconWrap(HomeFilled) 
    },
    children: [
      {
        path: '/dashboard',
        name: 'Dashboard',
        component: () => import('@/views/home/index.vue'),
      },
      {
        path: '/request',
        name: 'Request',
        component: () => import('@/views/request/index.vue'),
      },
      {
        path: '/json',
        name: 'JSON',
        component: () => import('@/views/json/index.vue'),
      },
      {
        path: '/plugin',
        name: 'Plugin',
        meta: {
          fullPage: true
        },
        component: () => import('@/views/tools/PluginPage.vue'),
      },
      {
        path: '/iframe',
        name: 'Iframe',
        meta: {
          fullPage: true
        },
        component: () => import('@/views/tools/IframePage.vue'),
      }
    ]
  },
  { path: '/:pathMatch(.*)*', component: NotFound }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
