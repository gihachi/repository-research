<template>
  <el-container>
    <el-main>
      <el-input v-model="repo" />
       <input id="json-upload" type="file" @change="onFileChange">
      <el-button
        type="primary"
        @click="registerJson"
      >
        Register
      </el-button>
      <el-button
        type="primary"
        @click="downloadLabeledData"
      >
        Download Labeled Data
      </el-button>

      <nuxt-link to="/sequentialCheck">Sequential Check</nuxt-link>
    </el-main>
  </el-container>
</template>

<script>

export default {
  data () {
    return {
      repo: '',
      file: null
    }
  },
  methods: {
    registerJson () {
      this.$store.commit('expandedDataset/setRepo', this.repo)

      const fr = new FileReader()
      fr.onload = (e) => {
        const json = JSON.parse(e.target.result)
        this.$store.commit('expandedDataset/setDatasets', json)
        this.$message.success('loaded!!!')
      }
      fr.readAsText(this.file)
    },
    downloadLabeledData () {
      const labeledDataStr = JSON.stringify(this.$store.state.expandedDataset.labeledData, null, '\t')
      const a = document.createElement('a')
      document.body.appendChild(a)
      a.style = 'display:none'
      const blob = new Blob([labeledDataStr], { type: 'text/json' })
      const url = window.URL.createObjectURL(blob)
      a.href = url
      a.download = `${this.$store.state.expandedDataset.repo}.json`
      a.click()
      window.URL.revokeObjectURL(url)
      a.parentNode.removeChild(a)
    },
    onFileChange (e) {
      const files = e.target.files
      this.file = files[0]
    }
  }
}
</script>

<style>
</style>
