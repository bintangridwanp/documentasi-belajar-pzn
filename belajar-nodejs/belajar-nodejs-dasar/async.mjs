function contohPromise(){
    return Promise.resolve("Luffy");
}

const name = await contohPromise();
console.log(name);