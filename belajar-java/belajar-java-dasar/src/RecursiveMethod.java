public class RecursiveMethod {
    public static void main(String[] args) {

        // Memanggil dan mencetak hasil dari metode factorialLoop dengan argumen 10
        System.out.println(factorialLoop(10));
        // Memanggil dan mencetak hasil dari metode factorialRecursive dengan argumen 10
        System.out.println(factorialRecursive(10));

        // Memanggil metode loop dengan argumen 1000
        loop(1000);
    }

    // Metode untuk menghitung faktorial menggunakan loop
    static int factorialLoop(int value) {
        var result = 1;

        // Menggunakan loop untuk menghitung faktorial
        for (var counter = 1; counter <= value; counter++) {
            result *= counter;
        }

        return result; // Mengembalikan hasil faktorial
    }

    // Metode untuk menghitung faktorial menggunakan rekursi
    static int factorialRecursive(int value) {
        // Basis rekursi: jika value sama dengan 1, kembalikan 1
        if (value == 1) {
            return 1;
        } else {
            // Memanggil dirinya sendiri untuk menghitung faktorial
            return value * factorialRecursive(value - 1);
        }
    }

    // Metode untuk melakukan loop menggunakan rekursi
    static void loop(int value) {
        // Basis rekursi: jika value sama dengan 0, cetak "Selesai"
        if (value == 0) {
            System.out.println("Selesai");
        } else {
            // Mencetak nilai saat ini dan memanggil dirinya sendiri dengan nilai yang berkurang
            System.out.println("Loop " + value);
            loop(value - 1);
        }
    }
}
