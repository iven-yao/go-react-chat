import React, { useState } from "react";
import Cookies from "universal-cookie";
const cookie  = new Cookies()

function Main() {

    const logout = () => {
        cookie.remove("token", {path:"/"})
        window.location.href = "/login"
    }

    return (
        <div className="App h-screen flex flex-col justify-center items-center">
            <div className="w-3/4 md:w-1/2 h-1/2">
                <div className=" text-4xl font-bold">CHATROOM</div>
                <div className=" text-4xl font-bold">LOGIN SUCCESSFUL</div>
                <button type="button" className="rounded-md border-2 p-2 border-black" onClick={logout}>LOGOUT</button>
            </div>
        </div>
    )
}

export default Main;