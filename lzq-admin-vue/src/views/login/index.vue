<template>
  <div class="login-container" :style="{backgroundImage:`url(${bgImg })`}">
    <div class="login-weaper animated bounceInDown">
      <div class="login-left">
        <div class="bold t1">老知青·SaaS云系统</div>
        <div class="line" />
        <div class="bold t3">专于SaaS系统研发，使用Golang语言开发的好用SaaS云系统，项目用户学习、开源</div>
      </div>
      <div class="login-border">
        <div class="login-main">
          <h4 class="login-title">
            登录
            <!-- <top-lang></top-lang> -->
          </h4>
          <el-form ref="loginForm" class="login-form" status-icon :rules="loginRules" :model="loginForm" label-width="0">
            <el-form-item prop="tenantCode">
              <el-input
                v-model="loginForm.tenantCode"
                size="small"
                auto-complete="off"
                placeholder="租户编码"
                @keyup.enter.native="handleLogin"
              >
                <svg-icon slot="prefix" icon-class="user" class="el-input__icon input-icon" />
              </el-input>
            </el-form-item>

            <el-form-item prop="loginName">
              <el-input
                v-model="loginForm.loginName"
                size="small"
                auto-complete="off"
                placeholder="登录名"
                @keyup.enter.native="handleLogin"
              >
                <svg-icon slot="prefix" icon-class="user" class="el-input__icon input-icon" />
              </el-input>
            </el-form-item>

            <el-form-item prop="password">
              <el-input
                v-model="loginForm.password"
                size="small"
                :type="passwordType"
                auto-complete="off"
                placeholder="密码"
                @keyup.enter.native="handleLogin"
              >
                <i slot="suffix" class="el-icon-view el-input__icon" @click="showPwd" />
                <svg-icon slot="prefix" icon-class="password" class="el-input__icon input-icon" />
              </el-input>
            </el-form-item>
            <el-form-item prop="captchaValue">
              <el-row :span="24">
                <el-col :span="14">
                  <el-input
                    v-model="loginForm.captchaValue"
                    size="small"
                    auto-complete="off"
                    placeholder="验证码"
                    @keyup.enter.native="handleLogin"
                  >
                    <svg-icon slot="prefix" icon-class="validCode" class="el-input__icon input-icon" />
                  </el-input>
                </el-col>
                <el-col :span="10">
                  <div class="login-code">
                    <img
                      :src="captchaUrl"
                      class="login-code-img"
                      @click="refreshCode"
                    >
                  </div>
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item>
              <el-button
                :loading="loading"
                type="primary"
                size="small"
                class="login-submit"
                @click.native.prevent="handleLogin"
              >登录
              </el-button>
            </el-form-item>
          </el-form>
          <div class="login-menu">
            <!-- <a href="#" @click.stop="activeName='user'">{{ $t('login.userLogin') }}</a> -->
            <!--<a href="#" @click.stop="activeName='code'">{{ $t('login.phoneLogin') }}</a>-->
            <!-- <a href="#" @click.stop="activeName='third'">{{ $t('login.thirdLogin') }}</a> -->
          </div>
        </div>
      </div>
    </div>

    <el-dialog title="Or connect with" :visible.sync="showDialog">
      Can not be simulated on local, so please combine you own business simulation! ! !
      <br>
      <br>
      <br>
      <social-sign />
    </el-dialog>
  </div>
</template>

