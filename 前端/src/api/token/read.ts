import axios from "axios";
import {API} from "../../config/api";
import {ElMessage} from "element-plus";
const instance = axios.create({
    baseURL:API
})
// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    let token=localStorage.getItem("token");
    if(!token){ //这里可以做拦截操作
        console.error("token不存在")
       // doTo("login")
       // return
    }
    config.headers.token=token //把token 加到头里 方便后端验证权限
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
      //  doTo("login")  //根据不同code做各种操作
    }
    return response;
}, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    ElMessage.error(error.response.data.error||"系统错误") //提示错误消息
    return Promise.reject(error.response.data.error||"系统错误");
});

export default  instance