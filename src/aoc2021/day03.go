package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 16:15:46
  * p1 done - 16:39:50
  * p2 done - 17:13:50
  */

func main() {
    var input, _ = utils.Get_input(2021, 03);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    part1(part1_test_input[i]),
                    part1_test_output[i]);
            break;
        }
    }

    fmt.Printf("part1 minitest success: %t! \n", success);
    p1 := part1(input);
    fmt.Printf("part1: %s\n\n", p1);
    
    success = true;
    for i := range part2_test_input {
        result := part2(part2_test_input[i])
        if (result != part2_test_output[i]) {
            success = false;
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    result,
                    part2_test_output[i]);
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `00100
    11110
    10110
    10111
    10101
    01111
    00111
    11100
    10000
    11001
    00010
    01010`,
};
var part1_test_output = []string{ // (gamma=22, epsilon=9)
    `198`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    var size = len(inputs[0])
    histogram := make([][]int, size);
    for i := 0; i < size; i++  {
        histogram[i] = make([]int, 2);
    }
    for _, line := range inputs {
        for i := 0; i < size; i++  {
            col := histogram[i]
            var d, err = strconv.Atoi(string(line[i]));
            if err != nil {
                fmt.Printf("error = %s \n", err);
            }
            col[d] = col[d] + 1;
        }
    }

    var gamma, epsilon = 0, 0;
    var str1, str2 = "", "";
    for i := 0; i < size; i++  {
        gamma = gamma << 1;
        epsilon = epsilon << 1;
        zeroes := histogram[i][0];
        ones := histogram[i][1];
        if zeroes > ones {
            // gamma += 0;
            epsilon += 1;
            str1+="0";
            str2+="1";
        } else if zeroes < ones {
            gamma += 1;
            // epsilon += 0;
            str1+="1";
            str2+="0";
        } else {
            fmt.Println("EQUALS ERRORRRS");
        }
    }

    // return "";
    return strconv.Itoa(gamma*epsilon);
}

var part2_test_input = []string{
    `00100
    11110
    10110
    10111
    10101
    01111
    00111
    11100
    10000
    11001
    00010
    01010`,
};
var part2_test_output = []string{
    `230`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    fmt.Println("--------------");
    var pos = 0;
    o2 := make([]string, len(inputs));
    copy(o2,inputs);
    for len(o2) > 1 || pos > 99 {
        histogram := hist(o2);
        new_inputs := make([]string, 0);

        zeroes := histogram[pos][0];
        ones := histogram[pos][1];
        for i := 0; i < len(o2); i++  {
            line := o2[i];
            char := string(line[pos]);
            if zeroes > ones {
                if char == "0" {
                    new_inputs = append(new_inputs, line);
                }
            } else if ones >= zeroes {
                if char == "1" {
                    new_inputs = append(new_inputs, line);
                }
            } else {
                fmt.Println("????");
            }
        }

        o2 = new_inputs;
        pos++;
    }
    fmt.Println("o2:",o2);

    pos = 0;
    co2 := make([]string, len(inputs));
    copy(co2,inputs);
    for len(co2) > 1 || pos > 99 {
        histogram := hist(co2);
        new_inputs := make([]string, 0);

        zeroes := histogram[pos][0];
        ones := histogram[pos][1];
        for i := 0; i < len(co2); i++  {
            line := co2[i];
            char := string(line[pos]);
            if zeroes <= ones {
                if char == "0" {
                    new_inputs = append(new_inputs, line);
                }
            } else if ones < zeroes {
                if char == "1" {
                    new_inputs = append(new_inputs, line);
                }
            } else {
                fmt.Println("????");
            }
        }

        co2 = new_inputs;
        pos++;
    }
    fmt.Println("co2:",co2);
    bla1, err := strconv.ParseInt(o2[0],2, 64);
    if err != nil {
        fmt.Printf("error1 = %s \n", err);
    }
    bla2, err := strconv.ParseInt(co2[0],2, 64);
    if err != nil {
        fmt.Printf("error2 = %s \n", err);
    }
    fmt.Println(bla1,bla2);

    // return "";
    return strconv.Itoa(int(bla1)*int(bla2));
}

func hist(strs []string) [][]int {
    var size = len(strs[0]);
    histogram := make([][]int, size);
    for i := 0; i < size; i++  {
        histogram[i] = make([]int, 2);
    }
    for _, line := range strs {
        for i := 0; i < size; i++  {
            col := histogram[i]
            var d, err = strconv.Atoi(string(line[i]));
            if err != nil {
                fmt.Printf("error = %s \n", err);
            }
            col[d] = col[d] + 1;
        }
    }
    return histogram;
}