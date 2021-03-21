<template>
  <div>
    <TopNav title="SakuraYume·樱之梦"></TopNav>
    <LeftBar></LeftBar>
    <Container></Container>
    <add-damage-panel></add-damage-panel>
  </div>
</template>

<script>
import TopNav from '@/components/TopNav'
import LeftBar from "@/components/LeftDrawer";
import Container from "@/components/Container";
import AddDamagePanel from "@/components/AddDamagePanel";

export default {
  name: 'App',
  data: function(){
    return {
    }
  },
  components: {
    AddDamagePanel,
    Container,
    LeftBar,
    TopNav,
  },
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
