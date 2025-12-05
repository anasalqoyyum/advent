import { readFileSync } from 'fs'

export const part1 = () => {
  let total = 0
  const input = readFileSync('./2024/02/files/question.input', 'utf-8').split(
    '\n'
  )

  for (const line of input) {
    const clean = line.split(' ') as string[]
    let status = ''
    let isSafe = true

    clean.forEach((v, i, arr) => {
      if (i === arr.length - 1) return

      const res = Number(v) - Number(arr[i + 1])
      if (res === 0) {
        isSafe = false
      }

      if (res < 0) {
        if (i > 0 && status !== 'Increasing') isSafe = false
        status = 'Increasing'
        if (!(Math.abs(res) >= 1 && Math.abs(res) <= 3)) {
          isSafe = false
        }
      }

      if (res > 0) {
        if (i > 0 && status !== 'Decreasing') isSafe = false
        status = 'Decreasing'
        if (!(Math.abs(res) >= 1 && Math.abs(res) <= 3)) {
          isSafe = false
        }
      }
    })

    if (isSafe) total += 1
  }

  console.log('Result Part 1: ', total)
}

part1()
