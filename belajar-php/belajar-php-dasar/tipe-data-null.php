<?php

$name = "luffy D monkey";

$name = null;
$age = null;

echo $name;
echo $age;

echo is_null($name);

echo "\n";

//mengecek variable isi variable
echo "apakah variable name null: ";
$isNull = is_null($name);
var_dump($isNull);

echo "\n";


//menghapus variable
$contoh = "luffy";
echo $contoh;

echo "\n";

unset ($contoh);
echo $contoh;

echo "\n";

//mengecek suatu variable ada atau tidak
$contoh = "D";
echo $contoh;
$contoh = null;

echo "\n";

var_dump(isset($contoh));

