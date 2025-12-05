import { readFileSync } from 'fs'

export const part1 = () => {
  let total = 0
  const input = readFileSync('./2023/01/files/input.txt', 'utf-8').split('\n')

  for (const line of input) {
    let final = 0
    const clean = line.replace(/[^0-9]/g, '')
    const lineLen = clean.length

    if (lineLen === 1) {
      final = Number(clean.concat(clean))
    } else {
      final = Number(clean.charAt(0).concat(clean.charAt(lineLen - 1)))
    }

    total = total + final
  }

  console.log('Result Part 1: ', total)
}
