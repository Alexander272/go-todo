export type IdResponse = {
    id: string
    message: string
}

export type Token = {
    token: TokenData
    userId: string
    name: string
    role: string
}

type TokenData = {
    accessToken: string
    exp: number
}
