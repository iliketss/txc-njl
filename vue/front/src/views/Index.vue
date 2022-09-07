<template>
  <div class="el_goma">
    <el-carousel :interval="4000" type="card" height="450px">
      <el-carousel-item v-for="item in 6" :key="item" >
        <el-image class="go-img" :src="src" />
      </el-carousel-item>
    </el-carousel>

  </div>
</template>

<script lang="ts">


import {defineComponent} from "vue";

export default defineComponent({
  setup(){
    const src ="../assets/1.png"

    return {
    src
    }
  }
})
</script>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}
.go-img{

}
.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}
.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>