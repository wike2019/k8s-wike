import request from "./read";

export  function  pvcAllByNs(ns:string){
    return request.get("/v1/pvc/all",{params:{ns}})
}
