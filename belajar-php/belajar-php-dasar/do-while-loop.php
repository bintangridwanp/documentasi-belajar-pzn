<?php

    $a = false;
    $b = 1;

    do{
        echo "Perulangan ke: " . $b . PHP_EOL;
        $b = $b + 1;
    } while (!$a == true && $b <= 10 );