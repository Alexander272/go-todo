import api from './api'

class TodoService {
    async getTodos(listId) {
        try {
            const res = await api.get(`/${listId}/todo/`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async getAllTodo() {
        try {
            const res = await api.get(`/todo/all/`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async getTodoById(id) {
        try {
            const res = await api.get(`/todo/${id}/`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async createTodo(todo) {
        try {
            const res = await api.post(`/todo/`, todo)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async updateTodo(id, todo) {
        try {
            const res = await api.put(`/todo/${id}/`, todo)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async removeTodo(id) {
        try {
            const res = await api.delete(`/todo/${id}/`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }
}

export default new TodoService()
