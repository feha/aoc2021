package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 21:48:57
  * p1 done - 22:09:12
  * p2 done - 22:48:22
  */

func main() {
    input, _ := utils.Get_input(2021, 14);
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

var part1_test_input = []string{
    `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
};
var part1_test_output = []string{
    `1588`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, " \n"), "\n\n"));

    template := inputs[0];
    rule_strs := strings.Split(inputs[1], "\n");

    rules := map[string]string{};
    for _, rule := range rule_strs {
        pair := strings.Split(rule, " -> ");
        lhs := pair[0];
        rhs := pair[1];
        rules[lhs] = rhs;
    }

    for i:=0; i < 10; i++ {
        template = step(template, rules);
    }
    template_asarr := strings.Split(template, "");
    hist := map[string]uint{};
    for _, c := range template_asarr {
        hist[c]++;
    }

    mce, lce := uint(0), uint(0);
    for _, count := range hist {
        lce = count;
        break;
    }
    for _, count := range hist {
        if mce < count {
            mce = count;
        }
        if lce > count {
            lce = count;
        }
    }

    result := mce - lce;
    return strconv.Itoa(int(result));
}

func step(template string, rules map[string]string) string {
    template_asarr := strings.Split(template, "");
    result := "";
    right := "";
    for i:=0; i < len(template_asarr)-1; i++ {
        left := template_asarr[i];
        right = template_asarr[i+1];
        add := rules[left + right];
        result += left + add;
    }
    result += right;
    // result += template_asarr[len(template_asarr)-1];
    return result;
}

var part2_test_input = []string{
    `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
};
var part2_test_output = []string{
    `2188189693529`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, " \n"), "\n\n"));

    template := inputs[0];
    rule_strs := strings.Split(inputs[1], "\n");

    rules := map[string]string{};
    for _, rule := range rule_strs {
        pair := strings.Split(rule, " -> ");
        lhs := pair[0];
        rhs := pair[1];
        rules[lhs] = rhs;
    }

    steps := 40;

    // memoization := map[Memo]string{};
    // template_asarr := strings.Split(template, "");
    // str := "";
    // for i:=0; i < len(template_asarr)-1; i++ {
    //     left := template_asarr[i];
    //     right := template_asarr[i+1];
    //     pair := left + right;
    //     add := recurse(pair, rules, steps, memoization);
    //     str += left + add;
    // }
    // str += template_asarr[len(template_asarr)-1];

    // hist := map[string]uint{};
    // for _, c := range strings.Split(str, "") {
    //     hist[c]++;
    // }

    memoization := map[Memo]map[string]uint{};
    template_asarr := strings.Split(template, "");
    hist := map[string]uint{};
    for i:=0; i < len(template_asarr)-1; i++ {
        left := template_asarr[i];
        right := template_asarr[i+1];
        hist[left]++;
        pair := left + right;
        sub_hist := recurse(pair, rules, steps, memoization);
        for k,v := range sub_hist {
            hist[k]+=v;
        }
    }
    hist[template_asarr[len(template_asarr)-1]]++;

    mce, lce := uint(0), uint(0);
    for _, count := range hist {
        lce = count;
        break;
    }
    for _, count := range hist {
        if mce < count {
            mce = count;
        }
        if lce > count {
            lce = count;
        }
    }

    // fmt.Println(template, str, mce, lce);
    result := mce - lce;
    return strconv.Itoa(int(result));
}

type Memo struct {
    pair string;
    steps int;
}

func recurse(pair string, rules map[string]string, steps int, memoization map[Memo]map[string]uint) map[string]uint {
    if steps == 0 {
        return map[string]uint{};
    }

    memo := Memo{pair: pair, steps: steps};
    if m, ok := memoization[memo]; ok {
        return m;
    }

    asarr := strings.Split(pair, "");
    left := asarr[0];
    right := asarr[1];
    add := rules[pair];

    hist := map[string]uint{};
    hist[add]++;

    sub_hist1 := recurse(left + add, rules, steps-1, memoization);
    sub_hist2 := recurse(add + right, rules, steps-1, memoization);
    for k,v := range sub_hist1 {
        hist[k]+=v;
    }
    for k,v := range sub_hist2 {
        hist[k]+=v;
    }
    // result := recurse(left + add, rules, steps-1, memoization) + add + recurse(add + right, rules, steps-1, memoization);
    memoization[memo] = hist;
    // return result;
    return hist;
}