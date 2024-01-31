import http from "../http-common";
import User from "../types/User";

const getAll = () => {
    return http.get<Array<User>>("/user");
}

const get = (id: number) => {
    return http.get<User>(`/user/${id}`)
}

const create = (data: User) => {
    return http.post<User>("/user", data)
}

const update = (id: number, data: User) => {
    return http.put<User>(`/user/${id}`, data)
}

const remove = (id: number) => {
    return http.delete<any>(`/user/${id}`)
}

const UserServices = {
    getAll,
    get,
    create,
    update,
    remove
}

export default UserServices;