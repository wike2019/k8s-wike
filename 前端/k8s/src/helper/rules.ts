export function requireRules(msg){
    return {required: true, trigger: 'change', message:msg }
}
function isInArr(arr,msg){
    return (rule, value, callback) => {
        if(arr.indexOf(value)==-1){
            callback(new Error(msg))
        }
        callback()
    }
}
export function inArrayWithMsg(arr,msg){
    return  { validator:isInArr( arr,msg), trigger: 'change' }
}
