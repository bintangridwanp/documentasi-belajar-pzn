import readline from "readline";

function askQuestion(query) {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });

  return new Promise((resolve) => {
    rl.question(query, (answer) => {
      resolve(answer);
      rl.close();
    });
  });
}


// Contoh penggunaan
(async () => {
  try {
    const name = await askQuestion("Apa nama Anda? ");
    console.log(`Halo, ${name}!`);
    const age = await askQuestion("Berapa umur Anda? ")
    console.log(`${age}`)
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
})();
