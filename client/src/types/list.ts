export type List = {
    id: string
    categoryId?: string
    title: string
    description: string
    createdAt: number
}

export type NewList = {
    categoryId: string
    title: string
    description: string
}

export type UpdateList = {
    id: string
    categoryId?: string
    title?: string
    description?: string
}
