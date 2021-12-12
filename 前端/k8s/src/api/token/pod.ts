import request from "./read";
export  function  getPodsByNs(ns:string,page:number){
    return request.get("/v1/pods",{params:{ns,current:page}})
}

export  function  getPodsContainers(ns,name){
return request.get("/v1/pods/containers",{params:{ns,name:name}})
}


