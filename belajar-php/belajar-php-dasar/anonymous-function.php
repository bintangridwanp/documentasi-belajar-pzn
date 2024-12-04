<?php

$sapa = function($name){
    return "halo selamat malam $name" . PHP_EOL;
};

echo $sapa("luffy");

function selamat_tinggal(string $name, $upper_lower){
    $nama_lengkap = $upper_lower($name);
    echo "selamt tinggal $nama_lengkap" . PHP_EOL;
}

selamat_tinggal("nami", function($name){
    return strtoupper($name);
});

$nama_awal = "luffy";
$nama_tengah = "D";
$nama_akhir = "monkey";

$sapa_onepiece = function () use ($nama_awal) {
    echo "Hello $nama_awal" . PHP_EOL;
};

$sapa_onepiece();

$sapa_onepiece = function () use ($nama_awal, $nama_tengah, $nama_akhir) {
    echo "Hello $nama_awal $nama_tengah $nama_akhir" . PHP_EOL;
};

$sapa_onepiece();