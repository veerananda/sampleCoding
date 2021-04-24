package main

import (
        "fmt"
        "os"
        "io/ioutil"
        "strings"
        "log"
        "bufio"
        "reflect"
        )

func checkScrambleExistence(one string,two string) bool {
fmt.Println(one)
fmt.Println(two)
map1:=make(map[string]int)
map2:=make(map[string]int)
length:=len(one)
for i:=0;i<length-1;i++ {
var s1 string=string(one[i])
var s2 string=string(two[i])
value1,ok:=map1[s1]
if ok {
map1[s1] = value1+1
} else {
map1[s1] = 1
}
value2,ok:=map2[s2]
if ok {
map2[s2] = value2+1
} else {
map2[s2] = 1
}
}
res1:=reflect.DeepEqual(map1,map2)
return res1
}


func main(){
var filename1 string
var filename2 string
fmt.Println("enter file dictionary file")
fmt.Scanln(&filename1)
fmt.Println("enter file that contains the long string")
fmt.Scanln(&filename2)
var result bool
var count int=0
data1,err:=ioutil.ReadFile(filename2)
if err!=nil {
        log.Panicf("failed reading : %s", err)
}
v1:=len(data1)
data2,err:=os.Open(filename1)
if err!=nil {
        log.Panicf("failed reading : %s", err)
}
scanner:=bufio.NewScanner(data2)
scanner.Split(bufio.ScanLines)
for scanner.Scan(){
result=false
data3:=scanner.Text()
fmt.Println(data3)
if strings.Contains(string(data1),string(data3)) {
        count++
        continue
}
v2:=len(data3)
for i:=0;i<v1-v2+1;i++{
if data1[i] == data3[0]{
if data1[i+v2-1] == data3[v2-1]{
chopString:=data1[i:i+v2]
result=checkScrambleExistence(string(data3),string(chopString))
fmt.Println(result)
}
}
if result==true {
        count++
        break
}

}
}
fmt.Println("case #1 :",count)
}
