package main;

import (
    "aoc/libs/utils"
    "fmt"
    "os"
    "strings"
    "strconv"
);

/**
  * Start - 16:15:46
  * p1 done - 16:39:50
  * p2 done - 17:13:50
  */

func main() {
    input, _ := utils.Get_input(2021, 03);
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
var part1_test_output = []string{
    `198`, // (gamma=22, epsilon=9)
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    gamma, epsilon := gamma_epsilon(inputs);

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
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    o2, err := strconv.ParseInt(
            filter(inputs, func(lhs int, rhs int)bool {return lhs > rhs}),
            2, 64);
    if err != nil {
        fmt.Printf("error1 = %s \n", err);
    }
    co2, err := strconv.ParseInt(
            filter(inputs, func(lhs int, rhs int)bool {return lhs <= rhs}),
            2, 64);
    if err != nil {
        fmt.Printf("error2 = %s \n", err);
    }

    return strconv.Itoa(int(o2)*int(co2));
}

func gamma_epsilon(inputs []string) (int, int) {
    zeroes, ones := hist(inputs);

    gamma, epsilon := 0, 0;
    for i := 0; i < len(inputs[0]); i++  {
        gamma = gamma << 1;
        epsilon = epsilon << 1;
        zeroes := zeroes[i];
        ones := ones[i];
        if zeroes >= ones {
            // gamma += 0;
            epsilon += 1;
        } else if zeroes <= ones {
            gamma += 1;
            // epsilon += 0;
        } else {
            fmt.Fprintf(os.Stderr, "ERROR - Undefined behaviour for this state!");
        }
    }

    return gamma, epsilon;
}

func filter(inputs []string, zero_cond func(int,int)bool) string {
    pos := 0;
    candidates := make([]string, len(inputs));
    copy(candidates, inputs);
    for len(candidates) > 1 || pos > 99 {
        h0, h1 := hist(candidates);
        pass := make([]string, 0);

        zeroes := h0[pos];
        ones := h1[pos];
        for _, candidate := range candidates  {
            switch string(candidate[pos]) {
            case "0":
                if zero_cond(zeroes, ones) {
                    pass = append(pass, candidate);
                }
            case "1":
                if !zero_cond(zeroes, ones) {
                    pass = append(pass, candidate);
                }
            default:
                fmt.Fprintf(os.Stderr, "ERROR - Unreachable reached!");
            }
        }
        candidates = pass;

        pos++;
    }
    return candidates[0];
}

func hist(strs []string) ([]int, []int) {
    size := len(strs[0]);
    zeroes, ones := make([]int, size), make([]int, size);
    for i := 0; i < size; i++  {
        for _, str := range strs {
            d, err := strconv.Atoi(string(str[i]));
            if err != nil {
                fmt.Printf("error = %s \n", err);
            }
            zeroes[i] += 1-d; // 1 = 1-(0)
            ones[i] += d;
        }
    }
    return zeroes, ones;
}