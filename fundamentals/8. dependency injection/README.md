# Dependency Injection

For capturing this concept the Learn Go With tests suggest that we test a function like this:
```go
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```
But how can we test this? It doesn't return a value, it prints to the standard output.

Here's where dependency injection comes into play. We can inject (pass in) a dependency of printing.

It doesn't care the where it prints, so we can pass something we can control where it prints.

Now, let's go to [code](greet_test.go)