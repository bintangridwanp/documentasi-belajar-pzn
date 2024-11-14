<?php

$nilai_string = (string)100;
var_dump($nilai_string);

$nilai_integer = (int)"100";
var_dump($nilai_integer);

$nilai_floating = (float)"1.10";
var_dump($nilai_floating);

$name = "Luffy";
echo $name[0] . PHP_EOL;
echo $name[1] . PHP_EOL;
echo $name[2] . PHP_EOL;
echo $name[3] . PHP_EOL;
echo $name[4] . PHP_EOL;
//echo $name[5] . PHP_EOL; Error

echo "Hallo aku $name, Selamat belajar ya..." . PHP_EOL;

$pesan = "Luffy ";

echo "oke siap {$pesan}D mongkey" . PHP_EOL;
