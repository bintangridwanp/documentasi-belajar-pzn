public class Method {
    public static void main(String[] args) {

        // Memanggil metode sayHelloWorld tiga kali
        sayHelloWorld();
        sayHelloWorld();
        sayHelloWorld();

    }

    // Mendefinisikan metode sayHelloWorld yang tidak mengembalikan nilai (void)
    static void sayHelloWorld() {
        // Mencetak "Hello World 1", "Hello World 2", dan "Hello World 3" setiap kali metode dipanggil
        System.out.println("Hello World 1");
        System.out.println("Hello World 2");
        System.out.println("Hello World 3");
    }
}
