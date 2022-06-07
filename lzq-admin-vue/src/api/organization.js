import request from '@/utils/request'

export function getCompanyAndDeptList(query) {
  return request({
    url: '/api/app/dept/companyAndDeptList',
    method: 'get'
  })
}

export function createCompany(data) {
  return request({
    url: '/api/app/company/create',
    method: 'post',
    data
  })
}

export function createDept(data) {
  return request({
    url: '/api/app/dept/create',
    method: 'post',
    data
  })
}

export function updateCompany(data) {
  return request({
    url: '/api/app/company/update',
    method: 'put',
    data
  })
}

export function updateDept(data) {
  return request({
    url: '/api/app/dept/update',
    method: 'put',
    data
  })
}

export function deleteCompany(id) {
  return request({
    url: `/api/app/company/delete?id=${id}`,
    method: 'delete'
  })
}

export function deleteDept(id) {
  return request({
    url: `/api/app/dept/delete?id=${id}`,
    method: 'delete'
  })
}
