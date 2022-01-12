export  class cache {
    set(key,data,express=1000*60*60){
        let dataContent={
            data:data,
            ctime:(new Date()).getTime(),//时间戳，同Date.now()
            express:express//默认设置过期时间一个小时
        }
        localStorage.setItem(key,JSON.stringify(dataContent));
    }
    get(key){
        let dataContent=JSON.parse(localStorage.getItem(key));
        let nowTime=(new Date()).getTime();
        if(nowTime-dataContent.ctime>=dataContent.express){
            localStorage.removeItem(key);
            return null;
        }else{
            return dataContent.data;
        }
    }
    remove(key){
        localStorage.removeItem(key);
    }
}