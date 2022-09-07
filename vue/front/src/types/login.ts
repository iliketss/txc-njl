import type {FormInstance}from 'element-plus';
import {ref} from "vue";

export interface LoginFormInt{
    name:string,
    password:string
}
export class InitData{
    loginForm:LoginFormInt={
        name:"",
        password:""
    }
    loginFormRef = ref<FormInstance>()
}
