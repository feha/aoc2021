package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 11:47:07
  * p1 done - 11:53:05
  *   p1 delta - 11:53:05-11:47:07 = 00:05:58
  * p2 done - 11:54:55
  *   p2 delta - 11:54:55-11:53:05 = 00:01:50
  * total time - 11:54:55-11:47:07 = 00:07:48
  * total time - 00:05:58+00:01:50 = 00:07:48
  */

func main() {
    var input, _ = utils.Get_input(2021, 02);
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
        if (part2(part2_test_input[i]) != part2_test_output[i]) {
            success = false;
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    part2(part2_test_input[i]),
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
    `forward 5
    down 5
    forward 8
    up 3
    down 8
    forward 2`,
};
var part1_test_output = []string{
    `150`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);

    x, y := 0, 0;
    for _, str := range inputs {
        args := strings.Split(str, " ");
        cmd := args[0];
        n, err := strconv.Atoi(args[1]);
        if err != nil {
            fmt.Printf("error = %s \n", err);
        }
        
        switch cmd {
        case "forward":
            x+= n;
        case "down":
            y+= n;
        case "up":
            y-= n;
        default:
            fmt.Printf("error: cmd=%s \n", cmd);
        }
    }

    // return "";
    return strconv.Itoa(x*y);
}

var part2_test_input = []string{
    `forward 5
    down 5
    forward 8
    up 3
    down 8
    forward 2`,
};
var part2_test_output = []string{
    `900`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // var nums, _ = utils.StrToInt_array(inputs);

    x, y, aim := 0, 0, 0;
    for _, str := range inputs {
        args := strings.Split(str, " ");
        cmd := args[0];
        n, err := strconv.Atoi(args[1]);
        if err != nil {
            fmt.Printf("error = %s \n", err);
        }
        
        switch cmd {
        case "forward":
            x += n;
            y += aim*n;
        case "down":
            aim += n;
        case "up":
            aim -= n;
        default:
            fmt.Printf("error: cmd=%s \n", cmd);
        }
    }

    // return "";
    return strconv.Itoa(x*y);
}
