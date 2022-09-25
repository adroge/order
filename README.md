# Order

The generic ordering mechanism for slices in Go.

- [Order](#order)
	- [Usage](#usage)
		- [Comparison Functions](#comparison-functions)
		- [Compare and ComparePv](#compare-and-comparepv)
		- [Compare for Descending Values](#compare-for-descending-values)
		- [Sorting Operations](#sorting-operations)
	- [Examples](#examples)
		- [Sorting a Simple List](#sorting-a-simple-list)
		- [Sorting a List of Structures](#sorting-a-list-of-structures)
		- [Checking if a List is Sorted](#checking-if-a-list-is-sorted)

## Usage

The unit tests are a good source to see working examples. With that being said,
the basic usage is as follows.

order.By(...) will create on order structure. This structure can then be used
to sort a list. A set of comparison functions are passed as arguments to `By`.

### Comparison Functions

These functions can compare whatever you want, but must return an `int`.

| value returned | Means   |
|----------------|---------|
| < 0            | less    |
| 0              | equal   |
| > 0            | greater |

### Compare and ComparePv

`order.Compare` and `order.ComparePv` are two shortcut functions that will
allow the comparison of values without having to write the same functions over
and over again. These two functions aren't necessary, but they more than
suffice for what is needed in almost all cases.

The unit tests have examples of comparison without `order.Compare`.

An example of where you would want to make your own comparison functions would
be if some type of fuzzy comparison was needed.

### Compare for Descending Values

If you want a descending sorted list, inside the comparison function, multiply
the value by `-1`. There is also an example of this in the unit tests.

```go
descending := func(left, right float64) int {
		return -1 * order.Compare(left, right)
	}
```

### Sorting Operations

| Operation | Description           |
|-----------|-----------------------|
| Sort      | Unstable sort         |
| Stable    | Stable sort (slower)  |
| IsSorted  | If the list is sorted |

## Examples

### Sorting a Simple List

Ordering is as simple as this for a list of built-in types.

```go
list := []string{"b", "a"}
stringValue := func(l, r string) int {return order.Compare(l,r)}
order.By(stringValue).Sort(list)
```

Or single line.

```go
list := []string{"b", "a"}
order.By(func(l, r string) int {return order.Compare(l,r)}).Sort(list)
```

### Sorting a List of Structures

It's as simple as this for a list of structures.

```go
age := func(l, r Person) int {return order.Compare(l.age, r.age)}
name := func(l, r Person) int {return order.Compare(l.name, r.name)}
order.By(age, name).Sort(sliceOfPeopleStructures)
```

### Checking if a List is Sorted

```sh
list := []string{"a", "b"}
sorted := order.By(func(l, r string) int {return order.Compare(l,r)}).IsSorted(list)
```
