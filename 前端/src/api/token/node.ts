import request from "./read";
export  function  getNodes(){
    return request.get("/v1/nodes")
}

export  function  getNodeDetail(nodeName){
    return request.get("/v1/nodes/detail",{params:{node:nodeName}})
}
export  function  saveNode(data){
    return request.put("/v1/nodes",data)
}