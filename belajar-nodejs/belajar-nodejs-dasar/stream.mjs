import fs from "fs";

const writer = fs.createWriteStream("target.log");

writer.write("luffy\n");
writer.write("D\n");
writer.write("mongkey\n");
writer.end();

// Tunggu sampai penulisan selesai sebelum mulai membaca
writer.on("finish", () => {
    const reader = fs.createReadStream("target.log");
    reader.addListener("data", (data) => {
        console.info(data.toString());
    });
});
