import { constantRoutes, componentsMap } from '@/router'
// import { fetchList } from '@/api/menu'
import Layout from '@/layout'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param menus asyncRoutes
 */
export function filterAsyncMenus(menus) {
  const res = []
  menus.forEach(route => {
    const tmpmodel = { ...route }
    const tmp = {
      path: tmpmodel.path,
      name: tmpmodel.code,
      meta: {
        title: tmpmodel.name,
        icon: tmpmodel.icon
      },
      hidden: tmpmodel.isHidden
    }
    
    if (tmpmodel.isBranch) {
      tmp.component = Layout
      if (tmpmodel.path === '/') {
        tmp.redirect = '/dashboard'
      } else {
        tmp.redirect = 'noRedirect'
      }
      tmp.children = filterAsyncMenus(tmpmodel.children)
    } else { 
      var mapKey = tmpmodel.policy.replace('.', '_')
      tmp.component = componentsMap[mapKey]
    }
    res.push(tmp)
  })

  return res
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      res.push(tmp)
    }
  })

  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  }
}

const actions = {
  generateRoutes({ commit }, menus) {
    return new Promise(resolve => {
      const accessedRoutes = filterAsyncMenus(menus)
      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
