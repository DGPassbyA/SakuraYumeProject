<template>
  <div>
    <div class="mdui-drawer" id="left-drawer">
      <ul class="mdui-list">
        <router-link to="/boss">
          <li class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons">move_to_inbox</i>
            <div class="mdui-list-item-content">BOSS进度</div>
          </li>
        </router-link>

        <li class="mdui-list-item mdui-ripple" mdui-dialog="{target: '#addDamageDialog'}">
<!--          <a href="javascript:;" mdui-dialog="{target: '#addDamageDialog'}">-->
            <i class="mdui-list-item-icon mdui-icon material-icons">star</i>
            <div class="mdui-list-item-content">出刀&报刀</div>
<!--          </a>-->
        </li>

        <router-link to="/history">
        <li class="mdui-list-item mdui-ripple">
          <i class="mdui-list-item-icon mdui-icon material-icons">send</i>
          <div class="mdui-list-item-content">攻略历史</div>
        </li>
        </router-link>

        <router-link to="/guild">
        <li class="mdui-list-item mdui-ripple">
          <i class="mdui-list-item-icon mdui-icon material-icons">drafts</i>
          <div class="mdui-list-item-content">公会详情</div>
        </li>
        </router-link>


        <li class="mdui-subheader">娱乐类</li>

        <li class="mdui-list-item mdui-ripple">
          <i class="mdui-list-item-icon mdui-icon material-icons">email</i>
          <div class="mdui-list-item-content">网站链接</div>
        </li>

        <li class="mdui-list-item mdui-ripple">
          <i class="mdui-list-item-icon mdui-icon material-icons">delete</i>
          <div class="mdui-list-item-content">公主相册</div>
        </li>

        <!--      <li class="mdui-list-item mdui-ripple">-->
        <!--        <i class="mdui-list-item-icon mdui-icon material-icons">error</i>-->
        <!--        <div class="mdui-list-item-content">Spam</div>-->
        <!--      </li>-->

      </ul>
    </div>
  </div>

</template>

<script>
export default {
  name: "LeftBar"
}
</script>

<style scoped>

</style>