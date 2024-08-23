public class WhileLoop {
    public static void main(String[] args) {

        // Inisialisasi variabel counter dengan nilai awal 1
        var counter = 1;

        // Perulangan while yang berjalan selama nilai counter <= 10
        while (counter <= 10) {
            // Mencetak "Perulangan" diikuti dengan nilai counter pada setiap iterasi
            System.out.println("Perulangan " + counter);

            // Menambah nilai counter dengan 1 pada setiap iterasi
            counter++;
        }

    }
}
