/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const authorizeRouter = {
  path: '/authorization',
  component: Layout,
  redirect: 'noRedirect',
  name: 'authorization',
  meta: {
    title: '权限管理',
    icon: 'lock'
  },
  children: [
    {
      path: 'menu-list',
      component: () => import('@/views/authorization/menu-list'),
      name: 'MenuList',
      meta: { title: '菜单管理' }
    },
    {
      path: 'permission-list',
      component: () => import('@/views/authorization/permission-list'),
      name: 'PermissionList',
      meta: { title: '操作权限' }
    },
    {
      path: 'role-list',
      component: () => import('@/views/authorization/role-list'),
      name: 'RoleList',
      meta: { title: '角色管理' }
    }
  ]
}
export default authorizeRouter
