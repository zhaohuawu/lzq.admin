<template>
  <el-form ref="dataForm" :rules="pwdRules" :model="passwordInfo" label-position="center" label-width="80px">
    <el-form-item label="旧密码" prop="password">
      <el-input v-model.trim="passwordInfo.password" :type="passw1">
        <i slot="suffix" class="el-input__icon el-icon-view" @click="showPass(1)" />
      </el-input>
    </el-form-item>
    <el-form-item label="新密码" prop="newPassword">
      <el-input v-model.trim="passwordInfo.newPassword" :type="passw2">
        <i slot="suffix" class="el-input__icon el-icon-view" @click="showPass(2)" />
      </el-input>
    </el-form-item>
    <el-form-item label="确认密码" prop="surePassword">
      <el-input v-model.trim="passwordInfo.surePassword" :type="passw3">
        <i slot="suffix" class="el-input__icon el-icon-view" @click="showPass(3)" />
      </el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="updateUser">保存</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { updateCurrentUserPassword } from '@/api/user'
export default {
  name: 'UpdatePassword',
  data() {
    return {
      passwordInfo: {
        password: '',
        newPassword: '',
        surePassword: ''
      },
      pwdRules: {
        password: [{
          required: true,
          message: '旧密码',
          trigger: 'blur'
        }],
        newPassword: [{
          required: true,
          message: '新密码',
          trigger: 'blur'
        }],
        surePassword: [{
          required: true,
          message: '确认密码',
          trigger: 'blur'
        }]
      },
      passw1: 'password',
      passw2: 'password',
      passw3: 'password'
    }
  },
  methods: {
    updateUser() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          updateCurrentUserPassword(this.passwordInfo).then(rsp => {
            this.$notify({
              title: 'Success',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
            
            this.$store.dispatch('user/logout')
            this.$router.push(`/login?redirect=${this.$route.fullPath}`)
          })
        }
      })
    },
    // 密码的隐藏和显示
    showPass(n) {
      // 点击图标是密码隐藏或显示
      if (n === 1) {
        if (this.passw1 === 'text') {
          this.passw1 = 'password'
        } else {
          this.passw1 = 'text'
        }
      } else if (n === 2) {
        if (this.passw2 === 'text') {
          this.passw2 = 'password'
        } else {
          this.passw2 = 'text'
        }
      } else {
        if (this.passw3 === 'text') {
          this.passw3 = 'password'
        } else {
          this.passw3 = 'text'
        }
      }
    }
  }
}

</script>
