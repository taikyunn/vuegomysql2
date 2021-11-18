import { createApp } from 'vue'
import App from './App.vue'
import Axios from 'axios'
Axios.defaults.baseURL = process.env.VUE_APP_API_ENDPOINT;

createApp(App).mount('#app')
