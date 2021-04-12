package main

import(
	"strings"
	"fmt"
)

func findFirstStringInBracket(str string) string { 

	if(str == ""){
		return "";
	}

	indexOpenBracket := strings.Index(str,"("); // get index open bracket character
	indexCloseBracket := strings.Index(str,")"); // get index close bracket character

	if(indexOpenBracket >= 0 && indexCloseBracket >= 0){ // if index open and close found
		return str[indexOpenBracket+1:indexCloseBracket]; // return string from open+1 untul indexclose
	}

	return ""; // if not found open bracket char or close bracket
}

func main() {
	fmt.Println(findFirstStringInBracket("Ini tidak muncul ( ini muncul ) ini tidak muncul"))
}