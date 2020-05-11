<template>
  <el-container>
    <el-main>
      <el-col :span="12">
        <commit
          :repo-name="repoName"
          :commit-hash="commitHash"
        />
      </el-col>
      <el-col :span="12">
        <el-form
          :model="value"
          id="label-form"
        >
          <el-form-item label="label">
            <el-select v-model="label">
              <el-option
                v-for="(item, index) in labelList"
                :key="index"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="description">
            <el-input v-model="description" />
          </el-form-item>
        </el-form>
      </el-col>
    </el-main>
  </el-container>
</template>

<script>
import Commit from './Commit'

export default {
  components: {
    Commit
  },
  props: {
    repoName: String,
    commitHash: String,
    labelList: Array,
    value: Object
  },
  computed: {
    label: {
      get () {
        return this.value.label
      },
      set (val) {
        this.updateValue({ label: val })
      }
    },
    description: {
      get () {
        return this.value.description
      },
      set (val) {
        this.updateValue({ description: val })
      }
    }
  },
  methods: {
    updateValue (diff) {
      this.$emit('input', Object.assign({}, this.value, diff))
    }
  }
}
</script>

<style>
#label-form{
  padding: 20px;
}
</style>
