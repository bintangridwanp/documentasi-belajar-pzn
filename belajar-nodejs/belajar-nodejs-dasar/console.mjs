import fs from "fs";
import { Console } from "console";

const file = fs.createWriteStream("application.log");

const log = new Console ({
    stdout: file,
    stderr: file,
})

log.info("Hello luffy");
log.error("Ups");

const mahasiswa = {
    name: "luffy D monkey",
    age: 19,
    nim: 12345
}

log.table(mahasiswa);