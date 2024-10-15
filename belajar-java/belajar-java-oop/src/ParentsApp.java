public class ParentsApp {
    public static void main(String[] args) {

    var father  = new Father("Luffy");
    father.sapaHalo("Koby");

    var mother = new Mother("Nami");
    mother.sapaHalo("Nico Robin");

        System.out.println(mother);
        System.out.println(mother.toString());
        System.out.println(father);
        System.out.println(father.toString());

    }
}
