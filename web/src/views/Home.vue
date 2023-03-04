<script setup>
import {CollectionTag, Menu as IconMenu, Notification, Position} from '@element-plus/icons-vue'
import {onMounted, ref} from "vue"
import {useRoute} from "vue-router"

onMounted(() => {
  const path = useRoute().fullPath
  if (path !== '/') {
    defaultActive.value = path.replace('/', '')
  }
})

const defaultActive = ref('tracker')
let showDialog = ref(false)

</script>

<template>
  <div>
    <el-container class="container">
      <el-header class="fs14 bs-light" style="border-bottom: 1px solid var(--el-border-color)">
        <div class="header">
          <h2 style="color: var(--el-text-color-primary);">Tracker Manager Go</h2>
          <span style="color: var(--el-text-color-regular);" @click="showDialog = true">关于</span>
        </div>
      </el-header>
      <el-container class="menu-container">
        <el-aside width="200px">
          <el-menu :default-active="defaultActive" :router="true" style="height: 100%">
            <el-menu-item index="tracker">
              <template #title>
                <el-icon>
                  <icon-menu/>
                </el-icon>
                Tracker列表
              </template>
            </el-menu-item>
            <el-menu-item index="source">
              <template #title>
                <el-icon>
                  <collection-tag/>
                </el-icon>
                数据源
              </template>
            </el-menu-item>
            <el-menu-item index="sync">
              <template #title>
                <el-icon>
                  <notification/>
                </el-icon>
                同步
              </template>
            </el-menu-item>
            <el-menu-item index="client">
              <template #title>
                <el-icon>
                  <position/>
                </el-icon>
                客户端
              </template>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-scrollbar :max-height="'calc(100vh - 60px)'" style="width: 100%">
          <el-main>
            <router-view/>
          </el-main>
        </el-scrollbar>
      </el-container>
    </el-container>
  </div>

  <el-dialog v-model="showDialog" title="关于" width="30%">
    <div class="about">
      <span>Tracker Manager Go</span>
    </div>
    <template #footer>
        <span>
          <el-button type="default" @click="showDialog = false">关闭</el-button>
        </span>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.menu-container {
  display: flex;
  height: calc(100% - 60px);
}

.header {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  h2 {
    margin-left: 3%;
  }

  span {
    cursor: pointer;
    user-select: none;
    margin: 0 5% 0 auto;
  }
}

.about {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>