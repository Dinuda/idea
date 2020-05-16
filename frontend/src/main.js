import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import hero from './components/hero.vue'
import sign_config from './components/sign_config.vue'
import register_student from './components/register_student.vue'
import register_investor from './components/investor_register.vue'
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
    path: '/register_student',
    name: 'register_student',
    component: register_student,
  } ,
  {
    path: '/register_investor',
    name: 'register_investor',
    component: register_investor,
  } ,
  {
    path: '/sign_config',
    name: 'sign_config',
    component: sign_config,
  } ,
]

const router = new VueRouter({
  routes // short for `routes: routes`
})


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

