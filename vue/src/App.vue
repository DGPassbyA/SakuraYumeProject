<template>
  <div>
    <TopNav v-bind:title="webTitle"></TopNav>
    <LeftBar></LeftBar>
    <Container @changeBoss="changeBoss"></Container>
    <AddDamagePanel :round="round" :name="name" :NHP="NHP" :clanName="clanName"></AddDamagePanel>
    <AddCookiePanel></AddCookiePanel>
  </div>
</template>

<script>
import TopNav from '@/components/TopNav'
import LeftBar from "@/components/LeftDrawer";
import Container from "@/components/Container";
import AddDamagePanel from "@/components/AddDamagePanel";
import AddCookiePanel from "@/components/AddCookiePanel";

export default {
  name: 'App',
  data: function(){
    return {
      round : 2,
      name : "飞龙",
      NHP : 6000000,
      webTitle: "SakuraYume·樱之梦",
      clanName: "樱之梦"
    }
  },
  components: {
    AddDamagePanel,
    Container,
    LeftBar,
    TopNav,
    AddCookiePanel,
  },
  methods: {
    changeBoss: function (round, name, NHP) {
      this.round = round;
      this.name = name;
      this.NHP = NHP;
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  width: 100%;
}
</style>
