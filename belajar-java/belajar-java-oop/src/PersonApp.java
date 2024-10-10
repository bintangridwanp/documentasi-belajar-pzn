public class PersonApp {
    public static void main(String[] args) {

        var person1 = new Person("luffy", 19, "east blue");
        System.out.println(person1.name);
        System.out.println(person1.age);
        System.out.println(person1.address);
        System.out.println(person1.country);

        System.out.println("");

        Person Person2 = new Person("zoro", 21, "east blue");
        System.out.println(Person2.name);
        Person2.hello("nami");

        System.out.println("");

        Person Person3;
        Person3 = new Person("nico robin", 23, "west blue");
        System.out.println(Person3.name);


    }
}
