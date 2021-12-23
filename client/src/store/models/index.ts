import { Models } from "@rematch/core"
import { category } from "./category"
import { todo } from "./todo"
import { user } from "./user"

export interface RootModel extends Models<RootModel> {
    user: typeof user
    category: typeof category
    todo: typeof todo
}

export const models: RootModel = {
    user,
    category,
    todo,
}
