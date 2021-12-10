package test;

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
  * p2 done - 17:21:44
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
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    segments := Sort_str("abcdefg");
    real_digits_strs := Get_real_digits();
    result := Part2_solution2(inputs, real_digits_strs, segments)

    return strconv.Itoa(result);
}

func Part2_solution1(lines []string, real_digits_strs []string, segments string) int {
    result := 0;
    for _, line := range lines {
        notes := strings.Split(line, " | ");
        unordered_digits := strings.Split(notes[0], " ");
        warped_output := strings.Split(notes[1], " ");

        warped_output_sorted := make([]string, len(warped_output));
        for i, str := range warped_output {
            warped_output_sorted[i] = Sort_str(str);
        }

        digits, _ := find_digits_and_connections(unordered_digits, segments);

        digits_sorted := make(map[int]string);
        for i, digit := range digits {
            digits_sorted[i] = Sort_str(digit);
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
    return result;
}

type Seg7 struct {
    a, b, c, d, e, f, g bool;
    to_str string;
}
func Seg7_Init(seg7 Seg7, segments string) (Seg7, error) {
    seg7.to_str = Sort_str(strings.ToLower(segments));
    for _, c := range strings.Split(strings.ToLower(segments), "") {
        // sensible struct-indexing using strings require the package 'reflect',
        // so let's just hardcode...
        switch c {
        case "a":
            seg7.a = true;
        case "b":
            seg7.b = true;
        case "c":
            seg7.c = true;
        case "d":
            seg7.d = true;
        case "e":
            seg7.e = true;
        case "f":
            seg7.f = true;
        case "g":
            seg7.g = true;
        default:
            msg := fmt.Sprintf("ERROR - Can't init Seg7 with segments %s! Segment '%c' is not a valid segment\n", segments, c);
            fmt.Fprintf(os.Stderr, msg);
            return seg7, errors.New(msg);
        }
    }
    return seg7, nil;
}
func CreateSeg7(segments string) (Seg7, error) {
    return Seg7_Init(Seg7{}, segments);
}
func Seg7_len(seg7 Seg7) int {
    return len(seg7.to_str);
}
func Seg7_has(seg7 Seg7, segment string) bool {
    for _, segment2 := range strings.Split(seg7.to_str, "") {
        if segment == segment2 {
            return true;
        }
    }
    return false;
}

func array_remove(arr []Seg7, e Seg7) []Seg7 {
    if len(arr) == 1 && arr[0] == e {
        return []Seg7{};
    }
    for i, e2 := range arr {
        if e == e2 {
            ret := make([]Seg7, 0);
            ret = append(ret, arr[:i]...);
            return append(ret, arr[i+1:]...);
        }
    }
    return arr;
    // if len(arr) == 1 && arr[0] == e {
    //     return []Seg7{};
    // }
    // for i, e2 := range arr {
    //     if (e == e2) {
    //         arr[i] = arr[len[arr]-1];
    //         return arr[:len[arr]-1];
    //     }
    // }
}

type Context struct {
    known map[Seg7]Seg7;
    unknown map[Seg7]bool;
    hist_real, hist_real_known map[int][]Seg7;
    hist_scrambled, hist_scrambled_known map[int][]Seg7;
}
func NewContext() Context {
    return Context{
        known: map[Seg7]Seg7{},
        unknown: map[Seg7]bool{},
        hist_real: map[int][]Seg7{},
        hist_real_known: map[int][]Seg7{},
        hist_scrambled: map[int][]Seg7{},
        hist_scrambled_known: map[int][]Seg7{},
    };
}

func add_known_seg7(seg7 Seg7, real_seg7 Seg7, context Context) error {
    if _, ok := context.known[seg7]; ok {
        msg := fmt.Sprintf("WARNING - Tried to add known Seg7 '%v'. known_seg7='%v'\n", seg7, context.known);
        fmt.Fprintf(os.Stderr, msg);
        return errors.New(msg);
    }
    context.known[seg7] = real_seg7;
    delete(context.unknown, seg7);
    l := Seg7_len(seg7); // == Seg7_len(real_seg7)
    if len(context.hist_real[l]) == 1 {
        delete(context.hist_real, l);
        delete(context.hist_scrambled, l);
    } else {
        context.hist_real[l] = array_remove(context.hist_real[l], real_seg7);
        context.hist_scrambled[l] = array_remove(context.hist_scrambled[l], seg7);
    }
    context.hist_real_known[l] = append(context.hist_real_known[l], real_seg7);
    context.hist_scrambled_known[l] = append(context.hist_scrambled_known[l], seg7);
    return nil;
}
func contains(super Seg7, sub Seg7)bool {
    for _, segment := range strings.Split(sub.to_str, "") {
        if !Seg7_has(super, segment) {
            return false;
        }
    }
    return true;
}
func match_unique_lengths(context Context) bool {
    success := false;
    for l, segs := range context.hist_scrambled {
        if len(segs) == 1 {
            err := add_known_seg7(segs[0], context.hist_real[l][0], context);
            if err == nil {
                success = true;
            }
        }
    }
    return success;
}
func match_unique_sub_seg7s(context Context) bool {
    success := false;
    for known, known_real := range context.known {
        for l, seg7s := range context.hist_scrambled {
            subsets := 0;
            supersets := 0;
            candidate_sub := known;
            candidate_super := known;
            for _, unknown := range seg7s {
                if contains(known, unknown) {
                    subsets++;
                    candidate_sub = unknown;
                } else if contains(unknown, known) {
                    supersets++;
                    candidate_super = unknown;
                }
            }
            if subsets == 1 {
                for _, unknown_real := range context.hist_real[l] {
                    if contains(known_real, unknown_real) {
                        err := add_known_seg7(candidate_sub, unknown_real, context);
                        if err == nil {
                            success = true;
                        }
                    }
                }
            }
            if supersets == 1 {
                for _, unknown_real := range context.hist_real[l] {
                    if contains(unknown_real, known_real) {
                        err := add_known_seg7(candidate_super, unknown_real, context);
                        if err == nil {
                            success = true;
                        }
                    }
                }
            }
        }
    }
    return success;
}

func Part2_solution2(lines []string, real_digits_strs []string, segments string) int {
    real_digits := map[Seg7]int{};
    for i, str := range real_digits_strs {
        seg7, err := CreateSeg7(str);
        if err != nil {
            fmt.Printf("error = %s \n", err);
        }
        real_digits[seg7]=i;
    }
    
    // foo, _ := CreateSeg7("abcefg");
    // bar, _ := CreateSeg7("abcefg");
    // foobar, _ := CreateSeg7("abc");
    // fmt.Println("=====", foo == bar, foo == foobar);

    result := 0;

    for _, line := range lines {
        notes := strings.Split(line, " | ");
        lhs := strings.Split(notes[0], " ");
        rhs := strings.Split(notes[1], " ");

        context := NewContext();

        for _, scrambled_seg7_str := range lhs {
            scrambled_seg7, err := CreateSeg7(scrambled_seg7_str);
            context.unknown[scrambled_seg7] = true;
            if err != nil {
                fmt.Printf("error = %s \n", err);
            }
        }

        // Find unknown real_digits with unique length and match to unknown seg7
        for seg7 := range real_digits {
            context.hist_real[Seg7_len(seg7)] = append(context.hist_real[Seg7_len(seg7)], seg7);
        }
        for seg7 := range context.unknown {
            context.hist_scrambled[Seg7_len(seg7)] = append(context.hist_scrambled[Seg7_len(seg7)], seg7);
        }

        last_len := 0;
        for len(context.unknown) > 0 && len(context.unknown) != last_len {
            last_len = len(context.unknown); // infinite loop protection

            match_unique_lengths(context);
            match_unique_sub_seg7s(context);

            // a := match_unique_lengths(context);
            // b := match_unique_sub_seg7s(context);
            // // possible additional operations:
            // // - find unknown seg7 that differ by one from a known seg7 (identifies a specific segment)
            // if a || b {
            //     fmt.Println("operations: ", a, b)
            // }
        }
        if len(context.unknown) > 0 {
            msg := fmt.Sprintf("---------------------------\nERROR - There are unknown segments remaining! unknown=%v\n\nknown = %v\n\nlen=%d\nreal_digits=%v\n---------------------------\n", context.unknown, context.known, len(real_digits), real_digits);
            fmt.Fprintf(os.Stderr, msg);
            continue;
        }
        
        new_rhs := 0;
        for _, scrambled_seg7_str := range rhs {
            scrambled_seg7, err := CreateSeg7(scrambled_seg7_str);
            if err != nil {
                fmt.Printf("error = %s \n", err);
            }
            d := real_digits[context.known[scrambled_seg7]];
            new_rhs = 10 * new_rhs + d;
        }
        result += new_rhs;
    }

    return result;
}

func Get_real_digits() []string {
    return []string {
        "abcefg",
        "cf",
        "acdeg",
        "acdfg",
        "bcdf",
        "abdfg",
        "abdefg",
        "acf",
        "abcdefg",
        "abcdfg",
    }
}

func Part2_dads_solution(lines []string, real_digits_strs []string, segments string) int {
    result := 0;

    real_digits := map[string]int{};
    for i, str := range real_digits_strs {
        real_digits[str]=i;
    }

    permutations := generate_permutations(segments);
    possible_translations := make([]map[string]string, 0);
    for _, str := range permutations {
        translation := make(map[string]string);
        for i, c := range str {
            segment := string(rune(segments[i]));
            // equivalent
            // translation[segment] = string(rune(c)); // segment -> many
            translation[string(rune(c))] = segment; // many -> segment
        }
        possible_translations = append(possible_translations, translation);
    }

    for _, line := range lines {
        notes := strings.Split(line, " | ");
        lhs := strings.Split(notes[0], " ");
        rhs := strings.Split(notes[1], " ");

        var valid_translation map[string]string;
        for _, transl := range possible_translations {
            all := true;
            for _, warped_digit := range lhs {
                translated_digit := "";
                for _, c := range warped_digit {
                    translated_digit += transl[string(rune(c))];
                }
                _, ok := real_digits[Sort_str(translated_digit)];
                all = all && ok;
            }
            if all {
                valid_translation = transl;
                break;
            }
        }

        new_rhs := 0;
        for _, warped_digit := range rhs {
            digit := "";
            for _, c := range warped_digit {
                digit += valid_translation[string(rune(c))];
            }
            new_rhs = 10 * new_rhs + real_digits[Sort_str(digit)];
        }
        result += new_rhs;
    }

    return result;
}

func generate_permutations(set string) []string {
    l := len(set);
    if l == 0 {
        return []string{};
    } else if l == 1 {
        return []string{set};
    }

    permutations := make([]string, 0);
    for _, e := range strings.Split(set, "") {
        sub := generate_permutations(set_difference(set, e));
        for _, p := range sub {
            permutations = append(permutations, e + p);
        }
    }
    return permutations;
}

func Sort_str(str string) string {
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

func find_digits_and_connections(unordered_digits []string, segments string) (map[int]string, map[string]string){

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