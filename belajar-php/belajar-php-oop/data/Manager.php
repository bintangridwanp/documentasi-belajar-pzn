<?php

class Manager {

    var string $nama;
    var string $title;

    public function  _construct(string $nama, string $title){
        $this -> nama = $nama;
        $this -> title = $title;
    }

    function sapaHalo(string $nama) : void {
        echo "Halo {$this -> nama}, saya {$nama}" . PHP_EOL;
    }
}

class VP extends Manager {
    
}