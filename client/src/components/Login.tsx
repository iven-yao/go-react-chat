import React, { useState } from "react";
import LoginForm from "./LoginForm";
import SignUpForm from "./SignUpForm";

export const activeTab = "p-2 rounded-t-xl border-stone-600 bg-white z-10 bg-stone-400";
export const inactiveTab = "p-2 rounded-t-xl bg-stone-300  border-b-0 text-stone-100 border-stone-600 hover:text-black";

function Login() {

    const [isLogin, setIsLogin] = useState(true);
    
    const showLoginTab = () => setIsLogin(true);
    const showSignUpTab = () => setIsLogin(false);

    return (
        <div className="App h-screen flex flex-col justify-center items-center bg-gradient-to-t from-stone-500 to-sky-200">
            <div className="w-3/4 md:w-1/3 h-1/2">
                <div className=" text-6xl font-extrabold translate-y-2">nimble-chat</div>
                <div className="grid grid-cols-2 gap-0 cursor-pointer font-bold">
                    <div className={isLogin?activeTab:inactiveTab} onClick={showLoginTab}>Login</div>
                    <div className={isLogin?inactiveTab:activeTab} onClick={showSignUpTab}>Sign up</div>
                </div>
                {isLogin?
                    <LoginForm />
                    :
                    <SignUpForm showLogin={showLoginTab}/>
                }
                
            </div>
        </div>
    )
}

export default Login;