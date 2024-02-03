import React, { useEffect, useRef, useState } from "react";
import Cookies from "universal-cookie";
import {ChatsRecv, WebSocketMessage} from "../types/Chat";
import useWebSocket from "react-use-websocket";
import {ArrowUturnLeftIcon, MinusCircleIcon, PauseCircleIcon, XCircleIcon} from "@heroicons/react/24/solid"
import ChatBubble from "./ChatBubble";
import ChatService from "../services/ChatService";

const cookie  = new Cookies()
function Main() {

    const [message, setMessage] = useState<string>("");
    const [chatMap, setChatMap] = useState<Map<number, ChatsRecv>>(new Map());

    const messageEndRef = useRef<HTMLDivElement>(null)
    const scrollToBottom = () => {
        if(messageEndRef.current) {
            messageEndRef.current.scrollIntoView({behavior: "smooth"});
        }
    }

    const {sendMessage: sendWsMessage, lastMessage} = useWebSocket('ws://localhost:8080/api/ws',{},true);
    const sendMessage = (type: string, text: string, vote: number, chatid: number): void => {
        const msg: WebSocketMessage = {
            type: type,
            token: cookie.get("token"),
            message: text,
            vote: vote,
            chatid: chatid
        }

        sendWsMessage(JSON.stringify(msg))
    }

    useEffect(() => {
        
        if (lastMessage !== null) {
            const recv: ChatsRecv = JSON.parse(lastMessage.data)
            updateChatMap(recv.ID, recv);

            if(recv.type === "NEWMESSAGE") {
                scrollToBottom()
            }
            
        }
    },[lastMessage])

    const logout = () => {
        cookie.remove("token", {path:"/"});
        cookie.remove("username", {path:"/"});
        window.location.href = "/login";
    }

    const updateChatMap = (key:number, value: ChatsRecv) => {
        setChatMap(map => new Map(map.set(key, value)));
    }

    useEffect(() => {
        // fetch all old messages
        ChatService.getAll()
            .then((res) => {
                // eslint-disable-next-line
                res.data.map((chat:any) => {
                    var recv:ChatsRecv = chat
                    updateChatMap(recv.ID, recv);
                    scrollToBottom()
                })
            })
            .catch((err) => {
                console.log(err)
            });
    },[])

    const upvote = (id:number) => {
        sendMessage("VOTE", "", 1, id)
    }

    const downvote = (id:number) => {
        sendMessage("VOTE", "", -1, id)
    }

    const send = () => {
        if(message.length !== 0) { 
            sendMessage("MESSAGE",message,0, 0)
            setMessage("")
        }
    }

    const onEnterPress = (e:React.KeyboardEvent<HTMLTextAreaElement>) => {
        if(e.key === "Enter" && e.shiftKey === false) {
            e.preventDefault();
            send();
        }
    }

    return (    
        <div className="h-screen flex flex-col">
            <div className="bg-white rounded-t-lg flex justify-start px-2">
                <button type="button" onClick={logout} className="">
                    <XCircleIcon className="h-6 w-6 text-red-500 hover:text-red-700"/>
                </button>
                <button type="button" disabled className="">
                    <MinusCircleIcon className="h-6 w-6 text-yellow-500"/>
                </button>
                <button type="button" disabled className="">
                    <PauseCircleIcon className="h-6 w-6 text-green-500"/>
                </button>
                <div className="mt-2 mx-2 flex items-end z-10">
                    <div className="bg-gray-200 border h-0.5 w-1 -skew-y-[45deg] translate-y-0.5"></div>
                    <div className="rounded-t-lg bg-gray-200 pb-2 px-3 shadow-gray-500 shadow-md font-bold">
                        nimble-chat
                    </div>
                    <div className="bg-gray-200 border h-0.5 w-1 skew-y-[45deg] translate-y-0.5 "></div>
                </div>
                <div className="mt-2 flex items-end -translate-x-4">
                    <div className="bg-gray-200 border h-0.5 w-1 -skew-y-[45deg] translate-y-0.5"></div>
                    <button className="font-extrabold text-slate-700 rounded-t-lg bg-gray-200 pb-2 px-3 shadow-gray-500 shadow-md hover:bg-gray-400">+</button>
                    <div className="bg-gray-200 border h-0.5 w-1 skew-y-[45deg] translate-y-0.5 "></div>
                </div>
            </div>
            <div className="flex-grow bg-black text-white p-2 overflow-y-scroll z-50">
                
                {Array.from(chatMap).map((chat) => <ChatBubble key={chat[0]} chat={chat[1]} upvote={upvote} downvote={downvote}/>)}
                <div className="h-20" />
                <div ref={messageEndRef} />
            </div>
            <div className="p-2">
                <textarea placeholder="Enter Messages..." name="message" className="h-full w-full bg-white text p-1" value={message} onChange={(e) => setMessage(e.target.value)} onKeyDown={onEnterPress}/>
                <button className="absolute bottom-5 right-5 border rounded-lg p-2 bg-gray-200 shadow-md shadow-gray-500" onClick={() => send()}>
                    <ArrowUturnLeftIcon className="w-5" />
                </button>
            </div>
        </div>
    )
}

export default Main;