<script>
import SocialSign from './components/SocialSignin'
import bgImg from '@/assets/images/login-bg1.jpg'
import { getCaptcha } from '@/api/user'
export default {
  name: 'Login',
  components: { SocialSign },
  data() {
    return {
      loginForm: {
        loginName: 'admin',
        password: '123456',
        tenantCode: 'lzq',
        loginType: 1,
        captchaKey: '',
        captchaValue: ''
      },
      loginRules: {
        loginName: [{ required: true, message: '请输入登录名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
        tenantCode: [{ required: true, message: '请输入租户编码', trigger: 'blur' }],
        captchaValue: [{ required: true, message: '请输入验证码', trigger: 'blur' }]
      },
      bgImg: bgImg,
      captchaUrl: '',
      passwordType: 'password',
      capsTooltip: false,
      loading: false,
      showDialog: false,
      redirect: undefined,
      otherQuery: {}
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        const query = route.query
        if (query) {
          this.redirect = query.redirect
          this.otherQuery = this.getOtherQuery(query)
        }
      },
      immediate: true
    }
  },
  created() {
    // window.addEventListener('storage', this.afterQRScan)
    this.refreshCode()
  },
  mounted() {
    if (this.loginForm.loginName === '') {
      this.$refs.loginName.focus()
    } else if (this.loginForm.password === '') {
      this.$refs.password.focus()
    }
  },
  destroyed() {
    // window.removeEventListener('storage', this.afterQRScan)
  },
  methods: {
    checkCapslock(e) {
      const { key } = e
      this.capsTooltip = key && key.length === 1 && (key >= 'A' && key <= 'Z')
    },
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      // this.$nextTick(() => {
      //   this.$refs.password.focus()
      // })
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm)
            .then(() => {
              this.$router.push({ path: this.redirect || '/', query: this.otherQuery })
              this.loading = false
            })
            .catch(() => {
              this.loading = false
            })
          setTimeout(() => {			
            this.loading = false
          }, 5000)		// 设置时间，这里是5秒
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    getOtherQuery(query) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {})
    },
    refreshCode() {
      getCaptcha().then(response => {
        this.loginForm.captchaKey = response.key
        this.captchaUrl = response.captchaUrl
      })
    }
    // afterQRScan() {
    //   if (e.key === 'x-admin-oauth-code') {
    //     const code = getQueryObject(e.newValue)
    //     const codeMap = {
    //       wechat: 'code',
    //       tencent: 'code'
    //     }
    //     const type = codeMap[this.auth_type]
    //     const codeName = code[type]
    //     if (codeName) {
    //       this.$store.dispatch('LoginByThirdparty', codeName).then(() => {
    //         this.$router.push({ path: this.redirect || '/' })
    //       })
    //     } else {
    //       alert('第三方登录失败')
    //     }
    //   }
    // }
  }
}
</script>

<style lang="scss">

  .login-container {
  display: flex;
  align-items: center;
  position: relative;
  width: 100%;
  height: 100%;
  margin: 0 auto;
  background-size: 100% 100%;
}

.logo {
  position: fixed;
  top: 40px;
  left: 40px;
}

.login-weaper {
  display: flex;
  justify-content: space-around;
  width: 100%;
  .el-input-group__append {
    border: none;
  }
}

.login-left,
.login-border {
  position: relative;
  min-height: 400px;
  align-items: center;
  display: flex;
}

