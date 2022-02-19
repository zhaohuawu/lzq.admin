import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import Element from 'element-ui'
import './styles/element-variables.scss'

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import './icons' // icon
import './permission' // permission control
import './utils/error-log' // error log

import * as filters from './filters' // global filters
import vueWaves from './directive/waves'
Vue.use(vueWaves)

Vue.directive('enterNumber', {
  bind: function(el, { value = 2 }) {
    el = el.nodeName === 'INPUT' ? el : el.children[0]
    var RegStr = value === 0 ? `^[\\+\\-]?\\d+\\d{0,0}` : `^[\\+\\-]?\\d+\\.?\\d{0,${value}}`
    el.addEventListener('keyup', function() {
      el.value = el.value.match(new RegExp(RegStr, 'g'))
      el.dispatchEvent(new Event('input'))
    })
  }
})

Vue.directive('focus', {
  // 当被绑定的元素插入到 DOM 中时……
  inserted: function(el) {
    el.focus()
  }
})

Vue.use(Element, {
  size: Cookies.get('size') || 'medium' // set element-ui default size
})

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
