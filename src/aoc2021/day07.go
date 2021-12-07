package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 14:18:50
  * p1 done - 14:33:20
  * p2 done - 14:36:45
  */

func main() {
    input, _ := utils.Get_input(2021, 07);
    // fmt.Printf("Input: %s \n", input);

    success := true;
    for i := range part1_test_input {
        result := part1(part1_test_input[i])
        if (result != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    result,
                    part1_test_output[i]);
            break;
        }
    }

    fmt.Printf("part1 minitest success: %t! \n", success);
    if success {
        p1 := part1(input);
        fmt.Printf("part1: %s\n\n", p1);
    }
    
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

    if success {
        p2 := part2(input);
        fmt.Printf("part2: %s\n", p2);
    }
}

const separator string = ",";

var part1_test_input = []string{
    `16,1,2,0,4,2,7,1,2,14`,
};
var part1_test_output = []string{
    `37`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator+"\n"), separator));
    nums, err := utils.StrToInt_array(inputs);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    min, max := get_minmax(nums);
    fuels := make(map[int]int);
    // result := 0;
    for i:=min; i < max; i++ {
        // fuel := 0;
        for _, n := range nums {
            fuels[i] += utils.Abs(n-i);
// 
        }
    }

    map_max := 0;
    for _, n := range fuels {
        if map_max < n {
            map_max = n;
        }
    }
    result := map_max;
    for _, n := range fuels {
        if result > n {
            result = n;
        }
    }

    return strconv.Itoa(result);
}

func get_minmax(nums []int) (int, int) {
    min, max := nums[0], 0;
    for _, n := range nums {
        if max < n {
            max = n
        }
        if min > n {
            min = n
        }
    }
    return min, max;
}

var part2_test_input = []string{
    `16,1,2,0,4,2,7,1,2,14`,
};
var part2_test_output = []string{
    `168`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator+"\n"), separator));
    nums, err := utils.StrToInt_array(inputs);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    min, max := get_minmax(nums);
    fuels := make(map[int]int);
    // result := 0;
    for i:=min; i < max; i++ {
        // fuel := 0;
        for _, n := range nums {
            error := utils.Abs(n-i);
            cost := 0;
            for step_cost := 1; step_cost <= error; step_cost++ {
                cost += step_cost;
            }
            fuels[i] += cost;
    // 
        }
    }

    map_max := 0;
    for _, n := range fuels {
        if map_max < n {
            map_max = n;
        }
    }
    result := map_max;
    for _, n := range fuels {
        if result > n {
            result = n;
        }
    }

    return strconv.Itoa(result);
}
