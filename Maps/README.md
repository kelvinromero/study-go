# Maps

With Arrays and Slices data is stored in a linear structure and accessed by integers that represent positions.

With maps we use `keys`.

Something particular about go, it taks two types to declare a map:
- After the keyword that defines that it is a map
- We need the key type and the value type
```go
dictionary := map[string]string{"test": "this is a test"}
```

## Pointers
Another very interesting thing here is that you can modify a map without passing an address (&myMap),
because a map value is a pointer to a runtime.hmap structure.

## Initalization
[//]: # (TODO: https://go.dev/blog/maps)

