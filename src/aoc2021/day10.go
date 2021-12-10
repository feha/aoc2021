package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "os"
    "sort"
);

/**
  * Start - 17:22:33
  * p1 done - 17:41:15
  * p2 done - 18:05:56
  */

func main() {
    input, _ := utils.Get_input(2021, 10);

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
    `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`,
};
var part1_test_output = []string{
    `26397`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    result := 0;
    for _, line := range inputs {
        buffer := "";
        for _, c := range strings.Split(line, "") {
            if !is_closing(c) {
                buffer += c;
            } else {
                current := strings.Split(buffer, "")[len(buffer)-1];
                if get_matching(current) != c {
                    result+= get_points(c);
                    break;
                } else {
                    buffer = buffer[:len(buffer)-1];
                }
            }
        }
    }

    return strconv.Itoa(result);
}
func get_matching(c string) string {
    mapping := map[string]string {
        "(": ")",
        ")": "(",
        "[": "]",
        "]": "[",
        "{": "}",
        "}": "{",
        "<": ">",
        ">": "<",
    };
    return mapping[c]
}
func is_closing(c string) bool {
    switch c {
    case ")", "]", "}", ">":
        return true;
    default:
        return false;
    }
}
func get_points(c string) int {
    switch c {
    case ")":
        return 3;
    case "]":
        return 57;
    case "}":
        return 1197;
    case ">":
        return 25137;
    default:
        fmt.Fprintf(os.Stderr, "ERROR - Undefined behaviour for this state!");
    }
    return 0;
}

var part2_test_input = []string{
    `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`,
};
var part2_test_output = []string{
    `288957`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // valid := strings.Split(")]}>", "");
    
    scores := []int{};
    for _, line := range inputs {
        buffer := "";
        corrupted := false;
        for _, c := range strings.Split(line, "") {
            if !is_closing(c) {
                buffer += c;
            } else {
                current := strings.Split(buffer, "")[len(buffer)-1];
                if get_matching(current) != c {
                    corrupted = true;
                    break;
                } else {
                    buffer = buffer[:len(buffer)-1];
                }
            }
        }
        if !corrupted {
            as_arr := strings.Split(buffer, "");
            correction := "";
            for i := len(buffer)-1; i >= 0; i-- {
                correction += get_matching(as_arr[i]);
            }
            scores = append(scores, get_autocomplete_points(correction));
        }
    }
    
    sort.Ints(scores)
    
    result := scores[len(scores)/2];
    return strconv.Itoa(result);
}

func get_autocomplete_points(str string) int {
    if len(str) == 0 {
        return 0;
    }
    score := 0;
    for _, c := range strings.Split(str, "") {
        score *= 5;
        switch c {
        case ")":
            score += 1;
        case "]":
            score += 2;
        case "}":
            score += 3;
        case ">":
            score += 4;
        default:
            fmt.Fprintf(os.Stderr, "ERROR - Undefined behaviour for this state! c='%s' - str='%s'\n", c, str);
            return -1;
        }
    }
    return score;
}