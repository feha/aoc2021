package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "os"
    "errors"
    "sort"
);

/**
  * Start - 15:05:01
  * p1 done - 15:22:07
  * p2 done -  17:21:44
  */

func main() {
    input, _ := utils.Get_input(2021, 8);
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
    `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`,
    `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`,
};
var part1_test_output = []string{
    `0`,
    `26`,
};
func part1(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    // nums, _ := utils.StrToInt_array(inputs);

    result := 0;
    for _, line := range lines {
        notes := strings.Split(line, " | ");
        warped_digits_second := strings.Split(notes[1], " ");
        result += count_easy_digits(warped_digits_second);
    }

    return strconv.Itoa(result);
}

func count_easy_digits(warped_digits []string) int {
    count := 0;
    for _, str := range warped_digits {
        switch len(str) {
        case 2:
            count++; // 1
        case 4:
            count++; // 4
        case 3:
            count++; // 7
        case 7:
            count++; // 8
        default:
        }
    }
    return count;
}

var part2_test_input = []string{
    `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`,
    `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`,
};
var part2_test_output = []string{
    `5353`,
    `61229`,
};
func part2(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));


    result := 0;
    for _, line := range lines {
        notes := strings.Split(line, " | ");
        unordered_digits := strings.Split(notes[0], " ");
        warped_output := strings.Split(notes[1], " ");

        warped_output_sorted := make([]string, len(warped_output));
        for i, str := range warped_output {
            warped_output_sorted[i] = sort_str(str);
        }

        digits, _ := find_digits_and_connections(unordered_digits);

        digits_sorted := make(map[int]string);
        for i, digit := range digits {
            digits_sorted[i] = sort_str(digit);
        }

        inverse_digits := inverse_mapping(digits_sorted);

        output_str := "";
        for _, warped_digit := range warped_output_sorted {
            digit := inverse_digits[warped_digit];
            output_str += strconv.Itoa(digit);
        }
        output, err := strconv.Atoi(output_str);
        if err != nil {
            fmt.Printf("error = %s \n", err);
        }

        result += output;
    }

    return strconv.Itoa(result);
}

func sort_str(str string) string {
    runes := []rune(str);
    sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] });
    return string(runes);
}

// can't believe I have to implement a set-difference myself..
func set_difference(a string, b string) string {
    result_set := "";
    for _, e := range a {
        // exists := false;
        // for _, e_b := range b {
        //     if e_a == e_b {
        //         exists = true;
        //     }
        // }
        e2 := string(rune(e));
        if !strings.Contains(b, e2) {
            result_set += e2;
        }
    }
    return result_set;
}
// or union
func set_union(a string, b string) string {
    return a + set_difference(b, a);
}

func add_known(segment string, val string, connections map[string]string, known string, unknown string) (string, string, error) {
    msg := "";
    if len(val) != 1 {
        msg = fmt.Sprintf("ERROR - Tried to add known connection '%s' to segment %s, but it is under- or over-constrained!\n", val, segment);
        fmt.Fprintf(os.Stderr, msg);
    }
    if strings.Contains(known, val) {
        msg = fmt.Sprintf("WARNING - Tried to add known connection '%s' to segment %s, but it was already known! known = %s\n", val, segment, known);
    }
    if len(msg) > 0 {
        return known, unknown, errors.New(msg);
    }

    connections[segment] = val
    return known + val, set_difference(unknown, val), nil;
}

