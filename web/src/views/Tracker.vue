<script setup>

import {onMounted, reactive} from "vue"
import {TrackerService} from "@/http/tracker.js";
import {ElMessage} from "element-plus";
import useClipboard from "vue-clipboard3";

onMounted(() => {
  loadData()
})

const data = reactive({
  loading: false,
  showContent: false,
  content: '',
})


const loadData = () => {
  data.content = ''
  data.loading = true
  TrackerService.getTracker()
      .then(resp => {
        data.loading = false
        if (resp.data && resp.data !== '') {
          data.content = resp.data
          data.showContent = true
        }
      })
}

const handlerCopy = async () => {
  if (!data.content || data.content === '') {
    ElMessage.warning('请先同步数据')
    return
  }
  const {toClipboard} = useClipboard()
  try {
    await toClipboard(data.content)
    ElMessage.success('复制成功')
  } catch (e) {
    ElMessage.error('复制失败')
  }
}
</script>

<template>
  <div v-loading="data.loading" class="main-container">
    <div style="margin-top: 5px;">
      <el-button type="primary" @click="handlerCopy">复制</el-button>
    </div>
    <div>
      <el-input v-show="data.showContent" v-model="data.content" autosize type="textarea"/>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.main-container {
  display: flex;
  flex-direction: column;

  div:last-child {
    margin: 5px 0;
  }
}
</style>