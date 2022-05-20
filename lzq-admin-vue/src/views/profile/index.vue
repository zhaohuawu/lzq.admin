<template>
  <div class="app-container">
    <div v-if="user">
      <el-row :gutter="20">

        <el-col :span="6" :xs="24">
          <user-card :user="user" />
        </el-col>

        <el-col :span="18" :xs="24">
          <el-card>
            <el-tabs v-model="activeTab">
              <el-tab-pane label="登录历史" name="LoginLogs">
                <LoginLogs />
              </el-tab-pane>
              <el-tab-pane label="个人资料" name="account">
                <account :user-info="userInfo" />
              </el-tab-pane>
              <el-tab-pane label="修改密码" name="updatepassword">
                <UpdatePassword />
              </el-tab-pane>
            </el-tabs>
          </el-card>
        </el-col>

      </el-row>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import UserCard from './components/UserCard'
import LoginLogs from './components/LoginLogs'
import UpdatePassword from './components/UpdatePassword'
import Account from './components/Account'
import { getCurrentUserInfo } from '@/api/user'
export default {
  name: 'Profile',
  components: { UserCard, LoginLogs, UpdatePassword, Account },
  data() {
    return {
      user: {},
      activeTab: 'LoginLogs',
      userInfo: {
        email: '',
        headImgUrl: '',
        id: '',
        mobile: '',
        sex: '',
        userName: ''
      }
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'avatar',
      'roles'
    ])
  },
  created() {
    this.init()
  },
  methods: {
    async init() {
      this.user = {
        name: this.name,
        role: this.roles.join(' | '),
        email: 'admin@test.com',
        avatar: this.avatar
      }
      this.getUserInfo()
    },
    async getUserInfo() {
      getCurrentUserInfo().then(rsp => {
        this.userInfo = Object.assign({}, rsp)
      })
    }
    
  }
}
</script>
