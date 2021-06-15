import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import firebase from 'firebase/app'

Vue.config.productionTip = false

const config = {
  apiKey: "AIzaSyCvazJ-xdQk6C8GJEKyUZ1j1zmNTxd9ZC0",
  authDomain: "nistagram-frontend-6db28.firebaseapp.com",
  projectId: "nistagram-frontend-6db28",
  storageBucket: "nistagram-frontend-6db28.appspot.com",
  messagingSenderId: "408901934217",
  appId: "1:408901934217:web:cbebaf9a82807242827a58",
  measurementId: "G-WNNYN4VH1J"
}
firebase.initializeApp(config)

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
