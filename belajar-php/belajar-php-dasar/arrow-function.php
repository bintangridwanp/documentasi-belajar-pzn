<?php

$nama = "luffy D monkey";
$title = "kapten bajak laut";

$anonymous_fuction = function () use ($nama, $title){
    echo "Hello $nama adalah seorang $title" . PHP_EOL;
};

$anonymous_fuction();

$one_piece = fn() => "Hello $nama adalah seorang $title" . PHP_EOL;

echo $one_piece();
