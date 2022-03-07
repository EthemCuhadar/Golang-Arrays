package main

import (
    "fmt"
)

func main() {
	
	// defined an array which has name of array_name
    var array_name[3]string 
	
	
	// Assigning array elements
    array_name[0] = "car"
    array_name[1] = "cat"
    array_name[2] = "dog"
	
	// Print array elements
    fmt.Printf("%T\n", array_name)                        // Output: [3]string
    fmt.Println("The first element is: ", array_name[0])  // Output: The first element is: car
    fmt.Println("The second element is: ", array_name[1]) // Output: The second element is: cat
    fmt.Println("The third element is: ", array_name[2])  // Output: The third element is: dog
	
	// array_name[3] = "bike"
	// fmt.Println("The fourth element is: ", array_name[3]) // Output: Error (Array length is 3)
	
	
	// Assign first array with 4 integer length
	array1 := [4]int{10, 20, 30, 40}
	fmt.Println(array1)                 // Output: [10, 20, 30, 40]
	fmt.Printf("%T\n", array1)          // Output: [4]int
	
	// Assign second array with 4 integer length
	array2 := [4]int{10, 20, 30, 40}
	fmt.Println(array2)                 // Output: [10, 20, 30, 40]
	fmt.Printf("%T\n", array2)          // Output: [4]int
	
	// // Assign third array with 5 integer length
	array3 := [5]int{10, 20, 30, 40}
	fmt.Println(array3)                 // Output: [10, 20, 30, 40, 0]
	fmt.Printf("%T\n", array3)          // Output: [5]int
	
	// Check array equalities
	fmt.Println(array1 == array2)       // Output: true
	//fmt.Println(array2 == array3)     // Output: Error (They are not same type arrays)
	
	// Check equalities for array elements
	fmt.Println(array1[0] == array2[0]) // Output: true
	fmt.Println(array2[0] == array3[0]) // Output: true

} 
