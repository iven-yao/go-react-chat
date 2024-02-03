export default interface User {
    ID: number|null,
    username: string,
    password: string
}

export interface LoginSucess {
    message: string
    token: string
    ID: number
    username: string
}