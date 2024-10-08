import http from "../http-common";
import User, { LoginSucess } from "../types/User";

const register = (data: User) => {
    return http.post<User>("/user/register", data);
}

const login = (data: User) => {
    return http.post<LoginSucess>("/user/login", data);
}

const test = () => {
    return http.get("/user/test");
}

const UserServices = {
    register,
    login,
    test
}

export default UserServices;