import { createRouter, createWebHistory } from 'vue-router'
import MotorView from '../views/MotorView.vue'

const routes = [
  {
    path: '/',
    name: 'MotorView',
    component: MotorView,
  },

  {
    path: '/process',
    name: 'Process',
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/Process.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
