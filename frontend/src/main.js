import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import hero from './components/hero.vue'
import register from './components/register.vue'


``
Vue.config.productionTip = false
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'hero',
    component: hero,
  } ,
  {
    path: '/register',
    name: 'register',
    component: register,
  } ,
]

const router = new VueRouter({
  routes // short for `routes: routes`
})


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

