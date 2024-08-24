public class Scope {
    public static void main(String[] args) {
        // Memanggil metode sayHello dengan argumen "Eko"
        sayHello("Scope");
        // Memanggil metode sayHello dengan argumen string kosong
        sayHello("");
    }

    // Metode untuk menyapa berdasarkan nama
    static void sayHello(String name) {
        // Variabel hello bersifat lokal, menyimpan pesan sapaan
        String hello = "Hello " + name;

        // Memeriksa apakah nama tidak kosong atau blank
        if (!name.isBlank()) {
            // Variabel hi bersifat lokal dalam blok if, menyimpan sapaan alternatif
            String hi = "Hi " + name;
            // Mencetak pesan sapaan alternatif
            System.out.println(hi);
        }

        // Mencetak pesan sapaan dasar
        System.out.println(hello);
    }
}
