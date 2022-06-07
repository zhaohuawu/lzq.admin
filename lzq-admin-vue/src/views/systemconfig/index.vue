<template>
  <div class="app-container">
    <el-tabs :tab-position="tabPosition" type="border-card">
      <el-tab-pane label="七牛云配置"><qiniuconfig :extra-value="extraValue" /></el-tab-pane>
      <el-tab-pane label="本地配置">配置管理</el-tab-pane>
    </el-tabs>
 
  </div>
</template>
<script>
import qiniuconfig from './components/QiniuConfig'
import { getSysConfigInfo } from '@/api/sysconfig'
export default {
  name: 'SystemConfig',
  components: { qiniuconfig },
  data() {
    return {
      tabPosition: 'top', // top,right,bottom,left
      extraValue: {},
      sysconfig: {
        code: 'QiNiuStorage',
        configType: 'ExtraQiNiuConfig'
      }
    }
  },
  created() {
    this.getInfo()
  },
  methods: {
    getInfo() {
      getSysConfigInfo(this.sysconfig.configType, this.sysconfig.code).then(response => {
        this.extraValue = response.extraValue
      })
    }
  }
}
</script>
