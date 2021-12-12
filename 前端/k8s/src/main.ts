import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';


//加载根组件
import App from './App.vue'
//加载必要样式
import './style/index.less'
const app = createApp(App)
//引入router

// @ts-ignore
import router from './router/index.ts'

app.use(ElementPlus)

app.use(router)

app.mount('#root')
