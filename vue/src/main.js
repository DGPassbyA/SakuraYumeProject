import Vue from 'vue'
import VueCookies from 'vue-cookies'
import App from './App.vue'
import router from "./routers"
import axios from "axios"
Vue.config.productionTip = false

Vue.use(VueCookies)
axios.defaults.crossDomain = true;
axios.defaults.withCredentials = true;
//remember to delete
new Vue({
  render: h => h(App),
  router: router,
}).$mount('#app')