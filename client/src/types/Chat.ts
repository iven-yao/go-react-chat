export interface ChatSendOut {
    message: string
}

export interface ChatsRecv {
    id: number,
    message: string,
    username: string,
    created_at: EpochTimeStamp,
    votes: number
}