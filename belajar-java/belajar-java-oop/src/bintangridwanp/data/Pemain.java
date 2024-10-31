package bintangridwanp.data;

public class Pemain implements ManchesterUnited, OwnerMu {

    public void Defends(){
        System.out.println("Maguaeri adalah bek terbaik");
    }

    public void Gelandang(){
        System.out.println("Bruno fernandes gelandang paling bapuk");
    }

    public void Striker() {
        System.out.println("Ronaldo adalah striker paling keren");
    }

    public String Kiper() {
        return "Lord baginda onana";
    }

    public void Pelatih(){
        System.out.println("Eric Ten Hag OUT!!!");
    }

    public void Pemilik(){
        System.out.println("Keluarga Glazer bukan Owner Manchester United lagi!!!");
    }

    public boolean Bermain(){
        return true;
    }


}
