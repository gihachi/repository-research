<template>
  <el-container>
    <el-main>
      <el-card>
        <div
          slot="header"
          class="clearfix"
        >
          <p id="commit-title">
            {{ title }}
          </p>
          <p
            class="line-and-space"
            id="commit-description"
          >
            {{ description }}
          </p>
        </div>
        <el-collapse
          v-model="activeCollaps"
          @change="fetchDiff"
        >
          <el-collapse-item
            title="Commit Hash"
            name="commitHash"
          >
            {{ commitHash }}
          </el-collapse-item>
          <el-collapse-item
            title="Author Info"
            name="authorInfo"
          >
            <el-table :data="authorInfo">
              <el-table-column
                prop="authorName"
                label="Author Name"
              />
              <el-table-column
                prop="authorDate"
                label="Author Date"
              />
            </el-table>
          </el-collapse-item>
          <el-collapse-item
            title="Commiter Info"
            name="commiterInfo"
          >
            <el-table :data="commiterInfo">
              <el-table-column
                prop="commiterName"
                label="Commiter Name"
              />
              <el-table-column
                prop="commiterDate"
                label="Commiter Date"
              />
            </el-table>
          </el-collapse-item>
          <el-collapse-item
            title="Diff"
            name="diff">
            <div
              v-html="diffHtml"
              v-loading="loadingDiff"
            />
          </el-collapse-item>
        </el-collapse>
      </el-card>
    </el-main>
  </el-container>
</template>

<script>
import * as Diff2Html from 'diff2html'
import 'diff2html/bundles/css/diff2html.min.css'

export default {
  props: {
    repoName: String,
    commitHash: String
  },
  data () {
    return {
      diffs: '',
      authorInfo: [],
      commiterInfo: [],
      isMerge: false,
      title: '',
      description: '',
      activeCollaps: ['authorInfo', 'commitHash'],
      loadingDiff: true
    }
  },
  computed: {
    diffHtml () {
      return Diff2Html.html(this.diffs, {
        drawFileList: true,
        matching: 'lines',
        outputFormat: 'line-by-line'
      })
    }
  },
  async created () {
    const res = await this.$axios.get('/commit', {
      params: {
        repo: this.repoName,
        commit: this.commitHash
      }
    }).catch((err) => {
      return err.response
    })

    if (res.status === 200) {
      const authorInfoItem = {
        authorName: res.data.author,
        authorDate: res.data.authorDate
      }

      this.authorInfo.push(authorInfoItem)

      const commiterInfoItem = {
        commiterName: res.data.commiter,
        commiterDate: res.data.commiterDate
      }

      this.commiterInfo.push(commiterInfoItem)

      const messages = res.data.message.split('\n')
      this.title = messages[0]

      if (messages.length > 1) {
        this.description = messages.slice(1, messages.length).join('\n')
      }
    }
  },
  methods: {
    async fetchDiff (activeArr) {
      if (activeArr.includes('diff') && this.diffs === '') {
        const res = await this.$axios.get('/diff', {
          params: {
            repo: this.repoName,
            commit: this.commitHash
          }
        }).catch((err) => {
          return err.response
        })

        if (res.status === 200) {
          this.diffs = res.data.diff
          this.loadingDiff = false
        }
      }
    }
  }
}
</script>

<style>
#commit-title{
  font-size: 18px;
  font-weight: bold;
}
#commit-description{
  line-height: 1em;
}
.line-and-space{
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
