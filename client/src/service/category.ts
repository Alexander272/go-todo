import api from "./api"
import { Category, CategoryWithLists, NewCategory } from "../types/category"
import { IdResponse } from "../types/response"
import axios from "axios"

export default class CategoryService {
    static async get(): Promise<{ data: CategoryWithLists[] }> {
        try {
            const res = await api.get("/categories/lists/")
            return res.data
        } catch (error: any) {
            if (axios.isAxiosError(error)) {
                throw error.response?.data
            }
            throw error?.message
        }
    }

    static async create(category: NewCategory): Promise<{ data: IdResponse }> {
        try {
            const res = await api.post("/categories/", category)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async update(category: Category): Promise<{ data: IdResponse }> {
        try {
            const res = await api.patch(`/category/${category.id}/`, category)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async remove(id: string) {
        try {
            const res = await api.delete(`/category/${id}`)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }
}
