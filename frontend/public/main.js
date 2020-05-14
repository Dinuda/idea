import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import vueScrollto from 'vue-scrollto'
import Toasted from 'vue-toasted';
 


import VueQuillEditor from 'vue-quill-editor'

``
Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(Toasted, {
  position: 'top-center',
  duration: 2000,
  keepOnHover: true,
  fullWidth: true,
  fitToScreen: true
})
Vue.use(vueScrollto)

const routes = [
  
]

const router = new VueRouter({
  routes // short for `routes: routes`
})


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

