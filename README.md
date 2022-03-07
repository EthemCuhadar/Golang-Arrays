# Golang

![Golang Image](golang.png)

---------------------------------------------------------------------

## Arrays

Arrays in Go programming language is similar to other programming languages like Java, Python and C++. They are used to store collection of data of the **same type**. In Go programming language this is very important to store same types. Otherwise, an error will be occur. The indexing of array starts with 0 and continues. In Go programming language **var** keyword is used to define an array. Furthermore, **array name**, **array size** and **type of the array** should be defined as well.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    var array_name[3]string // defined an array which has name of array_name

    array_name[0] = "car"
    array_name[1] = "cat"
    array_name[2] = "dog"

    fmt.Printf("Type of the array_name is: %v\n", reflect.TypeOf(array_name))
    fmt.Println("The first element is: ", array_name[0])
    fmt.Println("The second element is: ", array_name[1])
    fmt.Println("The third element is: ", array_name[2])

}    
```

```console
Type of the array_name is: [3]string
The first element is:  car
The second element is:  cat
The third element is:  dog
```

* Note that the size of the array is defined at the beginning. Therefore, it cannot be changed. For example, adding a new element to **index 3** will cause an error.

```go
array_name[3] = "bike"
fmt.Println("The fourth element is: ", array_name[3])
```

```
# command-line-arguments
.\Arrays.go:21:12: invalid array index 3 (out of bounds for 3-element array)
.\Arrays.go:22:51: invalid array index 3 (out of bounds for 3-element array)
```

* Moreover, the elements of an array can be defined at the beginning. The example can be seen below.

```go
new_array := [4]int{10, 20, 30, 40}
fmt.Println(new_array)
```

```console
[10 20 30 40]
```

* Therefore, there are two decleration of arrays in Go programming language.
1. First, decleare with **var** keyword.

2. Second, decleare with short-hand decleration.
* Now let's look at array equalities below.

```go
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
```

```console
[10 20 30 40]
[4]int
[10 20 30 40]
[4]int
[10 20 30 40 0]
[5]int
```

* As you have noticed, arrays could be different types with lengths and variables types of elements.

```go
fmt.Println(array1 == array2)
```

```console
true
```

```go
fmt.Println(array2 == array3)
```

```console
# command-line-arguments
.\Arrays.go:45:21: invalid operation: array2 == array3 (mismatched types [4]int and [5]int)
```

* Hence, **array2** and **array3** are not same type and an error occured.

* Finally, let's check elements of **array2** and **array3** are the same or not.

```go
fmt.Println(array2[0] == array3[0])
```

```console
true
```

* The answer is **true**. Because, while defining the **array2** and **array3**, it is stated that array elements will be **integer**. Therefore, even if the types of arrays are not the same, element type would be so.
