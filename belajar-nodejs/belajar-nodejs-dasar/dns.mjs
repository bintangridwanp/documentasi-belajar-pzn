import dns from "dns/promises";

const address = await dns.lookup("www.progammerzamannow.com");

console.info(address);
