import axios from 'axios'
import store from '../store'
import authService from './auth.service'

const api = axios.create({
    withCredentials: true,
    baseURL: '/api/v1',
})

api.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${store.getters['auth/getToken']}`
    return config
})

api.interceptors.response.use(
    config => {
        return config
    },
    async error => {
        const originalRequest = error.config

        if (error.response.status === 401 && error.config && !error.config._isRetry) {
            originalRequest._isRetry = true
            try {
                const data = await authService.refresh()
                // const res = await api.post(`/auth/refresh`)
                store.dispatch('auth/refresh', data)
                return api.request(originalRequest)
            } catch (e) {
                return Promise.reject(e)
            }
        }
        return Promise.reject(error)
    }
)

export default api
