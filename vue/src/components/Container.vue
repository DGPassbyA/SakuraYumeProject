<template>
  <div class="mdui-container doc-container">
    <router-view :key="new Date().getTime()" @changeBoss="changeBoss"></router-view>
  </div>
</template>

<script>
// import Boss from "@/components/Boss";
// import AddDamagePanel from "@/components/AddDamagePanel";
// import HistoryPanel from "@/components/HistoryPanel";
export default {
  name: "Container",
  data: function (){
    return {
      nowBoss:{
        "round" : 2,
        "name" : "飞龙",
        "NHP" : 6000000
      }
    }
  },
  methods: {
    changeBoss: function (round, name, NHP) {
      this.nowBoss.round = round;
      this.nowBoss.name = name;
      this.nowBoss.NHP = NHP;
      this.$emit("changeBoss",this.nowBoss.round, this.nowBoss.name, this.nowBoss.NHP);
    }
  },
}
</script>

<style scoped>

</style>