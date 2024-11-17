<?php

$mahasiswa = [
    "nama" => "luffy d monkey",
    "umur" => 19,
    "nim" => 12345,
    "fakultas" => "informatika",
    "jurusan" => "rekayasa perangkat lunak"
];
$cekMahasiswa = $mahasiswa["nim"] ?? "Nothing";

echo $cekMahasiswa . PHP_EOL;