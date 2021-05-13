<template>
  <div>
    <h1 class="doc-title mdui-text-color-theme">BOSS进度</h1>
    <p>当前进度：{{boss[0].round}}周目 - {{boss[0].name}}</p>
    <div class="mdui-panel" id="panel">
      <div class="mdui-panel-item" v-for="item in boss" :key="item.id" v-bind:id="'pi-'+ boss.indexOf(item)" v-bind:class="item.NHP == 0?'':'mdui-panel-item-open'">
        <div class="mdui-panel-item-header"  v-on:click="openPanel" v-bind:num="boss.indexOf(item)">
          <div class="mdui-panel-item-title">{{item.name}}</div>
          <div class="mdui-panel-item-summary">{{item.NHP == 0 ? '已击破':'攻略中：'+ item.NHP + '/' + item.AHP}}</div>
          <i class="mdui-panel-item-arrow mdui-icon material-icons">keyboard_arrow_down</i>
        </div>
        <div class="mdui-progress">
          <div class="mdui-progress-determinate" v-bind:style="{width: 100- ((item.AHP - item.NHP)/item.AHP*100) + '%'}"></div>
        </div>
        <div class="mdui-panel-item-body">
          <ul class="mdui-list">
            <li v-for="i in item.history" :key="i.id" class="mdui-list-item mdui-ripple mdui-container">
              <span class="mdui-col-xs-4 mdui-col-sm-4">{{i.player}}</span>
              <span class="mdui-col-xs-4 mdui-col-sm-4">{{i.damage}}</span>
              <span class="mdui-col-xs-4 mdui-col-sm-4">{{i.time}}</span>
            </li>
          </ul>
<!--          <div class="mdui-panel-item-actions">-->
<!--            <button class="mdui-btn mdui-ripple" mdui-panel-item-close>cancel</button>-->
<!--            <button class="mdui-btn mdui-ripple">save</button>-->
<!--          </div>-->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import mdui from 'mdui'
import axios from 'axios'
import {historyToBoss} from '../assets/js/autoload.js'
export default {
  name: "Boss",
  data: function (){
    return {
      panel: new mdui.Panel("#panel"),
      num: 0,
      boss:[
        {
          "round" : 0,
          "name" : "",
          "AHP" : 0,
          "NHP" : 0,
          "history" : [
            {
              "player" : "player",
              "time" : "0000-0-00 00:00",
              "damage" : 0
            },
          ]
        }
      ]
    }
  },
  methods: {
    openPanel:function (e){
      let i = e.target;
      while(i.getAttribute("num") == null){
        i = i.parentNode;
      }
      let pi = document.getElementById("pi-" + i.getAttribute("num"));
      this.panel.toggle(pi);
    },
    fillBoss: function (resolve, response) {
      this.boss = historyToBoss(resolve.data, response.data)
      //不能直接传对象，不然会死循环？？？
      this.$emit("changeBoss",this.boss[0].round,this.boss[0].name,this.boss[0].NHP)
    }
  },
  mounted () {
    let month = new Date().getMonth()+1;
    month = month < 10 ? "0" + month : month + ""
    new Promise((resolve, reject)=>{
      axios
      .get('https://127.0.0.1:8081/api/getHistory?clanName=樱之梦&time=2021'+month,{
        withCredentials: true
      })
      .then(response=>{
        resolve(response.data)
      })
      .catch(error=>{
        reject(error)
      })
    }).then(resolve=>{
      axios.get('https://127.0.0.1:8081/api/boss.json', {
        withCredentials: true
      }).then(response=>{
        this.fillBoss(resolve, response)
      })
    }).catch(error=>{
      console.log(error);
    })
  }
}
</script>

<style scoped>

</style>