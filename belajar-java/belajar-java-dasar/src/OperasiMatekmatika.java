public class OperasiMatekmatika {
    public static void main(String[] args) {

        var a = 10;  // Deklarasi variabel 'a' dengan nilai 10
        var b = 500; // Deklarasi variabel 'b' dengan nilai 500

        // Operasi penjumlahan
        System.out.println(a + b); // Menampilkan hasil penjumlahan a + b (10 + 500)

        // Operasi pengurangan
        System.out.println(a - b); // Menampilkan hasil pengurangan a - b (10 - 500)

        // Operasi perkalian
        System.out.println(a * b); // Menampilkan hasil perkalian a * b (10 * 500)

        // Operasi pembagian
        System.out.println(a / b); // Menampilkan hasil pembagian a / b (10 / 500)

        // Operasi modulus (sisa bagi)
        System.out.println(a % b); // Menampilkan hasil modulus a % b (10 % 500)

        var c = 20; // Deklarasi variabel 'c' dengan nilai 20

        // Operasi penambahan dengan assignment (c = c + 10)
        c += 10;
        System.out.println(c); // Menampilkan nilai 'c' setelah penambahan (20 + 10)

        // Operasi pengurangan dengan assignment (c = c - 10)
        c -= 10;
        System.out.println(c); // Menampilkan nilai 'c' setelah pengurangan (30 - 10)

        // Operasi perkalian dengan assignment (c = c * 10)
        c *= 10;
        System.out.println(c); // Menampilkan nilai 'c' setelah perkalian (20 * 10)

        // Operasi pembagian dengan assignment (c = c / 10)
        c /= 10;
        System.out.println(c); // Menampilkan nilai 'c' setelah pembagian (200 / 10)

        // Operasi modulus dengan assignment (c = c % 10)
        c %= 10;
        System.out.println(c); // Menampilkan nilai 'c' setelah modulus (20 % 10)

        var d = 10; // Deklarasi variabel 'd' dengan nilai 10
        var e = 20; // Deklarasi variabel 'e' dengan nilai 20

        // Operasi increment (d = d + 1)
        ++d;
        System.out.println(d); // Menampilkan nilai 'd' setelah increment (10 + 1)

        // Operasi decrement (e = e - 1)
        --e;
        System.out.println(e); // Menampilkan nilai 'e' setelah decrement (20 - 1)

        // Operasi negasi (membalik nilai boolean)
        System.out.println(!true); // Menampilkan hasil dari negasi true (false)
    }
}
