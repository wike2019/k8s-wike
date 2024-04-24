import axios from "axios";
import {API} from "../../config/api";
import {ElMessage} from "element-plus";
const instance = axios.create({
    baseURL:API
})
// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
    return response;
}, function (error) {
    ElMessage.error(error.response.data.error||"系统错误") //提示错误消息
    return Promise.reject(error.response.data.error||"系统错误");
});

export default  instance