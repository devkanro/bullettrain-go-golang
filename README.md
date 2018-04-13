# Go car for the [Bullettrain](https://github.com/bullettrain-sh/bullettrain-go-core) shell prompt

## Features:

- Displaying only when needed
- Go version display


**Callword**: `go`

**Template variables**:

* `.Icon`: the car's icon
* `.Info`: the Go version text

**Template colours**:

* `c`: the car's colour
* `cs`: the car symbol's colour


## Car options

| Environment variable            | Description                                     | Default value                                   |
|:--------------------------------|:------------------------------------------------|:------------------------------------------------|
| BULLETTRAIN_CAR_GO_SHOW         | Whether the car needs to be shown all the time. | false                                           |
| BULLETTRAIN_CAR_GO_TEMPLATE     | The car's template.                             | `{{.Icon \| printf "%s " \| cs}}{{.Info \| c}}` |
| BULLETTRAIN_CAR_GO_PAINT        | Colour override for the car's paint.            | black:123                                       |
| BULLETTRAIN_CAR_GO_SYMBOL_ICON  | Icon displayed on the car.                      | ``                                             |
| BULLETTRAIN_CAR_GO_SYMBOL_PAINT | Colour override for the car's symbol.           | black:123                                       |

# Contribute

Even reporting your use case will greatly help us to figure out/improve
this product, so feel free to reach out in the Issues section.
