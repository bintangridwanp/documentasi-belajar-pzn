public class Break {
    public static void main(String[] args) {

        // Inisialisasi variabel counter dengan nilai awal 1
        var counter = 1;

        // Perulangan while yang berjalan terus menerus (infinite loop)
        while (true) {
            // Mencetak "Perulangan" diikuti dengan nilai counter pada iterasi saat ini
            System.out.println("Perulangan " + counter);
            // Menambah nilai counter dengan 1
            counter++;

            // Memeriksa apakah nilai counter lebih besar dari 10
            if (counter > 10) {
                // Jika kondisi terpenuhi, perulangan dihentikan dengan menggunakan break
                break;
            }
        }

        // Mencetak pesan setelah perulangan dihentikan
        System.out.println("Perulangan berhenti");
    }
}
