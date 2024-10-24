public class OrangTua {
    String name;

    void halo() {
        System.out.println("Hallo Orang Tua" + name);
    }
}

class Anak extends OrangTua {
    void halo(){
        System.out.println("Hallo Anak " + super.name);
        System.out.println(super.name + " adalah anak dari Orang tua");
    }
}
