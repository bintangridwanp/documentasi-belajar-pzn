function contohPromise(){
    return Promise.resolve("Luffy");
}

async function run(){
    const name = await contohPromise();
    console.log(name);
}

run()
