import request from "../read";
export  function  secretListByNs(ns:string,page:number){
    return request.get("/v1/secret",{params:{ns,current:page}})
}

export  function  secretCreate(data){
    return request.post("/v1/secret",data)
}
export  function  secretUpdate(data){
    return request.post("/v1/secret/update",data)
}
export  function secretDel(ns,name){
    return request.delete("/v1/secret",{params:{ns,name}})
}
export  function secretDetail(ns,name){
    return request.get("/v1/GetSecret",{params:{ns,name}})
}

export  function  secretAllByNs(ns:string){
    return request.get("/v1/secret/all",{params:{ns}})
}
