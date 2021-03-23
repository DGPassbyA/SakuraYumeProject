import Vue from 'vue'
import VueCookies from 'vue-cookies'
import App from './App.vue'
import router from "./routers"
import axios from "axios"
Vue.config.productionTip = false

Vue.use(VueCookies)
axios.defaults.crossDomain = true;
axios.defaults.withCredentials = true;
//test with cookie
axios.get("http://localhost:8081/getToken", {
  withCredentials: true,
}).then(response => (
  console.log(response)
))
//remember to delete
new Vue({
  render: h => h(App),
  router: router,
}).$mount('#app')