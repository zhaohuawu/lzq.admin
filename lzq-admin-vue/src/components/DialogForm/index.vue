<template>
  <el-dialog ref="dialog-form" :title="title" :visible.sync="dialogFormVisible" :width="width">
    <el-form ref="dataForm" :rules="rules" :model="formData" label-position="left" label-width="80px">
      <el-form-item
        v-for="item in formItems" 
        :key="item.key"
        :prop="item.prop" 
        :label="item.label" 
        :width="item.width"
      >

        <!--radio-->
        <el-radio-group v-if="item.type==='radio'" v-model="formData[item.prop]" @change="item.change">
          <el-radio v-for="r in item.radioes" :key="r[item.radioValue]" :label="r[item.radioValue]">r[item.radioName]</el-radio>
        </el-radio-group>
        <!--avatar-->
        <div v-else-if="item.type==='avatar'">
          <el-avatar :size="100" :src="item.imageLink" />
          <el-button type="primary" style="position:absolute;bottom:15px;margin-left: 20px;text-align:center;" icon="el-icon-upload" @click="avatarShow()">
            Upload Avatar
          </el-button>
              
          <div style="float:left">
            <image-cropper
              v-show="item.imagecropperShow"
              :key="item.imagecropperKey"
              :width="item.imgWidth"
              :height="item.imgHeight"
              field="file"
              :url="item.httpWebApi"
              lang-type="en"
              @close="avatarclose"
              @crop-upload-success="cropSuccess"
            />
          </div>
        </div>
        <!--select-->
        <el-select v-else-if="item.type==='select'" v-model="formData[item.prop]" :placeholder="item.placeholder" clearable @change="item.change">
          <el-option
            v-for="(s, o) in item.selectOptions"
            :key="o"
            :label="s[item.optionName]"
            :value="s[item.optionValue]"
          />
        </el-select>
        <el-date-picker
          v-else-if="['data', 'datetimerange', 'datetime'].indexOf(item.type) !== -1"
          v-model="value[i.name]"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          value-format="yyyy-MM-dd HH:mm:ss"
        />
        <!--password-->
        <el-input v-else-if="item.type==='password'" v-model="formData[item.prop]" :placeholder="item.placeholder" show-password />
        <el-input v-else v-model="formData[item.prop]" :placeholder="item.placeholder" />
            
      </el-form-item>

    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogClose()">取消</el-button>
      <el-button type="primary" @click="dialogSave()">确认</el-button>
    </div>
  </el-dialog>
</template>

<script>
import ImageCropper from '@/components/ImageCropper'

export default {
  name: 'DialogForm',
  components: { ImageCropper },
  props: {
    title: { type: String, default: '新增' },
    width: { type: String, default: '750px' },
    // 表单数据
    formData: { type: Object, default: null },
    // 表单项
    formItems: { type: Array, default: () => [] },
    dialogFormVisible: { type: Boolean, default: false },
    rules: { type: Object, default: null }
  },
  data() {
    return {
    }
  },
  methods: {
    avatarShow() {
      this.$emit('avatarShow')
    },
    cropSuccess(resData) {
      this.$emit('cropSuccess', resData)
    },
    avatarclose() {
      this.$emit('avatarclose')
    },
    dialogSave() {
      this.$emit('dialogSave')
    },
    dialogClose() {
      this.$emit('dialogClose')
    }
  }
}
</script>
