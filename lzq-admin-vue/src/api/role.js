import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/app/role/roleList',
    method: 'get',
    params: query
  })
}

export function getEnableRoles() {
  return request({
    url: '/api/app/role/roles',
    method: 'get'
  })
}

export function createRole(data) {
  return request({
    url: '/api/app/role/role',
    method: 'post',
    data
  })
}

export function updateRole(data) {
  return request({
    url: '/api/app/role/role',
    method: 'put',
    data
  })
}

export function updateRoleStatus(data) {
  return request({
    url: '/api/app/role/roleStatus',
    method: 'put',
    data
  })
}

export function deleteRole(id) {
  return request({
    url: `/api/app/role/role?id=${id}`,
    method: 'delete'
  })
}

// 获取角色已授权功能
export function getRolePermissionDatas(roleId) {
  return request({
    url: `api/app/authorize/rolePermissionDatas/${roleId}`,
    method: 'get'
  })
}

// 功能授权
export function grantPermissions(data) {
  return request({
    url: '/api/app/authorize/grantPermissions',
    method: 'post',
    data
  })
}
