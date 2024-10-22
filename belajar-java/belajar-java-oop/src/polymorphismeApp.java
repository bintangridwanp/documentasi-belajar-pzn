public class polymorphismeApp {
    public static void main(String[] args) {

        Nenek nenek = new Nenek("Luffy");
        nenek.sapaHalo("Choper");

        nenek = new Father("zoro");
        nenek.sapaHalo("nami");

        nenek = new Mother("nico robin");
        nenek.sapaHalo("kobi");

        sapaHalo(new Nenek("satu"));
        sapaHalo(new Father("dua"));
        sapaHalo(new Mother("tiga"));

    }
        static void sapaHalo(Nenek nenek){
            System.out.println("Halo " + nenek.name);
        }
}
