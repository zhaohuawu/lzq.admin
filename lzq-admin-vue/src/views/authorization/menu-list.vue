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
      <el-form ref="dataForm" :rules="temp.isBranch ? branchRules : menuRules" :model="temp" label-position="left" label-width="80px" style="">
        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="菜单编号" prop="code">
              <el-input v-model="temp.code" />
            </el-form-item> 
          </el-col>
          <el-col :span="2" />
          <el-col :span="11">
            <el-form-item label="菜单名称" prop="name">
              <el-input v-model="temp.name" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="排序" prop="rank">
              <el-input-number v-model="temp.rank" :min="0" />
            </el-form-item>
          </el-col>
          <el-col :span="2" />
          <el-col :span="11">
            <el-form-item label="图标" prop="icon">
              <el-input v-model="temp.icon" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="父菜单" prop="parentId">
              <el-select v-model="temp.parentId" class="filter-item" filterable clearable>
                <el-option v-for="item in branchMenus" :key="item.id" :label="item.name" :value="item.id" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="2" />
          <el-col :span="11">
            <el-form-item label="组件路径" prop="path">
              <el-input v-model="temp.path" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row type="flex" class="row-bg">
          <el-col :span="11">
            <el-form-item label="是否隐藏">
              <el-radio-group v-model="temp.isHidden">
                <el-radio :label="true">是</el-radio>
                <el-radio :label="false">否</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="2" />
          <el-col :span="11">
            <el-form-item label="菜单容器">
              <el-radio-group v-model="temp.isBranch">
                <el-radio :label="true">是</el-radio>
                <el-radio :label="false">否</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <div v-if="temp.isBranch === false">
          <el-form-item label="菜单路径" prop="url">
            <el-input v-model="temp.url" />
          </el-form-item>
          <el-form-item label="权限策略" prop="policy">
            <el-input v-model="temp.policy" width="100%" placeholder="规则：{容器编号}.{菜单编码}，示例，主数据公司菜单的Policy为：MasterData.Company" />
        
          </el-form-item>
        </div>
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
import { fetchList, getMenuList, createMenu, updateMenu, deleteMenu } from '@/api/menu'
import TreeTable from '@/components/TreeTable'
// import Pagination from '@/components/Pagination'
import store from '@/store'

export default {
  name: 'MenuList',
  components: { TreeTable },
  filters: {
    isHiddenFilter(isHidden) {
      const map = {
        true: '是',
        false: '否'
      }
      return map[isHidden]
    }
  },
  data() {
    return {
      isCanAdd: (store.getters.superAdmin || store.getters.permissions.indexOf('Authorization.MenuList:Operation.Create') > -1),
      list: [],
      tableCols: [
        { label: '菜单名称', prop: 'name', align: 'left', width: '180' },
        { label: '菜单编码', prop: 'code', align: 'center', width: '180' },
        { label: '权限策略', prop: 'policy', align: 'center', width: 'auto' },
        { label: 'Icon', prop: 'Icon', type: 'Icon', align: 'center', width: '60' },
        { label: '是否隐藏', prop: 'isHidden', type: 'Filter', align: 'center', width: '80', filters: {
          true: '是',
          false: '否'
        }},
        { label: '排序', prop: 'rank', align: 'center', width: '80' },
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
        icon: '',
        parentId: '',
        isBranch: true,
        isHidden: true,
        url: '',
        policy: '',
        path: '',
        moduleId: '0cbba6df-9078-4f4a-9b9c-8d6312d1b184'
      },
      branchMenus: null,
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑菜单',
        create: '新增菜单'
      
      },
      menuRules: {
        code: [{ required: true, message: '菜单编码必填', trigger: 'blur' }],
        name: [{ required: true, message: '菜单名称必填', trigger: 'blur' }],
        path: [{ required: true, message: '组件路径必填', trigger: 'blur' }],
        url: [{ required: true, message: '菜单路径必填', trigger: 'blur' }],
        policy: [{ required: true, message: '权限策略必填', trigger: 'blur' }]
      },
      branchRules: {
        code: [{ required: true, message: '菜单编码必填', trigger: 'blur' }],
        name: [{ required: true, message: '菜单名称必填', trigger: 'blur' }],
        path: [{ required: true, message: '组件路径必填', trigger: 'blur' }]
      },
      deleteText: null,
      dialogDeleteVisible: false
    }
  },
  created() {
    this.getList()
    this.getMenuList()
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
      getMenuList('Branch').then(response => {
        this.branchMenus = response
      })
    },
    handleFilter() {
      this.listQuery.page = 1
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
      }
    },
    cancelEdit(row) {
      row.title = row.originalTitle
      row.edit = false
      this.$message({
        message: 'The title has been restored to the original value',
        type: 'warning'
      })
    },
    confirmEdit(row) {
      row.edit = false
      row.originalTitle = row.title
      this.$message({
        message: 'The title has been edited',
        type: 'success'
      })
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
          this.temp.moduleId = '0cbba6df-9078-4f4a-9b9c-8d6312d1b184'
          createMenu(this.temp).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '新增菜单成功',
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
          this.temp.moduleId = '0cbba6df-9078-4f4a-9b9c-8d6312d1b184'
          const tempData = Object.assign({}, this.temp)
          updateMenu(tempData).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '修改菜单成功',
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
      this.deleteText = '您确认要删除该菜单：' + row.name
    },
    deleteData() {
      deleteMenu(this.temp.id).then(() => {
        this.dialogDeleteVisible = false
        this.$notify({
          title: 'Success',
          message: '删除菜单成功',
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
