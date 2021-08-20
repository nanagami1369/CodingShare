if (process.env.CODING_SHARE_API_URL == null) {
  console.error('Error 環境変数「CODING_SHARE_API_URL」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_CODING_SHARE_API_URL = process.env.CODING_SHARE_API_URL
}

module.exports = {
  publicPath: '/CodingShare',
  outputDir: '../docs',
}
