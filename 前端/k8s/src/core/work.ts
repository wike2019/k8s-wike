

/**
 * 数据转换计算方法
 */
const formatRawData = () => {
    return "default task over";
}


/**
 * WWorker class
 */
class DWorker {
    private worker:any;
    private flagMapping:any;
    constructor(worker = formatRawData,BASE_DATASETS={
    }) {
        // 依赖的全局变量声明，如果 BASE_DATASETS 非字符串形式，可调用 JSON.stringify 等方法进行处理，保证变量的正常声明
        if(!BASE_DATASETS.domain){
            // @ts-ignore
            BASE_DATASETS.domain=window.location.protocol+"//"+window.location.host
        }
        BASE_DATASETS=JSON.stringify(BASE_DATASETS)
        const CONSTS = `const base = ${BASE_DATASETS};`;

        // 转换计算方法声明
        const formatFn = `const formatFn = ${worker.toString()};`;

        /**
         * 内部 onmessage 处理
         */
        const onMessageHandlerFn = `self.onmessage = ({ data: { data, flag, coverage } }) => {
      console.log('Message received from main script');
      const {method} = data;
      if (data.data && method === 'format') {
        self.postMessage({
          data: formatFn(data.data),
          flag
        });
      }
      console.log('Posting message back to main script');
    }`;

        /**
         * 返回结果
         * @param {*} param0
         */
        const handleResult = ({ data: { data, flag } }) => {
            const resolve = this.flagMapping[flag];

            if (resolve) {
                resolve(data);
                delete this.flagMapping[flag];
            }
        }

        const blob = new Blob([`(()=>{${CONSTS}${formatFn}${onMessageHandlerFn}})()`]);
        this.worker = new Worker(URL.createObjectURL(blob));
        this.worker.addEventListener('message', handleResult);

        this.flagMapping = {};
        // @ts-ignore
        URL.revokeObjectURL(blob);
    }

    /**
     * 动态调用
     */
    send(data:any) {
        const w = this.worker;
        const flag = new Date().toString();
        w.postMessage({
            data,
            flag,
        });

        return new Promise((res) => {
            this.flagMapping[flag] = res;
        })
    }

    close() {
        this.worker.terminate();
    }
    showCase(){
        console.log(`
let result={
 name:wike  //全局数据
}
const worker = new DWorker(function(data){
   //importScripts("https://test.com/test.js") //导入脚本命令
    //importScripts(base.domain+"/test.js") //导入脚本命令
   console.log(base.name)
   //do some thing
},result);

worker.send({
  method: 'format', 
  data: result
}).then((data) => {
  // 处理结果
})        
        `)
    }
}

export {DWorker};
