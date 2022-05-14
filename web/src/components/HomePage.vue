<template class="box">
  <div class="container" id="HomePage">
    <div class="box" v-for="listData in listDatas" :key="listData">
      <div class="img">
        <img :src="listData.imgUrl" :alt="listData.imgAlt" :title="listData.imgTitle"
             @click.prevent="clickFn(listData)"/>
      </div>
    </div>
  </div>
</template>

<script>
const listDatas = [
  {
    imgUrl: "https://cdn.icon-icons.com/icons2/2648/PNG/128/weixin_icon_160690.png",
    imgAlt: "微信",
    imgTitle: "微信",
    nftUrl: "https://weixin.sogou.com/"
  },
  {
    imgUrl: "https://cdn.icon-icons.com/icons2/3455/PNG/128/favorite_folder_star_file_icon_219500.png",
    imgAlt: "微博",
    imgTitle: "微博",
    nftUrl: "https://weibo.com/"
  },
  {
    imgUrl: "https://cdn.icon-icons.com/icons2/3455/PNG/128/vector_file_design_icon_219513.png",
    imgAlt: "QQ",
    imgTitle: "QQ",
    nftUrl: "https://www.qq.com/"
  },
  {
    imgUrl: "https://cdn.icon-icons.com/icons2/3455/PNG/128/warning_file_alert_paper_document_icon_219508.png",
    imgAlt: "微信",
    imgTitle: "微信",
    nftUrl: "https://weixin.sogou.com/"
  },
  {
    imgUrl: "https://cdn.icon-icons.com/icons2/3455/PNG/128/vector_file_design_icon_219513.png",
    imgAlt: "QQ",
    imgTitle: "QQ",
    nftUrl: "https://www.qq.com/"
  },
  {
    imgUrl: "https://cdn.icon-icons.com/icons2/3455/PNG/128/warning_file_alert_paper_document_icon_219508.png",
    imgAlt: "微信",
    imgTitle: "微信",
    nftUrl: "https://weixin.sogou.com/"
  }
]
export default {
  name: 'HomePage',
  el: '#HomePage',
  methods: {
    clickFn: function (listData) {
      this.$router.push({
        path: "/nft",
        query: {
          src: listData.nftUrl
        }
      })
      return {
        src: listData.nftUrl
      }
    }
  },
  data() {
    return {
      listDatas: listDatas
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/*body{*/
/*  width: 100%;*/
/*  height: 100vh;*/
/*  display: flex;*/
/*  justify-content: center;*/
/*  align-items: center;*/
/*  background-color: #efeeee;*/
/*}*/
.container {
  width: 100%;
  height: 95vh;
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  align-items: center;
}

.container .box {
  display: flex;
  justify-content: space-around;
  flex-direction: column;
  align-items: center;
  align-content: center;
  width: 100px;
  height: 140px;
  margin: 20px;
}

.container .box .img {
  width: 100px;
  height: 100px;
  box-shadow: 18px 18px 30px rgba(0, 0, 0, 0.1),
  -18px -18px 30px rgba(255, 255, 255, 1);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #efeeee;
  transition: box-shadow .2s ease-out;
  position: relative;
}

.container .box .img img {
  width: 60px;
  transition: width 0.2s ease-out;
}

.container .box p {
  color: slategrey;
}

.container .box .img:hover {
  box-shadow: 0px 0px 0px rgba(0, 0, 0, 0.2),
  0px 0px 0px rgba(255, 255, 255, 0.8),
  inset 18px 18px 30px rgba(0, 0, 0, 0.1),
  inset -18px -18px 30px rgba(255, 255, 255, 1);
  transition: box-shadow .2s ease-out;
}

.container .box .img:hover img {
  width: 58px;
  transition: width 0.2s ease-out;
}
</style>
