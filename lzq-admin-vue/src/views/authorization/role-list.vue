<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="filtersQuery.name" placeholder="角色名称" style="width: 200px;" clearable class="filter-item" />
      <el-button v-waves class="filter-item" type="primary" @click="handleFilter">
        搜索
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" @click="handleCreate">
        新增
      </el-button>
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
      @handleTagClose="handleTagClose"
    />
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.take" @pagination="getList" />
    
    <!-- 新增、编辑 -->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="750px">
      <el-form ref="dataForm" :rules="menuRules" :model="temp" label-position="left" label-width="80px">

        <el-form-item label="角色名称" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="temp.description"
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入内容"
          />
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

    <!-- 功能授权 -->
    <el-drawer ref="drawer" :title="roleName" :visible.sync="dialogRolePermissionVisible" direction="rtl">
      <div class="n-drawer-content">
      
        <el-tree
          ref="tree"
          :data="rolePermissionList"
          :default-checked-keys="haveCheckedKeys"
          :props="defaultProps"
          show-checkbox
          node-key="id"
          class="permission-tree n-drawer-body"
          :expand-on-click-node="false"
        >
          <span slot-scope="{ node, data }" class="custom-tree-node">
            <span>{{ node.label }}</span>
            <span v-if="data.type==='Menu'"><i class="el-icon-menu" /></span>
            <span v-if="data.type==='Permission'"><i class="el-icon-s-operation" /></span>
          </span>
          
        </el-tree>
        <div class="n-drawer-footer">
          <el-button @click="dialogRolePermissionVisible=false">取 消</el-button>
          <el-button type="primary" :loading="loading" @click="confirmRole">{{ loading ? '提交中 ...' : '确 定' }}</el-button>
        </div>
      </div>
    </el-drawer>

    <!-- <el-dialog :title="roleName" :visible.sync="dialogRolePermissionVisible">
      <el-tree
        ref="tree"
        :data="rolePermissionList"
        :default-checked-keys="haveCheckedKeys"
        :props="defaultProps"
        show-checkbox
        node-key="id"
        class="permission-tree"
        :expand-on-click-node="false"
      >
        <span slot-scope="{ node, data }" class="custom-tree-node">
          <span>{{ node.label }}</span>
          <span v-if="data.type==='Menu'"><i class="el-icon-menu" /></span>
          <span v-if="data.type==='Permission'"><i class="el-icon-s-operation" /></span>
        </span>
      
      </el-tree>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogRolePermissionVisible=false">取 消</el-button>
        <el-button type="primary" @click="confirmRole">确 定</el-button>
      </span>
    </el-dialog> -->
    
  </div>

</template>

<script>
import { fetchList, 
  createRole, 
  updateRole, 
  deleteRole, 
  updateRoleStatus,
  getRolePermissionDatas, 
  grantPermissions } from '@/api/role'
import { deleteuserDatePrivilege } from '@/api/authorize'
import OperateTable from '@/components/OperateTable'
import Pagination from '@/components/Pagination'

