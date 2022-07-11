<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="filtersQuery.dictValue" placeholder="字典名称/编码" style="width: 200px;" clearable class="filter-item" />
      <el-button v-waves class="filter-item" type="primary" @click="handleFilter">
        搜索
      </el-button>
      <LzqButton 
        text="新增"
        type="primary"
        policy="Infrastructure.SystemDictionary:Operation.Create"
        btnstyle="margin-left: 10px;display: inline-block;vertical-align: middle;margin-bottom: 10px;"
        :btnclick="handleCreate"
      />
      <LzqButton 
        text="清除缓存"
        type="primary"
        policy="Infrastructure.SystemDictionary:Operation.Create"
        btnstyle="margin-left: 10px;display: inline-block;vertical-align: middle;margin-bottom: 10px;"
        :btnclick="refreshDict"
      />
      
    </div>
    <OperateTable
      size="medium"
      :is-index="false"
      :is-border="false"
      :is-handle="true"
      :table-data="list" 
      :table-cols="tableCols"
      @operationSelect="operationSelect"
    />
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.take" @pagination="getList" />
    
    <!-- 新增、编辑 -->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="750px">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="80px">

        <el-form-item label="字典名称" prop="dictValue">
          <el-input v-model="temp.dictValue" />
        </el-form-item>
        <el-form-item label="字典编码" prop="dictCode">
          <el-input v-model="temp.dictCode" />
        </el-form-item>
        <el-form-item label="描述" prop="remark">
          <el-input
            v-model="temp.remark"
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

    <el-drawer ref="drawer" title="字典配置" :visible.sync="drawerListVisible" direction="rtl" size="50%">
      <LzqButton 
        text="新增"
        type="primary"
        policy="Infrastructure.SystemDictionary:Operation.Create"
        btnstyle="margin-left: 10px;display: inline-block;vertical-align: middle;margin-bottom: 10px;"
        :btnclick="handleCreateChildren"
      />
      <OperateTable
        size="medium"
        :is-index="false"
        :is-border="false"
        :is-handle="true"
        :table-data="childrenList" 
        :table-cols="childrenTableCols"
        @operationSelect="operationSelect"
      />
    </el-drawer>

    <!-- 新增、编辑 -->
    <el-dialog :title="textMap[dialogChildrenStatus]" :visible.sync="dialogChildrenFormVisible" width="750px">
      <el-form ref="dataFormChildren" :rules="childrenRules" :model="childrenTemp" label-position="left" label-width="80px">
        <el-form-item label="字典编码" prop="dictCode">
          <el-input v-model="childrenTemp.dictCode" disabled />
        </el-form-item>
        <el-form-item label="字典键" prop="dictKey">
          <el-input v-model="childrenTemp.dictKey" />
        </el-form-item>
        <el-form-item label="字典值" prop="dictValue">
          <el-input v-model="childrenTemp.dictValue" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="childrenTemp.sort" controls-position="right" :min="1" />
        </el-form-item>
        <el-form-item label="描述" prop="remark">
          <el-input
            v-model="childrenTemp.remark"
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入内容"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogChildrenFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogChildrenStatus==='createChildren'?createChildrenData():updateChildrenData()">
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
import { fetchList, 
  createDict, 
  createChildernDict,
  updateDict, 
  deleteDict, 
  updateDictStatus,
  refreshDict } from '@/api/systemdict'
import OperateTable from '@/components/OperateTable'
import Pagination from '@/components/Pagination'
import LzqButton from '@/components/LzqButton'