func find_digits_and_connections(unordered_digits []string) (map[int]string, map[string]string){
    segments := "abcdefg";

    // gives 1,4,7,8 known
    digits, digits_by_length := numToPossibleSegments(unordered_digits);
    connections := init_connections(segments);

    unknown := segments;
    known := "";
    // fmt.Println("- known=", known, "- unknown=", unknown, "- connections=", connections, "- digits=", digits, "---");

    // 7 and 1 are known and only have one difference
    known, unknown, _ = add_known("a", set_difference(digits[7], digits[1]),
            connections, known, unknown);

    // 4-1 gives B & D, limiting their possible connections to 2.
    connections["b"] = set_difference(digits[4], digits[1]);
    connections["d"] = set_difference(digits[4], digits[1]);

    // which lets us get B and 5 by checking against 2, 3 and 5
    for _, str := range digits_by_length[5] {
        diff := set_difference(connections["b"], str);
        if len(diff) == 1 {
            // 2 & 3 eats D, but not B. 
            known, unknown, _ = add_known("b", diff, connections, known, unknown);
        } else if len(diff) == 0 {
            digits[5] = str; // 5 is only 5-segmented digit that eats both B and D
        } else {
            fmt.Fprintf(os.Stderr, "ERROR - Undefined behaviour for this state!");
        }
    }
    // And in turn remove B from D's possible connections
    known, unknown, _ = add_known("d", set_difference(connections["d"], connections["b"]),
            connections, known, unknown);

    // 6 and 7 are known and only have one difference
    known, unknown, _ = add_known("c", set_difference(digits[7], digits[5]),
            connections, known, unknown);
    
    // With C known, we can find F from 1
    known, unknown, _ = add_known("f", set_difference(digits[1], connections["c"]),
            connections, known, unknown);
    
    // Now the only unknown in 5, is G
    known, unknown, _ = add_known("g", set_difference(digits[5], known),
            connections, known, unknown);

    // Which leaves the last unknown as E
    known, unknown, _ = add_known("e", unknown,
    connections, known, unknown);

    // And finally I now know the digits mapping
    digits[0] = set_difference(digits[0], connections["d"]);
    // digits[1] = digits[1];
    digits[2] = set_difference(digits[2], connections["b"] + connections["f"]);
    digits[3] = set_difference(digits[3], connections["b"] + connections["e"]);
    // digits[4] = digits[4];
    // digits[5] = digits[5];
    digits[6] = set_difference(digits[6], connections["c"]);
    // digits[7] = digits[7];
    // digits[8] = digits[8];
    digits[9] = set_difference(digits[9], connections["e"]);

    // fmt.Println("- known=", known, "- unknown=", unknown, "- connections=", connections, "- digits=", digits, "---");

    return digits, connections;
}

func numToPossibleSegments(unordered_digits []string) (map[int]string, map[int][]string) {
    digits := make(map[int]string);
    by_length := make(map[int][]string);
    for _, str := range unordered_digits {
        by_length[len(str)] = append(by_length[len(str)], str);
        switch len(str) {
        case 2:
            digits[1] = str;
        case 3:
            digits[7] = str;
        case 4:
            digits[4] = str;
        case 7:
            digits[8] = str;
        case 5:
            digits[2] = set_union(digits[2], str);
            digits[3] = set_union(digits[3], str);
            digits[5] = set_union(digits[5], str);
        case 6:
            digits[0] = set_union(digits[0], str);
            digits[6] = set_union(digits[6], str);
            digits[9] = set_union(digits[9], str);
        default:
            fmt.Fprintf(os.Stderr, "ERROR - Undefined behaviour for this state!");
        }
    }
    return digits, by_length;
}
// func get_easy_mapping(warped_digits []string) map[string]int {
//     mapping := make(map[string]int);
//     for _, str := range warped_digits {
//         switch len(str) {
//         case 2:
//             mapping[str] = 1;
//         case 4:
//             mapping[str] = 4;
//         case 3:
//             mapping[str] = 7;
//         case 7:
//             mapping[str] = 8;
//         default:
//         }
//     }
//     return mapping;
// }
func inverse_mapping(mapping map[int]string) map[string]int {
    inversion := make(map[string]int);
    for k, v := range mapping {
        inversion[v]=k;
    }
    return inversion;
}
func init_connections(segments string) map[string]string {
    mapping := make(map[string]string);
    for _, c := range segments {
        mapping[string(rune(c))] = segments;
    }
    return mapping;
}