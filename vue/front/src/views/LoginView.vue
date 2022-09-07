<template>
  <div class="login-box">
    <el-form
        ref="loginFormRef"
        :model="loginForm"
        status-icon
        :rules="rules"
        label-width="70px"
        class="loginForm"
    >
      <el-form-item label="账号：" prop="name">
        <el-input v-model="loginForm.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="密码：" prop="password">
        <el-input
            v-model="loginForm.password"
            type="password"
            autocomplete="off"
        />
      </el-form-item>
      <el-form-item>
        <el-button class="sub-btn" type="primary" @click="submitForm()"
        >登录</el-button
        >
      </el-form-item>
    </el-form>

  </div>
</template>

<script lang="ts">

import { InitData } from "@/types/login";
import {defineComponent, reactive,toRefs} from "vue";
import {login} from "@/http/api";
import {useRouter} from "vue-router";

export default defineComponent({
  setup: function () {
    const data = reactive(new InitData())
    const router = useRouter()
    const rules = {
      name: [
        {required: true, message: '请输入账号', trigger: 'blur'},
        {min: 2, max: 24, message: '账号长度大于3小于等于24', trigger: 'blur'},
      ],
      password: [
        {required: true, message: '请输入密码', trigger: 'blur'},
        {min: 3, max: 12, message: '密码长度3-12', trigger: 'blur'},
      ]
    }

    const submitForm = () => {
      data.loginFormRef?.validate((valid) => {
        if (valid) {
          login(data.loginForm).then(res => {
              console.log(res)
              /*localStorage.setItem("token", res?.token)
              localStorage.setItem("user",res.user)*/
              router.push("/")
          })
        }
      })
    }

    return {
      ...toRefs(data),
      rules,
      submitForm
    }
  }
})
</script>

<style scoped>
  .login-box{
    width: 100%;
    height: 100%;
    background: url("../assets/bg.jpeg");
    box-sizing: border-box;
    padding-top: 200px;

  }
  .loginForm{
    margin: 0 auto;
    width: 400px;
    padding: 20px;
    background: #ffff;
    border-radius: 20px;
  }
  .sub-btn{
    width: 100%;
  }
</style>