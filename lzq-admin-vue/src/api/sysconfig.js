import request from '@/utils/request'

export function getToken() {
  return request({
    url: '/qiniu/upload/token', // 假地址 自行替换
    method: 'get'
  })
}

export function getSysConfigInfo(configType, code) {
  return request({
    url: '/api/app/sysconfig/getInfo',
    method: 'get',
    params: { configType: configType, code: code }
  })
}

export function qnUpdate(data) {
  return request({
    url: '/api/app/sysconfig/qnupdate',
    method: 'post',
    data
  })
}
