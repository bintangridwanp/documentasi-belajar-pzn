import { constants } from "buffer";
import { Console } from "console";
import { after } from "node:test";
import util from "util";

const name = "luffy D monkey";
const age = 19;

console.info(`name: ${name} age: ${age}`);
console.info(util.format("name: %s age: %i", name, age));

const mahasiswa = {
    name : "luffy d monkey",
    nim : 123456
};

console.info(`mahasiswa: ${JSON.stringify(mahasiswa)}`);
console.info(util.format("mahasiswa: %j", mahasiswa));