export default {
  name: 'SystemDictList',
  components: { OperateTable, Pagination, LzqButton },
  data() {
    return {
      list: [],
      tableCols: [
        { label: '状态', prop: 'statusText', align: 'center', width: '60', key: 'id' },
        { label: '字典名称', prop: 'dictValue', align: 'center', width: '250' },
        { label: '字典编码', prop: 'dictCode', align: 'center', width: '250' },
        { label: '描述', prop: 'Remark', align: 'center' },
        { label: '操作', type: 'Button', prop: 'operation', width: '200' }
      ],
      total: 0,
      listLoading: true,
      filtersQuery: {
        dictCode: null,
        dictValue: null
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
        dictValue: '',
        dictCode: '',
        remark: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑字典',
        create: '新增字典',
        updateChildren: '编辑字典配置',
        createChildren: '新增字典配置'
      },
      rules: {
        dictCode: [{ required: true, message: '字典编码必填', trigger: 'blur' }],
        dictValue: [{ required: true, message: '字典名称必填', trigger: 'blur' }]
      },
      childrenList: [],
      childrenTableCols: [
        { label: '状态', prop: 'statusText', align: 'center', width: '50', key: 'id' },
        { label: '字典编码', prop: 'dictCode', align: 'center', width: '150' },
        { label: '字典键', prop: 'dictKey', align: 'center', width: '150' },
        { label: '字典值', prop: 'dictValue', align: 'center', width: '150' },
        { label: '排序', prop: 'sort', align: 'center', width: '50' },
        { label: '描述', prop: 'Remark', align: 'center' },
        { label: '操作', type: 'Button', prop: 'operation', width: '150' }
      ],
      childrenTemp: {
        id: undefined,
        dictValue: '',
        dictCode: '',
        dictKey: '',
        sort: 0,
        remark: ''
      },
      childrenRules: {
        dictCode: [{ required: true, message: '字典编码必填', trigger: 'blur' }],
        dictKey: [{ required: true, message: '字典键必填', trigger: 'blur' }],
        dictValue: [{ required: true, message: '字典值必填', trigger: 'blur' }]
      },
      maxSort: 0,
      drawerListVisible: false,
      dialogChildrenFormVisible: false,
      dialogChildrenStatus: '',
      deleteText: null,
      dialogDeleteVisible: false
    }
  },
  created() {
    this.handleFilter()
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
      var filterArray = []
      
      if (this.filtersQuery.dictValue !== '' && this.filtersQuery.dictValue !== null) {
        filterArray.push('["dictValue","contains","' + this.filtersQuery.dictValue + '"]')
        filterArray.push('["dictCode","contains","' + this.filtersQuery.dictValue + '","or"]')
      } 
      filterArray.push('["dictKey","=","ParentDict"]')

      if (filterArray.length > 0) {
        this.listQuery.filter = '[' + filterArray.join(',') + ']'
      }
      this.getList()
    },
    operationSelect(operationName, row) {
      if (operationName === 'Modify') {
        if (row.dictKey === 'ParentDict') {
          this.handleUpdate(row)
        } else {
          this.handleUpdateChildren(row)
        }
      } else if (operationName === 'Delete') {
        this.handleDelete(row)
      } else if (operationName === 'DisOrEnable') {
        this.updateStatus(row)
      } else if (operationName === 'DictConfig') {
        this.handleDictConfig(row)
      }
    },
    resetTemp() {
      this.temp = {
        id: undefined
      }
      this.childrenTemp = {
        id: undefined,
        dictCode: this.childrenTemp.dictCode,
        sort: this.maxSort + 1
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
          createDict(this.temp).then(() => {
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
      console.log(this.temp)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      console.log(this.dialogFormVisible, this.temp)
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateDict(tempData).then(() => {
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
    handleCreateChildren() {
      this.resetTemp()
      this.dialogChildrenStatus = 'createChildren'
      this.dialogChildrenFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataFormChildren'].clearValidate()
      })
    },
    createChildrenData() {
      this.$refs['dataFormChildren'].validate((valid) => {
        if (valid) {
          createChildernDict(this.childrenTemp).then(() => {
            this.dialogChildrenFormVisible = false
            this.$notify({
              title: 'Success',
              message: '新增成功',
              type: 'success',
              duration: 2000
            })
            this.handleDictConfig(this.childrenTemp)
          })
        }
      })
    },
    handleUpdateChildren(row) {
      this.childrenTemp = Object.assign({}, row) // copy obj
      console.log(this.childrenTemp)
      this.dialogChildrenStatus = 'updateChildren'
      this.dialogChildrenFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataFormChildren'].clearValidate()
      })
    },
    updateChildrenData() {
      this.$refs['dataFormChildren'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.childrenTemp)
          updateDict(tempData).then(() => {
            this.dialogChildrenFormVisible = false
            this.$notify({
              title: 'Success',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
            this.handleDictConfig(this.childrenTemp)
          })
        }
      })
    },
    handleDelete(row) {
      this.temp = Object.assign({}, row)
      this.dialogDeleteVisible = true
      this.deleteText = '您确认要删除该字典：' + row.dictValue
    },
    deleteData() {
      deleteDict(this.temp.id).then(() => {
        this.dialogDeleteVisible = false
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
        const t = {
          dictCode: this.childrenTemp.dictCode
        }
        this.handleDictConfig(t)
      })
    },
    updateStatus(row) {
      updateDictStatus(row.id).then(() => {
        this.$notify({
          title: 'Success',
          message: '修改成功',
          type: 'success',
          duration: 2000
        })
        if (row.dictKey === 'ParentDict') {
          this.getList()
        } else {
          const t = {
            dictCode: row.dictCode
          }
          this.handleDictConfig(t)
        }
      })
    },
    handleDictConfig(row) {
      this.resetTemp()
      this.childrenTemp.dictCode = row.dictCode
      this.drawerListVisible = true
      
      const tempData = {
        filter: `[["dictCode","=","${row.dictCode}"]]`,
        sort: '[{"selector":"sort","desc":false}]'
      }
      fetchList(tempData).then(response => {
        this.childrenList = response.data.filter(item => { return item.dictKey !== 'ParentDict' })
        if (this.childrenList.length > 0) {
          this.maxSort = this.childrenList[this.childrenList.length - 1].sort
        } else {
          this.maxSort = 0
        }
      })
    },
    refreshDict() {
      refreshDict().then(() => {
        this.$notify({
          title: 'Success',
          message: '清除缓存成功',
          type: 'success',
          duration: 2000
        })
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
