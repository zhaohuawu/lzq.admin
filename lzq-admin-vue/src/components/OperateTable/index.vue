<template>
  <el-table
    ref="cesTable"
    v-loading="loading"
    :data="tableData" 
    :size="size"
    :border="isBorder"
    :default-selections="defaultSelections" 
    @select="select"
    @select-all="selectAll"
    @sort-change="sortChange"
  >
    <el-table-column v-if="isSelection" type="selection" align="center" />
    <el-table-column v-if="isIndex" type="index" :label="indexLabel" align="center" width="50" />
    <!-- 数据栏 -->
    <el-table-column
      v-for="(item,index) in tableCols" 
      :key="index"
      :prop="item.prop" 
      :label="item.label" 
      :width="item.width"
      :align="item.align" 
      :sortable="item.sortable"
      :render-header="item.require?renderHeader:null"
    >
      <template slot-scope="scope">
        <!-- html -->
        <span v-if="item.type==='Html'" v-html="item.html(scope.row)" />
        <!-- 按钮 -->
        <span v-if="item.type==='Button'&& scope.row[item.prop] !== null && scope.row[item.prop] !== ''">
          <span v-if="JSON.parse(scope.row[item.prop]).length > 2">
            <span v-for="(btn, n) in JSON.parse(scope.row[item.prop])" :key="btn.name">
              <el-button v-if="n < 1" type="text" size="medium" @click="handleClick(btn.name, scope.row)">
                {{ btn.text }}
              </el-button>
            </span>
            
            <el-dropdown style="margin-left: 10px;">
              <span class="el-dropdown-link">
                更 多<i class="el-icon-arrow-down el-icon--right" />
              </span>
              <el-dropdown-menu slot="dropdown">
                <span v-for="(btn, m) in JSON.parse(scope.row[item.prop])" :key="btn.name">
                  <el-dropdown-item v-if="m > 0" style="color: #409EFF;" @click.native="handleClick(btn.name, scope.row)">
                    {{ btn.text }}
                  </el-dropdown-item>
                </span>
              </el-dropdown-menu>
            </el-dropdown>
          </span>
          <span v-else>
            <el-button v-for="btn in JSON.parse(scope.row[item.prop])" :key="btn.name" type="text" size="medium" @click="handleClick(btn.name, scope.row)">
              {{ btn.text }}
            </el-button>
          </span>
        </span>
        <!-- 输入框 -->
        <el-input
          v-if="item.type==='Input'"
          v-model="scope.row[item.prop]"
          :size="size" 
          :disabled="btn.isDisabled && btn.isDisabled(scope.row)"
          @focus="item.focus && item.focus(scope.row)"
        />
        <!-- 下拉框 -->
        <el-select
          v-if="item.type==='Select'"
          v-model="scope.row[item.prop]"
          :size="size"
          :props="item.props"
          :disabled="btn.isDisabled && btn.isDisabled(scope.row)" 
          @change="item.change && item.change(scope.row)"
        >
          <el-option v-for="op in item.options" :key="op[item.props.value]" :label="op[item.props.label]" :value="op[item.props.value]" />
        </el-select>
     
        <!-- 评价 -->
        <el-rate
          v-if="item.type==='Rate'"
          v-model="scope.row[item.prop]"
          :disabled="btn.isDisabled && btn.isDisabled(scope.row)"
          @change="item.change && item.change(scope.row)"
        />
        <!-- 开关 -->
        <el-switch
          v-if="item.type==='Switch'"
          v-model="scope.row[item.prop]"
          :disabled="btn.isDisabled && btn.isDisabled(scope.row)"
          @change="item.change && item.change(scope.row)"
        />
        <!-- 图像 -->
        <!-- { label: '图片', prop: 'img', align: 'center', type: 'Image', width: '120', style: 'height: 80px' } -->
        <img v-if="item.type==='Image'" :src="scope.row[item.prop]" :style="item.style" @click="item.handle && item.handle(scope.row)">
        <!-- 滑块 -->
        <el-slider
          v-if="item.type==='Slider'"
          v-model="scope.row[item.prop]" 
          :disabled="btn.isDisabled && btn.isDisabled(scope.row)"
          @change="item.change && item.change(scope.row)"
        />
        <!-- Filters -->
        <span v-if="item.type==='Filter'">
          {{ (item.filters && item.filters[scope.row[item.prop]]) }}
        </span>
        <span v-if="item.type==='Tag'">
          <el-tag :disable-transitions="false" :type="item.color">
            {{ scope.row[item.prop] }}
          </el-tag>
        </span>
        <!--TagFilter 过滤器Tag-->
        <!--  { label: '结果', prop: 'httpStatusCode', type: 'TagFilter', align: 'center', width: '100', sortable: 'custom', filters: { 200: '成功', 500: '失败' }, colorFilters: { 200: 'success', 500: 'danger' }} -->
        <span v-if="item.type==='TagFilter'">
          <el-tag v-if="item.colorFilters" :disable-transitions="false" :type="(item.colorFilters && item.colorFilters[scope.row[item.prop]])">
            {{ (item.filters && item.filters[scope.row[item.prop]]) }}
          </el-tag>
          <el-tag v-else :disable-transitions="false" :type="item.color">
            {{ (item.filters && item.filters[scope.row[item.prop]]) }}
          </el-tag>
        </span>
        <span v-if="item.type==='TagList'">
          <el-tag
            v-for="(tag, i) in scope.row[item.prop]"
            :key="i"
            :disable-transitions="false"
            :type="item.color"
          >
            {{ tag[item.showProp] }}
          </el-tag>
        </span>
        <span v-if="item.type==='TagListWithClosable'">
          <el-tag
            v-for="(tag,i) in scope.row[item.prop]"
            :key="i"
            closable
            :disable-transitions="false"
            :type="item.color"
            @close="handleTagClose(item.prop,tag,scope.row)"
          >
            {{ tag[item.showProp] }}
          </el-tag>
        </span>
        <span v-if="item.type==='DateTime'">
          {{ parseTime(scope.row[item.prop]) }}
        </span>
        <span v-if="item.type==='Dialog'" style="color:#1890ff;cursor:pointer;" @click="handleClick(item.btnName, scope.row)">
          {{ scope.row[item.prop] }}
        </span>
        <!-- 默认 -->
        <span
          v-if="!item.type" 
          :style="item.itemStyle && item.itemStyle(scope.row)" 
          :class="item.itemClass && item.item.itemClass(scope.row)"
        >{{ (item.formatter && item.formatter(scope.row)) || scope.row[item.prop] }}</span>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>

