<template>
  <el-form ref="dataForm" :rules="rules" :model="userInfo" label-position="center" label-width="80px">
    <el-form-item label="账号名称" prop="userName">
      <el-input v-model="userInfo.userName" />
    </el-form-item>
    <!-- <el-form-item label="头像" prop="headImgUrl" style="">
          <el-avatar :size="100" :src="image" />
          <el-button type="primary" style="position:absolute;bottom:15px;margin-left: 20px;text-align:center;" icon="el-icon-upload" @click="imagecropperShow=true">
            Upload Avatar
          </el-button>
          
          <div style="float:left">
            <image-cropper
              v-show="imagecropperShow"
              :key="imagecropperKey"
              :width="300"
              :height="300"
              field="file"
              :url="httpWebApi"
              lang-type="en"
              @close="imagecropperShow=false"
              @crop-upload-success="cropSuccess"
            />
          </div>
        </el-form-item> -->
    <el-form-item label="手机号码" prop="mobile">
      <el-input v-model="userInfo.mobile" />
    </el-form-item>
    <el-form-item label="邮箱" prop="email">
      <el-input v-model="userInfo.email" />
    </el-form-item>
    <el-form-item label="性别">
      <el-radio-group v-model="userInfo.sex">
        <el-radio label="Secrecy">保密</el-radio>
        <el-radio label="Male">男</el-radio>
        <el-radio label="Female">女</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="updateUser">保存</el-button>
    </el-form-item>
  </el-form>
  
</template>

<script>
import { updateSysUser } from '@/api/user'
export default {
  name: 'Account',
  props: {
    userInfo: {
      type: Object,
      default: () => {
        return {
          email: '',
          headImgUrl: '',
          id: '',
          mobile: '',
          sex: '',
          userName: '',
          roleIds: []
        }
      }
    }
  },
  data() {
    return {
      rules: {
        loginName: [{ required: true, message: '登录账号必填', trigger: 'blur' }],
        userName: [{ required: true, message: '用户名称必填', trigger: 'blur' }]
        // password: [{ validator: validatePass, required: true, trigger: 'blur' }],
        // surePassword: [{ validator: validatePass2, required: true, trigger: 'blur' }]
      }
    }
  },
  methods: {
    updateUser() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.userInfo.roleIds = ['00000000-0000-0000-0000-000000000000']
          updateSysUser(this.userInfo).then(() => {
            this.$notify({
              title: 'Success',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    }
  }
}
</script>
