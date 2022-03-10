import { login, logout, getInfo } from '@/api/user'
import { getMenuRouterList, getGrantedPermissions } from '@/api/authorize'

import { getToken, setToken, removeToken } from '@/utils/auth'
import router, { resetRouter } from '@/router'

const state = {
  token: getToken(),
  name: '',
  superAdmin: false,
  avatar: '',
  introduction: '',
  roles: [],
  menus: [],
  permissions: []
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_SUPERADMIN: (state, superAdmin) => {
    state.superAdmin = superAdmin
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_MENUS: (state, menus) => {
    state.menus = menus
  },
  SET_PERMISSIONS: (state, permissions) => {
    state.permissions = permissions
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { loginName, password, tenantCode, loginType } = userInfo
    return new Promise((resolve, reject) => {
      login({ loginName: loginName.trim(), password: password, tenantCode: tenantCode, loginType: loginType }).then(response => {
        const aToken = response.tokenType + ' ' + response.accessToken
        console.log('SET_TOKEN', aToken)
        commit('SET_TOKEN', aToken)
        setToken(aToken)
        resolve()
      }).catch(error => {
        console.log('loginError', error)
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        const data = {
          roles: ['admin'],
          introduction: 'I am a super administrator',
          avatar: response.headImgLink,
          name: response.userName,
          superAdmin: response.superAdmin
        }
        if (!data) {
          reject('Verification failed, please Login again.')
        }
        const { roles, name, avatar, introduction, superAdmin } = data

        // roles must be a non-empty array
        if (!roles || roles.length <= 0) {
          reject('getInfo: roles must be a non-null array!')
        }

        commit('SET_ROLES', roles)
        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        commit('SET_INTRODUCTION', introduction)
        commit('SET_SUPERADMIN', superAdmin)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        removeToken()
        resetRouter()

        // reset visited views and cached views
        // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
        dispatch('tagsView/delAllViews', null, { root: true })

        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  getMenus({ commit }) {
    return new Promise((resolve, reject) => {
      getMenuRouterList().then(response => {
        const data = response
        if (!data) {
          reject('获取用户已授权菜单失败')
        }

        commit('SET_MENUS', data)
        resolve(data)
      }).catch(error => {
        console.log('menuError', error)
        reject(error)
      })
    })
  },

  getGrantedPermissions({ commit }) {
    return new Promise((resolve, reject) => {
      getGrantedPermissions().then(response => {
        const data = response
        if (!data) {
          reject('获取用户已授权菜单失败')
        }

        commit('SET_PERMISSIONS', data)
        resolve(data)
      }).catch(error => {
        console.log('permissionError', error)
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      resolve()
    })
  },

  // dynamically modify permissions
  changeRoles({ commit, dispatch }, role) {
    return new Promise(async resolve => {
      const token = role + '-token'

      commit('SET_TOKEN', token)
      setToken(token)

      const { roles } = await dispatch('getInfo')

      resetRouter()

      // generate accessible routes map based on roles
      const accessRoutes = await dispatch('permission/generateRoutes', roles, { root: true })

      // dynamically add accessible routes
      router.addRoutes(accessRoutes)

      // reset visited views and cached views
      dispatch('tagsView/delAllViews', null, { root: true })

      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
