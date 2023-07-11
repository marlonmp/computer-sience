package datastructure

// The array data structure allow us to store a fix amount of elements
// for a one data type in spesific, and we can acces to all of those data
// one by one with its spesific index

func array() {
	// here we define an array with not limit spesified
	a := []string{"a", "b", "c"}

	// the `a` variable looks like this:
	// ["a", "b", "c"]

	// we can add more items to it
	a = append(a, "d", "e", "f") // ["a", "b", "c", "d", "e", "f"]

	// and we can remove itmes from it
	a = append(a[:2], a[4:]...)

	/// ["a", "b", "e", "f"]

	// and we remove values from it

	// here we define an array with a limit spesified.
	// We can't add more values to it, this type or array is called slice
	s := [4]string{"x", "y", "z"}

	// the representation of `s` looks like this:
	// ["x", "y", "z", ""]

	// and we can add values on a spesific cell
	s[3] = "0" // ["x", "y", "z", "0"]

	// and we can iter this data
	// with the range operator we get the index `i`, and the item `i`
	// in any cycle
	for i, item := range a {
		println("i", i, "; item", item)
	}

	// or we can iter all items by the index
	for i := 0; i < len(s); i++ {
		println("i", i, "; item[i]", s[i])
	}
}
