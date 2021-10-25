import { createApp } from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import App from './App.vue'
import router from './router'
import store from './store'
import {
    faCaretDown,
    faCog,
    faListUl,
    faPlus,
    faSave,
    faTimes,
} from '@fortawesome/free-solid-svg-icons'

library.add(faPlus, faCog, faListUl, faCaretDown, faTimes, faSave)

createApp(App)
    .use(store)
    .use(router)
    .mount('#app')
