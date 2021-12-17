
export  function getData(data,keys,obj){
    let arr=keys.split('.')
    for (let i =0;i<arr.length-1;i++){
        obj=obj[arr[i]]
    }
    obj[arr[arr.length-1]]=data
}