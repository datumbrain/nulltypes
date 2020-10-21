# nulltypes  [![Build Status](https://api.travis-ci.com/datumbrain/nulltypes.svg?branch=master)](https://travis-ci.com/github/datumbrain/nulltypes) [![Go Report Card](https://goreportcard.com/badge/github.com/datumbrain/nulltypes)](https://goreportcard.com/report/github.com/datumbrain/nulltypes) ![License](https://img.shields.io/badge/license-MIT-blue.svg)

<img align="right" src="https://github.com/ashleymcnamara/gophers/raw/master/MovingGopher.png" width="120">

`nulltypes` is a golang module that provides an alternative to nullable data types from `database/sql` with proper JSON marshalling and unmarshalling.

It also provides a wrapper for `time.Time` to format time to use with `timestamp` of SQL databases, i.e. `mysql`, `postgres`.

The default database time zone is set to `UTC`, but it can easily be changed with:

```go
nulltypes.DatabaseLocation, _ = time.LoadLocation([YOUR_TIME_ZONE])
```

## Import

```go
import "github.com/datumbrain/nulltypes"
```

## Usage

Here is an example usage with _GORM_.

```go
package models

type User struct {
	ID              uint `gorm:"primary_key"`
	Name            string
	Address         nulltypes.NullString
	CreationDate    time.Time `gorm:"autoCreateTime;default:current_timestamp"`
	UpdationDate    nulltypes.NullTime
	TerminationDate nulltypes.NullTime
	ManagerID       nulltypes.NullInt64 `gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
}
```

```go
user := User{
	ID:           0,
	Name:         "John Doe",
	Address:      nulltypes.String("221B Baker Street"),
	CreationDate: time.Now(),
	UpdationDate: nulltypes.Now(),
	ManagerID:    nulltypes.Int64(5),
}
```

## Author

[Faizan Khalid](https://github.com/iamfaizankhalid)
