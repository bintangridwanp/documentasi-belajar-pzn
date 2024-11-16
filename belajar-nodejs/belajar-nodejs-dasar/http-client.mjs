import https from "https";

const endpoint = "https://eo8ax959k5g80up.m.pipedream.net"
const request = https.request(endpoint, {
    method: "POST",
    headers: {
        "Content-Type": "application/json",
        "Accept": "applicationc/json"
    }
}, (response) => {
    response.addListener("data", (data) => {
        console.info(`Receive data: ${data.toString()}`);
    })
}) 

const body = JSON.stringify({
    name: "luffy d mongkey",
    age: 19,
})

request.write(body);
request.end();
