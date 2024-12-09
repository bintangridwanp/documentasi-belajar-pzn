<?php

function sapaHalo(string $nama, callable $filter){
    $nama_lengkap = call_user_func($filter, $nama);
    echo "Halo $nama_lengkap" . PHP_EOL;
}

sapaHalo("luffy", "strtolower");
sapaHalo("nami", "strtoupper");
sapaHalo("zoro", function(string $name) : string {
    return strtoupper($name);
});
sapaHalo("choper", fn($name) => strtoupper($name));