import {createApp} from 'vue'
import {createPinia} from "pinia"

import router from './router'
import App from './App.vue'

import "normalize.css"
import "./styles/transitions.pcss"
import "./styles/edit_page.pcss"

createApp(App)
	.use(createPinia())
	.use(router)
	.mount('#app')