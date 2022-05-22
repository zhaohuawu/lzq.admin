<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="filtersQuery.name" placeholder="菜单名称" style="width: 200px;" clearable class="filter-item" />
      <el-button v-waves class="filter-item" type="primary" @click="handleFilter">
        搜索
      </el-button>
      <el-button v-if="isCanAdd" class="filter-item" style="margin-left: 10px;" type="primary" @click="handleCreate">
        新增
      </el-button>
    </div>

    <TreeTable
      size="medium"
      :is-selection="false"
      :is-index="false"
      :is-border="true"
      :is-handle="true"
      :table-data="list" 
      :table-cols="tableCols" 
      @operationSelect="operationSelect"
    />

    <!-- <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.take" @pagination="getList" /> -->

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="750">
      <el-form ref="dataForm" :rules="menuRules" :model="temp" label-position="left" label-width="80px" style="">
        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="权限编号" prop="code">
              <el-input v-model="temp.code" @input="policyValue" />
            </el-form-item> 
          </el-col>
          <el-col :span="2" />
          <el-col :span="11">
            <el-form-item label="权限名称" prop="name">
              <el-input v-model="temp.name" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="权限组" prop="permissionGroup">
              <el-select v-model="temp.permissionGroup" class="filter-item" clearable @change="policyValue">
                <el-option v-for="item in permissionGroups" :key="item.key" :label="item.value" :value="item.key" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="2" />
          <el-col :span="11">
            <el-form-item label="排序" prop="rank">
              <el-input-number v-model="temp.rank" :min="0" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="所属菜单" prop="menuId">
              <el-select v-model="temp.menuId" class="filter-item" filterable clearable>
                <el-option v-for="item in branchMenus" :key="item.id" :label="item.name" :value="item.id" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="权限策略" prop="policy">
          <el-input v-model="temp.policy" width="100%" placeholder="规则：{权限组编号}.{权限编号}，示例：页面访问为:View.Access，编辑为：Operation.Edit" />
        
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

    <el-dialog
      title="提示"
      :visible.sync="dialogDeleteVisible"
      width="30%"
    >
      <span>{{ deleteText }}</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogDeleteVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteData()">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import { fetchList, getPermissionGroupList, createPermission, updatePermission, deletePermission } from '@/api/permission'
import { getMenuList } from '@/api/menu'
import TreeTable from '@/components/TreeTable'
// import Pagination from '@/components/Pagination'
import store from '@/store'

export default {
  name: 'PermissionList',
  components: { TreeTable },
  data() {
    return {
      isCanAdd: (store.getters.superAdmin || store.getters.permissions.indexOf('Authorization.PermissionList:Operation.Create') > -1),
      list: [],
      tableCols: [
        { label: '菜单名称', prop: 'menuName', align: 'left', width: '180' },
        { label: '权限名称', prop: 'name', align: 'center', width: '200' },
        { label: '权限编码', prop: 'code', align: 'center', width: '120' },
        { label: '排序', prop: 'rank', align: 'center', width: '80' },
        { label: '权限策略', prop: 'policy', align: 'center', width: '200' },
        { label: '验证权限策略', prop: 'actualPolicy', align: 'center' },
        { label: '操作', type: 'Button', prop: 'operation', align: 'center', width: '200' }
      ],
      total: 0,
      columns: null,
      listLoading: true,
      filtersQuery: {
        name: null
      },
      listQuery: {
        requireTotalCount: true,
        page: 1,
        skip: 0,
        take: 0,
        filter: null
      },
      temp: {
        id: undefined,
        code: '',
        name: '',
        rank: 0,
        menuId: '',
        permissionGroup: 'View',
        policy: '',
        concurrencyStamp: ''
      },
      branchMenus: null,
      permissionGroups: null,
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑操作权限',
        create: '新增操作权限'
      },
      menuRules: {
        code: [{ required: true, message: '菜单编码必填', trigger: 'blur' }],
        name: [{ required: true, message: '菜单名称必填', trigger: 'blur' }],
        menuId: [{ required: true, message: '所属菜单必填', trigger: 'blur' }],
        permissionGroup: [{ required: true, message: '权限组必填', trigger: 'blur' }],
        rank: [{ required: true, message: '排序必填', trigger: 'blur' }],
        policy: [{ required: true, message: '权限策略必填', trigger: 'blur' }]
      },
      deleteText: null,
      dialogDeleteVisible: false
    }
  },
  created() {
    this.getList()
    this.getMenuList()
    this.getPermissionGroupList()
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
    getMenuList() {
      getMenuList('Menu').then(response => {
        this.branchMenus = response
      })
    },
    getPermissionGroupList() {
      getPermissionGroupList().then(response => {
        this.permissionGroups = response
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      console.log(this.filtersQuery.name)
      if (this.filtersQuery.name !== '' && this.filtersQuery.name !== null) {
        this.listQuery.filter = '[["menuName","contains","' + this.filtersQuery.name + '"]]'
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
          createPermission(this.temp).then(() => {
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
          updatePermission(tempData).then(() => {
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
      deletePermission(this.temp.id).then(() => {
        this.dialogDeleteVisible = false
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

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
