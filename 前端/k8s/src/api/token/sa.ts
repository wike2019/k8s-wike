
import request from "./read";

export  function  getSaList(ns){
    console.log(ns)
    return request.get("/v1/sa",{params:{ns}})
}

export  function  SaDel(ns,name){
    return request.delete("/v1/sa",{params:{ns,name}})
}
export  function  SACreate(data){
    console.log(data)
    return request.post("/v1/sa",data)
}
