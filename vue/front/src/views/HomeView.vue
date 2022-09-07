<template>
  <div class="common-layout">
    <el-container>
      <el-header height="80px">
        <el-row :gutter="20">
          <el-col :span="4"> <img class="logo" src="../assets/logo.png">
          </el-col>
          <el-col class="center" :span="16">

            <el-menu
                router
                :default-active='1'
                class="el-menu-demo"
                mode="horizontal"
                background-color="#545c64"
                text-color="#fff"
                active-text-color="#ffd04b"
            >
              <el-menu-item v-for="i in list" :key="i.path" :index="i.path">
                <span>{{i.meta.title}}</span>
              </el-menu-item>
            </el-menu>
          </el-col>
          <el-col class="btn-text" :span="4"><el-button type="text">jkl</el-button></el-col>
        </el-row>
        <div>

        </div>
      </el-header>
      <el-main>
        <router-view/>
      </el-main>
      <el-footer>Footer</el-footer>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
  import {useRouter} from "vue-router";

  const router =useRouter()
  const list=router.getRoutes().filter(v=>v.meta.isShow)
  console.log(list)

</script>


<style lang="scss" scoped>
  .logo{
    height: 80px;
  }
  h2{
    text-align: center;
    color: #ffff;
    height: 80px;
    line-height: 80px;
  }
  .btn-text{
    text-align: right;
    height: 80px;
    line-height: 80px;
  }
  .center{
    height: 80px;
    line-height: 80px;
  }
  .el-menu-demo{
    text-align: center;
    height: 80px;
    line-height: 80px;
  }
  .el-header{
    background: #545c64;
  }
  .el-main{
    height: calc(100vh - 160px);
  }
  .el-footer{
    text-align: center;
    background: #545c64;
    height: 80px;
    line-height: 80px;
  }
</style>