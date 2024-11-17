<?php

$a = 1; 

while ($a <= 100){

    if ($a % 2 == 0){
        echo "Ini adalah bilangan genap: " . $a . PHP_EOL;
    }
    $a = $a + 1;
    continue;
    echo "Ini adalah bilangan ganjil: " . $a . PHP_EOL;
}

