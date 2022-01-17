
import request from "../read";

//获取namespace列表 无分页
export  function  getEventList(ns,name,kind,uid){
    return request.get("/v1/event",{params:{ns,name,kind,uid}})
}
