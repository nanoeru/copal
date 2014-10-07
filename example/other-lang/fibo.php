<?php

    // by satosystems
    // http://d.hatena.ne.jp/satosystems/20121228/1356655565
    function fib1($n){
        if ($n < 2){ return $n; }
        return fib1($n - 2) + fib1($n - 1);
	}

echo fib1(30);
