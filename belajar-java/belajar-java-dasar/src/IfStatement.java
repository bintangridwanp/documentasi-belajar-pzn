public class IfStatement {
    public static void main(String[] args) {

        var nilai = 80; // Deklarasi variabel 'nilai' dengan nilai 80
        var absen = 80; // Deklarasi variabel 'absen' dengan nilai 80

        // Mengecek apakah nilai dan absen >= 75, jika ya, cetak "Selamat anda lulus"
        if (nilai >= 75 && absen >= 75) {
            System.out.println("Selamat anda lulus");
        } else {
            // Jika salah satu atau keduanya kurang dari 75, cetak "Silahkan Coba Lagi Tahun Depan"
            System.out.println("Silahkan Coba Lagi Tahun Depan");
        }

        // Kondisi berjenjang untuk menentukan grade berdasarkan nilai dan absen
        if (nilai >= 80 && absen >= 80) {
            System.out.println("Nilai Anda A"); // Cetak "Nilai Anda A" jika nilai dan absen >= 80
        } else if (nilai >= 70 && absen >= 70) {
            System.out.println("Nilai Anda B"); // Cetak "Nilai Anda B" jika nilai dan absen >= 70
        } else if (nilai >= 60 && absen >= 60) {
            System.out.println("Nilai Anda C"); // Cetak "Nilai Anda C" jika nilai dan absen >= 60
        } else if (nilai >= 50 && absen >= 50) {
            System.out.println("Nilai Anda D"); // Cetak "Nilai Anda D" jika nilai dan absen >= 50
        } else {
            System.out.println("Nilai Anda E"); // Cetak "Nilai Anda E" jika nilai dan absen < 50
        }

    }
}
