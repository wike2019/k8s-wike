
import request from "../read";

//获取namespace列表 无分页
export  function  getNsList(page){
    return request.get("/v1/ns?page="+page)  //获取列表
}
//创建namespace
export  function  createNs(data){
    return request.post("/v1/ns",data)  //创建
}
//删除namespace
export  function  deleteNs(name){
    return request.delete("/v1/ns",{params:{name}}) //删除
}

//获取namespace列表 无分页
export  function  getNsALL(){
    return request.get("/v1/ns/all")  //获取列表
}