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
        if(nenek instanceof Father){
            Father father = (Father) nenek;
            System.out.println("Halo father " + father.name);
        } else if (nenek instanceof Mother) {
            Mother mother = (Mother) nenek;
            System.out.println("Halo mother " + mother.name);
        } else {
            System.out.println("Halo non father and non mother " + nenek.name);
        }

    }
}
