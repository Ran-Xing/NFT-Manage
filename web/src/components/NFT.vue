<template>
  <div class="nft" id="nft">
    <div class="nft-image">
      <iframe no-referrer
              sandbox="allow-same-origin allow-scripts allow-popups allow-forms allow-pointer-lock"
              style="
                width: 100%;
                height:95vh;
                border-width: 0px;
                " class="iframe" id="iframe"
              :src="nftUrl">
      </iframe>
    </div>
  </div>
</template>

<script>
export default {
  name: "NFT",
  el: "#nft",
  methods: {
    getNftUrl() {
      console.log()
      return this.$route.query.src;
    }
  },
  data() {
    return {
      nftUrl: this.getNftUrl()
    }
  },
}
</script>

<style scoped>

</style>
