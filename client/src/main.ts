import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from "axios";
import './assets/styles.css'

axios.defaults.baseURL = "http://localhost:3000/api/"
axios.defaults.withCredentials = true

createApp(App).use(store).use(router).mount('#app')
