import React, { useState } from "react";
import LoginForm from "./LoginFrom";
import SignUpForm from "./SignUpForm";

const activeTab = "p-2 bg-black text-white rounded-t-xl border-black";
const inactiveTab = "p-2 rounded-t-xl bg-gray-300  border-b-0 text-gray-600";
function Login() {

    const [isLogin, setIsLogin] = useState(true);
    
    const showLoginTab = () => setIsLogin(true);
    const showSignUpTab = () => setIsLogin(false);

    return (
        <div className="App h-screen flex flex-col justify-center items-center">
            <div className="w-3/4 md:w-1/2 h-1/2">
                <div className=" text-4xl font-bold">CHATROOM</div>
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