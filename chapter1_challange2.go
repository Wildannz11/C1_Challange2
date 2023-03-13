package main

import (
	"fmt"
)

func main () {
	for i := 0 ; i <= 4 ; i++ {
		fmt.Println("Nilai i = ", i)
	}

	// keys := []string{"U+0421", "U+0410", "U+0428", "U+0410", "U+0420", "U+0412", "U+041E"}
	keys := "САШАРВО"
    // values := []string{"C", "A", "Ⅲ", "A", "P", "B", "O"}

	// unicode := make(map[string]string)

    // for i := range keys {
    //     unicode[keys[i]] = values[i]
    // }
	
	p := 0
	for j := 0 ; j <= 10; j++ {
		if j == 5 {

			for _, k := range keys {
				// fmt.Printf("Character %v '%v' starts at byte position %v \n", k, k,p)
				fmt.Printf("Character %U '%c' starts at byte position %v \n", k, k, p)
				// p= p+2
				p += 2 
			}	
			continue
		}
		
		fmt.Println("Nilai j = ", j)
	} 
}
