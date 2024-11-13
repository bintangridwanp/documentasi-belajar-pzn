import net from "net";

const client = net.createConnection({
    port: 3000,
    host: "localhost"
});

client.addListener("data", (data) => {
    console.info(`menerima data dari server ${data.toString()}`);
});

setInterval(() => {
    client.write("luffy\r\n");
}, 5000)