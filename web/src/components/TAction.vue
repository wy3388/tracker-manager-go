<script setup>
const props = defineProps({
  showAddButton: {
    type: Boolean,
    default: true,
  },
  showDelButton: {
    type: Boolean,
    default: false
  }
})
const emit = defineEmits(['handlerAdd', 'handlerDel'])

const handlerAdd = () => {
  emit('handlerAdd')
}

const handlerDel = () => {
  emit('handlerDel')
}
</script>

<template>
  <div>
    <el-button type="primary" v-show="props.showAddButton" @click="handlerAdd">新增</el-button>
    <el-button type="danger" v-show="props.showDelButton" @click="handlerDel">删除</el-button>
  </div>
</template>

<style scoped lang="scss">

</style>