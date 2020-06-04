export const state = () => ({
  repo: null,
  datasets: [],
  index: 0,
  labeledData: {}
})

export const getters = {
  getRepo: argState => argState.repo,
  getDatasets: argState => argState.datasets,
  getIndex: argState => argState.datasets.index,
  getLabeleData: argState => argState.dataset.index
}

export const mutations = {
  setRepo: (argState, repo) => {
    argState.repo = repo
  },
  addDataset: (argState, dataset) => {
    argState.datasets.push(dataset)
  },
  setDatasets: (argState, datasets) => {
    argState.datasets = datasets
    argState.index = 0
    argState.labeledData = {}
  },
  incrementIndex: (argState) => {
    argState.index++
  },
  setLabeledData: (argState, labeledData) => {
    argState.labeledData[labeledData.key] = labeledData.value
  }
}
