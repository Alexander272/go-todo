import { createModel } from "@rematch/core"
import { RootModel } from "./index"

type AppState = {
    loading: boolean
    error: null | string
}

export const app = createModel<RootModel>()({
    state: {
        loading: false,
        error: null,
    } as AppState,

    reducers: {
        setLoading(state, payload: boolean) {
            state.loading = payload
            return state
        },
        setError(state, payload: string | null) {
            state.error = payload
            return state
        },
    },
})
