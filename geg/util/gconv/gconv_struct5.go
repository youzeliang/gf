package main

import (
    "gitee.com/johng/gf/g/util/gconv"
    "gitee.com/johng/gf/g"
    "fmt"
)

func main() {
    type Score struct {
        Name   string
        Result int
    }
    type User struct {
        Scores []Score
    }

    user   := new(User)
    scores := map[string]interface{}{
        "Scores" : []interface{}{
            map[string]interface{}{
                "Name"   : "john",
                "Result" : 100,
            },
        },
    }

    if err := gconv.Struct(scores, user); err != nil {
        fmt.Println(err)
    } else {
        g.Dump(user)
    }
}