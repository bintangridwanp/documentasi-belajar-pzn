<?php

/* 
    Perbedaan utama antara require dan include dalam PHP adalah ketika file yang disertakan tidak ditemukan,
    require akan menghentikan eksekusi script, sedangkan include hanya akan menampilkan pesan error dan melanjutkan eksekusi script.
*/

require "./folder-sample/function-sample.php";

sapaHalo(10, "luffy");

