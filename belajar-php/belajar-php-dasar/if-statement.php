<?php

$nilai = 0;
$absen = 0;

if($nilai >= 85 && $absen >= 85){
    echo "Anda mendapatkan mutu A";
} else if ($nilai >= 75 && $absen >= 75){
    echo "Anda mendapatkan mutu B" . PHP_EOL;
} else if ($nilai >= 65 && $absen >= 65) {
    echo "Anda mendapatkan mutu C". PHP_EOL;
} else if ($nilai >= 50 && $absen >= 50){
    echo "Anda mendapatkan mutu D" . PHP_EOL;
} else {
    echo "Maaf, anda tidak lulus" . PHP_EOL;
}

//code alternatif if statement 

$jari_satu = "gajah";
$jari_dua = "semut";

if ($jari_satu == "gajah" && $jari_dua == "semut") :
    echo "semut pemenangnya" . PHP_EOL;
elseif($jari_satu = "manusia" && $jari_dua == "semut") : 
    echo "manusia pemenangnya" . PHP_EOL;
elseif($jari_satu == "gajah" && $jari_dua = "manusia") : 
    echo "gajah pemenangnya" . PHP_EOL;
else : 
    echo "seri untuk keduanya";
endif;
