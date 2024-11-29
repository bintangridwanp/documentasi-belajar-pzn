<?php

$one_piece = ["luffy", "d", "monkey"];

for($i = 0; $i < count($one_piece); $i++){
    echo $one_piece[$i] . PHP_EOL;
}

foreach($one_piece as $i){
    echo $i . PHP_EOL;
}

// tipe data map

$manchester_united = [
    "kiper" => "onana",
    "cb" => "dalot",
    "dmf" => "bruno"
];

foreach ($manchester_united as $key => $value){
    echo "$key : $value" . PHP_EOL;
}
