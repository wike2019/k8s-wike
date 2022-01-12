import request from "../read";
export  function  configmapListByNs(ns:string,page:number){
    return request.get("/v1/ConfigMap",{params:{ns,current:page}})
}

export  function  configmapUpdate(data){
    return request.put("/v1/ConfigMap",data)
}

export  function configmapDel(ns:string,name:string){
    return request.delete("/v1/ConfigMap",{params:{ns,name}})
}
export  function configmapDetail(ns:string,name:string){
    return request.get("/v1/GetConfigMap",{params:{ns,name}})
}

export  function  configmapCreate(data){
    return request.post("/v1/ConfigMap",data)
}

export  function  configmapAllByNs(ns:string){
    return request.get("/v1/ConfigMap/all",{params:{ns}})
}
