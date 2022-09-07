import axios from "axios";
import { ElMessage } from 'element-plus'
enum MSGS {
    "登录成功"=200,
    "密码错误",
    "账号错误",
    "请求异常"

}
const $http=axios.create({
    baseURL:'http://localhost:8080',
    timeout:2000,
    headers:{
        "Content-Type":"application/json;charset=utf-8"
    }
})

$http.interceptors.request.use(config=>{
    config.headers=config.headers||{}
    if (localStorage.getItem("token")){
        config.headers.token=localStorage.getItem("token")||''
    }
    return config
})

$http.interceptors.response.use(res=>{
    const code:number = res.data.code
    if (code!=200){
        ElMessage.error(MSGS[code])
        return Promise.reject(res.data)
    }else {
        ElMessage.success(MSGS[code])
        return res.data
    }

},error => {
    console.log(error)
})

export default $http