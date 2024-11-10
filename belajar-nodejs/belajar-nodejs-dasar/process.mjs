import process, { exitCode } from "process";

process.addListener("exit", (exitCode) => {
    console.info(`nodeJS exit with code ${exitCode}`)
})

console.info(process.version); //mengecek versi nodeJs
console.info(process.argv); //mengecek argumen dari proces nodeJS
console.info(process.report); //mengecek report dari setiap process nodeJS
console.info(process.env); //mengecek enviorment dari file nodeJS

process.exit(1);

console.info("code ini tidak di eksekusi karena sudah di process.exit");