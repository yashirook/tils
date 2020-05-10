// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import firebase from 'firebase'

Vue.config.productionTip = false
let app

const config = {
  apiKey: 'AIzaSyCA0pJsBPeqGo6emqNzZDSPyvpZgP4IEas',
  authDomain: 'elegant-zodiac-264511.firebaseapp.com',
  databaseURL: 'https://elegant-zodiac-264511.firebaseio.com',
  projectId: 'elegant-zodiac-264511',
  storageBucket: 'elegant-zodiac-264511.appspot.com',
  messagingSenderId: '638824694275'
}
firebase.initializeApp(config)

firebase.auth().onAuthStateChanged(user => {
  /* eslint-disable no-new */
  if (!app) {
    new Vue({
      el: '#app',
      router,
      components: { App },
      template: '<App/>'
    })
  }
})
