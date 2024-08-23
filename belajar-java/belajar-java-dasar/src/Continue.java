public class Continue {
    public static void main(String[] args) {

        // Perulangan for yang dimulai dari 1 hingga 100
        for (var counter = 1; counter <= 100; counter++) {
            // Jika counter adalah bilangan genap, perulangan dilanjutkan ke iterasi berikutnya tanpa menjalankan kode di bawah ini
            if (counter % 2 == 0) {
                continue; // Melompati iterasi jika kondisi terpenuhi
            }

            // Mencetak nilai counter jika merupakan bilangan ganjil
            System.out.println("Perulangan Ganjil " + counter);
        }

    }
}
