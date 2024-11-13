import net from 'net';

const server = net.createServer(function(client){
    console.log("client sudah terhubung");
    client.on("data", function(data){
        console.info(`menerima data cari client: ${data.toString()}`);
        client.write(`Halo ${data.toString()}\r\n`);
    })
});

server.listen(3000, "localhost")   