<?php

var_dump(join(",", [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]));
var_dump(implode(",", [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]));
var_dump(explode(" ", "luffy d monkey"));
var_dump(strtolower("LUFFY D MONKEY"));
var_dump(strtoupper("luffy d monkey"));
var_dump(trim("                       luffy                      "));
var_dump(substr("luffy d monkey", 0, 5));
