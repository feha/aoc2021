package test;

import (
    "aoc/libs/utils"
    "aoc/src/aoc2021/test"
    "strings"
    "strconv"
    "testing"
    "sort"
    "math/rand"
);

const separator string = "\n";
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
func TestSolutionDad(t *testing.T) {
    segments := test.Sort_str("abcdefg");
    real_digits_strs := test.Get_real_digits();

    for i := range part2_test_input {
        inputs := utils.Trim_array(strings.Split(strings.Trim(part2_test_input[i], separator), separator));
        result := strconv.Itoa(test.Part2_dads_solution(inputs, real_digits_strs, segments))
        if (result != part2_test_output[i]) {
            t.Errorf("Test [%d] - failed with input %s:\nresult %s != expected %s \n",
                    i,
                    part2_test_input[i],
                    result,
                    part2_test_output[i]);
        }
    }
}

func TestSolution2(t *testing.T) {

    num_tests := 10;
    test_size := 10;
    
    segments := test.Sort_str("abcdefg");
    num_digits := 10;

    p_map := map[string]bool{};
    all_perms(segments, p_map);
    permutations := []string{};
    for p, _ := range p_map {
        permutations = append(permutations, p);
    }

    nofilter := func(_ string)bool {return true;};
    for i:=0; i < num_tests; i++ {
        // Generate input
        real_digits_strs := generate_digits(num_digits, permutations, nofilter);
        // t.Logf("afsasfasf%v",real_digits_strs);
        
        lines := []string{};
        for i:=0; i < test_size; i++ {
            hist := map[int]int{};
            for _, s := range real_digits_strs {
                hist[len(s)]++;
            }
            fun := func(p string)bool {
                if hist[len(p)] <= 0 {
                    return false;
                }
                hist[len(p)]--;
                return true;
            }
            lhs := generate_digits(num_digits, permutations, fun);
            rhs := strings.Join(generate_digits(4, lhs, nofilter), " ");
            lines = append(lines, strings.Join(lhs, " ") + " | " + rhs);
        }

        // test
        truth := test.Part2_dads_solution(lines, real_digits_strs, segments);
        result := test.Part2_solution2(lines, real_digits_strs, segments);
        if (result != truth) {
            t.Errorf("Failed with input %s:\nresult %d != expected %d \n",
                    strings.Join(lines,"\n"),
                    result,
                    truth);
        } else {
            t.Logf("\n____________\n| Success! |\n‾‾‾‾‾‾‾‾‾‾‾‾\n");
        }
    }
}

func generate_digits(num_digits int, permutations []string, filter func(string)bool) []string {
    result := []string{};
    generated := map[int]bool{};
    for i:=0; i < num_digits; i++ {
        n := rand.Int63n(int64(len(permutations)));
        if _, ok := generated[int(n)]; ok {
            i--;
            continue
        }
        p := permutations[n];
        if !filter(p) {
            i--;
            continue
        }
        generated[int(n)]=true
        result = append(result, p);
    }
    return result;
}

var input_TestPermutation = []string{
    `ab`,
    `abc`,
    `abcd`,
    `abcde`,
    `abcdef`,
    `abcdefg`,
    `abcdefgh`,
    `abcdefghi`,
    `abcdefghij`,
};
var output_TestPermutation = []string{
    `ab,b,a`,
    `abc,ab,ac,bc,a,b,c`,
    `abcd,abc,abd,acd,bcd,ab,ac,bc,ad,bd,cd,a,b,c,d`,
};
func TestPermutation(t *testing.T) {
    var output_TestPermutation_count = make([]int, len(input_TestPermutation));

    for i, input := range input_TestPermutation {
        output_TestPermutation_count[i] = count_perms(input);

        var output []string;
        if i < len(output_TestPermutation) {
            output = strings.Split(output_TestPermutation[i], ",");
            sort.Strings(output);
        }
        
        permutations := map[string]bool{};
        all_perms(input, permutations);
        result := []string{};
        for p, _ := range permutations {
            result = append(result, p);
        }
        sort.Strings(result);
        count := len(result);
        if (count != output_TestPermutation_count[i]) {
            t.Errorf("Test [%d] - Wrong length! Failed with input %s:\ncount %d != expected %d\nresult %s\n",
                    i,
                    input,
                    count,
                    output_TestPermutation_count[i],
                    result);
        }
        if i < len(output_TestPermutation) &&
                strings.Join(result, ",") != strings.Join(output, ",") {
            t.Errorf("Test [%d] - Failed with input %s:\nresult %s != expected %s \n",
                    i,
                    input,
                    result,
                    output_TestPermutation[i]);
        }
    }
}

func all_perms(elements string, res map[string]bool) {
    if _, ok := res[elements]; ok || len(elements) == 0 {
        return;
    }

    res[elements] = true;
    arr := strings.Split(elements, "");
    for i, _ := range arr {
        guaranteed_new_arr := []string{};
        guaranteed_new_arr = append(guaranteed_new_arr, arr[:i]...);
        guaranteed_new_arr = append(guaranteed_new_arr, arr[1+i:]...);
        all_perms(strings.Join(guaranteed_new_arr, ""), res);
    }
}
func count_perms(elements string) int {
    n := len(elements);
    sum := 0;
    for k := range elements {
        sum += factorial(n) / (factorial(k) * factorial(n - k));
    }
    return sum;
}
func factorial(n int) int {
    product := 1;
    for i:=2; i <= n; i++ {
        product *= i;
    }
    return product;
}