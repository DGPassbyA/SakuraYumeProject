import Vue from 'vue';
import Boss from "@/components/Boss.vue";
import HistoryPanel from "@/components/HistoryPanel";
import VueRouter from 'vue-router';
import GuildInfo from "@/components/GuildInfo";
Vue.use(VueRouter);
//
const routes = [
    {path:'/', name: 'boss', component: Boss},
    {path:'/history', name: 'history', component: HistoryPanel},
    {path:'/guild', name: 'guild', component: GuildInfo},
    {path:'/boss', redirect:{name:'boss'}},
];
const router = new VueRouter({
    mode: 'history',
    base: __dirname,
    routes: routes
});
export default router;