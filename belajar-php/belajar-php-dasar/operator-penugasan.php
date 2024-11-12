<?php

$a = 10;
$b = 20;
$c = 30;

// penjumlahan
$sum = 0;
$sum += $a;
$sum += $b;
$sum += $c;

var_dump($sum); // hasilnya 60

// pengurangan
$sum = 0;
$sum -= $a;
$sum -= $b;
$sum -= $c;

var_dump($sum); // hasilnya -60

// perkalian
$sum = 1; // diinisialisasi dengan 1 untuk perkalian
$sum *= $a;
$sum *= $b;
$sum *= $c;

var_dump($sum); // hasilnya 6000

// pembagian
$sum = $a; // diinisialisasi dengan $a agar pembagian masuk akal
$sum /= $b;
$sum /= $c;

var_dump($sum); // hasilnya 0.016666...

// modulo
$sum = $a; // inisialisasi awal
$sum %= $b;
$sum %= $c;

var_dump($sum); // hasilnya 10, karena $a < $b dan $c
