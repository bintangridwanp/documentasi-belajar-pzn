public class DoWhileLoop {
    public static void main(String[] args) {

        // Inisialisasi variabel counter dengan nilai awal 100
        var counter = 1;

        // Perulangan do-while yang akan dijalankan setidaknya sekali
        do {
            // Mencetak "Perulangan" diikuti dengan nilai counter pada iterasi saat ini
            System.out.println("Perulangan " + counter);
            // Menambah nilai counter dengan 1
            counter++;
        } while (counter <= 10); // Kondisi yang diperiksa setelah blok do dijalankan

    }
}
