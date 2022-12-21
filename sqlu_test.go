package sqlu

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
	"time"
	// _ "github.com/go-sql-driver/mysql"
)

func TestArgs(t *testing.T) {

	sql.Register("testdriver", &testDriver{})

	db, err := sql.Open("testdriver", "root:root@tcp(127.0.0.1:3306)/testdb?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}

	m := Model{}

	err = db.QueryRow(`
		SELECT
			-- Float64
			NULL, 2.5,
			NULL, 3.5,
			-- Float32
			NULL, 4.5,
			NULL, 5.5,
			-- Int64
			NULL, 4,
			NULL, 5,
			-- Int32
			NULL, 6,
			NULL, 7,
			-- Int16
			NULL, 8,
			NULL, 9,
			-- CustomTypeInt64
			NULL, 2,
			NULL, 3,
			-- Byte
			NULL, 10,
			NULL, 11,
			-- Bool
			NULL, true,
			NULL, true,
			-- String
			NULL, 'Hello',
			NULL, 'World',
			-- CustomString
			NULL, 'Hi',
			NULL, 'There',
			-- Time
			NULL, TIMESTAMP('2022-01-02 03:04:05'),
			NULL, TIMESTAMP('2022-06-07 08:09:10'),
			-- Json
			NULL, JSON_OBJECT('Value', 22),
			NULL, JSON_OBJECT('Value', 33)
		LIMIT 1
	`).Scan(Args(
		//
		&m.Float64Null, &m.Float64NotNull,
		&m.Float64PtrNull, &m.Float64PtrNotNull,
		//
		&m.Float32Null, &m.Float32NotNull,
		&m.Float32PtrNull, &m.Float32PtrNotNull,
		//
		&m.Int64Null, &m.Int64NotNull,
		&m.Int64PtrNull, &m.Int64PtrNotNull,
		//
		&m.Int32Null, &m.Int32NotNull,
		&m.Int32PtrNull, &m.Int32PtrNotNull,
		//
		&m.Int16Null, &m.Int16NotNull,
		&m.Int16PtrNull, &m.Int16PtrNotNull,
		//
		&m.CustomTypeIntNull, &m.CustomTypeIntNotNull,
		&m.CustomTypeIntPtrNull, &m.CustomTypeIntPtrNotNull,
		//
		&m.ByteNull, &m.ByteNotNull,
		&m.BytePtrNull, &m.BytePtrNotNull,
		//
		&m.BoolNull, &m.BoolNotNull,
		&m.BoolPtrNull, &m.BoolPtrNotNull,
		//
		&m.StringNull, &m.StringNotNull,
		&m.StringPtrNull, &m.StringPtrNotNull,
		//
		&m.CustomStringNull, &m.CustomStringNotNull,
		&m.CustomStringPtrNull, &m.CustomStringPtrNotNull,
		//
		&m.TimeNull, &m.TimeNotNull,
		&m.TimePtrNull, &m.TimePtrNotNull,
		//
		&m.JsonNull, &m.JsonNotNull,
		&m.JsonPtrNull, &m.JsonPtrNotNull,
	)...)
	if err != nil {
		t.Error(err)
	}

	d, err := json.Marshal(&m)
	if err != nil {
		t.Error(err)
	}

	exp := `{"Float64Null":0,"Float64NotNull":2.5,"Float64PtrNull":null,"Float64PtrNotNull":3.5,"Float32Null":0,"Float32NotNull":4.5,"Float32PtrNull":null,"Float32PtrNotNull":5.5,"Int64Null":0,"Int64NotNull":4,"Int64PtrNull":null,"Int64PtrNotNull":5,"Int32Null":0,"Int32NotNull":6,"Int32PtrNull":null,"Int32PtrNotNull":7,"Int16Null":0,"Int16NotNull":8,"Int16PtrNull":null,"Int16PtrNotNull":9,"CustomTypeIntNull":0,"CustomTypeIntNotNull":2,"CustomTypeIntPtrNull":null,"CustomTypeIntPtrNotNull":3,"ByteNull":0,"ByteNotNull":10,"BytePtrNull":null,"BytePtrNotNull":11,"BoolNull":false,"BoolNotNull":true,"BoolPtrNull":null,"BoolPtrNotNull":true,"StringNull":"","StringNotNull":"Hello","StringPtrNull":null,"StringPtrNotNull":"World","CustomStringNull":"","CustomStringNotNull":"Hi","CustomStringPtrNull":null,"CustomStringPtrNotNull":"There","TimeNull":"0001-01-01T00:00:00Z","TimeNotNull":"2022-01-02T03:04:05Z","TimePtrNull":null,"TimePtrNotNull":"2022-06-07T08:09:10Z","JsonNull":{"Value":0},"JsonNotNull":{"Value":22},"JsonPtrNull":null,"JsonPtrNotNull":{"Value":33}}`
	if string(d) != exp {
		t.Errorf("\nGot:\n%s\nExp:\n%s\n", d, exp)
	}
}

func ExampleZeroAsNull_int64() {

	var a int64

	a = 3
	fmt.Println(ZeroAsNull(a) == int64(3))
	fmt.Println(ZeroAsNull(&a) == int64(3))

	a = 0
	fmt.Println(ZeroAsNull(a) == nil)
	fmt.Println(ZeroAsNull(&a) == int64(0))
	fmt.Println(ZeroAsNull((*int64)(nil)) == nil)

	// Output:
	// true
	// true
	// true
	// true
	// true
}

func ExampleZeroAsNull_string() {

	var a string

	a = "a"
	fmt.Println(ZeroAsNull(a) == "a")
	fmt.Println(ZeroAsNull(&a) == "a")

	a = ""
	fmt.Println(ZeroAsNull(a) == nil)
	fmt.Println(ZeroAsNull(&a) == "")
	fmt.Println(ZeroAsNull((*string)(nil)) == nil)

	// Output:
	// true
	// true
	// true
	// true
	// true
}

type Model struct {
	//
	Float64Null       float64
	Float64NotNull    float64
	Float64PtrNull    *float64
	Float64PtrNotNull *float64
	//
	Float32Null       float32
	Float32NotNull    float32
	Float32PtrNull    *float32
	Float32PtrNotNull *float32
	//
	Int64Null       int64
	Int64NotNull    int64
	Int64PtrNull    *int64
	Int64PtrNotNull *int64
	//
	Int32Null       int32
	Int32NotNull    int32
	Int32PtrNull    *int32
	Int32PtrNotNull *int32
	//
	Int16Null       int16
	Int16NotNull    int16
	Int16PtrNull    *int16
	Int16PtrNotNull *int16
	//
	CustomTypeIntNull       MyInt64Type
	CustomTypeIntNotNull    MyInt64Type
	CustomTypeIntPtrNull    *MyInt64Type
	CustomTypeIntPtrNotNull *MyInt64Type
	//
	ByteNull       byte
	ByteNotNull    byte
	BytePtrNull    *byte
	BytePtrNotNull *byte
	//
	BoolNull       bool
	BoolNotNull    bool
	BoolPtrNull    *bool
	BoolPtrNotNull *bool
	//
	StringNull       string
	StringNotNull    string
	StringPtrNull    *string
	StringPtrNotNull *string
	//
	CustomStringNull       MyStringType
	CustomStringNotNull    MyStringType
	CustomStringPtrNull    *MyStringType
	CustomStringPtrNotNull *MyStringType
	//
	TimeNull       time.Time
	TimeNotNull    time.Time
	TimePtrNull    *time.Time
	TimePtrNotNull *time.Time
	//
	JsonNull       ModelNested
	JsonNotNull    ModelNested
	JsonPtrNull    *ModelNested
	JsonPtrNotNull *ModelNested
}

type ModelNested struct {
	Value int
}

type MyInt64Type int64
type MyStringType string

type testDriver struct{}
type testConn struct{}

func (d *testDriver) Open(name string) (driver.Conn, error) {
	return &testConn{}, nil
}
func (c *testConn) Close() error {
	return nil
}
func (c *testConn) Begin() (_ driver.Tx, err error) {
	return c, fmt.Errorf("Begin method not implemented")
}
func (c *testConn) Prepare(query string) (driver.Stmt, error) {
	return nil, fmt.Errorf("Prepare method not implemented")
}
func (c *testConn) Commit() error {
	return fmt.Errorf("Commit method not implemented")
}
func (c *testConn) Rollback() error {
	return fmt.Errorf("Rollback method not implemented")
}
func (c *testConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	values := []any{
		// Float64
		nil, float64(2.5),
		nil, float64(3.5),
		// Float32
		nil, float32(4.5),
		nil, float32(5.5),
		// Int64
		nil, int64(4),
		nil, int64(5),
		// Int32
		nil, int32(6),
		nil, int32(7),
		// Int16
		nil, int16(8),
		nil, int16(9),
		// CustomTypeInt64
		nil, int64(2),
		nil, int64(3),
		// Byte
		nil, byte(10),
		nil, byte(11),
		// Bool
		nil, true,
		nil, true,
		// String
		nil, "Hello",
		nil, "World",
		// CustomString
		nil, "Hi",
		nil, "There",
		// Time
		nil, time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC),
		nil, time.Date(2022, 6, 7, 8, 9, 10, 0, time.UTC),
		// Json
		nil, "{\"Value\": 22}",
		nil, "{\"Value\": 33}",
	}
	columns := make([]string, len(values))
	return &results{columns, values}, nil
}

type results struct {
	columns []string
	values  []any
}

func (r *results) Columns() []string {
	return r.columns
}
func (r *results) Close() error {
	return nil
}
func (r *results) Next(dest []driver.Value) error {
	for i, v := range r.values {
		dest[i] = v
	}
	return nil
}
