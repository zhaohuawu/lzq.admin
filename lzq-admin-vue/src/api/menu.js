import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/app/menu/menusList',
    method: 'get',
    params: query
  })
}

export function getMenuList(menuType) {
  return request({
    url: '/api/app/menu/menuList',
    method: 'get',
    params: { menuType }
  })
}

export function createMenu(data) {
  return request({
    url: '/api/app/menu/menu',
    method: 'post',
    data
  })
}

export function updateMenu(data) {
  return request({
    url: '/api/app/menu/menu',
    method: 'put',
    data
  })
}

export function deleteMenu(id) {
  return request({
    url: `/api/app/menu/menu?MenuId=${id}`,
    method: 'delete'
  })
}
