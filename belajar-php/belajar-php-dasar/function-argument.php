<?php

function sapa($firstname, $middlename, $lastname = "data default"){
    echo "selamat pagi $firstname $middlename $lastname" . PHP_EOL;
}

sapa("luffy", "D", "monkey");
sapa("luffy", "D");

function penjumlahan(int $a, int $b){
    $hasil = $a + $b;
    echo $hasil . PHP_EOL;
}

penjumlahan(10, 14);

function jumlah(...$nilai){

    $hasil = 0;
    foreach($nilai as $angka){
        $hasil = $hasil + $angka;
    }
    echo "Hasil dari: " . implode(", ", $nilai) . " = $hasil" . PHP_EOL;
}

jumlah(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

$nilai = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
jumlah(...$nilai);
