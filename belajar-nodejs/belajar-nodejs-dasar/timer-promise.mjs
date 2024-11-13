import timer from "timers/promises"

// console.info(new Date());
// const waktu = await timer.setTimeout(5000, "waktu sudah berjalan selama 5 detik");
// console.info(new Date());

// console.info(waktu);

for await (const startTime of timer.setInterval(1000, new Date())) {
    console.info(`Waktu sekarang menunjukan ${startTime}`);
}