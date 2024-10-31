package bintangridwanp.application;

import bintangridwanp.data.ManchesterUnited;
import bintangridwanp.data.Pemain;

public class ManchesterUnitedApp {
    public static void main(String[] args) {

        ManchesterUnited manchester = new Pemain();

        manchester.Defends();
        manchester.Gelandang();
        manchester.Striker();
        System.out.println(manchester.Kiper());

        System.out.println(" ");
        manchester.Pelatih();

        System.out.println(" ");
        manchester.Pemilik();

    }
}
