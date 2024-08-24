public class MethodParameter {
    public static void main(String[] args) {

        // Memanggil metode sayHello dengan berbagai pasangan nama sebagai argumen
        sayHello("Eko", "Kurniawan");
        sayHello("Budi", "Nugraha");
        sayHello("Joko", "Nugroho");

    }

    // Mendefinisikan metode sayHello dengan dua parameter: firstName dan lastName
    static void sayHello(String firstName, String lastName) {
        // Mencetak pesan "Hello" diikuti dengan firstName dan lastName yang diberikan sebagai argumen
        System.out.println("Hello " + firstName + " " + lastName);
    }
}
