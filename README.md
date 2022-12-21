# sqlu

A very simple util for golang's `database/sql` package.

# Example

```go
import "github.com/ofabricio/sqlu"

func main() {

    u := User{}

    db.QueryRow(`
        SELECT
            33, 'John', JSON_OBJECT('Country', 'Home Sweet Home')
        LIMIT 1
    `).Scan(sqlu.Args(&u.ID, &u.Name, &u.Address)...)

    fmt.Println(u) // {33 John {Home Sweet Home}}
}

type User struct {
    ID      int64
    Name    string
    Address struct {
        Country string
    }
}
```
