<?php

function findFirstStringInBracket($str){

    if($str === ""){
        return "";
    }

    $characterAfterOpenBracket = strstr($str,'(');
    $characterBeforeCloseBracket = strstr($characterAfterOpenBracket,')',true);
    $openBracketRemoved = ltrim($characterBeforeCloseBracket,'(');

    return $openBracketRemoved;
}

echo findFirstStringInBracket("Ini tidak muncul ( ini muncul ) ini tidak muncul");