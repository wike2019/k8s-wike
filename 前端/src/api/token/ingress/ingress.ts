import request from "../read";
export  function  ingressListByNs(ns:string,page:number){
    return request.get("/v1/ingress",{params:{ns,current:page}})
}

export  function  ingressCreate(data){

    return request.post("/v1/ingress",data)
}

export  function  ingressDel(ns:string,name:string){
    return request.delete("/v1/ingress",{params:{ns,name}})
}

export  function  ingressItem(ns:string,name:string){
    return request.get("/v1/ingress/item",{params:{ns,name}})
}

export  function  ingressUpdate(data){

    return request.put("/v1/ingress",data)
}