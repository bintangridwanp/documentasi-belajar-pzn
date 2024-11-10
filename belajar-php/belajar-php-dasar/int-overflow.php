<?php
// PHP memiliki batasan integer tergantung pada sistem (platform 32-bit atau 64-bit)
// Pada platform 64-bit, batas integer adalah 9223372036854775807

$maxInt = PHP_INT_MAX;
echo "Nilai maksimal integer: $maxInt\n";

// Menambahkan 1 ke nilai maksimum integer akan menyebabkan overflow
$overflow = $maxInt + 1;
echo "Nilai setelah overflow: $overflow\n";
?>
