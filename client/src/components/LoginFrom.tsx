import React, { useState } from "react";
import User from "../types/User";
import UserServices from "../services/UserService";
import { FieldValues, useForm } from "react-hook-form";

function LoginForm() {
    const {register, handleSubmit, formState: {errors}} = useForm()
    const [errMsg, setErrMsg] = useState<string>("")

    function login(data: FieldValues) {
        console.log(data)
        var user : User = {ID:null, username: data.username, password: data.password}
        UserServices.login(user)
            .then((res) => { 
                console.log(res.data);
                if(res.status === 200) {
                    alert("login success");
                    // todo
                }
            })
            .catch((err) => {
                console.error(err.response);
                if(err.response.status === 401) {
                    setErrMsg(err.response.data.message)
                }
            })
        
    }

    return (
        <div className="bg-black text-white p-5 rounded-b-md">
            <form onSubmit={handleSubmit((data) => login(data))}>
                <div className="grid grid-cols-1 md:grid-cols-3 gap-2 mb-2">
                    <label className="flex justify-start md:py-2">Username</label>
                    <input type="text" className=" rounded-md p-2 md:col-span-2 text-black" {...register('username', {required: true})}>
                    </input>
                    <label className="flex justify-start md:py-2">Password</label>
                    <input type="password" className="rounded-md p-2 md:col-span-2 text-black" {...register('password', {required: true})}>
                    </input>
                </div>
                <div className="text-red-600">
                    {Object.keys(errors).length > 0 && Object.keys(errors).map(
                        key => <div>Error: {key} is required</div>
                    )}
                    {errMsg.length > 0 && <div>Error: {errMsg}</div>}
                </div>
                <div className="flex justify-end">
                    <button className="bg-white rounded-md px-2 py-1 text-black hover:bg-blue-400" type="submit">Login</button>
                </div>
            </form>
        </div>
    );
}

export default LoginForm;