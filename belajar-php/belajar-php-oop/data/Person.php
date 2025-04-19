<?php

class Person
{
    const PENULIS = "Eko";

    var string $nama;
    var ?string $nomerhp = null;
    var string $alamat;
    // Default value for negara
    var string $negara = "Indonesia";

    function __construct(string $nama, string $alamat){
        $this -> nama = $nama;
        $this -> alamat = $alamat;
    }

    function sapaHalo(?string $alamat){
        if (is_null($alamat)){
            echo "Halo {$this -> nama}, saya tidak tahu alamat anda" . PHP_EOL;
        } else {
            echo "Halo {$this -> nama}, saya tahu alamat anda ternyata di {$this -> alamat}" . PHP_EOL;
        }
    }

    function informasi(){
        echo "Penulis : " . self::PENULIS . PHP_EOL;
    }

    function __destruct(){
        echo "Objek {$this -> nama} sudah dihapus" . PHP_EOL;
    }

}

