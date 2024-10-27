package bintangridwanp.application;

import bintangridwanp.data.Makanan;
//bisa menggunakan tanda * untuk menginport langsung semuanya.

public class MakananApp {
    public static void main(String[] args) {

        Makanan makanan = new Makanan("nasi goreng", 14500);
        System.out.println(makanan.nama + makanan.harga);


    }
}
