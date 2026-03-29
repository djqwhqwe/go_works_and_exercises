package main


import ("fmt")

func concatStrings(xs ...string) string {
	var ret string = ""
	if len(xs) == 0 {
		return ""
	} else if len(xs) == 1 {
		return xs[0]
	} else {
		for _, value := range xs{
			ret+=value
		}
		return ret
	}
}

func main() {
	fmt.Println(concatStrings())
}