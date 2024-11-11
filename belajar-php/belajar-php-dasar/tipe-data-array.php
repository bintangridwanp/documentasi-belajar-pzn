<?php

$angka = array(1, 2, 3, 4, 5, 0.1, 0.2, 0.3, 0.4, 0.5);
var_dump($angka);
print_r($angka);

$onePiece = ["luffy", "zoro", "chopper", "nami", "nico robin", "brok", "sanji", "franky", "jimbei"];
foreach($onePiece as $item){
    echo $item . "\n";
}

//operasi array

//menecek nilai array menggunakan index
var_dump($onePiece[3]);
print_r($onePiece[4]);

//mengganti nilai array
$onePiece[0] = "budiono siregar";
var_dump($onePiece);

//menghapus index array
unset($onePiece[0]);
var_dump($onePiece);

//menambahkan data array
$onePiece[] = "budi";
var_dump($onePiece);

//mengecek berapa panjang total array.
var_dump(count($onePiece));


