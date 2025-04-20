<?php

require_once 'data/Manager.php';

$manager = new Manager("Budi", "Manager");
$manager->nama = "Budi";
$manager->sapaHalo("Andi");

$vp = new VP("Siti", "VP");
$vp->nama = "Siti";
$vp->sapaHalo("Andi");
