import { readFileSync } from 'fs'

const letterDigit: Record<string, string> = {
  one: '1',
  two: '2',
  three: '3',
  four: '4',
  five: '5',
  six: '6',
  seven: '7',
  eight: '8',
  nine: '9',
}

export const part2 = () => {
  let total = 0
  const input = readFileSync('./2023/01/files/input.txt', 'utf-8').split('\n')

  for (let line of input) {
    let final = 0
    Object.keys(letterDigit).forEach((key) => {
      const occurance = line.split(key).length - 1
      for (let i = 1; i <= occurance; i++) {
        line = line.replace(
          key,
          `${key.charAt(0)}${letterDigit[key]}${key.charAt(key.length - 1)}`
        )
      }
    })

    const clean = line.replace(/[^0-9]/g, '')
    const lineLen = clean.length
    if (lineLen === 1) {
      final = Number(clean.concat(clean))
    } else {
      final = Number(clean.charAt(0).concat(clean.charAt(lineLen - 1)))
    }
    total = total + final
  }

  console.log('Result Part 2: ', total)
}

// 53652 -> Wrong
// 54119 -> Wrong
// 54303 -> Wrong
// 54094 -> Correct
