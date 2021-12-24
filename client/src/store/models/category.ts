import { createModel } from "@rematch/core"
import { toast } from "react-toastify"
import { RootModel } from "."
import CategoryService from "../../service/category"
import ListService from "../../service/list"
import { Category, CategoryWithLists, NewCategory } from "../../types/category"
import { List, NewList, UpdateList } from "../../types/list"

interface ICategoryState {
    loading: boolean
    listLoad: boolean
    categories: CategoryWithLists[]
    selectedCategories: string[]
}

export const category = createModel<RootModel>()({
    state: {
        loading: false,
        listLoad: false,
        categories: [
            {
                id: "123",
                title: "first category",
                lists: [
                    {
                        id: "i1",
                        title: "first list",
                        description: "mock list",
                        createdAt: 1640342763,
                    },
                    {
                        id: "i2",
                        title: "seconde list",
                        description: "mock list",
                        createdAt: 1640342763,
                    },
                    {
                        id: "i3",
                        title: "third list",
                        description: "mock list",
                        createdAt: 1640342763,
                    },
                ],
            },
            {
                id: "234",
                title: "second category",
                lists: [
                    {
                        id: "i1",
                        title: "first list",
                        description: "mock list",
                        createdAt: 1640342763,
                    },
                ],
            },
            { id: "345", title: "third category", lists: [] },
        ],
        selectedCategories: ["123"],
    } as ICategoryState,

    reducers: {
        setLoading(state, payload: boolean) {
            state.loading = payload
            return state
        },
        setListLoading(state, payload: boolean) {
            state.listLoad = payload
            return state
        },
        selectCategory(state, payload: string) {
            state.selectedCategories.push(payload)
            return state
        },
        unselectCategory(state, payload: string) {
            state.selectedCategories = state.selectedCategories.filter(catId => catId !== payload)
            return state
        },
        setCategory(state, payload: CategoryWithLists[]) {
            state.categories = payload
            return state
        },
        addCategory(state, payload: Category) {
            state.categories.push({ ...payload, lists: [] })
            return state
        },
        upgradeCategory(state, payload: Category) {
            state.categories = state.categories.map(cat => {
                if (cat.id === payload.id) cat = { ...cat, ...payload }
                return cat
            })
            return state
        },
        deleteCategory(state, payload: string) {
            state.categories = state.categories.filter(cat => cat.id !== payload)
            return state
        },

        addList(state, payload: List) {
            let index = state.categories.findIndex(cat => cat.id === payload.categoryId)
            if (index > -1) {
                state.categories[index].lists.push(payload)
            }
            return state
        },
        upgradeList(state, payload: UpdateList) {
            let index = state.categories.findIndex(cat => cat.id === payload.categoryId)
            if (index > -1) {
                state.categories[index].lists = state.categories[index].lists.map(list => {
                    if (list.id === payload.id) list = { ...list, ...payload }
                    return list
                })
            }
            return state
        },
        deleteList(state, payload: { categoryId: string; listId: string }) {
            let index = state.categories.findIndex(cat => cat.id === payload.categoryId)
            if (index > -1) {
                state.categories[index].lists = state.categories[index].lists.filter(
                    list => list.id !== payload.listId
                )
            }
            return state
        },
    },

    effects: dispatch => {
        const { category } = dispatch
        return {
            async getCategories() {
                category.setLoading(true)
                try {
                    const res = await CategoryService.get()
                    category.setCategory(res.data)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setLoading(false)
                }
            },
            async newCategory(payload: NewCategory) {
                category.setLoading(true)
                try {
                    const res = await CategoryService.create(payload)
                    category.addCategory({ id: res.data.id, title: payload.title })
                    toast.success(res.data.message)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setLoading(false)
                }
            },
            async updateCategory(payload: Category) {
                category.setLoading(true)
                try {
                    const res = await CategoryService.update(payload)
                    category.upgradeCategory(payload)
                    toast.success(res.data.message)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setLoading(false)
                }
            },
            async removeCategory(categoryId: string) {
                category.setLoading(true)
                try {
                    await CategoryService.remove(categoryId)
                    category.deleteCategory(categoryId)
                    toast.success("Deleted")
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setLoading(false)
                }
            },

            async newList(payload: NewList) {
                category.setListLoading(true)
                try {
                    const res = await ListService.create(payload)
                    category.addList({
                        id: res.data.id,
                        ...payload,
                        createdAt: +(Date.now() / 1000).toFixed(0),
                    })
                    toast.success(res.data.message)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setListLoading(false)
                }
            },
            async updateList(payload: UpdateList) {
                category.setListLoading(true)
                try {
                    const res = await ListService.update(payload)
                    category.upgradeList(payload)
                    toast.success(res.data.message)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setListLoading(false)
                }
            },
            async removeList(payload: { categoryId: string; listId: string }) {
                category.setListLoading(true)
                try {
                    await ListService.remove(payload.listId)
                    category.deleteList(payload)
                    toast.success("Deleted")
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    category.setListLoading(false)
                }
            },
        }
    },
})
