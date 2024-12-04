<?php

function penjumlahan(int $a, int $b){
    $hasil = $a + $b;
    return $hasil;
}

$total = penjumlahan(10, 10);
var_dump($total);

function kalkulus( int $nilai){
    if ($nilai >= 90){
        return "Selamat kamu mendapatkan nilai A";
    } else if ($nilai >= 80){
        return "Selamat kamu mendapatkan nilai B";
    } else if ($nilai >= 70){
        return "Selamat kamu mendapatkan nilai C";
    } else if ($nilai >= 60){
        return "Selamat kamu mendapatkan nilai D";
    } else {
        return "Sayang sekali, kamu harus mengulangin nya semester depan.";
    }
}

$nilai_kalkulus = kalkulus(100);
var_dump($nilai_kalkulus);

function perulangan($i) : string // untuk memberitahu bahawa return atau kembaliannya adalah string.
{
    for($a = 1; $a <= $i; $a++){
        echo "Perulangan ke " . $a . PHP_EOL;
    }
    return "Perulangan selesai";
}

$ulang = perulangan(100);
var_dump($ulang);