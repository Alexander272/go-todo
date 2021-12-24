import { Token } from "../types/response"
import { ISignIn, ISignUp } from "../types/user"
import api from "./api"

export default class AuthService {
    static async signIn(data: ISignIn): Promise<{ data: Token }> {
        try {
            const res = await api.post("/auth/sign-in/", data)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async signUp(data: ISignUp) {
        try {
            const res = await api.post("/auth/sign-up/", data)
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async refresh(): Promise<{ data: Token }> {
        try {
            const res = await api.post("/auth/refresh/")
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }

    static async signOut() {
        try {
            const res = await api.post("/auth/sign-out/")
            return res.data
        } catch (error: any) {
            throw error.response.data
        }
    }
}
