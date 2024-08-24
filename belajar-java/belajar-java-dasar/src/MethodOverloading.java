public class MethodOverloading {
    public static void main(String[] args) {

        // Memanggil metode sayHello() tanpa argumen
        sayHello();
        // Memanggil metode sayHello(String name) dengan satu argumen
        sayHello("Luffy");
        // Memanggil metode sayHello(String firstName, String lastName) dengan dua argumen
        sayHello("luffy", "D");
    }

    // Metode sayHello() tanpa parameter
    static void sayHello(){
        System.out.println("Hello");
    }

    // Metode sayHello(String name) dengan satu parameter
    static void sayHello(String name){
        System.out.println("Hello " + name);
    }

    // Metode sayHello(String firstName, String lastName) dengan dua parameter
    static void sayHello(String firstName, String lastName){
        System.out.println("Hello " + firstName + " " + lastName);
    }
}
