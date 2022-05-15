<template>
  <div style="text-align: center">
    <el-form ref="dataForm" :model="extraValue" label-position="center" style="width: 70%" label-width="150px">
      <el-form-item label="七牛云域名" prop="baseUrl">
        <el-input v-model.trim="extraValue.baseUrl" />
      </el-form-item>
      <el-form-item label="公钥" prop="accessKey">
        <el-input v-model.trim="extraValue.accessKey" />
      </el-form-item>
      <el-form-item label="私钥" prop="secretKey">
        <el-input v-model.trim="extraValue.secretKey" :type="passw">
          <i slot="suffix" :class="icon" @click="showPass" />
        </el-input>
      </el-form-item>
      <el-form-item label="空间名称" prop="bucket">
        <el-input v-model.trim="extraValue.bucket" />
      </el-form-item>
      <el-form-item label="主目录" prop="directory">
        <el-input v-model.trim="extraValue.directory" />
      </el-form-item>
      <el-form-item label="存储区域" prop="area">
        <el-input v-model.trim="extraValue.area" />
      </el-form-item>
     
      <el-form-item>
        <el-button type="primary" @click="submit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { qnUpdate } from '@/api/sysconfig'
export default {
  name: 'QiniuConfig',
  props: {
    extraValue: {
      type: Object,
      default: () => {
        return {
          accessKey: '',
          area: '',
          baseUrl: '',
          bucket: '',
          secretKey: '',
          directory: ''
        }
      }
    }
  },
  data() {
    return {
      icon: 'el-input__icon el-icon-view',
      passw: 'password'
    }
  },
  methods: {
    // 密码的隐藏和显示
    showPass() {
      // 点击图标是密码隐藏或显示
      if (this.passw === 'text') {
        this.passw = 'password'
        // 更换图标
        this.icon = 'el-input__icon el-icon-view'
      } else {
        this.passw = 'text'
        this.icon = 'el-input__icon el-icon-loading'
      }
    },
    submit() {
      const sysconfig = {
        code: 'QiNiuStorage',
        configType: 'ExtraQiNiuConfig',
        extraValue: this.extraValue
      }
      qnUpdate(sysconfig).then(response => {
        this.$message({
          message: '保存成功',
          type: 'success',
          duration: 5 * 1000
        })
      })
    }
  }
}
</script>
