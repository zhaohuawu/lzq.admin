<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="filtersQuery.userName" placeholder="登录人" style="width: 200px;" clearable class="filter-item" />
      <el-select v-model="filtersQuery.httpStatusCode" placeholder="状态" clearable class="filter-item">
        <el-option v-for="(item,index) in httpStatus" :key="index" :label="item.text" :value="item.code" />
      </el-select>
     
      <el-date-picker
        v-model="filtersQuery.executionTime"
        type="datetimerange"
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        clearable 
        class="filter-item"
      />
      <el-button v-waves class="filter-item" type="primary" @click="handleFilter">
        搜索
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
    />
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.take" @pagination="getList" />
    
    <!-- 接口请求详情 -->
    <el-drawer ref="drawer" title="接口请求详情" :visible.sync="dialogFormVisible" direction="rtl" size="50%">
      <div class="n-drawer-content">
        <json-editor ref="jsonEditor" v-model="jsonValue" :is-read-only="true" />
      </div>
    </el-drawer>

  </div>

</template>

<script>
import { getList } from '@/api/auditlog'
import { dateFormat } from '@/utils/date'
import OperateTable from '@/components/OperateTable'
import Pagination from '@/components/Pagination'
import JsonEditor from '@/components/JsonEditor'

export default {
  name: 'AuditLogActionList',
  components: { OperateTable, Pagination, JsonEditor },
  data() {
    return {
      list: [],
      tableCols: [
        { label: '服务模块', prop: 'serviceModuleCode', align: 'center', width: '100' },
        { label: '租户', prop: 'tenantName', align: 'center', width: '100' },
        { label: '操作人', prop: 'userName', align: 'center', width: '100' },
        { label: 'IP地址', prop: 'clientIpAddress', align: 'center', width: '120', sortable: 'custom' },
        { label: '状态', prop: 'httpStatusCode', type: 'TagFilter', align: 'center', width: '80', sortable: 'custom', filters: { 200: '成功', 500: '失败' }, colorFilters: { 200: 'success', 500: 'danger' }},
        { label: '登录设备', prop: 'browserInfo', type: 'Dialog', btnName: 'Info', align: 'center' },
        { label: '登录时间', prop: 'executionTime', type: 'DateTime', align: 'center', width: '150', sortable: 'custom' },
        { label: '系统', prop: 'fromSource', align: 'center', width: '100' }
      ],
      total: 0,
      listLoading: true,
      filtersQuery: {
        userName: '',
        httpMethod: '',
        httpStatusCode: null,
        executionTime: []
      },
      httpMethods: ['GET', 'POST', 'PUT', 'DELETE'],
      httpStatus: [
        { 
          code: 200,
          text: '成功' 
        },
        { 
          code: 500,
          text: '失败' 
        }
      ],
      listQuery: {
        requireTotalCount: true,
        page: 1,
        skip: 0,
        take: 10,
        filter: null,
        sort: null
      },
      dialogFormVisible: false,
      jsonValue: null
    }
  },
  created() {
    this.listQuery.sort = `[{"selector":"executionTime","desc":true}]`
    this.handleFilter()
  },
  methods: {
    getList() {
      this.listLoading = true
      this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.take
      const tempData = Object.assign({}, this.listQuery)
      getList(tempData).then(response => {
        this.list = response.data
        this.total = response.totalCount
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      var filterArray = ['["actionType","=","Login"]']
      if (this.filtersQuery.userName !== '' && this.filtersQuery.userName !== null) {
        filterArray.push('["userName","contains","' + this.filtersQuery.userName + '"]')
      }
      if (this.filtersQuery.httpStatusCode !== '' && this.filtersQuery.httpStatusCode !== null) {
        filterArray.push('["httpStatusCode","=","' + this.filtersQuery.httpStatusCode + '"]')
      }
      if (this.filtersQuery.executionTime.length > 0) {
        filterArray.push('["executionTime",">=","' + dateFormat(this.filtersQuery.executionTime[0]) + '"]')
        filterArray.push('["executionTime","<","' + dateFormat(this.filtersQuery.executionTime[1]) + '"]')
      }
      if (filterArray.length > 0) {
        this.listQuery.filter = '[' + filterArray.join(',') + ']'
      }
      this.getList()
    },
    operationSelect(operationName, row) {
      if (operationName === 'Info') {
        this.showInfo(row)
      }
    },
    showInfo(row) {
      console.log(row)
      this.jsonValue = JSON.parse(JSON.stringify(row))
      this.dialogFormVisible = true
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
