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
  {
    path: '/processes',
    name: 'EndedProcesses',
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/EndedProcesses.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
