package bintangridwanp.data;

class Buku {
    int id;
    String name;
}

 class Horor extends Buku {
    final void buku(int id, String name){
        System.out.println(01 +  "Nenek Sihir");
    }
}

// Ini error karena class horor sudah final class
class NenekSihir extends Horor {

    // Error karena void sudah di set final di class Horor.
    //    void buku(int id, String name){
    //        System.out.println(01 +  "Nenek Sihir");
    //    }

}


