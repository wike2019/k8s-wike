import axios from "axios";
import {doTo} from "../../router";
import {API} from "../../config/api";
const instance = axios.create({
    baseURL:API
})
// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    let token=localStorage.getItem("token");
    if(!token){
        console.log("token不存在")
        doTo("login")
        return
    }
    config.headers.token=token
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    if (response.data.code==401){
        doTo("login")
    }
    return response;
}, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
});

export default  instance