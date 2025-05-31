# go-defaults

[![Go Report Card](https://goreportcard.com/badge/github.com/lthphuw/go-defaults)](https://goreportcard.com/report/github.com/lthphuw/go-defaults) [![Go CI](https://github.com/lthphuw/go-defaults/actions/workflows/go.yml/badge.svg)](https://github.com/lthphuw/go-defaults/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/lthphuw/go-defaults/graph/badge.svg?token=GQ28QBZFBZ)](https://codecov.io/gh/lthphuw/go-defaults)

`go-defaults` provides functionality to parse and set default values for struct fields based on their "default" tags
<a name="readme-top"></a>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Features</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#supportedunsupported-field-type">Supported/Unsupported Field Type</a></li>
    <!-- <li><a href="#contributing">Contributing</a></li> -->
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <!-- <li><a href="#acknowledgments">Acknowledgments</a></li> -->
  </ol>
</details>

## About The Project

The `defaults` package simplifies the process of **setting default values** for **struct fields** in `Go` when they are **unset** (i.e., at their zero value or nil for pointers). Instead of **manually** checking each field for its zero value and assigning defaults, defaults uses struct tags to automatically parse and set default values, making your code cleaner and more maintainable. This project was inspired by the need for a streamlined, reusable solution to initialize struct fields with predefined defaults, especially in scenarios involving configuration structs, data models, or nested structures.

## Key Features

- **Automatic Default Value Setting**: The Defaults function processes a non-nil pointer to a struct, setting default values for exported fields based on their "default" tags (configurable via the package-level Tag variable).

- **Recursive Struct Handling**: Nested structs are processed recursively, applying default values to their fields.

- **Broad Type Support**: Supports a wide range of field types, including numeric types, strings, booleans, maps, slices, arrays, and their pointers. See [Supported/Unsupported Field Types](#supportedunsupported-field-type)

- **Robust Error Handling**: Returns errors for invalid inputs, unexported fields, empty tags (for non-struct fields), parsing failures, out-of-range values, or unsupported types.

## Getting Started

Note that we only supports the two most recent minor versions of Go.

```
go get github.com/lthphuw/go-defaults@latest
```

## Usage

With `go-defaults`, you can define `defaults` in `struct tags` and let the `Defaults` function handle initialization:

```go
type Config struct {
    Port int    `default:"8080"`
    Host string `default:"localhost"`
}

func main() {
    var config Config
    // Load from env
    // ...

    // Set defaults
    if err := defaults.Defaults(&config); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Port: %d, Host: %s\n", config.Port, config.Host)
    // Output: Port: 8080, Host: localhost
}
```

This approach reduces boilerplate code, improves readability, and ensures consistent default value handling across your application.

Here’s a another complete example demonstrating various field types:

```go
package main

import (
    "fmt"
    defaults "github.com/lthphuw/go-defaults"
)

type Nested struct {
    Value int `default:"100"`
}

type Example struct {
    ID      int            `default:"123"`
    Name    *string        `default:"hello"`
    Data    map[string]any `default:"{\"key\":\"value\",\"num\":42}"`
    Numbers [3]int         `default:"[1,2,3]"`
    Nested  Nested
}

func main() {
    var s Example
    if err := defaults.Defaults(&s); err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("ID: %d, Name: %v, Data: %v, Numbers: %v, Nested.Value: %d\n",
        s.ID, *s.Name, s.Data, s.Numbers, s.Nested.Value)
}
```

Output:

```bash
ID: 123, Name: hello, Data: map[key:value num:42], Numbers: [1 2 3], Nested.Value: 100
```

## Supported/Unsupported Field Type

### Supported Field Types and Example Tags

The Defaults function supports the following field types with their respective "default" tag formats:

- Numeric Types:

  - `int`: default:"123"

  - `int8`: default:"127"

  - `int16`: default:"32767"

  - `int32`: default:"2147483647"

  - `int64`: default:"9223372036854775807"

  - `uint`: default:"123"

  - `uint8`: default:"255"

  - `uint16`: default:"65535"

  - `uint32`: default:"4294967295"

  - `uint64`: default:"18446744073709551615"

  - `float32`: default:"3.14"

  - `float64`: default:"3.14159265359"

  - `complex64`: default:"1+2i"

  - `complex128`: default:"1.5+2.5i"

- `string`: default:"hello"

- `bool`: default:"true"

- `map` (e.g., map[string]any): default:"{\"key\":\"value\",\"num\":42}"

- `slice` (e.g., []string): default:"[\"a\",\"b\",\"c\"]"

- `array` (e.g., [3]int): default:"[1,2,3]"

`struct`: triggers recursive default setting for nested fields.

**Pointers** to Above Types:

- `\*int`: default:"123"

- `\*string`: default:"hello"

- `\*map[string]any`: default:"{\"key\":\"value\"}"

- `\*[]string`: default:"[\"a\",\"b\"]"

- `\*[3]int`: default:"[1,2,3]"

- `\*struct`: triggers recursive default setting for nested fields.
- ...

### Unsupported Field Types

The following types are not supported by `Defaults`:

- `Function types` (e.g., func(), \*func())

- `Channels` (e.g., chan int)

- `Interfaces`

- `Unsafe pointers` (e.g., unsafe.Pointer)

- Any other types not listed above

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

## Contact

Maintained by [Phu](https://github.com/lthphuw).  
For questions, feedback, or support, please contact: <a href="mailto:lthphuw@gmail.com">lthphuw@gmail.com</a>

<p align="right">(<a href="#readme-top">back to top</a>)</p>
