<?php

$data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

$function_map = fn (int $nilai) => $nilai * 10;
$hasil = array_map($function_map, $data);

var_dump($hasil);

rsort($data);
var_dump($data);

sort($data);
var_dump($data);

var_dump(array_keys($data));
var_dump(array_values($data));

$nama_lengkap = [
    "nama_awal" => "luffy",
    "nama_tengah" => "D",
    "nama_akhir" => "monkey"
];

var_dump(array_keys($nama_lengkap));
var_dump(array_values($nama_lengkap));
