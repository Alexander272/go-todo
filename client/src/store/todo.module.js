import { useToast } from 'vue-toastification'
import todoService from '../services/todo.service'

const toast = useToast()

export const todoModule = {
    namespaced: true,
    state: () => ({
        error: null,
        loading: false,
        ready: false,
        todos: [],
        curTodo: null,
    }),
    getters: {
        isEmptyList(state) {
            return state.todos.length === 0
        },
        getFilteredTodo(state) {},
    },
    mutations: {
        setError(state, payload) {
            state.error = payload
        },
        setLoading(state, payload) {
            state.loading = payload
        },
        setReady(state, payload) {
            state.ready = payload
        },
        setTodos(state, payload) {
            state.todos = payload
        },
        addTodo(state, payload) {
            state.todos = [payload, ...state.todos]
        },
        updateTodo(state, payload) {
            state.todos = state.todos.map(todo => {
                if (todo.id === payload.id) todo = payload
                return todo
            })
        },
        setDone(state, payload) {
            state.todos = state.todos.map(todo => {
                if (todo.id === payload.id) todo = { ...todo, done: payload.done }
                return todo
            })
        },
        removeTodo(state, payload) {
            state.todos = state.todos.filter(todo => todo.id !== payload)
        },
    },
    actions: {
        async getTodos({ commit }, listId) {
            try {
                commit('setReady', false)
                const data = await todoService.getTodos(listId)
                if (data) commit('setTodos', data)
                else commit('setTodos', [])
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setReady', true)
            }
        },
        async setTodoDone({ commit }, todo) {
            try {
                commit('setLoading', true)
                await todoService.updateTodo({ id: todo.id, done: todo.done })
                commit('setDone', todo)
            } catch (error) {
                toast.error(error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async createTodo({ commit }, todo) {
            try {
                commit('setLoading', true)
                const data = await todoService.createTodo(todo)
                toast.success('Задача создана')
                commit('addTodo', { id: data.id, ...todo })
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async updateTodo({ commit }, todo) {
            try {
                commit('setLoading', true)
                await todoService.updateTodo(todo)
                toast.success('Задача обновлена')
                commit('updateTodo', todo)
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async removeTodo({ commit }, id) {
            try {
                commit('setLoading', true)
                await todoService.removeTodo(id)
                toast.success('Задача удалена')
                commit('removeTodo', id)
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
    },
}
