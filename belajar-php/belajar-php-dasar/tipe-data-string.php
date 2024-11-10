<?php

echo 'hello';
echo 'bintang ridwan pribadi';

echo "\n";

echo "hello";
echo "luffy \t D\t monkey\t";

echo <<<HEREDOC

    list anime terbaik menurut saya:
    1. One piece
    2. kimetsu no yaiba
    3. attack on titan
    4. jujutsu kaisen

    untuk list ini di tulis menggunakan HEREDOC supaya bisa tidak manual menggunakan echo " ";

HEREDOC;

echo <<<'NOWDOC'

    list anime terbaik menurut saya:
    1. One piece
    2. kimetsu no yaiba
    3. attack on titan
    4. jujutsu kaisen

    untuk list ini di tulis menggunakan NOWDOC berbeda dengan HEREDOC, di NOWDOC tidak memparsing nilai,
    dan NOWDOC mengggunkan echo <<<''; untuk membuatnya.

NOWDOC;




