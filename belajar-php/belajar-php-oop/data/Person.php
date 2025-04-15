<?php

class Person
{
    const PENULIS = "Eko";

    var string $nama;
    var ?string $nomerhp = null;
    var string $alamat;
    // Default value for negara
    var string $negara = "Indonesia";

    function sapaHalo(?string $alamat){
        if (is_null($alamat)){
            echo "Halo {$this -> nama}, saya tidak tahu alamat anda" . PHP_EOL;
        } else {
            echo "Halo {$this -> nama}, saya tahu alamat anda ternyata di {$this -> alamat}" . PHP_EOL;
        }
    }
}

