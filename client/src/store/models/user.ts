import { createModel } from "@rematch/core"
import { AxiosResponse } from "axios"
import jwt, { JwtPayload } from "jsonwebtoken"
import api from "../../http"
import { RootModel } from "./index"

export interface User {
    userId: number
    login: string
    role: string
    token: string
    expiresAt: number
    issuedAt: number
}

type UserState = {
    error: null | string
    token: {
        accessToken: string
        expiresAt: number | undefined
        issuedAt: number | undefined
    }
    role: string
    userId: string
    login: string
    isAuth: boolean
}

export const user = createModel<RootModel>()({
    state: {
        error: null,
        token: {},
        role: "",
        userId: "",
        login: "",
        isAuth: false,
    } as UserState,

    reducers: {
        setError(state, payload: string | null) {
            state.error = payload
            return state
        },
        setUser(state, payload: { accessToken: string }) {
            const token = jwt.decode(payload.accessToken) as JwtPayload
            if (token) {
                state.token.accessToken = payload.accessToken
                state.token.expiresAt = token.exp
                state.token.issuedAt = token.iat
                state.role = token.role
                state.login = token.login
                state.userId = token.userId
                state.isAuth = !!payload.accessToken
            }
            return state
        },
        clearUser(state, payload) {
            state.token.accessToken = ""
            state.token.expiresAt = undefined
            state.token.issuedAt = undefined
            state.role = ""
            state.login = ""
            state.userId = ""
            state.isAuth = false

            return state
        },
    },

    effects: dispatch => {
        const { user, app } = dispatch
        return {
            async login(payload: { login: string; password: string }) {
                user.setError(null)
                app.setLoading(true)
                try {
                    const { login, password } = payload
                    const res: AxiosResponse<{ accessToken: string }> = await api.post(
                        "/auth/sign-in",
                        {
                            login,
                            password,
                        }
                    )

                    user.setUser(res.data)
                } catch (error: any) {
                    console.log(error.response.data.message)
                    user.setError(error.response.data.message)
                } finally {
                    app.setLoading(false)
                }
            },
            async logout() {
                user.setError(null)
                app.setLoading(true)
                try {
                    await api.post("/auth/sign-out")
                    user.clearUser()
                } catch (error: any) {
                    console.log(error.response.data.message)
                    user.setError(error.response.data.message)
                } finally {
                    app.setLoading(false)
                }
            },
            async refresh() {
                user.setError(null)
                app.setLoading(true)
                try {
                    const res = await api.get("/auth/refresh")
                    user.setUser(res.data)
                } catch (error: any) {
                    user.clearUser()
                    if (error.response.status !== 403) {
                        user.setError(error.response.data.message)
                    }
                } finally {
                    app.setLoading(false)
                }
            },
        }
    },
})
