<template>
  <div>
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
    <pagination v-show="total>0" :total="total" :page.sync="pageParams.page" :limit.sync="pageParams.take" @pagination="getList" />
  </div>
</template>

<script>
import OperateTable from '@/components/OperateTable'
import Pagination from '@/components/Pagination'
import { getCurrentUserLogsList } from '@/api/auditlog'
export default {
  name: 'LoginLogs',
  components: { OperateTable, Pagination },
  data() {
    return {
      tableCols: [
        { label: '登录时间', prop: 'executionTime', type: 'DateTime', align: 'center', width: '150', sortable: 'custom' },
        { label: 'IP地址', prop: 'clientIpAddress', align: 'center', width: '120', sortable: 'custom' },
        { label: '结果', prop: 'httpStatusCode', type: 'TagFilter', align: 'center', width: '100', sortable: 'custom', filters: { 200: '成功', 500: '失败' }, colorFilters: { 200: 'success', 500: 'danger' }},
        { label: '登录设备', prop: 'browserInfo', align: 'center' },
        { label: '系统', prop: 'fromSource', align: 'center', width: '100' }
      ],
      pageParams: {
        requireTotalCount: true,
        page: 1,
        skip: 0,
        take: 10
      },
      list: [],
      total: 0
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.pageParams.skip = (this.pageParams.page - 1) * this.pageParams.take
      const tempData = Object.assign({}, this.pageParams)
      getCurrentUserLogsList(tempData).then(response => {
        this.list = response.data
        this.total = response.totalCount
      }).catch(error => {
        console.log(error)
      })
    },
    operationSelect(operationName, row) {
    },
    sortChange(data) {
      this.pageParams.page = 1
      if (data.order !== null) {
        this.pageParams.sort = `[{"selector":"${data.prop}","desc":${data.order === 'descending'}}]`
      } else {
        this.pageParams.sort = null
      }
      this.getList()
    }
  }
}
</script>

