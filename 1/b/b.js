const readline = require("readline")
const rl = readline.createInterface({
  input: process.stdin,
})
let lines = []
rl.on("line", (line) => {
  lines.push(line)
}).on("close", () => {
  const [a, b, c] = lines
  let res = +a + +b > +c && +a + +c > +b && +b + +c > +a ? "YES" : "NO"
  process.stdout.write(res)
})

// console.log(
//   getTemp(`10 20
// heat`)
// )

// console.log(
//   getTemp(`10 20
// freeze`)
// )

// console.log(
//   getTemp(`-10 -20
// freeze`)
// )

// console.log(
//   getTemp(`-10 -10
// freeze`)
// )

// console.log(
//   getTemp(`0 -1
// freeze`)
// )

// console.log(
//   getTemp(`0 -1
// heat`)
// )

// console.log(
//   getTemp(`0 -1
// fan`)
// )
