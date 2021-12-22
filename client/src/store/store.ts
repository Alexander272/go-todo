import { init, RematchDispatch, RematchRootState } from "@rematch/core"
import immerPlugin from "@rematch/immer"
import { models, RootModel } from "./models"

export const store = init<RootModel>({
    models,
    plugins: [immerPlugin()],
    redux: {
        devtoolOptions: {
            actionSanitizer: action => action,
            disabled: false,
        },
    },
})

export type Store = typeof store
export type Dispatch = RematchDispatch<RootModel>
export type RootState = RematchRootState<RootModel>
