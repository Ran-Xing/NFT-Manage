<template>
  <!-- 路由组件 -->
  <router-view></router-view>
  <!-- 全局导航栏 -->
  <NavigationBar/>
</template>

<script>
import NavigationBar from "@/components/NavigationBar";
export default {
  name: 'App',
  components: {
    NavigationBar,
  }
}
</script>

<style>
/*#app {*/
/*  font-family: Avenir, Helvetica, Arial, sans-serif;*/
/*  -webkit-font-smoothing: antialiased;*/
/*  -moz-osx-font-smoothing: grayscale;*/
/*  text-align: center;*/
/*  color: #2c3e50;*/
/*  margin-top: 60px;*/
/*}*/
</style>
