# Go rest API Example Use Fiber Framework

## Project Structure

```bash
.
├── app.go
├── config
│   └── config.go
├── database
│   ├── database.go
│   └── schema.go
├── go.mod
├── go.sum
├── handlers
│   ├── handler.go
│   └── user.go
├── middlewares
│   └── middleware.go
├── models
│   └── user.go
├── router
│   ├── router.go
│   ├── test.go
│   └── user.go
└── static // Optional
    ├── private
    │   └── 404.html
    └── public
        ├── css
        │   └── style.css
        ├── img
        │   ├── logo.svg
        │   └── meme.png
        ├── index.html
        └── js
            └── app.js
```

## Database Connection

## Parameter Placeholder Syntax

**Database**|MySQL | PostgreSQL | Oracle |
------|----- |             ----- |            ----- |
**DSN Format** | username:password@protocol(address)/dbname?param=value | host={host} port={port} user={user} password={password} dbname={dbname} sslmode=disable | Coming Soon |
**Open Connection** | sql.Open("mysql", dsn) | sql.Open("postgres", dsn) | Coming Soon|
**Driver** | _ "github.com/go-sql-driver/mysql" |  _ "github.com/lib/pq" | Coming Soon |

MySQL |               PostgreSQL |            Oracle |
----- |             ----- |            ----- |
WHERE col = ?  |    WHERE col = $1  |   WHERE col = :col |
VALUES(?, ?, ?) |   VALUES($1, $2, $3)  |  VALUES(:val1, :val2, :val3) |