.login-left {
  flex-direction: column;
  color: white;
  align-items: baseline;
  justify-content: center;
  .bold {
    font-weight: bold;
  }
  .t1,.t2 {
    font-size: 48px;
  }
  .line {
    height: 3px;
    width: 100px;
    background-image: linear-gradient(to right, #ffffffff, #ffffff00);
  }
  .t3 {
    margin-top: 40px;
    font-size: 22px;
  }
}

.login-left .img {
  width: 140px;
}

.login-time {
  position: absolute;
  left: 25px;
  top: 25px;
  width: 100%;
  color: #fff;
  font-weight: 200;
  opacity: 0.9;
  font-size: 18px;
  overflow: hidden;
}

.login-left .title {
  margin-top: 60px;
  text-align: center;
  color: #fff;
  font-weight: 300;
  letter-spacing: 2px;
  font-size: 25px;
}

.login-border {
  border-left: none;
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  color: #fff;
  background-color: #fff;
  width: 400px;
  box-shadow: -4px 5px 10px rgba(0, 0, 0, 0.4);
  // float: left;
  box-sizing: border-box;
}

.login-main {
  margin: 0 auto;
  width: 81%;
  box-sizing: border-box;
}

.login-main > h3 {
  margin-bottom: 20px;
}

.login-main > p {
  color: #76838f;
}

.login-title {
  color: #333;
  margin-bottom: 40px;
  font-weight: 500;
  font-size: 22px;
  text-align: center;
  letter-spacing: 4px;
}

.login-menu {
  margin-top: 40px;
  width: 100%;
  text-align: center;

  a {
    color: #999;
    font-size: 12px;
    margin: 0px 8px;
  }
}

.login-submit {
  width: 100%;
  height: 45px;
  border: 1px solid #409EFF;
  background: none;
  font-size: 18px;
  letter-spacing: 2px;
  font-weight: 300;
  color: #409EFF;
  cursor: pointer;
  margin-top: 30px;
  font-family: "neo";
  transition: 0.25s;
}

.login-form {
  margin: 10px 0;

  i {
    color: #333;
  }

  .el-form-item__content {
    width: 100%;
  }

  .el-form-item {
    margin-bottom: 12px;
  }

  .el-input {
    input {
      padding-bottom: 10px;
      text-indent: 5px;
      background: transparent;
      border: none;
      border-radius: 0;
      color: #333;
      border-bottom: 1px solid rgb(235, 237, 242);
    }

    .el-input__prefix {
      i {
        padding: 0 5px;
        font-size: 16px !important;
      }
    }
  }
}

.login-code {
  display: flex;
  align-items: center;
  justify-content: space-around;
  margin: 0 0 0 10px;
}

.login-code-img {
  margin-top: 2px;
  width: 120px;
  height: 38px;
  background-color: #fdfdfd;
  color: #333;
  font-size: 14px;
  font-weight: bold;
  letter-spacing: 5px;
  line-height: 38px;
  text-indent: 5px;
  text-align: center;
  cursor:pointer!important;
}

</style>

// <style lang="scss">
// /* 修复input 背景不协调 和光标变色 */
// /* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

// $bg:#283443;
// $light_gray:#fff;
// $cursor: #fff;

// @supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
//   .login-container .el-input input {
//     color: $cursor;
//   }
// }

// /* reset element-ui css */
// .login-container {
//   .el-input {
//     display: inline-block;
//     height: 47px;
//     width: 85%;

//     input {
//       background: transparent;
//       border: 0px;
//       -webkit-appearance: none;
//       border-radius: 0px;
//       padding: 12px 5px 12px 15px;
//       color: $light_gray;
//       height: 47px;
//       caret-color: $cursor;

//       &:-webkit-autofill {
//         box-shadow: 0 0 0px 1000px $bg inset !important;
//         -webkit-text-fill-color: $cursor !important;
//       }
//     }
//   }

//   .el-form-item {
//     border: 1px solid rgba(255, 255, 255, 0.1);
//     background: rgba(0, 0, 0, 0.1);
//     border-radius: 5px;
//     color: #454545;
//   }
// }
// </style>

// <style lang="scss" scoped>
// $bg:#2d3a4b;
// $dark_gray:#889aa4;
// $light_gray:#eee;

// .login-container {
//   min-height: 100%;
//   width: 100%;
//   background-color: $bg;
//   overflow: hidden;

//   .login-form {
//     position: relative;
//     width: 520px;
//     max-width: 100%;
//     padding: 160px 35px 0;
//     margin: 0 auto;
//     overflow: hidden;
//   }

//   .tips {
//     font-size: 14px;
//     color: #fff;
//     margin-bottom: 10px;

//     span {
//       &:first-of-type {
//         margin-right: 16px;
//       }
//     }
//   }

//   .svg-container {
//     padding: 6px 5px 6px 15px;
//     color: $dark_gray;
//     vertical-align: middle;
//     width: 30px;
//     display: inline-block;
//   }

//   .title-container {
//     position: relative;

//     .title {
//       font-size: 26px;
//       color: $light_gray;
//       margin: 0px auto 40px auto;
//       text-align: center;
//       font-weight: bold;
//     }
//   }

//   .show-pwd {
//     position: absolute;
//     right: 10px;
//     top: 7px;
//     font-size: 16px;
//     color: $dark_gray;
//     cursor: pointer;
//     user-select: none;
//   }

//   .thirdparty-button {
//     position: absolute;
//     right: 0;
//     bottom: 6px;
//   }

//   @media only screen and (max-width: 470px) {
//     .thirdparty-button {
//       display: none;
//     }
//   }
// }
// </style>
