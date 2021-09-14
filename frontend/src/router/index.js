import Vue from 'vue'
import VueRouter from 'vue-router'
import EndedProcesses from '../views/EndedProcesses.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'EndedProcesses',
    component: EndedProcesses,
  },

  {
    path: '/process',
    name: 'Process',
    component: () =>
      import(/* webpackChunkName: "Process" */ '../views/Process.vue'),
  },
  {
    path: '/test-page',
    name: 'TestPage',
    component: () =>
      import(/* webpackChunkName: "test-page" */ '../views/TestPage.vue'),
  },
  {
    path: '/configuration',
    name: 'ConfigurationView',
    component: () =>
      import(/* webpackChunkName: "configuration" */ '../views/ConfigurationView.vue'),
  },
  {
    path: '/process-viewer/:processName',
    name: 'EndedProcessViewer',
    component: () =>
      import(/* webpackChunkName: "process-viewer" */ '../views/EndedProcessViewer.vue'),
    props:true
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
