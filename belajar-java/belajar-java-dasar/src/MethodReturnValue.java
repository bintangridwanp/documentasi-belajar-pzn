public class MethodReturnValue {
    public static void main(String[] args) {

        // Memanggil metode sum dengan argumen 100 dan 100, menyimpan hasilnya dalam variabel result1
        var result1 = sum(100, 100);
        // Mencetak nilai result1 ke console
        System.out.println(result1);

        // Memanggil metode sum dengan argumen 200 dan 200, dan langsung mencetak hasilnya ke console
        System.out.println(sum(200, 200));

        // Memanggil metode hitung dengan operasi "+" dan mencetak hasilnya
        System.out.println(hitung(100, "+", 100));
        // Memanggil metode hitung dengan operasi "-" dan mencetak hasilnya
        System.out.println(hitung(200, "-", 100));
        // Memanggil metode hitung dengan operasi yang tidak dikenal dan mencetak hasilnya (akan menghasilkan 0)
        System.out.println(hitung(200, "salah", 100));
    }

    // Metode sum yang menerima dua parameter integer, menjumlahkannya, dan mengembalikan hasilnya
    static int sum(int value1, int value2) {
        var total = value1 + value2;
        return total; // Mengembalikan hasil penjumlahan
    }

    // Metode hitung yang menerima dua nilai integer dan satu operasi, mengembalikan hasil perhitungan berdasarkan operasi yang diberikan
    static int hitung(int value1, String operasi, int value2) {
        switch (operasi) {
            case "+":
                return value1 + value2; // Jika operasi adalah "+", kembalikan hasil penjumlahan
            case "-":
                return value1 - value2; // Jika operasi adalah "-", kembalikan hasil pengurangan
            default:
                return 0; // Jika operasi tidak dikenal, kembalikan nilai 0
        }
    }
}
