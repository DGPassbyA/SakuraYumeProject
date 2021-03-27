<template>
    <div class="mdui-dialog" id="addCookieDialog">
        <div class="mdui-dialog-title">{{this.hasCookie?"删除饼干":"添加饼干"}}</div>
<!--      <div class="mdui-dialog-content">You'll lose all photos and media!</div>-->
        <div v-if="!hasCookie" class="mdui-textfield" style="padding:10px 20px">
            <textarea v-model="cookie" class="mdui-textfield-input" rows="4" placeholder="你的饼干 >.< 给我！"></textarea>
        </div>
        <div v-if="hasCookie" class="mdui-dialog-content">你确定要删除饼干吗？QAQ</div>
        <div class="mdui-dialog-actions">
            <button class="mdui-btn mdui-ripple" mdui-dialog-close>取消</button>
            <button class="mdui-btn mdui-ripple" mdui-dialog-confirm v-on:click="manageCookie">确定</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import mdui from 'mdui';
export default {
  name: "AddCookiePanel",
  props:{
    title: String,
    hasCookie: Boolean,
  },
  data: function(){
      return {
          cookie: "",
      }
  },
  methods: {
      manageCookie: function(event){
        console.log("aa");
        if(this.$cookies.isKey("sakurayume_user_token")){
            this.hasCookie = true;
        }
        if(this.hasCookie === true) {
            this.deleteCookie()
        }else{
            this.checkCookie();
        }
      },
      deleteCookie: function(event){
          if(this.$cookies.isKey("sakurayume_user_token")){
              this.$cookies.remove("sakurayume_user_token");
              this.hasCookie = false;
          }
      },
      checkCookie: function(event) {
        let data = new FormData();
        console.log(this.cookie);
        data.append("token",this.cookie);
        this.cookie = "";
        axios
        .post('http://localhost:8081/parseToken',data, {
            withCredentials: true
        })
        .then(response=>(
            response.data.code == 0 ? this.hasCookie = true : mdui.snackbar({
                message: ' 无效的饼干 QAQ 再去问妈要一个吧 ',
                position: 'top',
            })
        ))
      }
  }
}
</script>

<style scoped>

</style>
