export interface ChatsRecv {
    type: string,
    ID: number,
    message: string,
    username: string,
    CreatedAt: string,
    upvotes: number,
    downvotes: number
}

export interface WebSocketMessage {
    type: string,
    message: string, 
    token: string,
    vote: number,
    chatid: number
}