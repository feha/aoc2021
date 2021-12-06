package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 14:07:08
  * p1 done - 14:21:32
  * p2 done - 14:47:48
  */

func main() {
    input, _ := utils.Get_input(2021, 06);
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
    `3,4,3,1,2`,
};
var part1_test_output = []string{
    `5934`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, "\n"), separator));
    nums, err := utils.StrToInt_array(inputs);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    result := sum(iterate_days2(fishes_to_schools(nums), 80, 7, 2));
    // result := len(iterate_days(nums, 80, 7, 2));

    return strconv.Itoa(result);
}

func iterate_days(fishes []int, days int, spawn_time int, gestation_period int) []int {
    for i:=0; i < days; i++ {
        fishes = iterate_fishes(fishes, spawn_time, gestation_period);
    }
    return fishes;
}

func iterate_fishes(fishes []int, spawn_time int, gestation_period int) []int {
    new_fishes := make([]int, 0);
    for _, t := range fishes {
        if t == 0 {
            new_fishes = append(new_fishes, spawn_time + gestation_period - 1);
            t = spawn_time;
        }
        new_fishes = append(new_fishes, t-1);
    }
    return new_fishes;
}

var part2_test_input = []string{
    `3,4,3,1,2`,
};
var part2_test_output = []string{
    `26984457539`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, "\n"), separator));
    nums, err := utils.StrToInt_array(inputs);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    result := sum(iterate_days2(fishes_to_schools(nums), 256, 7, 2));

    return strconv.Itoa(result);
}

func sum(schools map[int]int) int {
    sum := 0;
    for _, n := range schools {
        sum += n;
    }
    return sum;
}

func fishes_to_schools(fishes []int) map[int]int {
    // schools := map[int]int{};
    schools := make(map[int]int);
    for _, t := range fishes {
        schools[t]++;
    }
    return schools;
}

func iterate_days2(schools map[int]int, days int, spawn_time int, gestation_period int) map[int]int {
    for i:=0; i < days; i++ {
        schools = iterate_schools(schools, spawn_time, gestation_period);
    }
    return schools;
}

func iterate_schools(schools map[int]int, spawn_time int, gestation_period int) map[int]int {
    new_schools := make(map[int]int, 0);
    for t, n := range schools {
        if t == 0 {
            new_schools[spawn_time - 1] = new_schools[spawn_time - 1] + n;
            new_schools[spawn_time + gestation_period - 1] = n;
        } else {
            new_schools[t-1] = new_schools[t-1] + n;
        }
    }
    return new_schools;
}