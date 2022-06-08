<template>
  <el-button v-if="isshow && isCanOperated" :disabled="disabled" :type="type" :size="size" :style="btnstyle" @click="btnclick">
    {{ text }}
  </el-button>
</template>

<script>
import store from '@/store'
export default {
  name: 'LzqButton',
  props: {
    text: {
      type: String,
      default: '',
      require: true
    },
    disabled: {
      type: Boolean,
      default: false
    },
    isshow: {
      type: Boolean,
      default: true
    },
    policy: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: '' // primary、success、info、warning、danger
    },
    size: {
      type: String,
      default: '' // medium / small / mini
    },
    btnclick: {
      type: Function,
      default: function() { }
    },
    btnstyle: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      isCanOperated: this.policy === null || this.policy === '' || this.policy === undefined || store.getters.superAdmin || store.getters.permissions.indexOf(this.policy) > -1
    }
  },
  methods: {
    handleClick(type, data) {
      this.$emit('operationSelect', type, data)
    }
  }
}
</script>

