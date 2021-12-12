
import request from "./read";

export  function  getSvcList(ns){
    console.log(ns)
    return request.get("/v1/svc",{params:{ns}})
}