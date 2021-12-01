package utils;

import (
  "fmt"
  "strings"
  "strconv"
  "net/http"
  "io/ioutil"
);

const url_format string = "https://adventofcode.com/%d/day/%d/input";
const session string = "PASTE_HERE"; // TODO paste your session-cookie!!

func Get_input(year int, day int) string {
    var url = fmt.Sprintf(url_format, year, day);

    // Declare http client
    var client = &http.Client{};

    // Declare HTTP Method and Url
    var req, err = http.NewRequest("GET", url, nil);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }   
    
    // Set cookie
    req.Header.Set("Cookie", fmt.Sprintf("session=%s; count=x", session));

    resp, err := client.Do(req);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }
    // Read response
    data, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    return string(data);
}

func Trim_array(strs []string) []string {
    for i, str := range strs {
        strs[i] = strings.Trim(str, " ");
    }
    return strs;
}

func StrToInt_array(strs []string) ([]int, error) {
    var ints = make([]int, len(strs));
    for i, str := range strs {
        var d, err = strconv.Atoi(str);
        if err != nil {
            return nil, err;
        }
        ints[i] = d;
    }
    return ints, nil;
}

func StrToFloat_array(strs []string) ([]float64, error) {
    var floats = make([]float64,len(strs));
    for i, str := range strs {
        var f, err = strconv.ParseFloat(str, 64);
        if err != nil {
            return nil, err;
        }
        floats[i] = f;
    }
    return floats, nil;
}
