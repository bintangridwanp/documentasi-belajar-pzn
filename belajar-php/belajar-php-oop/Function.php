<?php

require_once 'data/Person.php';

$Person4 = new Person("Joko", "Solo");

$Person4 -> nama = "Joko";
$Person4 -> alamat = "Solo";
$Person4 -> negara = "Indonesia";
$Person4 -> sapaHalo("Solo");
$Person4 -> sapaHalo(null);

$Person4 -> informasi();




