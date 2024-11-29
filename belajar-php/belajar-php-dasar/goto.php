<?php

goto a;

echo "Hello budi" . PHP_EOL;

a:

echo "Hello joko" . PHP_EOL;

$sum = 1;

while (true) {
    echo "Hello siti ke: " . $sum . PHP_EOL;
    $sum++;

    if ($sum > 10){
        goto end;
    }

}

end:
echo "Program selesai" . PHP_EOL;