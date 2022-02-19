import request from '@/utils/request'

export function deleteuserDatePrivilege(userDataPrivilegeId) {
  return request({
    url: `/api/app/authorize/userRole?UserDataPrivilegeId=${userDataPrivilegeId}`,
    method: 'delete'
  })
}
