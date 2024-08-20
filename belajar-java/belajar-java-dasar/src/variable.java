public class variable {
    public static void main(String[] args) {

        // Deklarasi variabel tipe String tanpa inisialisasi
        String name;
        // Inisialisasi variabel name dengan nilai "Luffy D monkey"
        name = "Luffy D monkey";

        // Deklarasi dan inisialisasi variabel tipe int dengan nilai 30
        int age = 30;

        // Deklarasi dan inisialisasi variabel tipe String dengan nilai "Vinsmoke Sanji"
        String teman = "Vinsmoke Sanji";

        // Mencetak nilai variabel age (umur)
        System.out.println(age);

        // Mencetak string "Name: " diikuti dengan nilai variabel name
        System.out.println("Name: " + name);

        // Mencetak nilai variabel name (nama)
        System.out.println(name);

        // Mengubah nilai variabel name menjadi "Luffy taroa"
        name = "Luffy taroa";

        // Mencetak nilai variabel name setelah diubah
        System.out.println(name);

        // Menggunakan kata kunci var untuk mendeklarasikan variabel dengan inferensi tipe
        var cita_cita = "Menjadi raja bajak laut";

        var namaAwal = "Luffy";
        var namaTengah = "D";
        var namaAkhir = "monkey";

        // Mencetak nilai variabel namaAwal, namaTengah, namaAkhir, dan cita_cita
        System.out.println(namaAwal);
        System.out.println(namaTengah);
        System.out.println(namaAkhir);
        System.out.println(cita_cita);

        // Mendeklarasikan variabel final, yang nilainya tidak bisa diubah setelah diinisialisasi
        final String study = "belajar java dasar by pzn";

        // Mencetak nilai variabel study
        System.out.println(study);

    }
}
