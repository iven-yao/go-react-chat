import http from "../http-common";
import User from "../types/User";

const register = (data: User) => {
    return http.post<User>("/user/register", data);
}

const login = (data: User) => {
    return http.post<User>("/user/login", data);
}

const UserServices = {
    register,
    login
}

export default UserServices;