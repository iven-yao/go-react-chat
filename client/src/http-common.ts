import axios from "axios";
import Cookies from "universal-cookie";

export default axios.create({
    baseURL: "http://localhost:9090/api",
    headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + new Cookies().get("token")
    }
});