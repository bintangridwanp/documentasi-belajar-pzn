package bintangridwanp.application;

public class eqqualsApp {
    @Override
    public boolean equals(Object obj) {
        return super.equals(obj);
    }

    @Override
    public int hashCode() {
        return super.hashCode();
    }

    public static void main(String[] args) {

        String FirstName = "James";
        FirstName = FirstName + " " + "Bond";
        System.out.println(FirstName);

        String name = "James Bond";
        System.out.println(name);

        String name_dua = "James Bond";
        System.out.println(name_dua);

        System.out.println(FirstName == name);
        System.out.println(name == name_dua);

    }
}
