import { readFileSync } from 'fs'

export const part1 = () => {
  let total = 0
  const input = readFileSync('./2024/01/files/question.input', 'utf-8').split(
    '\n'
  )

  let leftSide: number[] = []
  let rightSide: number[] = []

  for (const line of input) {
    const clean = line.split('   ') as string[]
    leftSide.push(Number(clean[0]))
    rightSide.push(Number(clean[1]))
  }

  leftSide.sort((a, b) => a - b)
  rightSide.sort((a, b) => a - b)

  for (let index = 0; index < leftSide.length; index++) {
    total = total + Math.abs(leftSide[index] - rightSide[index])
  }

  console.log('Result Part 1: ', total)
}

part1()
