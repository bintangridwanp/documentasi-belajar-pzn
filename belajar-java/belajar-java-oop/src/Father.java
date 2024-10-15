public class Father {
    String name;
    String pekerjaan;

    Father(String name) {
        this.name = name;
    }
    Father(String nama, String pekerjaan) {
        this.name = nama;
        this.pekerjaan = pekerjaan;
    }

    void sapaHalo(String name){
        System.out.println(" Hallo "+name + " saya bapak nya " + this.name);
    }


}
