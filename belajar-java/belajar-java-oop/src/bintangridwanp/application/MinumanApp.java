package bintangridwanp.application;

import bintangridwanp.data.Juice;
import bintangridwanp.data.Minuman;

public class MinumanApp {
    public static void main(String[] args) {

        //Minuman minuman = new Minuman; ERROR
        Juice juice = new Juice();
        juice.nama = "Juice Buah Naga";
        System.out.println(Juice.nama);


    }
}
