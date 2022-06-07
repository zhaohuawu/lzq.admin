<template>
  <div class="app-container">
    <el-row :gutter="20">
      <el-col :span="8" :xs="24">
        <el-card style="margin-bottom:20px;">
          <div slot="header" class="clearfix o-el-header">
            <el-button v-if="isCanAdd" class="filter-item" style="margin-right: 30px;" type="primary" @click="handleCreate('Company')">新增公司</el-button>
            <el-button v-if="isCanAdd" class="filter-item" style="margin-right: 30px;" type="primary" @click="handleCreate('Dept')">新增部门</el-button>
          </div>
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
        </el-card>
      </el-col>
      <el-col :span="16" :xs="24">
        <el-card style="margin-bottom:20px;">
          <div slot="header" class="clearfix o-el-header">
            <span class="o-span">{{ title }}</span>
            <div class="btn-eb">
              <el-button v-if="isCanUpdate && !isUpdating" class="filter-item" style="margin-right: 20px;" type="primary" @click="isUpdating=true">修改</el-button>
              <el-button v-if="isCanDelete && !isUpdating" class="filter-item" style="margin-right: 10px;" type="primary" @click="handleDelete">删除</el-button>
              <el-button v-if="isCanUpdate && isUpdating" class="filter-item" style="margin-right: 20px;" type="primary" @click="isUpdating=false">取消</el-button>
              <el-button v-if="isCanUpdate && isUpdating" class="filter-item" style="margin-right: 20px;" type="primary" @click="handleUpdate">保存</el-button>
            </div>
          </div>
          <el-form ref="udataForm" :model="udept" label-position="center" label-width="80px">
  
            <el-form-item :label="udept.type==='Dept'&&udept.parentId!==null&&udept.parentId!==''?'上级部门':'所属公司'" prop="parentId">
              <tree-select
                ref="utree"
                v-model="uParentId"
                :tree-data="companyAndDepts"
                :props-tree="defaultProps"
                :node-key="'id'"
                :clearable="true"
                :searchable="true"
                :disabled="isUpdating?false:true"
                placeholder="选择上级公司/部门"
                @change="searchSelectUDept"
              />
            </el-form-item>
            <el-form-item :label="udept.type=='Company'?'公司名称':'部门名称'" prop="name">
              <el-input v-model="udept.name" :disabled="isUpdating?false:true" />
            </el-form-item>
            <el-form-item label="排序" prop="rank">
              <el-input-number v-model="udept.rank" controls-position="right" :min="0" :disabled="isUpdating?false:true" />
            </el-form-item>
            
            <el-form-item label="备注" prop="remark">
              <el-input
                v-model="udept.remark"
                :disabled="isUpdating?false:true"
                type="textarea"
                :autosize="{ minRows: 2, maxRows: 4}"
              />
            </el-form-item>
          </el-form>
        </el-card>
        
      </el-col>
       
      <!-- 新增、编辑 -->
      <el-dialog :title="textMap[dialogType]" :visible.sync="dialogFormVisible" width="600px">
        <el-form ref="cdataForm" :model="dept" label-position="center" label-width="80px">
          <el-form-item label="上级公司" prop="parentId">
            <tree-select
              v-model="cParentId"
              :tree-data="companyAndDepts"
              :props-tree="defaultProps"
              :clearable="true"
              :searchable="true"
              placeholder="选择上级公司/部门"
              @change="searchSelectChain"
            />
          </el-form-item>
          <el-form-item :label="dataType=='Company'?'公司名称':'部门名称'" prop="name">
            <el-input v-model="dept.name" />
          </el-form-item>
          <el-form-item label="排序" prop="rank">
            <el-input-number v-model="dept.rank" controls-position="right" :min="0" />
          </el-form-item>
          
          <el-form-item label="备注" prop="remark">
            <el-input
              v-model="dept.remark"
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
          <el-button type="primary" @click="createData()">
            确认
          </el-button>
        </div>
      </el-dialog>

    </el-row>
  </div>
</template>

<script>
import { getCompanyAndDeptList, createCompany, createDept, updateCompany, updateDept, deleteCompany, deleteDept } from '@/api/organization'
import TreeSelect from '@/components/TreeSelect'

