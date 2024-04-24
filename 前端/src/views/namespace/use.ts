import {reactive, ref} from "vue";

let dialogFormVisible=ref(false)
const formRef=ref(null) //refs引用
let form=reactive({
    metadata:{ name:""}
})
export  {dialogFormVisible,formRef,form}