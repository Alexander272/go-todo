import { List, NewList, UpdateList } from "../types/list"
import { IdResponse } from "../types/response"
import api from "./api"

export default class ListService {
    static async getById(id: string): Promise<{ data: List }> {
        try {
            const res = await api.get(`/lists/${id}`)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async create(list: NewList): Promise<{ data: IdResponse }> {
        try {
            const res = await api.post("/lists/", list)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async update(list: UpdateList): Promise<{ data: IdResponse }> {
        try {
            const res = await api.patch(`/lists/${list.id}`, list)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async remove(id: string) {
        try {
            const res = await api.delete(`/lists/${id}`)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }
}
