public class Father extends Nenek {
    String pekerjaan;

    Father(String name) {
        super(name);
    }
    Father(String nama, String pekerjaan) {
        super(nama);
        this.pekerjaan = pekerjaan;
    }

    void sapaHalo(String name){
        System.out.println(" Hallo "+name + " saya bapak nya " + this.name);
    }


}
