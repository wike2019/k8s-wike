
import mitt from 'mitt'
import {DWorker} from "./work";

import {cache} from "./cache";
import {fileDownLoad, isInArr, parseData} from "./helper";

const bus = mitt()
function timeFormat(str){
   return  str.replace("T"," ").replace("Z",'')
}
let core={
    bus,
    DWorker,
    mitt,
    cache,
    timeFormat,
    fileDownLoad,
    isInArr,
    parseData
}


export {core}