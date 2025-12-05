import { readFileSync } from 'fs'

const Games: Record<string, number> = {
  red: 12,
  green: 13,
  blue: 14,
}

export default function part1() {
  let total = 0
  const input = readFileSync('./2023/02/files/input.txt', 'utf-8').split('\n')

  for (const line of input) {
    let perSet: string[] = []
    let isNotPossible = false

    const [gameText, recordText] = line.split(':')
    const [, gameId] = gameText.split(' ')
    recordText.split(';').forEach((elm) => {
      elm.split(',').forEach((el) => {
        perSet.push(el)
      })
    })

    for (const color of perSet) {
      Object.keys(Games).forEach((key) => {
        if (color.includes(key)) {
          const clean = color.replace(/[^0-9]/g, '')
          const num = Games[key] - Number(clean)
          if (!isNotPossible) {
            isNotPossible = num < 0
          }
        }
      })
    }

    total = isNotPossible ? total : total + Number(gameId)
  }

  console.log('Result Part 1: ', total)
}

part1()
