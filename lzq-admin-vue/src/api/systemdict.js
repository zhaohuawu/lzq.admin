import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/app/systemDictionary/list',
    method: 'get',
    params: query
  })
}

export function createDict(data) {
  return request({
    url: '/api/app/systemDictionary/createDict',
    method: 'post',
    data
  })
}

export function createChildernDict(data) {
  return request({
    url: '/api/app/systemDictionary/create',
    method: 'post',
    data
  })
}

export function updateDict(data) {
  return request({
    url: '/api/app/systemDictionary/update',
    method: 'put',
    data
  })
}

export function updateDictStatus(id) {
  return request({
    url: '/api/app/systemDictionary/updateStatus?id=' + id,
    method: 'put'
  })
}

export function deleteDict(id) {
  return request({
    url: `/api/app/systemDictionary/delete?id=${id}`,
    method: 'delete'
  })
}

export function refreshDict() {
  return request({
    url: '/api/app/systemDictionary/refresh',
    method: 'post'
  })
}
