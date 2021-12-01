package main;

import (
    "fmt"
    "strings"
    "strconv"
    "aoc/libs/utils"
);

/**
  * Start - 13:16:28
  * p1 done - 13:23:00
  *   p1 delta - 13:23:00-13:16:28 = 00:06:32
  * p2 done - 13:28:30
  *   p2 delta - 13:28:30-13:23:00 = 00:05:30
  * total time - 13:28:30-13:16:28 = 00:12:02
  * total time - 00:06:32+00:05:30 = 00:12:02
  */

func main() {
    var input, _ = utils.Get_input(2021, 01);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n"
                    ,part1_test_input[i],
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
    `199
200
208
210
200
207
240
269
260
263`,
};
var part1_test_output = []string{
    `7`,
};
func part1(input string) string {
    var inputs = strings.Split(strings.Trim(input, separator), separator);
    var args = utils.Trim_array(inputs);
    var nums, _ = utils.StrToInt_array(args);

    var result = 0;
    var last = nums[0];
    for i:=1; i < len(nums); i++ {
        if last < nums[i] {
            result = result + 1;
        }
        last = nums[i];
    }

    return strconv.Itoa(result);
}

var part2_test_input = []string{
    `199
200
208
210
200
207
240
269
260
263`,
};
var part2_test_output = []string{
    `5`,
};
func part2(input string) string {
    var inputs = strings.Split(strings.Trim(input, separator), separator);
    var args = utils.Trim_array(inputs);
    var nums, _ = utils.StrToInt_array(args);

    var window_size = 3;
    var new_nums = make([]int, len(nums)-(window_size-1));
    for i:=0; i < len(new_nums); i++ {
        new_nums[i] = nums[i] + nums[i+1] + nums[i+2];
    }

    var result = 0;
    var last = new_nums[0];
    for i:=1; i < len(new_nums); i++ {
        if last < new_nums[i] {
            result = result + 1;
        }
        last = new_nums[i];
    }

    return strconv.Itoa(result);
}
