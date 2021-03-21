<template>
    <div class="mdui-dialog" id="addDamageDialog">
      <div class="mdui-dialog-title">出刀&报刀</div>
<!--      <div class="mdui-dialog-content">You'll lose all photos and media!</div>-->
      <div style="font-weight: 300;height:30px;" class="mdui-dialog-content">
        <div style="padding: 10px 0;">{{nowBoss.name}}（{{nowBoss.round}}周目） <br/>剩余血量：{{nowBoss.NHP}}</div>
        <div>
          <select class="mdui-select" mdui-select>
            <option value="1" selected>普通刀</option>
            <option value="2">尾刀</option>
            <option value="3">补时刀</option>
          </select>
        </div>
        <div style="padding-top: 10px;">
          <input class="mdui-textfield-input" type="text" placeholder="Your Damage"/>
        </div>
        <div class="mdui-textfield">
          <textarea class="mdui-textfield-input" placeholder="Note(optional)"></textarea>
        </div>
      </div>
      <div class="mdui-dialog-actions">
        <button class="mdui-btn mdui-ripple" mdui-dialog-close>取消</button>
        <button class="mdui-btn mdui-ripple" mdui-dialog-confirm>确定</button>
      </div>
  </div>
</template>

<script>
export default {
  name: "AddDamagePanel",
  data: function (){
    return {
      nowBoss:{
        "round" : 1,
        "name" : "狮鹫",
        "NHP" : 5000000
      }
    }
  }
}
</script>

<style scoped>

</style>