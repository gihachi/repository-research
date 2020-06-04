<template>
  <el-container>
    <el-main>
      <el-row>
        <el-col :span="12">
          <h2>Inducing</h2>
          <div
            v-for="(item, index) in expandedData.originalInducingCommitInfoList"
            :key="index"
          >
            <commit
              :repo-name="repoName"
              :commit-hash="item.commitHash"
            />
          </div>
        </el-col>
        <el-col :span="12">
          <h2>Fix</h2>
          <div
            v-for="(item, index) in expandedData.originalFixCommitInfoList"
            :key="index"
          >
            <commit
              :repo-name="repoName"
              :commit-hash="item.commitHash"
            />
          </div>
        </el-col>
      </el-row>
      <el-button
        type="primary"
        @click="nextButtonClicked"
      >
        next data
      </el-button>
      <el-tabs>
        <el-tab-pane label="Expanded Induce" name="induce">
          <el-collapse v-model="inducingCollapse" accordion>
            <el-collapse-item
              v-for="(item, index) in expandedInduceDatas"
              :key="index"
              :title="item.commitHash"
              :name="index"
            >
              <commit-label-form
                v-model="expandedInduceDatas[index]"
                :repo-name="repoName"
                :commit-hash="item.commitHash"
                :labelList="['miss-link', 'correct-link', 'unkown']"
              />
            </el-collapse-item>
          </el-collapse>
        </el-tab-pane>
        <el-tab-pane label="Expanded fix" name="fix">
          <el-collapse v-model="fixCollapse" accordion>
            <el-collapse-item
              v-for="(item, index) in expandedFixDatas"
              :key="index"
              :title="item.commitHash"
              :name="index"
            >
              <commit-label-form
                v-model="expandedFixDatas[index]"
                :repo-name="repoName"
                :commit-hash="item.commitHash"
                :labelList="['miss-link', 'correct-link', 'unkown']"
              />
            </el-collapse-item>
          </el-collapse>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<script>
import Commit from './Commit'
import CommitLabelForm from './CommitLabelForm'

export default {
  components: {
    Commit,
    CommitLabelForm
  },
  props: {
    repoName: String,
    expandedData: Object
  },
  data () {
    return {
      activeTab: 'induce',
      expandedInduceDatas: [],
      expandedFixDatas: [],
      inducingCollapse: [],
      fixCollapse: []
    }
  },
  created () {
    const expandedInduceList = this.expandedData.expandedInducingCommitInfoList
    for (let i = 0; i < expandedInduceList.length; i++) {
      this.expandedInduceDatas.push({
        commitHash: expandedInduceList[i].commitHash,
        label: '',
        description: ''
      })
    }
    const expandedFixList = this.expandedData.expandedFixCommitInfoList
    for (let i = 0; i < expandedFixList.length; i++) {
      this.expandedFixDatas.push({
        commitHash: expandedFixList[i].commitHash,
        label: '',
        description: ''
      })
    }
  },
  methods: {
    nextButtonClicked () {
      for (const formData of this.expandedInduceDatas) {
        if (formData.label === '') {
          this.$message.error(formData.commitHash + ' in induce is not labeled')
          return
        }
      }
      for (const formData of this.expandedFixDatas) {
        if (formData.label === '') {
          this.$message.error(formData.commitHash + ' in fix is not labeled')
          return
        }
      }
      console.log(this.expandedData.issueId)
      this.$emit('inputComplete', this.expandedInduceDatas, this.expandedFixDatas, this.expandedData.issueId)
    }
  }
}
</script>

<style>
</style>
