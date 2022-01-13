import axios from "axios";
import {API} from "../config/api";

//发送邮件使用
export  function  sendMail(email:string){
    return axios.post(API+"/v1/email",{email})
}
//登录使用
export  function  Login(email:string,code:string){
    return axios.post(API+"/v1/login",{email,code})
}