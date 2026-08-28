# yo-go

Go client for [yo](https://github.com/tamnd/yo), an embedded multi-model database that lives in one `.yo` file. Struct tags are the schema and there is no query language to learn.

## Status

Nothing to use yet, and the module says so out loud.

```go
db, err := yo.Open("app.yo")
// err: yo is not usable yet. This is a reserved placeholder at 0.0.1; see https://github.com/tamnd/yo
```

`Open` returns an error rather than panicking, because that is what a Go caller is prepared to handle, and the error is the same sentence every other language's placeholder carries. The version is held so the name is held.

The engine is at `M0`. The record plane and the file format are `M1` and in progress, so there is nothing for this binding to sit on top of yet. Watch the [milestones](https://github.com/tamnd/yo/milestones).

## Install

```bash
go get github.com/tamnd/yo-go
```

Go modules are paths, so there is no registry name to hold here, only this repository. The package clause is `yo`, so the import needs no alias:

```go
import yo "github.com/tamnd/yo-go"
```

## What this will be

```go
package main

import (
	"context"
	"fmt"

	yo "github.com/tamnd/yo-go"
)

type User struct {
	ID    uint64  `yo:"id,key"`
	Name  string  `yo:"name"`
	Score float64 `yo:"score,index,ord"`
}

var UserScore = yo.Field[User]("score")

func main() {
	ctx := context.Background()
	db, err := yo.Open("app.yo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	users, err := yo.Docs[User](db, "users")
	if err != nil {
		panic(err)
	}
	users.Put(ctx, User{ID: 1, Name: "Ada", Score: 99.5})
	for u, err := range users.OrderBy(UserScore.Desc()).Limit(10).All(ctx) {
		if err != nil {
			panic(err)
		}
		fmt.Println(u.Name, u.Score)
	}
}
```

Twenty-two lines, the longest of the tier-1 snippets, and it is that long on purpose: Go's import block and error returns are not optional, and a snippet that elides them is not a Go program. None of it works today.

## Planned support

| Item | Version |
|---|---|
| Go floor | 1.25 |
| Go current | 1.26.x |
| cgo | required for the default build |

The floor is 1.25 and is not being lowered. `iter.Seq2` is what makes the iteration surface a `range` loop instead of a `Next()`/`Err()` pair, and carrying both shapes doubles the API for the sake of a toolchain two releases old.

**No yo data is on the Go heap.** The engine's memory is invisible to the Go garbage collector, by construction, which is the direct answer to the GC tail latencies that motivated this project in the first place.

Field paths are resolved once, at handle construction, from struct tags. `yo.Field[User]("score")` returns a typed path and panics at construction if the field does not exist. That is deliberate and it is scoped: package-level initialisation is where a misspelling should be a startup failure rather than a request failure.

## Design

The full Go specification is `dx/07` in the project specification.

## Licence

Apache 2.0 or MIT, at your option.
