import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueThreeSixty from 'vue-360'
import 'vue-360/dist/css/style.css'
Vue.config.productionTip = false
Vue.use(VueThreeSixty)
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

