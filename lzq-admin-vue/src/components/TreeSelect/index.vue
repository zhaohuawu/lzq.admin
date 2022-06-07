<template>
  <el-select
    ref="option"
    class="select_tree"
    popper-class="select_tree_popper"
    :placeholder="placeholder"
    :value="value"
    :clearable="clearable"
    :disabled="disabled"
    @clear="clearHandle"
  >
    <li class="select_tree_search">
      <el-input v-model="filterText" :disabled="disabled" placeholder="输入关键字进行过滤" />
    </li>
    <el-option :value="currentSelect.value" :label="currentSelect.label" :disabled="disabled">
      <div class="select_tree_content">
        <el-tree
          ref="tree"
          class="select_tree_cmp"
          :accordion="accordion"
          :highlight-current="true"
          :data="treeData"
          :props="propsTree"
          :node-key="nodekey"
          empty-text="没有数据"
          :show-checkbox="showcheckbox"
          :check-on-click-node="false"
          :check-strictly="false"
          :default-expanded-keys="defaultExpandedKey"
          :filter-node-method="filterNode"
          @node-click="handleNodeClick"
          @check-change="handleCheckChange"
        />
      </div>
    </el-option>
  </el-select>
</template>
<script>
export default {
  name: 'TreeSelect',
  props: {
    treeData: { type: Array, default: () => [] },
    defaultExpandedKey: { type: Array, default: () => [] },
    clearable: { type: Boolean, default: true },
    accordion: { type: Boolean, default: true },
    showcheckbox: { type: Boolean, default: false },
    placeholder: { type: String, default: '' },
    value: { type: String, default: '' },
    nodekey: { type: String, default: 'id' },
    disabled: { Type: Boolean, default: false },
    propsTree: {
      type: Object,
      default() {
        return {
          label: 'label',
          children: 'children'
        }
      }
    },
    input: { type: Function, default: null },
    change: { type: Function, default: null } // 选择回调
  },
  data() {
    return {
      filterText: '', // 筛选文字
      currentSelect: {}
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
    // treeData(val) {
    //   if (this.value) {
    //     const selected = this.getActiveLeaf(val, this.value)
    //     if (selected) {
    //       this.currentSelect = {
    //         value: selected.value,
    //         label: selected.label
    //       }
    //     }
    //   }
    // },
    // currentSelect(val) {
    //   this.$refs.tree.setCheckedKeys([val.value])

    // }
  },
  methods: {
    getActiveLeaf(array, value) {
      const res = array.find(item => {
        if (item.value === value) {
          return true
        } else if (array.children) {
          this.getActiveLeaf(array.children, value)
        }
      })
      if (res) {
        return res
      }
    },
    handleNodeClick(data, node) {
      // console.log('el-tree:handleNodeClick', data, node)
      this.currentSelect = {
        value: node.key,
        label: node.label
      }
      if (node.parent.label !== null && node.parent.label !== '' && node.parent.label !== undefined) {
        this.currentSelect.label = node.parent.label + '/' + this.currentSelect.label
      }
      // console.log(this.currentSelect)
      this.$emit('input', node.key)
      this.$emit('change', data)
      if (!data.children) {
        this.$refs.option.blur()
      }
    },
    handleCheckChange(data, checked, indeterminate) {
      if (this.showcheckbox) {
        if (checked) {
          this.currentSelect = {
            value: data.value,
            label: data.label
          }
          this.$emit('input', data.value)
          this.$emit('change', data)
          if (!data.children) {
            this.$refs.option.blur()
          }
        } else {
          this.$emit('input', '')
          this.$emit('change', {})
        }
        // console.log(data, checked, indeterminate)
      }
    },
    filterNode(value, data) {
      if (!value) return true
      return data.name.indexOf(value) !== -1
    },
    clearHandle() {
      this.currentSelect = {}
      this.$emit('input', '')
    },
    setTreeCurrentKey(id) {
      if (id === null || id === '') {
        this.$refs.tree.setCurrentKey(null)
        this.clearHandle()
        return
      }
      this.$refs.tree.setCurrentKey(id)
      const cdata = this.$refs.tree.getCurrentNode()
      const cnode = this.$refs.tree.getNode(id)
      // console.log('CurrentNode', cdata, cnode) 
      if (cnode === null) {
        this.clearHandle()
        return
      }
      this.currentSelect = {
        value: cnode.key,
        label: cnode.label
      }
      if (cnode.parent.label !== null && cnode.parent.label !== '' && cnode.parent.label !== undefined) {
        this.currentSelect.label = cnode.parent.label + '/' + this.currentSelect.label
      }
      // console.log(this.currentSelect)
      this.$emit('input', cnode.key)
      this.$emit('change', cdata)
      if (!cdata.children) {
        this.$refs.option.blur()
      }
    }
  }
}
</script>
<style lang="scss" scoped>
.select_tree_popper {
  width: auto;
  
  .select_tree_search {
    padding: 8px;
    width: 100%;
    box-sizing: border-box;
    background-color: #fff;
    position: absolute;
    left: 0px;
    top: 0px;
    z-index: 100;
  }

  .el-input__inner {
    height: 30px;
    line-height: 30px;
  }

  .el-select-dropdown__item {
    height: auto;
    padding: 0;
  }

  .select_tree_content {
    background-color: #fff;
    padding-bottom: 6px;
    padding-top: 44px;
  }
  .el-checkbox{
    display: none;
  }

  .el-select-dropdown__item.selected {
    font-weight: normal;
  }
}
</style>
