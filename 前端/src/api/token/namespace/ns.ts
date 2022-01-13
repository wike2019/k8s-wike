
import request from "../read";

//获取namespace列表 无分页
export  function  getNsList(){
    return request.get("/v1/ns")
}
//创建namespace
export  function  createNs(data){
    return request.post("/v1/ns",data)
}
//删除namespace
export  function  deleteNs(name){
    return request.delete("/v1/ns",{params:{name}})
}