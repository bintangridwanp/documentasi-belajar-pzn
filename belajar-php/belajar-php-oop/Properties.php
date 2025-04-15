<?php

require_once 'data/Person.php';

$Person1 = new Person();
$Person1 -> nama = "Joko";
$Person1 -> alamat = "Solo";
$Person1 -> negara = "Indonesia";

$Person2 = new Person();
$Person2 -> nama = "Budi";
$Person2 -> alamat = "Jakarta";
$Person2 -> negara = "Indonesia";

$Person3  = new Person();
$Person3 -> nama = "Siti";
$Person3 -> alamat = "Bandung";

var_dump($Person1);
var_dump($Person2);
var_dump($Person3);

echo "Nama : "  . $Person1 -> nama . PHP_EOL;
echo "Alamat : " . $Person1 -> alamat . PHP_EOL;
echo "Negara : " . $Person1 -> negara . PHP_EOL;