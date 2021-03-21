<template>
  <div class="mdui-container doc-container">
    <router-view :key="new Date().getTime()"></router-view>
  </div>
</template>

<script>
// import Boss from "@/components/Boss";
// import AddDamagePanel from "@/components/AddDamagePanel";
// import HistoryPanel from "@/components/HistoryPanel";
export default {
  name: "Container",
  // components: {HistoryPanel, AddDamagePanel, Boss}
}
</script>

<style scoped>

</style>