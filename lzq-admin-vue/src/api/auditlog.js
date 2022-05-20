import request from '@/utils/request'

export function getList(query) {
  return request({
    url: '/api/app/auditLogAction/list',
    method: 'get',
    params: query
  })
}

export function getCurrentUserLogsList(query) {
  return request({
    url: '/api/app/auditLogAction/currentUserLogsList',
    method: 'get',
    params: query
  })
}
