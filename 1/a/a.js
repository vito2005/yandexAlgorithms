const readline = require("readline")
const fs = require("fs")
const path = require("path")
const rl = readline.createInterface({
  input: fs.createReadStream(path.join(__dirname, "input.txt")),
})
let lines = []
rl.on("line", (line) => {
  lines.push(line)
}).on("close", () => {
  const [troom, tcond] = lines[0].split(" ")
  const mode = lines[1]
  function getTemp({ troom, tcond, mode }) {
    switch (mode) {
      case "heat": {
        return (+troom <= +tcond && tcond) || troom
      }
      case "freeze": {
        return (+troom <= +tcond && troom) || tcond
      }
      case "auto": {
        return tcond
      }
      case "fan": {
        return troom
      }
    }
  }
  const res = getTemp({ troom, tcond, mode })
  process.stdout.write(res)
})
