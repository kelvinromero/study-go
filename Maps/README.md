# Maps

With Arrays and Slices data is stored in a linear structure and accessed by integers that represent positions.

With maps we use `keys`.

Something particular about go, it taks two types to declare a map:
- After the keyword that defines that it is a map
- We need the key type and the value type
```go
dictionary := map[string]string{"test": "this is a test"}
```


Another very interesting thing here is that you can modify a map without passing an address (&myMap).