export default {
  name: 'RoleList',
  components: { OperateTable, Pagination },
  data() {
    return {
      list: null,
      tableCols: [
        { label: '状态', prop: 'roleStatusText', align: 'center', width: '60', key: 'id' },
        { label: '角色名称', prop: 'name', align: 'center', width: '100' },
        { label: '已授权用户', prop: 'haveAuthorizeUser', type: 'Tag', showProp: 'loginName', align: 'center' },
        { label: '描述', prop: 'description', align: 'center', width: '200' },
        { label: '操作', type: 'Button', prop: 'operation', width: '200' }
      ],
      total: 0,
      listLoading: true,
      filtersQuery: {
        name: null
      },
      listQuery: {
        requireTotalCount: true,
        page: 1,
        skip: 0,
        take: 10,
        filter: null
      },
      temp: {
        id: undefined,
        name: '',
        description: ''
      },
      branchMenus: null,
      permissionGroups: null,
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑角色',
        create: '新增角色'
      },
      menuRules: {
        name: [{ required: true, message: '菜单名称必填', trigger: 'blur' }]
      },
      deleteText: null,
      dialogDeleteVisible: false,
      rolePermissionList: null,
      dialogRolePermissionVisible: false,
      roleName: null,
      roleId: null,
      haveCheckedKeys: [],
      perissionIds: [],
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      permissionList: []
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.take
      const tempData = Object.assign({}, this.listQuery)
      fetchList(tempData).then(response => {
        this.list = response.data
        this.total = response.totalCount
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      console.log(this.filtersQuery.name)
      if (this.filtersQuery.name !== '' && this.filtersQuery.name !== null) {
        this.listQuery.filter = '[["name","contains","' + this.filtersQuery.name + '"]]'
      } else {
        this.listQuery.filter = null
      }
      this.getList()
    },
    operationSelect(operationName, row) {
      if (operationName === 'Edit') {
        this.handleUpdate(row)
      } else if (operationName === 'Delete') {
        this.handleDelete(row)
      } else if (operationName === 'FunctionAuthorization') {
        this.handleFunctionAuthorization(row)
      } else if (operationName === 'EditStatus') {
        this.updateRoleStatus(row)
      }
    },
    policyValue() {
      this.temp.policy = this.temp.permissionGroup + '.' + this.temp.code
    },
    resetTemp() {
      this.temp = {
        id: undefined
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createRole(this.temp).then(() => {
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
      this.temp = Object.assign({}, row) // copy obj
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateRole(tempData).then(() => {
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
      this.deleteText = '您确认要删除该操作权限：' + row.name
    },
    deleteData() {
      deleteRole(this.temp.id).then(() => {
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
    handleFunctionAuthorization(row) {
      this.roleName = '功能授权-' + row.name
      this.roleId = row.id
      this.dialogRolePermissionVisible = true
      this.listLoading = true
      getRolePermissionDatas(row.id).then(response => {
        this.rolePermissionList = response.rolePermissionTree
        this.haveCheckedKeys = response.haveCheckedKeys
        this.perissionIds = response.perissionIds
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    functionAuthorizationData() {
      grantPermissions(this.temp.id).then(() => {
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
    async confirmRole() {
      const checkedKeys = this.$refs.tree.getCheckedKeys()
      // 找出操作授权项
      const ckeckedPermissionIds = []
      checkedKeys.forEach(i => {
        if (this.perissionIds.indexOf(i) > -1) {
          ckeckedPermissionIds.push(i)
        }
      })
      console.log(checkedKeys, ckeckedPermissionIds)
      
      const postData = []
      if (this.haveCheckedKeys.length > 0) {
        // 已授权取消授权的项
        this.haveCheckedKeys.forEach(perId => {
          if (ckeckedPermissionIds.indexOf(perId) === -1) {
            const grantPer = {
              roleId: this.roleId,
              permissionId: perId,
              isGranted: false
            }
            postData.push(grantPer)
          }
        })

        // 未授权进行授权项
        ckeckedPermissionIds.forEach(perId => {
          if (this.haveCheckedKeys.indexOf(perId) === -1) {
            const grantPer = {
              roleId: this.roleId,
              permissionId: perId,
              isGranted: true
            }
            postData.push(grantPer)
          }
        })
      } else {
        ckeckedPermissionIds.forEach(perId => {
          const grantPer = {
            roleId: this.roleId,
            permissionId: perId,
            isGranted: true
          }
          postData.push(grantPer)
        })
      }
      await grantPermissions(postData)
      
      this.dialogRolePermissionVisible = false
      this.$notify({
        title: '成功',
        message: '授权成功',
        type: 'success'
      })
    },
    updateRoleStatus(row) {
      const tempData = {
        concurrencyStamp: row.concurrencyStamp,
        id: row.id
      }
      updateRoleStatus(tempData).then(() => {
        this.$notify({
          title: 'Success',
          message: '修改成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
      })
    },
    handleTagClose(tag, row) {
      deleteuserDatePrivilege(tag.userDataPrivilegeId).then(() => {
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
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

  .n-drawer-content{
  display: flex;
  flex-direction: column;
  height: 100%;

  .n-drawer-body{
    flex: 1;
  }

  .n-drawer-footer{
    display: flex;
  }
}
}
</style>
