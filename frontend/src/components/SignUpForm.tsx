import React, { useState } from "react";
import { FieldValues, useForm } from 'react-hook-form';
import UserServices from "../services/UserService";
import User from "../types/User";

function SignUpForm({showLogin} : {showLogin:any}) {
    const {register, handleSubmit, formState: {errors}} = useForm()
    const [errMsg, setErrMsg] = useState<string>("")

    function signup(data: FieldValues) {
        var newUser : User = {ID:null, username: data.username, password: data.password}
        UserServices.register(newUser)
            .then((res) => {
                if(res.status === 200) {
                    alert("sign up success, please login with your account");
                    showLogin();
                }
            })
            .catch((err) => {
                if(err.response.status === 409) {
                    setErrMsg(err.response.data.message)
                } else {
                    setErrMsg("Bad Request")
                }
            })
        
    }

    return (
        <div className="shadow-lg shadow-stone-600 p-5 rounded-b-md bg-stone-400">
            <form onSubmit={handleSubmit((data) => signup(data))}>
                <div className="grid grid-cols-1 md:grid-cols-3 gap-2 mb-2">
                    <label className="flex justify-start md:py-2">Username</label>
                    <input type="text" placeholder="Enter Username" className="border-b border-stone-600 rounded-md p-2 md:col-span-2 text-black" {...register('username', {required: true})}>
                    </input>
                    <label className="flex justify-start md:py-2">Password</label>
                    <input type="password" placeholder="Enter Password" className="border-b border-stone-600 rounded-md p-2 md:col-span-2 text-black" {...register('password', {required: true})}>
                    </input>
                </div>
                <div className="text-red-600">
                    {Object.keys(errors).length > 0 && Object.keys(errors).map(
                        key => <div>Error: {key} is required</div>
                    )}
                    {errMsg.length > 0 && <div>Error: {errMsg}</div>}
                </div>
                <div className="flex justify-end">
                    <button className="rounded-md px-2 py-1 text-black hover:bg-stone-200  bg-stone-500 shadow-md" type="submit">Sign Up</button>
                </div>
            </form>
        </div>
    );
}

export default SignUpForm;