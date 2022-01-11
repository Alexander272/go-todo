import { createModel } from "@rematch/core"
import { List } from "../../types/list"
import { RootModel } from "./index"

export interface ITodo {}

type TodoState = {
    selectList: List | null
}

export const todo = createModel<RootModel>()({
    state: {
        selectList: null,
    } as TodoState,

    reducers: {
        setSelectList(state, payload: List) {
            state.selectList = payload
            return state
        },
    },
})
