import http from "../http-common";
import {ChatSendOut} from "../types/Chat";
const send = (data: ChatSendOut) => {
    return http.post("/chat", data);
}

const getAll = (chatroomId: number) => {
    return http.get(`/chat/${chatroomId}`);
}

const ChatServices = {
    send,
    getAll
}

export default ChatServices;