import React, { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import ChatService from "../services/ChatService";
import {ChatSendOut, ChatsRecv} from "../types/Chat";
const cookie  = new Cookies()

function Main() {

    const [message, setMessage] = useState<string>("");
    const [chatMap, setChatMap] = useState<Map<number, ChatsRecv>>(new Map());

    const logout = () => {
        cookie.remove("token", {path:"/"});
        window.location.href = "/login";
    }

    const updateChatMap = (key:number, value: ChatsRecv) => {
        setChatMap(map => new Map(map.set(key, value)));
    }

    useEffect(() => {
        // const intervalId = setInterval(() => {
            
        // }, 5000)

        // return () => clearInterval(intervalId);

        ChatService.getAll(1)
            .then((res) => {
                res.data.map((chat:any) => {
                    var c_recv:ChatsRecv = {
                        id : chat.ID,
                        message : chat.message,
                        username : chat.username,
                        votes : chat.upvotes - chat.downvotes,
                        created_at : chat.CreatedAt
                    }

                    updateChatMap(c_recv.id, c_recv);
                    
                })
                
                // console.log(chatMap)
            })
            .catch((err) => {
                console.log(err)
            });
        
    },[])

    const send = () => {
        if(message.length != 0) { 
            var chat : ChatSendOut = {message: message} 
            setMessage("")

            ChatService.send(chat)
                .then((res) => {
                    console.log(res.data)
                })
                .catch((err) => {
                    console.log(err)
                })
            

        }
    }

    return (
        <div className="App h-screen grid grid-cols-4">
            <div className="bg-blue-500 h-full">
                <button type="button" className="rounded-md border-2 p-2 border-black" onClick={logout}>LOGOUT</button>
            </div>
            <div className="col-span-3">
                <div className="h-3/4 bg-black text-white p-2">
                    {
                        Array.from(chatMap).map((chat) => (
                            <p key={chat[0]} className="text-left">{chat[1].username}:{chat[1].message}</p>
                        ))
                    }
                </div>
                <div className="h-1/4">
                    <textarea name="message" className="h-full w-full bg-green-400 text p-2" value={message} onChange={(e) => setMessage(e.target.value)} />
                    <button className="absolute bottom-5 right-5 border rounded-lg p-2 bg-pink-200" onClick={send}>SEND</button>
                </div>
            </div>
        </div>
    )
}

export default Main;