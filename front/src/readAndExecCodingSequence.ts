import { CodingSequence } from './models/CodingSequence'

export const readAndExecCodingSequence = (
  editor: CodeMirror.Editor,
  codingSequence: CodingSequence
): void => {
  if (codingSequence.changeData != null) {
    const { text, from, to, origin } = codingSequence.changeData
    editor.replaceRange(text, from, to, origin)
  }
  if (codingSequence.cursor != null) {
    editor.setCursor(codingSequence.cursor)
  }
}
