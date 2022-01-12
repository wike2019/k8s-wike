import request from "./read";
export  function  getList(ns:string,page:number){
    return request.get("/v1/deployments",{params:{ns,current:page}})
}
