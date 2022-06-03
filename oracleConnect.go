package svcc-Oracle

import (
    "fmt"
    "log"
    "os"
)

func OracleConnect() *sql.DB {

    user := os.Getenv("ORACLE_USER")
    password := os.Getenv("ORACLE_PASS")
    host := os.Getenv("ORACLE_HOST")
    sid := os.Getenv("ORACLE_SID")

    if user == "" || password == "" || host == "" || sid == "" {
        fmt.Println("Please set ORACLE_USER, ORACLE_PASS, ORACLE_HOST, and ORACLE_SID first")
        os.Exit(1)
    }

    db, err := sql.Open("godror", user+"/"+password+"@"+host+"/"+sid)
    if err != nil {
        log.Fatal(err)
    }

    db.SetMaxOpenConns(1)

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    return db
}
