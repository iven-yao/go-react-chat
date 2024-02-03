import React from "react";
import Cookies from "universal-cookie";
import {HandThumbDownIcon, HandThumbUpIcon} from "@heroicons/react/24/outline"
import { ChatsRecv } from "../types/Chat";
const cookie  = new Cookies()
function ChatBubble({chat, upvote, downvote} : {chat: ChatsRecv, upvote:(id:number) => void, downvote:(id: number) => void}) {

    return cookie.get("username") === chat.username ?
            <div className="p-1 m-1 flex text-xl justify-end">
                <div>
                    <p className="bg-blue-400 rounded-t-xl rounded-l-xl px-2 py-1 break-words max-w-lg">{chat.message}</p>
                    <div className="flex justify-end text-xs mt-1 text-gray-700">
                        <button disabled>
                            <HandThumbUpIcon className="h-4 w-4 mr-2"/>
                        </button>
                        {chat.upvotes - chat.downvotes}
                        <button disabled>
                            <HandThumbDownIcon className="h-4 w-4 ml-2"/>
                        </button>
                    </div>
                </div>
                <div className=" flex flex-col justify-end text-xs text-gray-500">{new Date(chat.CreatedAt).toLocaleTimeString([],{hour:'2-digit', minute:'2-digit'})}</div>
            </div>
            :
            <div className="p-1 m-1 flex text-xl flex-col">
                <div className="text-gray-500">{chat.username}: </div>
                <div className="flex">
                    <div>
                        <p className="bg-green-400 rounded-b-xl rounded-r-xl px-2 py-1 break-words max-w-lg ">{chat.message}</p>
                        <div className="flex justify-end text-xs mt-1">
                            <button onClick={() => upvote(chat.ID)}>
                                <HandThumbUpIcon className="h-4 w-4 mr-2" />
                            </button>
                            {chat.upvotes - chat.downvotes}
                            <button onClick={() => downvote(chat.ID)}>
                                <HandThumbDownIcon className="h-4 w-4 ml-2"/>
                            </button>
                        </div>    
                    </div>
                    <div className=" flex flex-col justify-end text-xs text-gray-500">{new Date(chat.CreatedAt).toLocaleTimeString([],{hour:'2-digit', minute:'2-digit'})}</div>
                </div>
            </div>
}

export default ChatBubble;