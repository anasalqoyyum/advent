import { readFileSync } from 'fs'

export default function part1() {
  let total = 0
  const input = readFileSync('./2023/02/files/input.txt', 'utf-8').split('\n')

  // For of Index and Val
  // Walkthrough the string to find Special Characters
  // If found -> Check str[idx-1], str[idx], str[idx+1] to see any number
  //             If number found walkthrough the number until '.' (it means it's the whole part number)
  //             Convert to number add it to total
  // If not found -> Continue
}

part1()
