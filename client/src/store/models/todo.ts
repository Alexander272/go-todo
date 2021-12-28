import { createModel } from "@rematch/core"
import { List } from "../../types/list"
import { RootModel } from "./index"

export interface ITodo {}

type TodoState = {
    loading: boolean
    selectList: List | null
}

export const todo = createModel<RootModel>()({
    state: {
        loading: false,
        selectList: null,
    } as TodoState,

    reducers: {
        setLoading(state, payload: boolean) {
            state.loading = payload
            return state
        },
        setSelectList(state, payload: List) {
            state.selectList = payload
            return state
        },
    },
})
