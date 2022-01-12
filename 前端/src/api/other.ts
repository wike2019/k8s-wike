///v1/deployments
import axios from "axios";
import {API} from "../config/api";


export  function  autoComplete(){
    return axios.get(API+"/v1/autoComplete")
}