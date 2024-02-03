import http from "../http-common";
const getAll = () => {
    return http.get("/chat");
}
const ChatServices = {
    getAll
}

export default ChatServices;