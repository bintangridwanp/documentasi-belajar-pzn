public class SwitchStatement {
    public static void main(String[] args) {

        var nilai = "E";

        switch (nilai) {
            case "A":
                System.out.println("A");
                break;
            case "B":
                System.out.println("B");
                break;
            case "C":
                System.out.println("C");
                break;
            case "D":
                System.out.println("D");
            default:
                System.out.println("Mungkin anda kurang beruntung");
                break;
        }

        switch (nilai) {
            case "A" -> System.out.println("A");
            case "B", "C" -> System.out.println("B dan C");
            case "D" -> System.out.println("D");
            default -> System.out.println("Mungkin anda kurang beruntung");
        }

        String ucapan;
        switch (nilai) {
            case "A" -> ucapan = "Wow, Anda Lulus Dengan Baik";
            case "B", "C" -> ucapan = "Nilai Anda Cukup Baik";
            case "D" -> ucapan = "Anda Tidak Lulus";
            default -> {
                ucapan = "Muungkin Anda Salah Jurusan";
            }
        }

        System.out.println(ucapan);

        ucapan = switch (nilai) {
            case "A": yield "A";
            case "B", "C": yield "B dan C";
            case "D": yield "D";
            default:
                yield "Mungkin anda kurang beruntung";
        };

        System.out.println(ucapan);

    }
}
