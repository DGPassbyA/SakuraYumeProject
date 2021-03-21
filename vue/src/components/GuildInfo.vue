<template>
  <div>
    <h1 class="doc-title mdui-text-color-theme">公会详情</h1>
    <div>
      <h3>公会成员</h3>
      <ul class="mdui-list">
        <li class="mdui-list-item mdui-ripple" v-for="item in members" :key="item.id">
          <i class="mdui-list-item-avatar mdui-icon material-icons">folder</i>
          <div class="mdui-list-item-content mdui-col-xs-6">
            <div style="font-size: 20px;" v-bind:style="{'color':item.position === '会长'?'#E91E63':'#3F51B5'}">{{item.id}}</div>
            <div style="font-weight:100;font-size:12px;padding-top: 3px;">{{item.position}}</div>
          </div>
        </li>
      </ul>
    </div>

  </div>
</template>

<script>
export default {
  name: "GuildInfo",
  data: function (){
    return {
      members: [
        {"id":"洛夜","position": "会长"},
        {"id": "aaa","position": "副会长"},
        {"id": "bbb","position": "会员"}
      ]
    }
  }
}
</script>

<style scoped>

</style>