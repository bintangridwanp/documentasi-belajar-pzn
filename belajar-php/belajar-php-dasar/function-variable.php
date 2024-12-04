<?php

function hello_budi(){
    echo "Hello budi" . PHP_EOL;
}

function hello_joko(){
    echo "Hello joko" . PHP_EOL;
}

$sapa = "hello_budi";
$sapa();

$sapa = "hello_joko";
$sapa();

function sapa_hallo(string $name, $upper_lower){
    $nama_lengkap = $upper_lower($name); //callback
    echo "hallo $nama_lengkap" . PHP_EOL;
}

function ucapan_hallo(string $name) : string {
    return "selamat $name";
}

sapa_hallo("luffy", "ucapan_hallo");
sapa_hallo("luffy", "strtoupper");
sapa_hallo("luffy", "strtolower");