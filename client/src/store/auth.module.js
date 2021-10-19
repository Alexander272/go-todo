import AuthService from '../services/auth.service'

export const authModule = {
    namespaced: true,
    state: () => ({
        error: null,
        message: null,
        token: {
            accessToken: '',
            expiresAt: undefined,
            issuedAt: undefined,
        },
        role: '',
        userId: '',
        name: '',
    }),
    getters: {
        isAuth(state) {
            return !!state.token.accessToken
        },
    },
    mutations: {
        setError(state, payload) {
            state.error = payload
        },
        setMessge(state, payload) {
            state.message = payload
        },
        setUser(state, payload) {
            state.token = payload.token
            state.role = payload.role
            state.userId = payload.userId
            state.name = payload.name
        },
        setToken(state, payload) {
            state.token = payload.token
        },
        clearUser(state) {
            state.token = {
                accessToken: '',
                expiresAt: undefined,
                issuedAt: undefined,
            }
            state.userId = ''
            state.name = ''
            state.role = ''
        },
    },
    actions: {
        async signIn({ commit }, user) {
            try {
                const data = await AuthService.signIn(user.email, user.password)
                const decode = AuthService.decodeToken(data.accessToken)
                if (!decode) {
                    commit('setError', 'authorization failed')
                } else {
                    commit('setUser', decode)
                }
            } catch (error) {
                commit('setError', error.message)
            }
        },
        async signUp({ commit }, user) {
            try {
                const data = await AuthService.signUp(user.login, user.email, user.password)
                if (data.status) {
                    commit('setMessage', '')
                } else {
                    commit('setError', 'registration failed')
                }
            } catch (error) {
                commit('setError', error.message)
            }
        },
        async signOut({ commit }) {
            try {
                await AuthService.signOut()
                commit('clearUser')
            } catch (error) {
                commit('setError', error.message)
            }
        },
        async refresh({ commit, getters }) {
            try {
                const data = await AuthService.refresh()
                const decode = AuthService.decodeToken(data.accessToken)
                if (!decode) commit('setError', 'refresh failed')
                if (getters.isAuth) {
                    commit('setToken', decode)
                } else {
                    commit('setUser', decode)
                }
            } catch (error) {
                commit('setError', error.message)
                commit('clearUser')
            }
        },
    },
}
