
import request from "../read";

export  function  getNsList(){
    return request.get("/v1/ns")
}
export  function  createNs(data){
    return request.post("/v1/ns",data)
}
export  function  deleteNs(name){
    return request.delete("/v1/ns",{params:{name}})
}