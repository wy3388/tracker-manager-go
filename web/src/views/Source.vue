<script setup>
import TAction from "@/components/TAction.vue"
import TTable from "@/components/TTable.vue"
import {nextTick, reactive, ref} from "vue";
import {SourceService} from "@/http/source.js";
import message from "@/utils/message.js";
import {ElMessageBox} from "element-plus";

const formRef = ref()
const tableRef = ref()

let data = reactive({
  id: '',
  loading: false,
  showDialog: false,
  title: '',
  form: {
    key: '',
    name: '',
    author: '',
    url: '',
    is_checked: 0
  }
})

const rules = reactive({
  key: [
    {required: true, message: '请输入key'}
  ],
  name: [
    {required: true, message: '请输入名称'}
  ],
  url: [
    {required: true, message: '请输入URL'}
  ],
  author: [
    {required: true, message: '请输入作者'}
  ]
})

const rowDbClick = (row) => {
  data.title = '修改'
  nextTick(() => {
    formRef.value.clearValidate()
    formRef.value.resetFields()
  })
  data.showDialog = true
  data.loading = true
  SourceService.getById(row.uuid)
      .then(resp => {
        data.id = resp.data.uuid
        data.form.key = resp.data.key
        data.form.name = resp.data.name
        data.form.author = resp.data.author
        data.form.is_checked = resp.data.is_checked
        data.form.url = resp.data.url
        data.loading = false
      })
}

const submitForm = async (el) => {
  if (!el) return
  await el.validate((valid) => {
    if (!valid) return
    if (data.id !== '') {
      data.form['uuid'] = data.id
      SourceService.updateById(data.form)
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
      SourceService.add(data.form)
          .then(resp => {
            message.show(resp, () => {
              data.showDialog = false
              tableRef.value.reloadData()
            })
          })
    }
  })
}

const handlerDel = (id) => {
  ElMessageBox.confirm('是否删除该条记录？', '删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    SourceService.deleteById(id)
        .then(resp => {
          message.show(resp, () => {
            tableRef.value.reloadData()
          })
        })
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
</script>

<template>
  <div style="display: flex;flex-direction: column">
    <t-action @handler-add="handlerAdd"/>
    <t-table ref="tableRef" :show-page="false" style="margin-top: 10px" url="source"
             @row-db-click="rowDbClick">
      <el-table-column type="selection" width="55"/>
      <el-table-column align="center" label="序号" type="index" width="80"/>
      <el-table-column align="center" label="KEY" prop="key"/>
      <el-table-column align="center" label="名称" prop="name"/>
      <el-table-column align="center" label="作者" prop="author"/>
      <el-table-column align="center" label="是否选中" prop="is_checked" width="80">
        <template #default="scope">
          <el-tag :type="scope.row.is_checked === 1 ? 'success' : 'danger'">
            {{ scope.row.is_checked === 1 ? '是' : '否' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="添加日期" prop="create_time"/>
      <el-table-column align="center" label="操作">
        <template #default="scope">
          <el-button size="small" type="default" @click="rowDbClick(scope.row)">修改</el-button>
          <el-button size="small" type="danger" @click="handlerDel(scope.row.uuid)">删除</el-button>
        </template>
      </el-table-column>
    </t-table>
  </div>

  <el-dialog v-model="data.showDialog" :title="data.title" width="30%">
    <el-form ref="formRef" v-loading="data.loading" :model="data.form" :rules="rules">
      <el-form-item label="KEY" prop="key">
        <el-input v-model="data.form.key"/>
      </el-form-item>
      <el-form-item label="名称" prop="name">
        <el-input v-model="data.form.name"/>
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input v-model="data.form.author"/>
      </el-form-item>
      <el-form-item label="URL" prop="url">
        <el-input v-model="data.form.url"/>
      </el-form-item>
      <el-form-item label="是否选中" prop="is_checked">
        <el-radio-group v-model="data.form.is_checked">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
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

<style lang="scss" scoped>
</style>