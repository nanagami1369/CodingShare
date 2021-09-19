const fs = require('fs')

if (process.env.CODING_SHARE_API_URL == null) {
  console.error('Error 環境変数「CODING_SHARE_API_URL」が定義されていません')
  process.exit(-1)
} else {
  process.env.VUE_APP_CODING_SHARE_API_URL = process.env.CODING_SHARE_API_URL
}

let CERTIFICATE_FILE_PATH = ''
let KEY_FILE_PATH = ''
if (process.env.CODING_SHARE_CERTIFICATE_FILE_PATH == null) {
  console.error(
    'Error 環境変数「CODING_SHARE_CERTIFICATE_FILE_PATH」が定義されていません'
  )
  process.exit(-1)
} else {
  CERTIFICATE_FILE_PATH = process.env.CODING_SHARE_CERTIFICATE_FILE_PATH
}

if (process.env.CODING_SHARE_KEY_FILE_PATH == null) {
  console.error(
    'Error 環境変数「CODING_SHARE_KEY_FILE_PATH」が定義されていません'
  )
  process.exit(-1)
} else {
  KEY_FILE_PATH = process.env.CODING_SHARE_KEY_FILE_PATH
}
module.exports = {
  publicPath: '/CodingShare',
  outputDir: '../docs',
  devServer: {
    https: {
      cert: fs.readFileSync(CERTIFICATE_FILE_PATH),
      key: fs.readFileSync(KEY_FILE_PATH),
    },
  },
}
