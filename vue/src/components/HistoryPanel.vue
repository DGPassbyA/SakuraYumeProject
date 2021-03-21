<template>
  <div>
    <h1 class="doc-title mdui-text-color-theme">攻略历史</h1>
    <ul class="mdui-list">
      <li class="mdui-list-item mdui-ripple mdui-p-y-2" v-for="item in history" :key="item.id">
        <div class="mdui-col-xs-6">
          <div class="mdui-text-color-pink mdui-typo-title-opacity">{{item.guild}}</div>
          <div class="mdui-p-t-2"><span class="mdui-text-color-blue-500 mdui-typo-headline-opacity">{{item.rank == 0 ? '???' : item.rank}}</span> 名 </div>
        </div>
        <div class="mdui-col-xs-6">
          <div>{{item.time}}</div>
          <div>进度：{{item.progress}}</div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "HistoryPanel",
  data: function (){
    return {
      history:[
        {
          "guild" :"樱之梦",
          "rank" : 0,
          "time" : "2020年十月 天秤座",
          "progress" : "攻略中",
        },
        {
          "guild" :"樱之梦",
          "time" : "2020年九月 处女座",
          "rank" : 20000,
          "progress" : "狮鹫（二王） 二周目",
        },
        {
          "guild" :"樱之梦",
          "time" : "2020年八月 狮子座",
          "rank" : 25000,
          "progress" : "飞龙（一王） 二周目",
        }
      ]
    }
  }
}
</script>

<style scoped>

</style>