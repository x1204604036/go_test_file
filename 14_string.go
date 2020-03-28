package main

import (
    "fmt"
    "unicode/utf8"
)

func printBytes(s string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x", s[i])
    }
}


//有字符的UTF-8编码占用了两个字节
func printChars(s string) {
    for i := 0; i<len(s); i++ {
        fmt.Printf("%c", s[i])
    }
}


func printCharsRunes(s string) {
    runes := []rune(s)
    for i := 0; i < len(runes); i++ {
        fmt.Printf("%c", runes[i])
    }
}


func length(s string){
    fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}


func mutate(s []rune) string {
    s[0] = 'a'
    return string(s)
}


func main() {
    //name := "hello world"
    //fmt.Println(name)
    //for i, v := range name {
    //    fmt.Printf("%s ", i, v, name[i])
    //}
    name := "Señor"
    printBytes(name)
    fmt.Println("\n")
    printChars(name)
    fmt.Println("\n")
    printCharsRunes(name)

    fmt.Println("start\n")
    for i, v := range(name){
        fmt.Printf("%c, %d\n", v, i)
    }

    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)

    byteSlice2 := []byte{67, 97, 102, 195, 169}
    str2 := string(byteSlice2)
    fmt.Println(str2)

    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str3 := string(runeSlice)
    fmt.Println(str3)


    word1 := "Señor"
    length(word1)
    word2 := "Pets"
    length(word2)

    //修改字符串  先将字符串切片，然后修改，再转换成字符串
    h := "hello"
    fmt.Println(mutate([]rune(h)))
}










