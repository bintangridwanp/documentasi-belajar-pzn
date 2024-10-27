package bintangridwanp.application;

import bintangridwanp.data.Buah;

public class BuahApp {
    public static void main(String[] args) {

        Buah buah = new Buah("Anggur", 15000);

        System.out.println(buah);
        System.out.println(buah.nama);
        System.out.println(buah.harga);

    }
}
