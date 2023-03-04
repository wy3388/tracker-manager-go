<script setup>
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import {reactive} from "vue"

const local = reactive(zhCn)
</script>

<template>
  <el-config-provider :locale="local">
    <router-view/>
  </el-config-provider>
</template>

<style scoped>
</style>
