package sqlu

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

func Args(vs ...any) args {
	a := make([]any, 0, len(vs))
	for _, v := range vs {
		vOf := reflect.ValueOf(v)
		sub := vOf.Type()
		for sub.Kind() == reflect.Pointer {
			sub = sub.Elem()
		}
		switch sub.Kind() {
		case reflect.Bool:
			a = append(a, &sqlNullBool{v: vOf})
		case reflect.Uint8:
			a = append(a, &sqlNullByte{v: vOf})
		case reflect.Int16:
			a = append(a, &sqlNullInt16{v: vOf})
		case reflect.Int32:
			a = append(a, &sqlNullInt32{v: vOf})
		case reflect.Int64:
			a = append(a, &sqlNullInt64{v: vOf})
		case reflect.Float32:
			a = append(a, &sqlNullFloat32{v: vOf})
		case reflect.Float64:
			a = append(a, &sqlNullFloat64{v: vOf})
		case reflect.Struct:
			switch reflect.Indirect(vOf).Interface().(type) {
			case time.Time, *time.Time:
				a = append(a, &sqlNullTime{v: vOf})
			default:
				a = append(a, &sqlNullJson{v: vOf})
			}
		default:
			a = append(a, &sqlNullString{v: vOf})
		}
	}
	return args(a)
}

type args []any

// #region Nullable Types

type sqlNullString struct {
	s sql.NullString
	v reflect.Value
}

func (n *sqlNullString) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.String).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.String).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullString) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullTime struct {
	s sql.NullTime
	v reflect.Value
}

func (n *sqlNullTime) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Time).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Time).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullTime) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullBool struct {
	s sql.NullBool
	v reflect.Value
}

func (n *sqlNullBool) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Bool).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Bool).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullBool) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullByte struct {
	s sql.NullByte
	v reflect.Value
}

func (n *sqlNullByte) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Byte).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Byte).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullByte) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullInt16 struct {
	s sql.NullInt16
	v reflect.Value
}

func (n *sqlNullInt16) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Int16).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Int16).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullInt16) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullInt32 struct {
	s sql.NullInt32
	v reflect.Value
}

func (n *sqlNullInt32) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Int32).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Int32).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullInt32) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullInt64 struct {
	s sql.NullInt64
	v reflect.Value
}

func (n *sqlNullInt64) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Int64).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Int64).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullInt64) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullFloat32 struct {
	s sql.NullFloat64
	v reflect.Value
}

func (n *sqlNullFloat32) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		f := float32(n.s.Float64)
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&f).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(f).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullFloat32) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullFloat64 struct {
	s sql.NullFloat64
	v reflect.Value
}

func (n *sqlNullFloat64) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		if dst := n.v.Elem(); dst.Kind() == reflect.Pointer {
			dst.Set(reflect.ValueOf(&n.s.Float64).Convert(dst.Type()))
		} else {
			dst.Set(reflect.ValueOf(n.s.Float64).Convert(dst.Type()))
		}
	}
	return nil
}

func (n sqlNullFloat64) Value() (driver.Value, error) {
	return n.s.Value()
}

type sqlNullJson struct {
	s sql.NullString
	v reflect.Value
}

func (n *sqlNullJson) Scan(value any) error {
	if err := n.s.Scan(value); err != nil {
		return err
	}
	if n.s.Valid {
		v := n.v.Interface()
		if err := json.Unmarshal([]byte(n.s.String), &v); err != nil {
			return err
		}
	}
	return nil
}

func (n sqlNullJson) Value() (driver.Value, error) {
	return n.s.Value()
}

// #endregion Nullable Types

func ZeroAsNulls(vs ...any) []any {
	for i, v := range vs {
		vs[i] = ZeroAsNull(v)
	}
	return vs
}

func ZeroAsNull(v any) any {
	vOf := reflect.ValueOf(v)
	if vOf.Kind() == reflect.Pointer {
		if vOf.IsNil() {
			return nil
		}
		return vOf.Elem().Interface()
	}
	if vOf.IsZero() {
		return nil
	}
	return vOf.Interface()
}
