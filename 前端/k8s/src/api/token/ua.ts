
import request from "./read";

export  function  getUaList(){
    return request.get("/v1/ua")
}

export  function  UaDel(user){
    return request.delete("/v1/ua",{params:{user:user}})
}
export  function  UaCreate(data){
    return request.post("/v1/ua",data)
}

export function genConfigFile(user,api) {
    return request.get( '/v1/clientConfig',{params:{user,api}} )
}
