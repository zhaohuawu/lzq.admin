<template>
  <div class="app-container">
    <el-row :gutter="20">
      <el-col :span="6" :xs="24">
        <el-input
          v-model="filterText"
          placeholder="请输入公司/部门名称"
          clearable
          size="small"
          prefix-icon="el-icon-search"
        />
        <el-tree
          ref="qtree"
          :data="companyAndDepts"
          :props="defaultProps"
          :expand-on-click-node="false"
          :filter-node-method="filterNode"
          default-expand-all
          highlight-current
          @node-click="handleNodeClick"
        />
      </el-col>
      <el-col :span="18" :xs="24">
        <div class="filter-container">
          <el-input v-model="filtersQuery.userName" placeholder="用户名称" style="width: 200px;" clearable class="filter-item" />
          <el-button v-waves class="filter-item" type="primary" @click="handleFilter">
            搜索
          </el-button>
          <LzqButton 
            text="新增"
            type="primary"
            policy="Infrastructure.SysUserList:Operation.Create"
            btnstyle="margin-left: 10px;display: inline-block;vertical-align: middle;margin-bottom: 10px;"
            :btnclick="handleCreate"
          />
        </div>
      
        <OperateTable
          size="medium"
          :is-selection="false"
          :is-index="false"
          :is-border="false"
          :is-handle="true"
          :table-data="list" 
          :table-cols="tableCols"
          @operationSelect="operationSelect"
          @sortChange="sortChange"
        />
        <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.take" @pagination="getList" />
      </el-col>
    </el-row>
    <!-- 新增、编辑 -->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="600px">
      <el-form ref="dataForm" :rules="menuRules" :model="temp" label-position="center" label-width="80px">
        <el-form-item label="上级公司" prop="parentId">
          <tree-select
            ref="utree"
            v-model="cParentId"
            :tree-data="companyAndDepts"
            :props-tree="defaultProps"
            :node-key="'id'"
            :clearable="true"
            :searchable="true"
            placeholder="选择上级公司/部门"
            @change="searchSelectChain"
          />
        </el-form-item>
        <el-form-item label="登录账号" prop="loginName">
          <el-input v-model="temp.loginName" />
        </el-form-item>
        <el-form-item label="账号名称" prop="userName">
          <el-input v-model="temp.userName" />
        </el-form-item>
        <div v-if="dialogStatus === 'create'">
          <el-form-item label="密码" prop="password">
            <el-input v-model="temp.password" type="password" autocomplete="off" />
          </el-form-item>
          <el-form-item label="确认密码" prop="surePassword">
            <el-input v-model="temp.surePassword" type="password" autocomplete="off" />
          </el-form-item>
        </div>
        <el-form-item label="角色名称" prop="roleId">
          <el-select v-model="temp.roleIds" multiple placeholder="请选择">
            <el-option
              v-for="item in roleList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="头像" prop="headImgUrl" style="">
          <el-avatar :size="100" :src="image" />
          <el-button type="primary" style="position:absolute;bottom:15px;margin-left: 20px;text-align:center;" icon="el-icon-upload" @click="imagecropperShow=true">
            Upload Avatar
          </el-button>
          
          <div style="float:left">
            <image-cropper
              v-show="imagecropperShow"
              :key="imagecropperKey"
              :width="300"
              :height="300"
              field="file"
              :url="httpWebApi"
              lang-type="en"
              @close="imagecropperShow=false"
              @crop-upload-success="cropSuccess"
            />
          </div>
        </el-form-item> -->
        <el-form-item label="手机号码" prop="mobile">
          <el-input v-model="temp.mobile" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="temp.email" />
        </el-form-item>
        <el-form-item label="性别">
          <el-radio-group v-model="temp.sex">
            <el-radio label="Secrecy">保密</el-radio>
            <el-radio label="Male">男</el-radio>
            <el-radio label="Female">女</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          确认
        </el-button>
      </div>
    </el-dialog>
 
    <!-- 删除 -->
    <el-dialog title="提示" :visible.sync="dialogDeleteVisible" width="30%">
      <span>{{ deleteText }}</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogDeleteVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteData()">确 定</el-button>
      </span>
    </el-dialog>

  </div>

