# go-x (extensions)

---

Go `x` is collection of utility functions that allow to simplify common operations.
Library uses loose coupling between the packages and does not have any dependencies
on 3rd party tools / libraries.

## why `x`?

---

The `"x"` is commonly used to indicate that a package or library is experimental or an extension of a certain functionality.

When developers create experimental or non-standard/extension packages, they might prefix them with "x" to indicate
that they are not yet fully developed and/or to differentiate them from standard packages and prevent
naming collisions with official packages.

## Packages

* `xhttp` utilities for facilitating writing JSON HTTP responses to the http.ResponseWriter.
* `xjson` utilities for marshaling/unmarshaling of data with generics support.
* `xmaps` utilities for working with maps with generics support.
* `xslices` utilities for working with slices with generics support.
* `xstrings` utilities for working with strings.
* `xtesting` utilities to work with tests and test fixtures.
* `httpbody` provides utilities to create/bind http.Request body.
* `ptr` utilities for converting literal type values to/from pointers inline.

## Installation

```shell
go get git.naspersclassifieds.com/olxeu/realestate/go-toolkit/x@latest
```

## Examples

### `httpbody`

---
* converting struct into JSON-encoded `io.ReaderCloser` stream
  #### go-x

  ```go
  package example

  func Example() {
    reader, err := httpbody.FromJSON(&payload{ID: "12345"})
  }

  ```

  #### SDK code

  ```go
  package example

  func Example()  {
    input := payload{ID: "12345"}
    data, err := json.Marshal(&input)
    if err != nil {
      return nil, fmt.Errorf("marshaling input: %w", err)
    }
    return io.NopCloser(bytes.NewReader(data)), nil
  }
  ```
---
* reading http.Body into JSON struct
  #### go-x

  ```go
  package example

  func Example() {
    type payload struct {
       ID string `json:"id"`
    }
    var res http.Response

    p, err := httpbody.BindJSON[payload](res.Body)
  }

  ```

  #### SDK code

  ```go
  package example

  func Example()  {
    type payload struct {
       ID string `json:"id"`
    }
    var res http.Response
    bytes, err := io.ReadAll(res.Body)
    if err != nil {
       // handle
    }
    var p payload
    if err := json.Unmarshal(bytes, &p); err != nil {
       // handle
    }
  }
  ```

### `xslices`

---
Examples

* mapping a slice of strings into a slice of ints:
  #### go-x

  ```go
  package example

  func Example() {
      input := []string{"a", "b", "c"}
      res := xslices.Map[[]string, []int](input, func(s string) int {
        return len(s) + 1
      })
      fmt.Println(res)
  }
  ```

  #### SDK code

  ```go
  package example

  func Example() {
      mapper := func(s string) int { return len(s) + 1 }
      input := []string{"a", "b", "c"}
      res := make([]int, 0, len(input))
      for _, v := range input {
        res = append(res, mapper(v))
      }
  }
  ```
---

* returning filtered slice
  #### go-x

  ```go
  package example

  func Example() {
      input := []string{"a", "b", "c"}
      res := xslices.Filter(input, func(s string) bool {
        return s == "a"
      })
      fmt.Println(res)
  }
  ```

  #### SDK code

  ```go
  package example

  func Example() {
      mapper := func(s string) int { return len(s) + 1 }
      input := []string{"a", "b", "c"}
      res := make([]int, 0, len(input))
      for _, v := range input {
        res = append(res, mapper(v))
      }
  }
  ```


### `xstrings`

---
* checking if string is alphanumeric

  #### go-x
  ```go
  package example

  func Example() {
    s := "abas1231"
    res := xstrings.IsAlphanumeric(s)
    fmt.Println(res)
  }
  ```
  #### SDK code
  ```go
  package example

  func Example() {
    s := "abas1231"
    if s == "" {
      return true
    }
    for _, c := range s {
      if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
        return false
      }
    }
  }
  ```
