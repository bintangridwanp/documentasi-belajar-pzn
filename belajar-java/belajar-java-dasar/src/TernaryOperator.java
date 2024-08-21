public class TernaryOperator {
    public static void main(String[] args) {

        var nilai = 70; // Deklarasi variabel 'nilai' dengan nilai 70

        // Menggunakan operator ternary untuk menentukan ucapan berdasarkan nilai
        String ucapan  = nilai >= 75 ? "Selamat anda lulus" : "Anda belum beruntung";

        System.out.println(ucapan); // Menampilkan ucapan berdasarkan hasil evaluasi operator ternary

    }
}
