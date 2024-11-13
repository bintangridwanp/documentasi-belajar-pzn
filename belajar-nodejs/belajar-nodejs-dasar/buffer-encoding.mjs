const buffer = Buffer.from("luffy D monkey", "utf8");

console.info(buffer.toString());
console.info(buffer.toString("hex"));
console.info(buffer.toString("base64"));

const bufferBase64 = Buffer.from("bHVmZnkgRCBtb25rZXk=", "base64");
console.log(bufferBase64.toString());
