public class MethodVariableArgument {
    public static void main(String[] args) {

        // Membuat array integer values dan memanggil metode sayCongrats dengan nama "Eko" dan array values
        int[] values = {80, 50, 50, 50, 80};
        sayCongrats("Eko", values);

        // Memanggil metode sayCongrats dengan nama "Budi" dan beberapa nilai langsung sebagai argumen
        sayCongrats("Budi", 80, 90, 76, 80);

    }

    // Metode sayCongrats dengan parameter name dan variable-length arguments (varargs) untuk values
    static void sayCongrats(String name, int... values) {
        var total = 0;

        // Perulangan untuk menjumlahkan semua nilai yang diberikan dalam values
        for (var value : values) {
            total += value;
        }

        // Menghitung nilai rata-rata dari total nilai
        var finalValue = total / values.length;

        // Mengecek apakah rata-rata nilai lebih besar atau sama dengan 75
        if (finalValue >= 75) {
            System.out.println("Selamat " + name + ", Anda Lulus");
        } else {
            System.out.println("Maaf " + name + ", Anda Tidak Lulus");
        }
    }
}
