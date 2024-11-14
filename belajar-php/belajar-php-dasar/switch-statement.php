<?php

$hari = "kamis";

switch ($hari){
    case "senin" :
        echo "praktikum std, pancasila, dpbo, keprofesian" . PHP_EOL;
        break;
    case "selasa" :
        echo "aka, std materi" . PHP_EOL;
        break;
    case "rabu" :
        echo "materi std" . PHP_EOL;
        break;
    case "kamis" : 
        echo "agama islam, imk, praktikum dpbo" . PHP_EOL;
        break;
    case "jum'at" : 
        echo "oak, kwn" . PHP_EOL;
        break ;
    default :
        echo "hari untuk bersantai" . PHP_EOL;
}

// code alternatif switch statement
switch ($hari) :
    case "senin" :
        echo "praktikum std, pancasila, dpbo, keprofesian" . PHP_EOL;
        break;
    case "selasa" :
        echo "aka, std materi" . PHP_EOL;
        break;
    case "rabu" :
        echo "materi std" . PHP_EOL;
        break;
    case "kamis" : 
        echo "agama islam, imk, praktikum dpbo" . PHP_EOL;
        break;
    case "jum'at" : 
        echo "oak, kwn" . PHP_EOL;
        break ;
    default :
        echo "hari untuk bersantai" . PHP_EOL;
    endswitch;