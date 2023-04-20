package main
import "fmt"

func main() {
	// languages := make(map[int]string)
	// languages[1] = "JavaScript"
	// languages[2] = "Python"
	// languages[3] = "Ruby"
	// languages[4] = "GoLang"
	// fmt.Println(languages)

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	languages["go"] = "GoLang"
	fmt.Println(languages)


	fmt.Println("JS short for", languages["js"])
	fmt.Println("Python short for", languages["py"])
	fmt.Println("Ruby short for", languages["rb"])

	// delete(languages, "rb")
	delete(languages, "rb")
	fmt.Println(languages)

	// Loops through the map
	for key, value := range languages {
		fmt.Printf("For each key %v,the value is %v\n", key, value)
	}

}