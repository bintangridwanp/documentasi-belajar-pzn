package bintangridwanp.application;

import bintangridwanp.data.Produk;

public class ProdukApp {
    public static void main(String[] args) {

        Produk produk = new Produk();
        produk.setId("Beras");
        System.out.println(produk.getId());
        System.out.println(produk.isTersedia());

        //produk.setId(null);
        System.out.println(produk.getId());
    }
}
