<template>
    <div class="mdui-appbar mdui-appbar-fixed mdui-appbar-scroll-hide">
      <div class="mdui-toolbar mdui-color-theme">
        <a href="javascript:;" class="mdui-btn mdui-btn-icon"><i class="mdui-icon material-icons" mdui-drawer="{target: '#left-drawer'}">menu</i></a>
        <a href="javascript:;" class="mdui-typo-title">{{title}}</a>
        <div class="mdui-toolbar-spacer"></div>
        <a href="javascript:;" class="mdui-btn mdui-btn-icon" v-on:click="changeHasCookie" mdui-dialog="{target: '#addCookieDialog'}"><i class="mdui-icon material-icons">add</i></a>
        <a href="javascript:;" class="mdui-btn mdui-btn-icon"><i class="mdui-icon material-icons">refresh</i></a>
        <a href="javascript:reload();" class="mdui-btn mdui-btn-icon"><i class="mdui-icon material-icons">more_vert</i></a>
      </div>
<!--      <div class="mdui-tab mdui-color-theme" mdui-tab>-->
<!--          <a href="#example3-tab1" class="mdui-ripple mdui-ripple-white">web</a>-->
<!--          <a href="#example3-tab1" class="mdui-ripple mdui-ripple-white">shopping</a>-->
<!--          <a href="#example3-tab1" class="mdui-ripple mdui-ripple-white">videos</a>-->
<!--      </div>-->
    </div>
</template>

<script>
export default {
  name: "NavBar",
  props:{
    title: String,
  },
  methods: {
    reload :function(e) {
      this.$$forceUpdate();
    },
    changeHasCookie: function(e){
      if(this.$cookies.isKey("sakurayume_user_token")) {
        this.$emit("changeHasCookie",true);
      }else{
        this.$emit("changeHasCookie",false);
      }
    }
  }
}
</script>

<style scoped>

</style>