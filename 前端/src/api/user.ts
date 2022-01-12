///v1/deployments
import axios from "axios";
import {API} from "../config/api";

export  function  sendMail(email:string){
    return axios.post(API+"/v1/email",{email})
}
export  function  Login(email:string,code:string){
    return axios.post(API+"/v1/login",{email,code})
}