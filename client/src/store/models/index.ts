import { Models } from "@rematch/core"
import { todo } from "./todo"
import { user } from "./user"

export interface RootModel extends Models<RootModel> {
    user: typeof user
    todo: typeof todo
}

export const models: RootModel = {
    user,
    todo,
}
