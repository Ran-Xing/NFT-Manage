<template>
  <div class="navigation" id="NavigationBar">
    <div class="left" @touchstart="leftOnTouchStart" @touchmove="OnTouchMove" @touchend="leftOnTouchEnd">
      <i class="iconfont" v-if="leftMain" id="left">&#xe8f2;</i>
      <i class="iconfont" v-if="leftOption" id="left">&#xe654;</i>
    </div>
    <div class="home" @touchstart="homeOnTouchStart" @touchmove="OnTouchMove" @touchend="homeOnTouchEnd">
      <i class="iconfont" v-if="homeMain" id="refresh">&#xe64f;</i>
      <i class="iconfont" v-if="homeOption" id="home">&#xe632;</i>
    </div>
    <div class="right" @touchstart="rightOnTouchStart" @touchmove="OnTouchMove" @touchend="rightOnTouchEnd">
      <!--      <i class="iconfont" v-if="main" id="right" onclick=history.go(+1)>&#xe8f1;</i>-->
      <i class="iconfont" v-if="rightMain" id="right">&#xe8f1;</i>
      <i class="iconfont" v-if="rightOption" id="menu">&#xe664;</i>
    </div>
  </div>
</template>

<script>
export default {
  name: "NavigationBar",
  data() {
    return {
      loop: 0,
      leftMain: true,
      leftOption: false,
      homeMain: true,
      homeOption: false,
      rightMain: true,
      rightOption: false,
    };
  },
  methods: {
    OnTouchMove() {
      // 清除定时器
      clearTimeout(this.loop)
      this.loop = 0
    },
    leftOnTouchStart() {
      this.loop = setTimeout(() => {
        this.loop = 0
        // 执行长按要执行的内容
        this.leftMain = false
        this.leftOption = true
        // 选择cookies
        alert("选择cookies")
      }, 500) // 定时的时间
      return false
    },
    leftOnTouchEnd() {
      // 清除定时器
      clearTimeout(this.loop)
      if (this.loop !== 0) {
        // 单击操作
        this.leftMain = true
        this.leftOption = false
        // 向前
        history.go(-1)
      }
    },
    homeOnTouchStart() {
      this.loop = setTimeout(() => {
        this.loop = 0
        // 执行长按要执行的内容
        this.homeMain = false
        this.homeOption = true
        // 跳转主页
        this.$router.push({
          path: "/",
        });
      }, 500) // 定时的时间
      return false
    },
    homeOnTouchEnd() {
      // 清除定时器
      clearTimeout(this.loop)
      if (this.loop !== 0) {
        // 单击操作
        this.homeMain = true
        this.homeOption = false
        // 刷新页面
        history.go(0)
      }
    },
    rightOnTouchStart() {
      this.loop = setTimeout(() => {
        this.loop = 0
        // 执行长按要执行的内容
        this.rightMain = false
        this.rightOption = true
        //  打开菜单
      }, 500) // 定时的时间
      return false
    },
    rightOnTouchEnd() {
      // 清除定时器
      clearTimeout(this.loop)
      if (this.loop !== 0) {
        // 单击操作
        this.rightMain = true
        this.rightOption = false
        // 向后
        history.go(+1)
      }
    }
  }
}
</script>

<style scoped>
.navigation {
  z-index: 99999;
  height: 5%;
  position: fixed;
  left: 1vh;
  bottom: 0vh;
  width: 96%;
  text-align: center;
  border-radius: 15vh;
  /*padding: 0.7vh;*/
  background: rebeccapurple;
}

.navigation > .left {
  display: block;
  float: left;
  margin-left: 3vh;
  line-height: 5vh;
  color: #ff6700;
  font-weight: bolder;
}

.navigation > .home {
  line-height: 5vh;
  display: inline-flex;
  color: #ff6700;
  font-weight: bolder;
}

.navigation > .right {
  display: block;
  float: right;
  margin-right: 3vh;
  line-height: 5vh;
  color: #ff6700;
  font-weight: bolder;
}
</style>
