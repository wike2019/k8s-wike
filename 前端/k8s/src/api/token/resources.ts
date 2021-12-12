
import request from "./read";

export  function  getRoleResources(){
    return request.get("/v1/resources")
}