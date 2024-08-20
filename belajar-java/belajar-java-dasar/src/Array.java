public class Array {
    public static void main(String[] args) {

        // Deklarasi array tipe String tanpa inisialisasi
        String[] iniArray;
        // Inisialisasi array iniArray dengan panjang 3
        iniArray = new String[3];
        // Mengisi elemen array dengan nama
        iniArray[0] = "Luffy";
        iniArray[1] = "D";
        iniArray[2] = "Monkey";

        // Mencetak elemen pertama, kedua, dan ketiga dari array iniArray
        System.out.println(iniArray[0]); // Output: Luffy
        System.out.println(iniArray[1]); // Output: D
        System.out.println(iniArray[2]); // Output: Monkey

        // Inisialisasi array tipe String dengan panjang 2
        String[] iniArray2 = new String[2];
        // Mengisi elemen array iniArray2 dengan nama
        iniArray2[0] = "Vinsmoke";
        iniArray2[1] = "Sanji";

        // Mencetak elemen pertama dan kedua dari array iniArray2
        System.out.println(iniArray2[0]); // Output: Vinsmoke
        System.out.println(iniArray2[1]); // Output: Sanji

        // Mengubah nilai elemen kedua dari array iniArray2 menjadi "zoro"
        iniArray2[1] = "zoro";

        // Mencetak elemen kedua dari array iniArray2 setelah diubah
        System.out.println(iniArray2[1]); // Output: zoro

        // Inisialisasi array tipe int dengan beberapa elemen
        int[] iniArray3 = new int[]{
                1, 2, 4, 5, 6, 7, 8, 9, // Array ini memiliki 8 elemen
        };

        // Inisialisasi array tipe long dengan tiga elemen
        long[] iniArray4 = {
                10L, 20L, 30L // Array ini memiliki 3 elemen
        };

        // Mencetak panjang array iniArray4 (jumlah elemen)
        System.out.println(iniArray4.length); // Output: 3

        // Deklarasi dan inisialisasi array dua dimensi tipe String
        String[][] nama = {
                {"luffy", "zoro"},  // Array pertama
                {"luffy", "nami"},  // Array kedua
                {"luffy", "sanji"}   // Array ketiga
        };

        // Mencetak elemen dari array dua dimensi
        System.out.println(nama[0][0]); // Output: luffy (elemen pertama dari array pertama)
        System.out.println(nama[1][0]); // Output: luffy (elemen pertama dari array kedua)

    }
}