</template>

<script>
import { getSysUserList, 
  createSysUser, 
  updateSysUser, 
  deleteSysUser, 
  updateSysUserStatus,
  getDefaultAvatar,
  get } from '@/api/user'
import { getEnableRoles } from '@/api/role'
import OperateTable from '@/components/OperateTable'
import Pagination from '@/components/Pagination'
// import ImageCropper from '@/components/ImageCropper'
import { getCompanyAndDeptList } from '@/api/organization'
import TreeSelect from '@/components/TreeSelect'
import LzqButton from '@/components/LzqButton'

export default {
  name: 'SysUserList',
  components: { OperateTable, Pagination, TreeSelect, LzqButton },
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.temp.surePassword !== '') {
          this.$nextTick(() => {
            this.$refs['dataForm'].validateField('surePassword')
          })
        }
        callback()
      }
    }
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.temp.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      companyAndDepts: [],
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      filterText: '',
      cParentId: '',
      tableCols: [
        { label: '状态', prop: 'statusName', align: 'center', width: '100', sortable: 'custom' },
        { label: '头像', prop: 'headImgLink', type: 'Image', align: 'center', width: '120', style: 'height: 100px' },
        { label: '登录账号', prop: 'loginName', align: 'center', width: '100', sortable: 'custom' },
        { label: '用户名', prop: 'userName', align: 'center', sortable: 'custom' },
        // { label: '角色名称', prop: 'roleName', align: 'center', width: '100' },
        // { label: '性别', prop: 'sexName', align: 'center', width: '60' },
        { label: '手机号码', prop: 'mobile', align: 'center' },
        // { label: '创建时间', prop: 'creationTime', type: 'DateTime', align: 'center', width: '150' },
        { label: '操作', type: 'Button', prop: 'operation', width: '120' }
      ],
      list: [],
      total: 0,
      listLoading: true,
      filtersQuery: {
        userName: null
      },
      listQuery: {
        requireTotalCount: true,
        page: 1,
        skip: 0,
        take: 10,
        filter: null,
        sort: null
      },
      pageFilter: [],
      roleList: [],
      temp: {
        id: undefined,
        loginName: '',
        userName: '',
        password: '',
        surePassword: '',
        roleIds: [],
        headImgUrl: null,
        sex: null,
        mobile: null,
        email: null,
        deptId: null,
        companyId: null
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑用户',
        create: '新增用户'
      },
      menuRules: {
        loginName: [{ required: true, message: '登录账号必填', trigger: 'blur' }],
        userName: [{ required: true, message: '用户名称必填', trigger: 'blur' }],
        password: [{ validator: validatePass, required: true, trigger: 'blur' }],
        surePassword: [{ validator: validatePass2, required: true, trigger: 'blur' }]
      },
      deleteText: null,
      dialogDeleteVisible: false,
      imagecropperShow: false,
      imagecropperKey: 0
    }
  },
  watch: {
    // 根据名称筛选部门树
    filterText(val) {
      this.$refs.qtree.filter(val)
    }
  },
  created() {
    var that = this
    this.getDeptList().then(function() {
      that.getList()
      that.getEnableRoles()
    })
  },
  methods: {
    async getDeptList() {
      await getCompanyAndDeptList().then(rsp => {
        this.companyAndDepts = rsp
      })
    },
    filterNode(value, data) {
      if (!value) return true
      return data.name.indexOf(value) !== -1
    },
    handleNodeClick(data, node) {
      this.pageFilter = []
      console.log('handleNodeClick2', data, node)
      const ids = []
      // 递归
      const dg = function(type, list) {
        if (list) {
          list.forEach(function(item) {
            if (type === 'Company' && item.type !== 'Company') {
              console.log(type, item.type)
            } else {
              ids.push(item.id)
              if (item.children) {
                dg(item.children)
              }
            }
          })
        }
      }

      ids.push(data.id)
      if (data.type === 'Dept') {
        dg('Dept', data.children)
        this.pageFilter.push(['deptId', 'in', ids.join(',')])
      } else {
        dg('Company', data.children)
        this.pageFilter.push(['companyId', 'in', ids.join(',')])
      }
      this.handleFilter()
    },
    searchSelectChain(obj) {
      this.temp.deptId = null
      this.temp.companyId = null
      if (obj.type === 'Company') {
        this.temp.companyId = obj.id
      } else {
        this.cParentId = obj.id
        this.temp.deptId = obj.id
        this.temp.companyId = obj.companyId
      }
    },
    getList() {
      this.listLoading = true
      this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.take
      const tempData = Object.assign({}, this.listQuery)
      getSysUserList(tempData).then(response => {
        this.list = response.data
        this.total = response.totalCount
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    getEnableRoles() {
      getEnableRoles().then(response => {
        this.roleList = response
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      const f = []
      if (this.filtersQuery.userName !== '' && this.filtersQuery.userName !== null) {
        f.push(['userName', 'contains', this.filtersQuery.userName])
      } 
      if (this.pageFilter.length > 0) {
        this.pageFilter.forEach((item, index) => {
          f.push(item)
        })
      }
      if (f.length > 0) {
        this.listQuery.filter = JSON.stringify(f)
      }
      this.getList()
    },
    operationSelect(operationName, row) {
      if (operationName === 'Edit') {
        this.handleUpdate(row)
      } else if (operationName === 'Delete') {
        this.handleDelete(row)
      } else if (operationName === 'EditStatus') {
        this.updateStatus(row)
      }
    },
    sortChange(data) {
      console.log('sortChange:', data)
      this.listQuery.page = 1
      if (data.order !== null) {
        this.listQuery.sort = `[{"selector":"${data.prop}","desc":${data.order === 'descending'}}]`
      } else {
        this.listQuery.sort = null
      }
      this.getList()
    },
    handleCreate() {
      console.log('handleCreate')
      this.temp = Object.assign({}, {}) // copy obj
      this.$nextTick(() => {
        this.$refs.utree.setTreeCurrentKey(null)
      })
      this.dialogStatus = 'create'
      this.getDefaultAvatar()
      // this.getEnableRoles()
      this.dialogFormVisible = true
      this.temp = Object.assign({}, null) // copy obj
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createSysUser(this.temp).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '新增成功',
              type: 'success',
              duration: 2000
            })
            this.getList()
          })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, {}) // copy obj
      const q = { id: row.id }
      get(q).then(response => {
        this.temp = Object.assign({}, response)
        this.$nextTick(() => {
          if (this.temp.deptId !== '' && this.temp.deptId !== null) {
            this.$refs.utree.setTreeCurrentKey(this.temp.deptId)
          } else if (this.temp.companyId !== '' && this.temp.companyId !== null) {
            this.$refs.utree.setTreeCurrentKey(this.temp.companyId)
          } else {
            this.$refs.utree.setTreeCurrentKey(null)
          }
        })
      })

      this.image = this.temp.headImgLink
      this.dialogStatus = 'update'
      // this.getEnableRoles()
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateSysUser(tempData).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
            this.getList()
          })
        }
      })
    },
    handleDelete(row) {
      this.temp = Object.assign({}, row)
      this.dialogDeleteVisible = true
      this.deleteText = '您确认要删除该用户：' + row.userName
    },
    deleteData() {
      deleteSysUser(this.temp.id, this.temp.concurrencyStamp).then(() => {
        this.dialogDeleteVisible = false
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
      })
    },
    updateStatus(row) {
      const tempData = {
        concurrencyStamp: row.concurrencyStamp,
        id: row.id
      }
      updateSysUserStatus(tempData).then(() => {
        this.$notify({
          title: 'Success',
          message: '修改成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
      })
    },
    cropSuccess(resData) {
      this.imagecropperShow = false
      this.imagecropperKey = this.imagecropperKey + 1
      console.log(resData)
      this.image = resData.src
      this.temp.headImgUrl = resData.path
    },
    getDefaultAvatar() {
      getDefaultAvatar().then(response => {
        this.image = response
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.app-container {
  .permission-tree {
    margin-bottom: 30px;
  }
}

.avatar {
    width: 200px;
    height: 200px;
    border-radius: 50%;
}

</style>
