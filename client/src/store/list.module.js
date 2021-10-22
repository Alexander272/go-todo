import listService from '../services/list.service'

export const listModule = {
    namespaced: true,
    state: () => ({
        error: null,
        message: null,
        loading: false,
        lists: [],
    }),
    getters: {},
    mutations: {
        setError(state, payload) {
            state.error = payload
        },
        setMessge(state, payload) {
            state.message = payload
        },
        setLoading(state, payload) {
            state.message = payload
        },
        setLists(state, payload) {
            state.lists = payload
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
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async createList({ commit }, list) {
            try {
                commit('setLoading', true)
                await listService.createList(list)
                commit('setMessage', 'Список задач создан')
            } catch (error) {
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async updateList({ commit }, id, list) {
            try {
                commit('setLoading', true)
                await listService.updateList(id, list)
                commit('setMessage', 'Список задач обновлен')
            } catch (error) {
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
        async removeList({ commit }, id) {
            try {
                commit('setLoading', true)
                await listService.removeList(id)
                commit('setMessage', 'Список задач удален')
            } catch (error) {
                commit('setError', error.message)
            } finally {
                commit('setLoading', false)
            }
        },
    },
}
