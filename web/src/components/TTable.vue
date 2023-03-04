<script setup>
import {onMounted, ref} from "vue"
import {request} from "@/http/index.js"

onMounted(() => {
  loadData()
})

const props = defineProps({
  url: {
    type: String,
    default: '',
  },
  method: {
    type: String,
    default: 'GET',
  },
  params: {
    type: Object,
    default: {
      page: '1',
      pageSize: '15',
    },
  },
  showPage: {
    type: Boolean,
    default: true,
  },
})

let tableData = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(15)
const loading = ref(false)
const tableRef = ref()

const loadData = () => {
  if (props.url) {
    if (props.showPage) {
      props.params['page'] = page.value.toString()
      props.params['pageSize'] = pageSize.value.toString()
    } else {
      delete props.params['page']
      delete props.params['pageSize']
    }
    loading.value = true
    request(props.url, props.params, props.method)
        .then(resp => {
          let data = resp.data
          loading.value = false
          if (props.showPage) {
            tableData.value.push(...data.data)
            total.value = data.total
            pageSize.value = data.pageSize
          } else {
            tableData.value.push(...data)
          }
        })
  }
}

const handlerSizeChange = (val) => {
  pageSize.value = val
  loadData()
}

const handlerCurrentChange = (val) => {
  page.value = val
  loadData()
}

const emit = defineEmits(['rowDbClick', 'selection'])
const rowDbClick = (row, column, event) => {
  emit('rowDbClick', row, column, event)
}
const reloadData = () => {
  tableData.value = []
  loadData()
}

const getSelectionRows = () => {
  return tableRef.value.getSelectionRows()
}

defineExpose({
  reloadData,
  getSelectionRows
})
</script>

<template>
  <div class="table-container">
    <el-table ref="tableRef" v-loading.fullscreen.lock="loading" :data="tableData" :empty-text="'暂无数据'"
              @row-dblclick="rowDbClick">
      <slot/>
    </el-table>
    <div class="page-container">
      <el-pagination v-if="showPage" :current-page="page"
                     :page-size="pageSize"
                     :page-sizes="[15,30,50,100]"
                     :total="total"
                     style="margin: 0 2% 0 auto"
                     @size-change="handlerSizeChange"
                     @current-change="handlerCurrentChange"/>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.table-container {
  display: flex;
  flex-direction: column;

  .page-container {
    display: flex;
    margin-top: 10px;
  }
}
</style>