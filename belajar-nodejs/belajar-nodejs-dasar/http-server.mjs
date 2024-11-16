import http from "http";

const server = http.createServer((request, response) => {
    console.info(request.method);
    console.info(request.headers);
    console.info(request.url);

    if (request.method === "POST"){
        request.addListener("data", (data) => {
            response.setHeader("Content Type", "application/json");
            response.write(data);
            response.end();
        }) 
    } else {
        if (request.url === "/zoro"){
            response.write("Hello zoro");
        } else {
            response.write("Hello luffy");
        }
        response.end();
    }
    
});

server.listen(5000);