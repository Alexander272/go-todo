import { List } from "./list"

export type Category = {
    id: string
    title: string
}

export type NewCategory = {
    title: string
}

export type CategoryWithLists = {
    id: string
    title: string
    lists: List[]
}
