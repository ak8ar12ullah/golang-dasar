package main

import "fmt"

func main(){
	var nama string;

	nama = "faridl akbar";

	fmt.Println(nama);

	nama = "faridl";

	fmt.Println(nama);

	var nama2 = "faridl m. akbar";

	fmt.Println(nama2);

	nama3 := "akbar";

	fmt.Println(nama3)

	var (
		firstName = "faridl";
		lastName = "akbar";
	)

	fmt.Println(firstName, lastName)
}