export default {
  name: 'Organization',
  components: { TreeSelect },
  data() {
    return {
      title: '公司信息',
      isCanAdd: true,
      isCanUpdate: true,
      isCanDelete: true,
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      filterText: '',
      companyAndDepts: [],
      activeTab: 'LoginLogs',
      dataType: '',
      dept: {
        companyId: '',
        name: '',
        parentId: '',
        rank: 0,
        remark: '',
        type: ''
      },
      isUpdating: false,
      cParentId: '',
      uParentId: '',
      udept: {
        companyId: '',
        id: '',
        name: '',
        parentId: '',
        rank: 0,
        remark: '',
        type: ''
      },
      cRules: [],
      dialogFormVisible: false,
      dialogType: '',
      textMap: {
        AddCompany: '新增公司',
        AddDept: '新增部门'
      }
    }
  },
  watch: {
    // 根据名称筛选部门树
    filterText(val) {
      this.$refs.qtree.filter(val)
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      getCompanyAndDeptList().then(rsp => {
        this.companyAndDepts = rsp
      })
    },
    filterNode(value, data) {
      if (!value) return true
      return data.name.indexOf(value) !== -1
    },
    handleNodeClick(data, node) {
      this.udept = Object.assign({}, {})
      this.udept = Object.assign({}, data)
      if (data.type === 'Dept') {
        this.title = '部门信息'
        if (data.parentId !== '' && data.parentId !== null) {
          this.$refs.utree.setTreeCurrentKey(this.udept.parentId)
        } else {
          this.$refs.utree.setTreeCurrentKey(this.udept.companyId)
        }
      } else {
        this.title = '公司信息'
        this.$refs.utree.setTreeCurrentKey(this.udept.parentId)
      }
    },
    handleCreate(type) {
      this.dept = Object.assign({}, null) // copy obj
      this.dialogFormVisible = true
      this.dataType = type
      if (type === 'Company') {
        this.dialogType = 'AddCompany'
      } else {
        this.dialogType = 'AddDept'
      }
    },
    createData() {
      if (this.dataType === 'Company') {
        this.$refs['cdataForm'].validate((valid) => {
          if (valid) {
            const tempData = Object.assign({}, this.dept)
            createCompany(tempData).then(() => {
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
      } else {
        this.$refs['cdataForm'].validate((valid) => {
          if (valid) {
            const tempData = Object.assign({}, this.dept)
            createDept(tempData).then(() => {
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
      }
    },
    handleUpdate() {
      if (this.udept.type === 'Company') {
        this.$refs['udataForm'].validate((valid) => {
          if (valid) {
            const tempData = Object.assign({}, this.udept)
            updateCompany(tempData).then(() => {
              this.$notify({
                title: 'Success',
                message: '修改成功',
                type: 'success',
                duration: 2000
              })
              this.isUpdating = false
              this.getList()
            })
          }
        })
      } else {
        this.$refs['udataForm'].validate((valid) => {
          if (valid) {
            const tempData = Object.assign({}, this.udept)
            updateDept(tempData).then(() => {
              this.$notify({
                title: 'Success',
                message: '修改成功',
                type: 'success',
                duration: 2000
              })
              this.isUpdating = false
              this.getList()
            })
          }
        })
      }
    },
    handleDelete() {
      if (this.udept.type === 'Company') {
        deleteCompany(this.udept.id).then(() => {
          this.$notify({
            title: 'Success',
            message: '删除成功',
            type: 'success',
            duration: 2000
          })
          this.isUpdating = false
          this.getList()
        })
      } else {
        deleteDept(this.udept.id).then(() => {
          this.$notify({
            title: 'Success',
            message: '删除成功',
            type: 'success',
            duration: 2000
          })
          this.isUpdating = false
          this.getList()
        })
      }
    },
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children
      }
    },
    searchSelectChain(obj) {
      this.dept.parentId = null
      if (this.dataType === 'Dept') {
        if (obj.type === 'Company') {
          this.dept.companyId = obj.id
        } else {
          this.cParentId = obj.id
          this.dept.parentId = obj.id
          this.dept.companyId = obj.companyId
        }
      } else {
        this.cParentId = obj.id
        this.dept.parentId = obj.id
      }
    },
    searchSelectUDept(obj) {
      this.udept.parentId = null
      if (this.dataType === 'Dept') {
        if (obj.type === 'Company') {
          this.udept.companyId = obj.id
        } else {
          this.uParentId = obj.id
          this.udept.parentId = obj.id
          this.udept.companyId = obj.companyId
        }
      } else {
        this.uParentId = obj.id
        this.udept.parentId = obj.id
      }
    }
  }
}

</script>
<style lang="scss" scoped>
.box-center {
  margin: 0 auto;
  display: table;
}

.text-muted {
  color: #777;
}

.user-profile {
  .user-name {
    font-weight: bold;
  }

  .box-center {
    padding-top: 10px;
  }

  .user-role {
    padding-top: 10px;
    font-weight: 400;
    font-size: 14px;
  }

  .box-social {
    padding-top: 30px;

    .el-table {
      border-top: 1px solid #dfe6ec;
    }
  }

  .user-follow {
    padding-top: 20px;
  }
}

.user-bio {
  margin-top: 20px;
  color: #606266;

  span {
    padding-left: 4px;
  }

  .user-bio-section {
    font-size: 14px;
    padding: 15px 0;

    .user-bio-section-header {
      border-bottom: 1px solid #dfe6ec;
      padding-bottom: 10px;
      margin-bottom: 10px;
      font-weight: bold;
    }
  }
}

.o-el-header {
    height: 30px;
    .o-span {
      line-height: 30px;
    }
    .btn-eb {
      display: inline;
      float: right;
    }
}

</style>
