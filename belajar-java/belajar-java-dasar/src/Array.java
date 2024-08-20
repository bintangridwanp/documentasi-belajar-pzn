public class Array {
    public static void main(String[] args) {

        String[] iniArray;
        iniArray = new String[3];
        iniArray[0] = "Luffy";
        iniArray[1] = "D";
        iniArray[2] = "Monkey";

        System.out.println(iniArray[0]);
        System.out.println(iniArray[1]);
        System.out.println(iniArray[2]);

        String[] iniArray2 = new String[2];
        iniArray2[0] = "Vinsmoke";
        iniArray2[1] = "Sanji";

        System.out.println(iniArray2[0]);
        System.out.println(iniArray2[1]);

        iniArray2[1] = "zoro";

        System.out.println(iniArray2[1]);

        int[] iniArray3 = new int[]{
                1, 2, 4, 5, 6, 7, 8, 9,
        };

        long[] iniArray4 = {
                10L, 20L, 30L
        };

        System.out.println(iniArray4.length);

        String[][] nama = {
                {"luffy", "zoro"},
                {"luffy", "nami"},
                {"luffy", "sanji"}
        };

        System.out.println(nama[0][0]);
        System.out.println(nama[1][0]);

    }
}
