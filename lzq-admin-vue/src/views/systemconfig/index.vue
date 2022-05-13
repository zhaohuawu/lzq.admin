<template>
  <div class="app-container">
    <el-tabs :tab-position="tabPosition" type="border-card">
      <el-tab-pane label="七牛云配置"><qiniuconfig :user="user" /></el-tab-pane>
      <el-tab-pane label="本地配置">配置管理</el-tab-pane>
    </el-tabs>
 
  </div>
</template>
<script>
import qiniuconfig from './components/QiniuConfig'
export default {
  name: 'SystemConfig',
  components: { qiniuconfig },
  data() {
    return {
      tabPosition: 'top', // top,right,bottom,left
      user: {}
    }
  }
}
</script>
