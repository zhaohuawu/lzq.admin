import request from '@/utils/request'

export function login(data) {
  return request({
    url: 'api/app/auth/login',
    method: 'post',
    data
    // baseURL: 'http://localhost:5015/'
  })
}

export function getInfo() {
  return request({
    url: '/api/app/sysUser/userInfo',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/api/app/auth/logOut',
    method: 'post'
  })
}

export function getSysUserList(query) {
  return request({
    url: '/api/app/sysUser/sysUserList',
    method: 'get',
    params: query
  })
}

export function createSysUser(data) {
  return request({
    url: '/api/app/sysUser/sysUser',
    method: 'post',
    data
  })
}

export function updateSysUser(data) {
  return request({
    url: '/api/app/sysUser/editSysUser',
    method: 'post',
    data
  })
}

export function updateSysUserStatus(data) {
  return request({
    url: '/api/app/sysUser/sysUserStatus',
    method: 'put',
    data
  })
}

export function deleteSysUser(id, concurrencyStamp) {
  return request({
    url: `/api/app/sysUser/user?id=${id}&concurrencyStamp=${concurrencyStamp}`,
    method: 'delete'
  })
}

export function updateSysUserPassword(data) {
  return request({
    url: '/api/app/sysUser/editUserPassword',
    method: 'post',
    data
  })
}

export function getDefaultAvatar() {
  return request({
    url: '/api/app/sysUser/defaultAvatar',
    method: 'get'
  })
}
