import $http from "@/http/index";
interface loginData{
    name:string,
    password:string
}
export const login =(data:loginData)=>$http({
    url:"/login",
    method:"post",
    data
})