import { createModel } from "@rematch/core"
import { toast } from "react-toastify"
import { RootModel } from "."
import AuthService from "../../service/auth"
import { Token } from "../../types/response"
import { ISignIn, ISignUp } from "../../types/user"

interface IUserState {
    loading: boolean
    token: {
        accessToken: string
        expiresAt: number
    }
    userId: string
    nickname: string
    role: string
    isAuth: boolean
}

export const user = createModel<RootModel>()({
    state: {
        loading: false,
        token: {
            accessToken: "",
            expiresAt: 0,
        },
        role: "",
        userId: "",
        nickname: "",
        isAuth: false,
    } as IUserState,

    reducers: {
        setLoading(state, payload: boolean) {
            state.loading = payload
            return state
        },
        setUser(state, payload: Token) {
            state.token.accessToken = payload.token.accessToken
            state.token.expiresAt = payload.token.exp
            state.userId = payload.userId
            state.nickname = payload.name
            state.role = payload.role
            state.isAuth = true
            return state
        },
        clearUser(state) {
            state.token.accessToken = ""
            state.token.expiresAt = 0
            state.role = ""
            state.nickname = ""
            state.userId = ""
            state.isAuth = false
            return state
        },
    },

    effects: dispatch => {
        const { user } = dispatch
        return {
            async signIn(payload: ISignIn) {
                user.setLoading(true)
                try {
                    const res = await AuthService.signIn(payload)
                    user.setUser(res.data)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    user.setLoading(false)
                }
            },

            async singUp(payload: ISignUp) {
                user.setLoading(true)
                try {
                    await AuthService.signUp(payload)
                    toast.success("Registration completed successfully")
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    user.setLoading(false)
                }
            },

            async singOut() {
                user.setLoading(true)
                try {
                    await AuthService.signOut()
                    user.clearUser()
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    user.setLoading(false)
                }
            },

            async refresh() {
                user.setLoading(true)
                try {
                    const res = await AuthService.refresh()
                    user.setUser(res.data)
                } catch (error: any) {
                    console.log(error.message)
                } finally {
                    user.setLoading(false)
                }
            },
        }
    },
})
