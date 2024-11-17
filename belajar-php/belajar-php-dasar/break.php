<?php

$a = 1;

while($a >= 1){
    echo "Perulangan ke: " . $a . PHP_EOL;
    $a = $a + 1;

    if($a > 10){
        break;
    }
}