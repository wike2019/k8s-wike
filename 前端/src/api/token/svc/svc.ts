
import request from "../read";

export  function  getSvcList(ns){
    return request.get("/v1/svc/all",{params:{ns}})
}
export  function  getSvcListByPage(ns,page=1){
    return request.get("/v1/svc",{params:{ns,page}})
}

export  function  removeSvc(ns,name){
    return request.get("/v1/svc",{params:{ns,name}})
}