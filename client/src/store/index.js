import { createStore } from 'vuex'
import { authModule } from './auth.module'
import { listModule } from './list.module'
import { todoModule } from './todo.module'

export default createStore({
    state: {
        loading: false,
        ready: false,
        error: null,
    },
    mutations: {
        setLoading(state, payload) {
            state.loading = payload
        },
        setError(state, payload) {
            state.error = payload
        },
        setReady(state, payload) {
            state.ready = payload
        },
    },
    actions: {},
    modules: {
        auth: authModule,
        lists: listModule,
        todo: todoModule,
    },
})
