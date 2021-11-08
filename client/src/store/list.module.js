import { useToast } from 'vue-toastification'
import listService from '../services/list.service'

const toast = useToast()

export const listModule = {
    namespaced: true,
    state: () => ({
        error: null,
        message: null,
        loading: false,
        lists: [],
    }),
    getters: {
        isEmptyList(state) {
            return state.lists.length === 0
        },
        getListById: state => id => {
            return state.lists.find(list => list.id === id)
        },
    },
    mutations: {
        setError(state, payload) {
            state.error = payload
        },
        setMessage(state, payload) {
            state.message = payload
        },
        setLoading(state, payload) {
            state.loading = payload
        },
        setLists(state, payload) {
            state.lists = payload
        },
        addList(state, payload) {
            state.lists = [payload, ...state.lists]
        },
        updateList(state, payload) {
            state.lists = state.lists.map(list => {
                if (list.id === payload.id) list = payload
                return list
            })
        },
        removeItem(state, payload) {
            state.lists = state.lists.filter(list => list.id !== payload)
        },
    },
    actions: {
        async getLists({ commit }) {
            try {
                commit('setLoading', true)
                const data = await listService.getLists()
                if (data) commit('setLists', data)
                else commit('setLists', [])
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async createList({ commit }, list) {
            try {
                commit('setLoading', true)
                const data = await listService.createList(list)
                toast.success('Список задач создан')
                commit('addList', { id: data.id, ...list })
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async updateList({ commit }, list) {
            try {
                commit('setLoading', true)
                await listService.updateList(list.id, list)
                toast.success('Список задач обновлен')
                // commit('setMessage', 'Список задач обновлен')
                commit('updateList', list)
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async removeList({ commit }, id) {
            try {
                commit('setLoading', true)
                await listService.removeList(id)
                toast.success('Список задач удален')
                // commit('setMessage', 'Список задач удален')
                commit('removeItem', id)
            } catch (error) {
                toast.error(error.message)
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
    },
}
