// Mendefinisikan dan mengekspor kelas Orang
export class Orang {

    // Konstruktor kelas yang menerima satu parameter dan menetapkan ke properti name
    constructor(name) {
        this.name = name;
    }

    // Metode sapaHalo yang menerima satu parameter dan menampilkan pesan di konsol
    sapaHalo(name) {
        console.info(`hallo ${name}, perkenalkan nama ku ${this.name}`);
    }

}
