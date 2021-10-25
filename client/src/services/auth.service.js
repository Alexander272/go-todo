import api from './api'
import jwt from 'jsonwebtoken'

class AuthService {
    async signIn(email, password) {
        try {
            const res = await api.post('/auth/sign-in/', { email, password })
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async signUp(name, email, password) {
        try {
            const res = await api.post('/auth/sign-up/', { name, email, password })
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async refresh() {
        try {
            const res = await api.post('/auth/refresh/')
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    async signOut() {
        try {
            const res = await api.post('/auth/sign-out/')
            return res.data
        } catch (error) {
            throw error.response.data
        }
    }

    decodeToken(accessToken) {
        const token = jwt.decode(accessToken)
        if (token) {
            return {
                token: {
                    accessToken: accessToken,
                    expiresAt: token.exp,
                    issuedAt: token.iat,
                },
                role: token.role,
                name: token.name,
                userId: token.userId,
            }
        }
        return null
    }
}

export default new AuthService()
