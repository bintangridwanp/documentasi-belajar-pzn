public class ForEach {
    public static void main(String[] args) {

        // Inisialisasi array String dengan beberapa elemen
        String[] names = {
                "Luffy", "D", "monkey",
                "Programmer", "Zaman", "Now"
        };

        // Perulangan for biasa, menggunakan indeks untuk mengakses setiap elemen array
        for (var i = 0; i < names.length; i++) {
            // Mencetak setiap elemen array berdasarkan indeks
            System.out.println(names[i]);
        }

        // Memisahkan hasil perulangan dengan pesan "FOR EACH"
        System.out.println("FOR EACH");

        // Perulangan for-each, yang secara langsung mengakses setiap elemen dalam array
        for (String name : names) {
            // Mencetak setiap elemen array tanpa perlu menggunakan indeks
            System.out.println(name);
        }
    }
}
