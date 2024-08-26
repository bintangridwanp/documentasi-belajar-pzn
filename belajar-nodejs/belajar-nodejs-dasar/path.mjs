import path from "path"

const file = "/Users/luffy/hello-world.txt";

console.info(path.dirname(file));
console.info(path.basename(file));
console.log(path.extname(file));
console.info(path.parse(file));
