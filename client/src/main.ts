import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import VueCookies from 'vue-cookies'

Vue.config.productionTip = false
Vue.use(VueCookies)
Vue.$cookies.config('60d')
// Vue.$cookies.config('60d','','',true)
// Vue.use(vuetify, {
//     theme: {
//         primary: "#f44336",
//         secondary: "#e57373",
//         accent: "#9c27b0",
//         error: "#f44336",
//         warning: "#ffeb3b",
//         info: "#2196f3",
//         success: "#4caf50"
//       }
// })

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
