import api from "./api"
import { Category, CategoryWithLists, NewCategory } from "../types/category"
import { IdResponse } from "../types/response"

class CategoryService {
    async get(): Promise<{ data: CategoryWithLists }> {
        try {
            const res = await api.get("/categories/lists/")
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async create(category: NewCategory): Promise<{ data: IdResponse }> {
        try {
            const res = await api.post("/categories/", category)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async update(category: Category): Promise<{ data: IdResponse }> {
        try {
            const res = await api.patch(`/category/${category.id}/`, category)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async remove(id: string) {
        try {
            const res = await api.delete(`/category/${id}`)
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }
}

export default new CategoryService()
