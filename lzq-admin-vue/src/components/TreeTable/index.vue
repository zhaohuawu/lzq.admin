<template>
  <el-table
    v-loading="loading"
    :data="tableData"
    style="width: 100%;margin-bottom: 20px;"
    :size="size"
    row-key="id"
    fit
    highlight-current-row
    :default-expand-all="expandAll"
    :border="isBorder"
    :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
  >

    <el-table-column
      v-for="(item,tIndex) in tableCols" 
      :key="tIndex"
      :prop="item.prop" 
      :label="item.label" 
      :width="item.width"
      :align="item.align" 
      :render-header="item.require?renderHeader:null"
    >
      <template slot-scope="scope">
        <!-- html -->
        <span v-if="item.type==='Html'" v-html="item.html(scope.row)" />
        <!-- 按钮 -->
        <span v-if="item.type==='Button' && scope.row[item.prop] !== null && scope.row[item.prop] !== ''">
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
        <!-- Icon -->
        <div v-if="item.type==='Icon'" class="icon-item">
          <span v-if="scope.row.icon!=null&&scope.row.icon!=''">
            <svg-icon :icon-class="scope.row.icon" class-name="disabled" />
          </span>
        
        </div>
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
        <img v-if="item.type==='Image'" :src="scope.row[item.prop]" @click="item.handle && item.handle(scope.row)">
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
        <!-- 默认 -->
        <span
          v-if="!item.type" 
          :style="item.style && item.style(scope.row)" 
          :class="item.class && item.item.class(scope.row)"
        >{{ (item.formatter && item.formatter(scope.row)) || scope.row[item.prop] }}</span>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'TreeTable',
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
    indexLabel: { type: String, default: '序号' },
    // 是否展开所有行
    expandAll: { type: Boolean, default: false }
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
    // tableRowClassName({rowIndex}) {
    //     if (rowIndex % 2 === 0) {
    //         return "stripe-row";
    //     }
    //     return "";
    // }
    renderHeader(h, obj) {
      return h('span', { class: 'ces-table-require' }, obj.column.label)
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
  /* .el-table th.is-center, .el-table td.is-center{
    text-align:left;
  } */
</style>
