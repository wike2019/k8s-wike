//公共请求
import {API} from "../../../config/api";
import request from "../read";

//yaml文件提示用
export  function  getAutoComplete(){
    return request.get(API+"/v1/autoComplete")
}

//获取API资源信息
export  function  getResources(){
    return request.get("/v1/resources")
}