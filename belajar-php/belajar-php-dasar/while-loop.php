<?php

$a = true;
$b = 1;

while ($a == true && $b <= 10){
    echo "Perulangan ke: " . $b . PHP_EOL;
    $b = $b + 1;
}