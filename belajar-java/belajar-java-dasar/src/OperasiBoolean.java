public class OperasiBoolean {
    public static void main(String[] args) {

        var nilai = 80; // Deklarasi variabel 'nilai' dengan nilai 80
        var absen = 100; // Deklarasi variabel 'absen' dengan nilai 100

        // Mengecek apakah nilai >= 75, hasilnya true
        var lulusNilai = nilai >= 75;

        // Mengecek apakah absen >= 75, hasilnya true
        var lulusAbsen = absen >= 75;

        // Operasi logika AND untuk mengecek apakah keduanya lulus
        var lulus = lulusNilai && lulusAbsen;
        System.out.println(lulus); // Menampilkan hasil akhir (true jika keduanya true)
    }
}
