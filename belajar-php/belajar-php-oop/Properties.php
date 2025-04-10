<?php

require_once 'data/Person.php';

$Person1 = new Person();
$Person1 -> nama = "Joko";
$Person1 -> alamat = "Solo";
$Person1 -> negara = "Indonesia";

var_dump($Person1);

echo "Nama : "  . $Person1 -> nama . PHP_EOL;
echo "Alamat : " . $Person1 -> alamat . PHP_EOL;
echo "Negara : " . $Person1 -> negara . PHP_EOL;