import { readFileSync } from 'fs'

export default function part2() {
  let total = 0
  const input = readFileSync('./day-2/files/input.txt', 'utf-8').split('\n')

  for (const line of input) {
    let perSet: string[] = []
    let prevRed = 0
    let prevGreen = 0
    let prevBlue = 0

    const [gameText, recordText] = line.split(':')
    const [, gameId] = gameText.split(' ')
    recordText.split(';').forEach((elm) => {
      elm.split(',').forEach((el) => {
        perSet.push(el)
      })
    })

    for (const color of perSet) {
      const realColor = color.replace(/\s[0-9]/g, '')
      const strNum = color.replace(/[^0-9]/g, '')
      const num = Number(strNum)

      if (realColor.includes('red')) {
        prevRed = prevRed < num ? num : prevRed
      } else if (realColor.includes('green')) {
        prevGreen = prevGreen < num ? num : prevGreen
      } else if (realColor.includes('blue')) {
        prevBlue = prevBlue < num ? num : prevBlue
      }
    }

    const power = prevRed * prevGreen * prevBlue
    total = total + power
  }

  console.log('Result Part 2: ', total)
}

part2()