export default {
  name: 'OperateTable',
  props: {
    // 表格型号：mini,medium,small
    size: { type: String, default: 'medium' },
    isBorder: { type: Boolean, default: true },
    loading: { type: Boolean, default: false },
    // 表格操作
    isHandle: { type: Boolean, default: false },
    tableHandles: { type: Array, default: () => [] },
    // 表格数据
    tableData: { type: Array, default: () => [] },
    // 表格列配置
    tableCols: { type: Array, default: () => [] },
    // 是否显示表格复选框
    isSelection: { type: Boolean, default: false },
    defaultSelections: { type: [Array, Object], default: () => null },
    // 是否显示表格索引
    isIndex: { type: Boolean, default: false },
    indexLabel: { type: String, default: '序号' }
    
  },
  data() {
    return {}
  },
  watch: {
    'defaultSelections'(val) {
      this.$nextTick(function() {
        if (Array.isArray(val)) {
          val.forEach(row => {
            this.$refs.cesTable.toggleRowSelection(row)
          })
        } else {
          this.$refs.cesTable.toggleRowSelection(val)
        }
      })      
    }
  },
  methods: {
    // 表格勾选
    select(rows, row) {
      this.$emit('select', rows, row)
    },
    // 全选
    selectAll(rows) {
      this.$emit('select', rows)
    },
    // 数据操作
    handleClick(type, data) {
      this.$emit('operationSelect', type, data)
    },
    handleTagClose(prop, tag, data) {
      this.$emit('handleTagClose', tag, data)
    },
    sortChange(data) {
      this.$emit('sortChange', data)
    },
    renderHeader(h, obj) {
      return h('span', { class: 'ces-table-require' }, obj.column.label)
    },
    parseTime(time, cFormat) {
      if (time === null || time === '') {
        return ''
      }
      if (time.indexOf('T') > -1) {
        time = time.replace('T', ' ')
      }
      const date = new Date(time)
      let fmt = cFormat || 'yyyy-MM-dd hh:mm:ss'
      if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
      }
      const o = {
        'M+': date.getMonth() + 1,
        'd+': date.getDate(),
        'h+': date.getHours(),
        'm+': date.getMinutes(),
        's+': date.getSeconds()
      }
      for (const k in o) {
        if (new RegExp(`(${k})`).test(fmt)) {
          const str = o[k] + ''
          fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? str : ('00' + str).substr(str.length))
        }
      }
      return fmt
    }
  }
}
</script>
<style>
.ces-table-require::before{
  content:'*';
  color:red;
}

.el-dropdown-link {
    cursor: pointer;
    color: #409EFF;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
</style>
