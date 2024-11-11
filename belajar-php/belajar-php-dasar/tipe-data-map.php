<?php
//tipe data map dan array sebenarnya sama saja hanya beda di key nya, index array di php itu bisa berubaha-rubah

//cara membuat map satu
$biodata = array(
    "nama" => "luffy",
    "umur" => 19,
    "alamat" => "east blue",
    "title" => "kapten",

);

//cara membuat map dua
$onepice = [
    "nama" => "luffy",
    "umur" => 19,
    "alamat" => "east blue",
    "title" => "kapten",

];


var_dump($biodata["nama"]);
var_dump($onepice["nama"]);


//cara membuar array di dlm array
$Ggmu = [
    "club" => "Manchester United",
    "manager" => "ruben",
    "pemain" => array(
        "kiper" => "onana",
        "bek" => "diogo dalot"
    )
];

var_dump($Ggmu["pemain"]["kiper"]);
