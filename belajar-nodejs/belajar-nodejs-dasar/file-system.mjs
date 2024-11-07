import fs from "fs/promises"

//digunakan untuk membaca file
const buffer = await fs.readFile("file-system.mjs");
console.info(buffer.toString());

//digunakan untuk membuat dan membaca file.
const tulis = await fs.writeFile("temp.pdf", "fmt.Println(Hello)");
console.info(tulis);

//membuat folder dengan mkdir
const  membuatFile = await fs.mkdir("belajar node js dasar");
console.info(membuatFile);