package main

import (
    "fmt"
    "os/exec"
    "flag"
    // "os"
    "strings"
)

func main() {
    fsystem := flag.String("f", "ext4", "file system")
    flag.Parse()
    b, err := exec.Command("df", "-H", "-t", *fsystem).Output()
    if err != nil {
        return
    }

    str := string(b)
    str_arr := strings.Split(str, "\n")
    data :=(str_arr[1])
    tmp_arr := strings.Split(data, " ")

    var data_arr []string
    for _, d := range tmp_arr {
        if(d != "") {
            data_arr = append(data_arr, d)
        }
    }

    // 0 /dev/sda2
    // 1 31G Size
    // 2 13G Use
    // 3 17G Ava
    // 4 44% Use%
    // 5 /

    st := fmt.Sprintf("Mount:%s Size:%s Use:%s Ava:%s Use%%:%s",
                       data_arr[0], data_arr[1], data_arr[2], data_arr[3], data_arr[4])

    fmt.Println(st)

}
