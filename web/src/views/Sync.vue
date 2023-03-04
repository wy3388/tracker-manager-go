<script setup>
import {SyncService} from "@/http/sync.js";
import message from "@/utils/message.js";
import TTable from "@/components/TTable.vue";
import {ref} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";

const tableRef = ref()
const loading = ref(false)

const handlerSync = () => {
  loading.value = true
  SyncService.sync()
      .then(resp => {
        loading.value = false
        message.show(resp, () => {
          tableRef.value.reloadData()
        })
      })
}

const handlerDel = () => {
  let rows = tableRef.value.getSelectionRows()
  let id = rows.map(row => row.uuid)
  if (id.length < 1) {
    ElMessage.warning('请选择数据');
    return
  }
  ElMessageBox.confirm('是否删除记录？', '删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    SyncService.deleteById(id)
        .then(resp => {
          message.show(resp, () => {
            tableRef.value.reloadData()
          })
        })
  })
}

const handlerClear = () => {
  ElMessageBox.confirm('是否清空缓存记录？', '删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    SyncService.clear().then(resp => {
      message.show(resp)
    })
  })
}

</script>

<template>
  <div v-loading.fullscreen.lock="loading">
    <div>
      <el-button type="primary" @click="handlerSync">同步</el-button>
      <el-button type="danger" @click="handlerDel">删除</el-button>
      <el-button type="warning" @click="handlerClear">清空缓存数据</el-button>
    </div>
    <t-table ref="tableRef" :show-page="false" style="margin-top: 5px;" url="sync">
      <el-table-column type="selection" width="55"/>
      <el-table-column align="center" label="序号" type="index" width="80"/>
      <el-table-column :show-overflow-tooltip="true" label="名称" prop="sourceName"/>
      <el-table-column label="成功数量" prop="successNum"/>
      <el-table-column label="失败数量" prop="errorNum"/>
      <el-table-column align="center" label="同步时间" prop="createTime" width="200"/>
    </t-table>
  </div>
</template>

<style lang="scss" scoped>

</style>