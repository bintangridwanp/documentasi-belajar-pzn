<?php

// Menggabungkan dua array
$a = [
    "name" => "luffy"
];

$b = [
    // Jika terdapat key yang sama antara array $a dan $b, nilai di array $a akan dipertahankan
    // dan nilai di array $b dengan key yang sama akan diabaikan.
    "name" => "zoro",
    "age" => 19
];

// var_dump($a + $b); // Menampilkan hasil penggabungan $a dan $b

// Mengecek kesamaan key-value antar array (==)
$c = [
    "buah_dua" => "durian",
    "buah_satu" => "mangga"
];

$d = [
    "buah_satu" => "mangga",
    "buah_dua" => "durian"
];
// Dalam pengecekan kesamaan dengan operator ==, yang penting adalah key dan value-nya sama,
// urutan elemen tidak diperhatikan.
// var_dump($c == $d); // true jika key-value sama, walaupun urutannya berbeda
// var_dump($c != $d); // true jika key-value tidak sama

// Mengecek kesamaan key-value dan urutan antar array (===)
$e = [
    "laptop_satu" => "Lenovo X270",
    "laptop_dua" => "Macbook Pro chip m2"
]; 

$f = [
    "laptop_satu" => "Lenovo X270",
    "laptop_dua" => "Macbook Pro chip m2"
]; 

// Dalam pengecekan kesamaan dengan operator ===, key, value, dan urutan elemen harus persis sama.
var_dump($e === $f); // true jika key, value, dan urutan elemen sama
var_dump($e !== $f); // true jika ada perbedaan pada key, value, atau urutan elemen
