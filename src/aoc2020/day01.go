package main;

import (
    "fmt"
    "strings"
    "strconv"
    "aoc/libs/utils"
);

/**
  * Start - 16:07:02
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input = utils.Get_input(2020, 1);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
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
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `1721
    979
    366
    299
    675
    1456`,
};
var part1_test_output = []string{
    `514579`,
};
func part1(input string) string {
    var inputs = strings.Split(strings.Trim(input, separator), separator);
    var args = utils.Trim_array(inputs);
    var nums, _ = utils.StrToInt_array(args);

    for i := range make([]int, len(nums)) {
        for j:=i+1; j < len(nums); j++ {
            x, y := nums[i], nums[j];
            if (x+y == 2020) {
                return strconv.Itoa(x*y);
            }
        }
    }

    return "";
}

var part2_test_input = []string{
    ``,
};
var part2_test_output = []string{
    ``,
};
func part2(input string) string {
    var inputs = strings.Split(strings.Trim(input, "separator"), "separator");

    // ...

    return "";
}
