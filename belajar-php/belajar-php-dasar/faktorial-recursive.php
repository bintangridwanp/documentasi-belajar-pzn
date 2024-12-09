<?php

function faktorialLoop(int $nilai) : int {
    $total = 1;
    for($i = 1; $i <= $nilai; $i++){
        $total = $total *$i;
    }

    return $total;
}

var_dump(faktorialLoop(5));
var_dump(1 * 2 * 3 * 4 * 5);

function faktorialRecursive(int $angka){
    if ($angka == 1){
        return 1;
    } else {
        return $angka * faktorialRecursive($angka - 1);
    }
}

var_dump(faktorialRecursive(5));

//funtin untuk cek berapa kuat laptop bisa menampung memory sampai overflow
function cekOverFlow(int $angka) : int 
{
    if ($angka == 1){
        echo "Perulangan selesai" . PHP_EOL;
    } else {
        echo "Perulangan ke . $angka" . PHP_EOL;
        cekOverFlow($angka - 1);
    }
}

var_dump(cekOverFlow(3000000));