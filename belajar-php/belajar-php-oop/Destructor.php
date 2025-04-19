<?php

require_once 'data/Person.php';

$Person1 = new Person("Joko", "Solo");

echo "Program Selesai" . PHP_EOL;
// Destructor akan dipanggil secara otomatis ketika objek dihapus

var_dump($Person1);