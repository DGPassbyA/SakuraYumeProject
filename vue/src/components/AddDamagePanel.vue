<template>
    <div class="mdui-dialog" id="addDamageDialog">
      <div class="mdui-dialog-title">出刀&报刀</div>
<!--      <div class="mdui-dialog-content">You'll lose all photos and media!</div>-->
      <div style="font-weight: 300;height:30px;" class="mdui-dialog-content">
        <div style="padding: 20px 0;">{{name}}（{{round}}周目） <br/>剩余血量：{{NHP}}</div>
        <div>
          <select v-model="type" class="mdui-select" mdui-select>
            <option value="0" selected>普通刀</option>
            <option value="1">尾刀</option>
            <option value="2">补时刀</option>
          </select>
        </div>
        <div style="padding-top: 10px;">
          <input v-model="dmg" class="mdui-textfield-input" id="add-damge" type="text" placeholder="Your Damage"/>
        </div>
        <!-- <div class="mdui-textfield">
          <textarea class="mdui-textfield-input" placeholder="Note(optional)"></textarea>
        </div> -->
      </div>
      <div class="mdui-dialog-actions">
        <button class="mdui-btn mdui-ripple" mdui-dialog-close>取消</button>
        <button class="mdui-btn mdui-ripple" mdui-dialog-confirm  v-on:click="addHisory">确定</button>
      </div>
  </div>
</template>

<script>
import axios from 'axios';
import mdui from 'mdui';
export default {
  name: "AddDamagePanel",
  props: ["round", "name", "NHP", "clanName"],
  data: function (){
    return {
      dmg: "",
      type: "0",
    }
  },
  methods: {
    addHisory: function (event){
      new Promise((resolve, reject)=>{
        axios.get('https://sakura.passbya.xyz:8081/api/boss.json', {
          withCredentials: true
        }).then(response=>{
          resolve(response)
        }).catch(error=>{
          reject(error);
        })
      }).then(response=>{
        var reg = new RegExp("^[0-9]*$");
        if(!reg.test(this.dmg)){
          mdui.snackbar({
                message: ' 请输入数字 ',
                position: 'top',
            })
          this.dmg = "";
        }
        let data = new FormData();
        let nowDate = new Date();
        let year = nowDate.getFullYear();
        let month = nowDate.getMonth()+1;
        let BossName = response.data.name;
        data.append("clanName",this.clanName);
        data.append("time", year + (month < 10 ? "0" + month : month + ""))
        data.append("round", this.round);
        data.append("dmg", this.dmg);
        data.append("boss", BossName.indexOf(this.name) + 1);
        data.append("flag", this.type);
        axios
          .post('https://127.0.0.1:8081/api/addHistory',data, {
            withCredentials: true
          })
          .then(function(response) {
                if(response.data.code === 0){
                  mdui.snackbar({
                    message: ' 出刀成功 ',
                    position: 'top',
                  })
                }else{
                  mdui.snackbar({
                    message: ' 出刀失败 ',
                    position: 'top',
                  })
                }
          })
      }).catch(error =>{
        console.log(error);
      })
    }
  }
}
</script>

<style scoped>

</style>