import request from '@/utils/request'

export function getMenuRouterList() {
  return request({
    url: '/api/app/authenticateChecker/grantedMenus',
    method: 'get'
  })
}

export function getGrantedPermissions() {
  return request({
    url: '/api/app/permissionChecker/grantedPermissions',
    method: 'get'
  })
}

export function deleteuserDatePrivilege(userDataPrivilegeId) {
  return request({
    url: `/api/app/authorize/userRole?userDataPrivilegeId=${userDataPrivilegeId}`,
    method: 'delete'
  })
}

