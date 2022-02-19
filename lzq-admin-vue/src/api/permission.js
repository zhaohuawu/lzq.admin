import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/app/permission/data',
    method: 'get',
    params: query
  })
}

export function getPermissionGroupList() {
  return request({
    url: '/api/app/permission/permissionGroup',
    method: 'get'
  })
}

export function createPermission(data) {
  return request({
    url: '/api/app/permission/permission',
    method: 'post',
    data
  })
}

export function updatePermission(data) {
  return request({
    url: '/api/app/permission/permission',
    method: 'put',
    data
  })
}

export function deletePermission(id) {
  return request({
    url: `/api/app/permission/permission?id=${id}`,
    method: 'delete'
  })
}
