function sapa(name){
    // tag debugger menjadi breakpoint
    debugger;
    return(`Hello ${name}`);
}

var name = "luffy";

debugger;
console.log(sapa(name))

// Perintah Debugger
// 1. cont, c : melanjutkan eksekusi code program.
// 2. next, n : melanjutkan step code berikutnya.
// 3. step, s : masuk ke dalam function.
// 4. out, o : keluar ke sebuah function.
// 5. pause : menghentikan running code.