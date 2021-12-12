
import request from "./read";

export  function  getList(){
    return request.get("/v1/nslist")
}