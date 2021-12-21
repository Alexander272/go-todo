export type Todo = {
    id: string
    listId: string
    title: string
    description: string
    createdAt: number
    completedAt: number
    startAt: number
    done: boolean
    priority: number
}

export type NewTodo = {
    listId: string
    title: string
    description: string
    startAt: number
    priority: number
}

export type UpdateTodo = {
    id: string
    listId?: string
    title?: string
    description?: string
    startAt?: number
    priority?: number
}
