import { createApp } from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import App from './App.vue'
import router from './router'
import store from './store'
import {
    faCaretDown,
    faCog,
    faExclamationTriangle,
    faListUl,
    faPen,
    faPlus,
    faSave,
    faTimes,
    faTrash,
} from '@fortawesome/free-solid-svg-icons'

library.add(
    faPlus,
    faCog,
    faListUl,
    faCaretDown,
    faTimes,
    faSave,
    faPen,
    faTrash,
    faExclamationTriangle
)

createApp(App)
    .use(store)
    .use(router)
    .use(Toast)
    .mount('#app')
