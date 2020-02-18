package main

	import (
		utility "github.com/omnified-testing/test-lambda-gen-repo-2020-02-18-11-22-49-045876-0000-gmt-m-0-004553115/utility"
	)
	
	func main() {
	}
	
	// Call : The executable API of the function, just calls the subpackage.
	// Input and output needs to match that of utility subpackage's Call function
	func Call(s string) string {
		return utility.Call(s)
	}