<script setup>
import TTable from "@/components/TTable.vue"
import {nextTick, reactive, ref} from "vue"
import TAction from "@/components/TAction.vue"
import {ClientService} from "@/http/client.js"
import message from "@/utils/message.js"
import {ElMessage, ElMessageBox} from "element-plus"
import {ClientHistoryService} from "@/http/client_history.js"

const formRef = ref()
const tableRef = ref()
const tableHistoryRef = ref()

const activeName = ref('client')
let data = reactive({
  id: '',
  showDialog: false,
  title: '',
  loading: false,
  form: {
    name: '',
    url: '',
    username: '',
    password: '',
  },
})

const rules = reactive({
  name: [
    {required: true, message: '请选择下载软件类型'},
  ],
  url: [
    {required: true, message: '请输入地址'},
  ],
})

const rowDbClick = (row) => {
  data.title = '修改'
  nextTick(() => {
    formRef.value.clearValidate()
    formRef.value.resetFields()
  })
  data.showDialog = true
  data.loading = true
  ClientService.getById(row.uuid)
      .then(resp => {
        data.id = resp.data.uuid
        data.form.name = resp.data.name
        data.form.url = resp.data.url
        data.form.username = resp.data.username
        data.form.password = resp.data.password
        data.loading = false
      })
}

const submitForm = async (el) => {
  if (!el) return
  await el.validate((valid) => {
    if (!valid) return
    if (data.id !== '') {
      data.form['uuid'] = data.id
      ClientService.updateById(data.form)
          .then(resp => {
            message.show(resp, () => {
              data.id = ''
              data.showDialog = false
              tableRef.value.reloadData()
            })
          })
    } else {
      delete data.form['uuid']
      delete data.form['create_time']
      ClientService.add(data.form)
          .then(resp => {
            message.show(resp, () => {
              data.showDialog = false
              tableRef.value.reloadData()
            })
          })
    }
  })
}

const handlerAdd = () => {
  data.title = '新增'
  nextTick(() => {
    formRef.value.clearValidate()
    formRef.value.resetFields()
  })
  data.showDialog = true
}

const handlerDel = (id) => {
  ElMessageBox.confirm('是否删除该条记录？', '删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    ClientService.deleteById(id)
        .then(resp => {
          message.show(resp, () => {
            tableRef.value.reloadData()
          })
        })
  })
}

const handlerSync = (id) => {
  ElMessageBox.confirm('是否开始同步客户端？', '同步', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    ClientService.sync(id)
        .then(resp => {
          message.show(resp, () => {
            tableHistoryRef.value.reloadData()
          })
        })
  })
}

const handlerHistoryDel = () => {
  let rows = tableHistoryRef.value.getSelectionRows()
  let id = rows.map(row => row.uuid)
  if (id.length < 1) {
    ElMessage.warning('请选择数据')
    return
  }
  ElMessageBox.confirm('是否删除记录？', '删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    ClientHistoryService.deleteByIds(id)
        .then(resp => {
          message.show(resp, () => {
            tableHistoryRef.value.reloadData()
          })
        })
  })
}
</script>

<template>

  <el-tabs v-model="activeName">
    <el-tab-pane label="客户端" name="client">
      <div style="display: flex;flex-direction: column;">
        <t-action @handler-add="handlerAdd"/>
        <t-table ref="tableRef" :show-page="false" style="margin-top: 5px;" url="client"
                 @row-db-click="rowDbClick">
          <el-table-column type="selection" width="55"/>
          <el-table-column align="center" label="序号" type="index" width="80"/>
          <el-table-column label="类型" prop="name" align="center"/>
          <el-table-column label="地址" prop="url" align="center"/>
          <el-table-column label="用户名" prop="username" align="center"/>
          <el-table-column align="center" label="密码" prop="password"/>
          <el-table-column align="center" label="添加时间" prop="create_time"/>
          <el-table-column label="操作" align="center">
            <template #default="scope">
              <el-button size="small" type="success" @click="handlerSync(scope.row.uuid)">同步</el-button>
              <el-button size="small" type="default" @click="rowDbClick(scope.row)">修改</el-button>
              <el-button size="small" type="danger" @click="handlerDel(scope.row.uuid)">删除</el-button>
            </template>
          </el-table-column>
        </t-table>
      </div>
    </el-tab-pane>
    <el-tab-pane label="同步历史" name="history">
      <div style="display: flex;flex-direction: column;">
        <t-action :show-add-button="false" :show-del-button="true" @handler-del="handlerHistoryDel"/>
        <t-table ref="tableHistoryRef" :show-page="false" style="margin-top: 5px;" url="clientHistory">
          <el-table-column type="selection" width="55"/>
          <el-table-column align="center" label="序号" type="index" width="80"/>
          <el-table-column label="类型" prop="client_name" align="center"/>
          <el-table-column label="同步内容" :show-overflow-tooltip="true" prop="tracker_content" align="center"/>
          <el-table-column align="center" label="同步时间" prop="create_time"/>
        </t-table>
      </div>
    </el-tab-pane>
  </el-tabs>

  <el-dialog v-model="data.showDialog" :title="data.title" width="30%">
    <el-form ref="formRef" v-loading="data.loading" :model="data.form" :rules="rules" label-width="60">
      <el-form-item label="类型" prop="name">
        <el-select v-model="data.form.name" style="width: 100%;">
          <el-option label="qBittorrent" value="qBittorrent">qBittorrent</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="地址" prop="url">
        <el-input v-model="data.form.url" placeholder="http://127.0.0.1:8080"/>
      </el-form-item>
      <el-form-item label="用户名" prop="username">
        <el-input v-model="data.form.username" placeholder="admin"/>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="data.form.password" placeholder="adminadmin"/>
      </el-form-item>
    </el-form>
    <template #footer>
        <span>
          <el-button type="primary" @click="submitForm(formRef)">确定</el-button>
          <el-button type="default" @click="data.showDialog = false">关闭</el-button>
        </span>
    </template>
  </el-dialog>
</template>

<style scoped lang="scss">

</style>