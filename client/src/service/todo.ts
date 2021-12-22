import { IdResponse } from "../types/response"
import { NewTodo, Todo, UpdateTodo } from "../types/todo"
import api from "./api"

export default class TodoService {
    static async get(listId: string): Promise<{ data: Todo[] }> {
        try {
            const res = await api.get(`/${listId}/todos`)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async create(todo: NewTodo): Promise<{ data: IdResponse }> {
        try {
            const res = await api.post("/todos/", todo)
            return res.data
        } catch (error: any) {
            throw error.respoonse.data
        }
    }

    static async update(todo: UpdateTodo): Promise<{ data: IdResponse }> {
        try {
            const res = await api.patch(`/todos/${todo.id}`, todo)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async remove(id: string) {
        try {
            const res = await api.delete(`/todos/${id}`)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }
}
