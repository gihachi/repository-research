<template>
  <el-container>
    <el-main>
      <expanded-data
        v-if="checkDataset"
        :repo-name="repoName"
        :expanded-data="checkDataset"
        :key="datasetIndex"
        @inputComplete="processInputData"
      />
      <div v-else>
        all datas checked
        <nuxt-link to="/">
          Go to index
        </nuxt-link>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import ExpandedData from '../components/commit/ExpandedData'
export default {
  components: {
    ExpandedData
  },
  computed: {
    checkDataset () {
      let nowIndex = this.$store.state.expandedDataset.index
      const expandedDataset = this.$store.state.expandedDataset.datasets
      console.log(nowIndex)
      console.log(expandedDataset)

      while (nowIndex < expandedDataset.length &&
        expandedDataset[nowIndex].numOfExpandedInducingCommit === 0 &&
        expandedDataset[nowIndex].numOfExpandedFixCommit === 0) {
        this.$store.commit('expandedDataset/incrementIndex')
        nowIndex = this.$store.state.expandedDataset.index
      }

      if (nowIndex >= expandedDataset.length) {
        return null
      } else {
        return expandedDataset[nowIndex]
      }
    },
    repoName () {
      return this.$store.state.expandedDataset.repo
    },
    datasetIndex () {
      return this.$store.state.expandedDataset.index
    }
  },
  methods: {
    processInputData (induceDatas, fixDatas, issueId) {
      console.log(issueId)
      const labeledData = {
        induce: induceDatas,
        fix: fixDatas
      }
      this.$store.commit('expandedDataset/setLabeledData', {
        value: labeledData,
        key: issueId
      })
      this.$store.commit('expandedDataset/incrementIndex')

      this.$router.push('/sequentialCheck')
    }
  }
}
</script>

<style>
</style>
