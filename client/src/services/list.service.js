import api from './api'

class ListService {
    async getLists() {
        try {
            const res = await api.get('/list/')
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async getListById(id) {
        try {
            const res = await api.get(`/list/${id}/`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async createList(list) {
        try {
            const res = await api.post('/list/', list)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async updateList(id, list) {
        try {
            const res = await api.patch(`/list/${id}/`, list)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async removeList(id) {
        try {
            const res = await api.delete(`/list/${id}/`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }
}

export default new ListService()
