package bintangridwanp.application;

import bintangridwanp.data.Binatang;
import bintangridwanp.data.Cat;

public class AnimalApp {
    public static void main(String[] args) {

        Binatang binatang = new Cat();
        binatang.name = "Kucing";
        binatang.suara();


    }
}
