import { createModel } from "@rematch/core"
import { RootModel } from "./index"

export interface Todo {
    
}

type TodoState = {
    error: null | string
    loading: false
    completed: number
    remaining: number
}

export const todo = createModel<RootModel>()({
    state: {
        error: null,
        loading: false,
        completed: 12,
        remaining: 34,
    } as TodoState,

    reducers: {
        
    },

    // effects: dispatch => {
    //     return {
            
    //         }
    //     }
    // },
})
