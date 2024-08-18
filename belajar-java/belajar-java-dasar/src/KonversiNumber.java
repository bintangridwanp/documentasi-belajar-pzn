public class KonversiNumber {

    public static void main(TipeDataString[] args) {

        // Deklarasi variabel tipe data byte dengan nilai 100
        byte iniByte = 100;

        // Implicit casting: Konversi dari byte ke short (lebih besar) secara otomatis
        short iniShort = iniByte;

        // Implicit casting: Konversi dari short ke int (lebih besar) secara otomatis
        int iniInt = iniShort;

        // Deklarasi variabel tipe data int dengan nilai 1000
        int iniInt2 = 1000;

        // Explicit casting: Konversi dari int ke byte (lebih kecil) secara manual
        // Nilai int 1000 di-cast ke byte, sehingga nilai akan berubah karena byte memiliki rentang nilai yang lebih kecil
        byte inibyte3 = (byte) iniInt2;

        // Explicit casting: Konversi dari int ke byte secara manual
        // Ini menghasilkan pemotongan nilai jika int lebih besar dari kapasitas byte
        byte iniByte2 = (byte) iniInt;

    }

}
