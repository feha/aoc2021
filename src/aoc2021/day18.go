package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "regexp"
);

/**
  * Start - 15:51:24
  * p1 done - 21:04:00
  * p2 done - 21:14:34
  */

func main() {
    input, _ := utils.Get_input(2021, 18);
    // fmt.Printf("Input: %s \n", input);

    success := true;
    for i := range parse_test_input {
        result := parse_test(parse_test_input[i])
        if (result != parse_test_input[i]) {
            success = false;
            fmt.Printf("parse_test failed with input %s: result %s != expected %s \n",
                    parse_test_input[i],
                    result,
                    parse_test_input[i]);
            return;
        }
    }

    success = true;
    for i := range explode_test_input {
        result := explode_test(explode_test_input[i])
        if (result != explode_test_output[i]) {
            success = false;
            fmt.Printf("explode_test failed with input %s: result %s != expected %s \n",
                    explode_test_input[i],
                    result,
                    explode_test_output[i]);
            return;
        }
    }

    success = true;
    for i := range split_test_input {
        result := split_test(split_test_input[i])
        if (result != split_test_output[i]) {
            success = false;
            fmt.Printf("split_test failed with input %s: result %s != expected %s \n",
                    split_test_input[i],
                    result,
                    split_test_output[i]);
            return;
        }
    }

    success = true;
    for i := range sum_test_input {
        result := sum_test(sum_test_input[i])
        if (result != sum_test_output[i]) {
            success = false;
            fmt.Printf("sum_test failed with input %s: result %s != expected %s \n",
                    sum_test_input[i],
                    result,
                    sum_test_output[i]);
            return;
        }
    }

    success = true;
    for i := range part1_test_input {
        result := part1(part1_test_input[i])
        if (result != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    result,
                    part1_test_output[i]);
            return;
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

var parse_test_input = []string{
    `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
};
func parse_test(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    strs := []string{};
    for _, line := range lines {
        strs = append(strs, parse(line).String());
    }
    return strings.Join(strs, "\n");
}
var explode_test_input = []string{
    `[[[[[9,8],1],2],3],4]`,
    `[7,[6,[5,[4,[3,2]]]]]`,
    `[[6,[5,[4,[3,2]]]],1]`,
    `[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]`,
    `[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]`,
    // `[[[[4,0],[5,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]`, // tests singular explode
};
var explode_test_output = []string{
    `[[[[0,9],2],3],4]`,
    `[7,[6,[5,[7,0]]]]`,
    `[[6,[5,[7,0]]],3]`,
    `[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]`,
    `[[3,[2,[8,0]]],[9,[5,[7,0]]]]`,
    // `[[[[4,0],[5,4]],[[0,[7,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]`, // tests singular explode
};
func explode_test(input string) string {
    n := parse(input);
    return explode(n).String();
}
var split_test_input = []string{
    `[0,10]`,
    `[0,11]`,
    `[[[[0,7],4],[15,[0,13]]],[1,1]]`,
    `[[[[0,7],4],[[7,8],[0,13]]],[1,1]]`,
};
var split_test_output = []string{
    `[0,[5,5]]`,
    `[0,[5,6]]`,
    `[[[[0,7],4],[[7,8],[0,13]]],[1,1]]`,
    `[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]`,
};
func split_test(input string) string {
    n := parse(input);
    return split(n).String();
}

var sum_test_input = []string{
    `[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]`,
    `[1,2]
[[3,4],5]`,
    `[1,1]
[2,2]
[3,3]
[4,4]`,
    `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
    `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
    `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]`,
    `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
    `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
};
var sum_test_output = []string{
    `[[[[0,7],4],[[7,8],[6,0]]],[8,1]]`,
    `[[1,2],[[3,4],5]]`,
    `[[[[1,1],[2,2]],[3,3]],[4,4]]`,
    `[[[[3,0],[5,3]],[4,4]],[5,5]]`,
    `[[[[5,0],[7,4]],[5,5]],[6,6]]`,
    `[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]`,
    `[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]`,
    `[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]`,
};
func sum_test(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    return sum(lines).String();
}

var part1_test_input = []string{
//     `[1,2]
// [[1,2],3]
// [9,[8,7]]
// [[1,9],[8,5]]
// [[[[1,2],[3,4]],[[5,6],[7,8]]],9]
// [[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]
// [[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]`, // first example they give... but where do they give the answer?
    `[9,1]`,
    `[1,9]`,
    `[[9,1],[1,9]]`,
    `[[1,2],[[3,4],5]]`,
    `[[[[0,7],4],[[7,8],[6,0]]],[8,1]]`,
    `[[[[1,1],[2,2]],[3,3]],[4,4]]`,
    `[[[[3,0],[5,3]],[4,4]],[5,5]]`,
    `[[[[5,0],[7,4]],[5,5]],[6,6]]`,
    `[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]`,
    `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
    `[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]`,
};
var part1_test_output = []string{
    // ``, // first example they give... but where do they give the answer?
    `29`,
    `21`,
    `129`,
    `143`,
    `1384`,
    `445`,
    `791`,
    `1137`,
    `3488`,
    `4140`,
    `4140`,
};
func part1(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    result := magnitude(sum(lines));

    return strconv.Itoa(result);
}

type SingleFisher interface {
    fmt.Stringer
    GetValue() int
    GetLeaves() []SingleFisher
    Add(SingleFisher) SingleFisher
}

type Pair struct {
    fst, snd SingleFisher;
}
func (this Pair) String() string {
    return "[" + this.fst.String() + "," + this.snd.String() + "]";
}
func (this Pair) GetValue() int {
    return this.fst.GetValue() + this.snd.GetValue();
}
func (this Pair) GetLeaves() []SingleFisher {
    safe_append := []SingleFisher{};
    safe_append = append(safe_append, this.fst.GetLeaves()...);
    safe_append = append(safe_append, this.snd.GetLeaves()...);
    return safe_append;
}

//? Or should I use `type Leaf struct {fst, snd int}` instead?
type Single int;
func (this Single) String() string {
    return strconv.Itoa(int(this));
}
func (this Single) GetValue() int {
    return int(this);
}
func (this Single) GetLeaves() []SingleFisher {
    return []SingleFisher{this};
}
func (this Single) Add(rhs SingleFisher) SingleFisher {
    switch t := rhs.(type) {
    case Single:
        return Single(int(this) + int(rhs.(Single)));
    case Pair:
        // adds rightwards: n+[fst,snd] -> [n+fst, snd]
        pair := rhs.(Pair);
        pair.fst = this.Add(pair.fst);
        return pair;
    default: 
        fmt.Printf("Error! - unexpected type %T\n", t)
    }
    return Single(-1);
}
func (this Pair) Add(rhs SingleFisher) SingleFisher {
    switch t := rhs.(type) {
    case Single:
        // adds leftwards: [fst,snd]+n -> [fst, snd+n]
        this.snd = this.snd.Add(rhs);
        return this;
    case Pair:
        fmt.Printf("Error! - Behavior undefined for Pair.Add(Pair)!\n")
    default: 
        fmt.Printf("Error! - unexpected type %T\n", t)
    }
    return Single(-1);
}
// type Leaf struct {
//     fst, snd int
// }
// func (this Leaf) String() string {
//     return "[" + strconv.Itoa(this.fst) + "," + strconv.Itoa(this.snd) + "]";
// }
// func (this Leaf) GetValue() int {
//     return this.fst + this.snd;
// }
// func (this Leaf) GetLeaves() []int {
//     safe_append := []int{};
//     safe_append = append(safe_append, this.fst.GetLeaves()...);
//     safe_append = append(safe_append, this.snd.GetLeaves()...);
//     return safe_append;
// }


func find_comma(arr []string) (int, []string, []string) {
    // c := arr[0];
    arr = arr[1:len(arr)-1]; // trim current pair

    count := 0;
    for i, c := range arr {
        switch c {
        case "[":
            count++;
        case "]":
            count--;
        case ",":
            if count == 0 {
                return i, arr[:i], arr[i+1:]; // i, [("..."),("...")]
            }
        default:
        }
    }
    return -1, []string{}, []string{};
}

func parse(str string) SingleFisher {
    return subparser(strings.Split(str, ""));
}
func subparser(arr []string) SingleFisher {
    re := regexp.MustCompile("^\\d+$");
    str := re.FindString(strings.Join(arr,""));
    if str != "" {
        n, _ := strconv.Atoi(str);
        return Single(n);
    }

    _, fst, snd := find_comma(arr);
    return Pair{fst: subparser(fst), snd: subparser(snd)};
}

func explode(n SingleFisher) SingleFisher {
    n, _, _ = explode_helper(n,0);
    return n;
}
func explode_helper(n SingleFisher, i int) (SingleFisher, Single, bool) {
    pair, is_pair := n.(Pair);
    if is_pair {
        if i == 3 { // level 3
            // any pairs at level 4 - We know it must directly contain Single's,
            //  as we assume parsing valid input (questionable),
            //  and reduce after every operation.
            // We use lookaheads as we must remember parent/outer-scope,
            //  since SingleFisher's don't contain pointers to their children
            //  (which means mutating them doesn't propagate to parent).
            // Also because we must mutate the sibling by adding to it,
            //  though, again, if we used pointers it could be done with GetLeaves.
            // With a reference to root,
            //  and singles knowing their index (or otherwise tracking it),
            //  we could also mutate the carry.
            lookahead_fst, fst_is_pair := pair.fst.(Pair);  // 4,5
            lookahead_snd, snd_is_pair := pair.snd.(Pair);  // 2,6
            if fst_is_pair {
                fst := lookahead_fst.fst.(Single);          // 4
                snd := lookahead_fst.snd.(Single);          // 5
                
                pair.fst = Single(0);                        // [0,_]
                if snd_is_pair {
                    lookahead_snd.fst = snd.Add(lookahead_snd.fst); // [(7),_] = 5 + [(2),_]
                    pair.snd = lookahead_snd;                       // [0, ([7,_])] = [7,_]
                } else {
                    pair.snd = pair.snd.Add(snd);
                }
                // [0, ([7,_])], 4, true
                return pair, fst, true; // carry left
            }
            if snd_is_pair {
                fst := lookahead_snd.fst.(Single);
                snd := lookahead_snd.snd.(Single);
                // lookahead_fst can't be Pair

                pair.fst = pair.fst.Add(fst);
                pair.snd = Single(0);

                return pair, snd, false; // carry right
            }

            // unneeded row, but skips some calls and conditionals
            return n, -1, false; // can't explode a single
        } 
        
        fst, carry, was_fst := explode_helper(pair.fst, i+1);
        if fst != pair.fst {
            pair.fst = fst;
            if carry > 0 && !was_fst {
                pair.snd = carry.Add(pair.snd);
                carry = Single(-1);
            }
            return pair, carry, true;
        }
        // don't explode more than once, to let us handle the carry
        snd, carry, was_fst := explode_helper(pair.snd, i+1)
        if snd != pair.snd {
            if carry > 0 && was_fst {
                pair.fst = pair.fst.Add(carry);
                carry = Single(-1);
            }
            pair.snd = snd;
            return pair, carry, false;
        }

        return pair, -1, false; // no change (also unneeded)
    }

    // fmt.Println("Warning: Tried to explode a Single!");
    return n, -1, false; // can't explode a single
}
func split(n SingleFisher) SingleFisher {
    single, ok := n.(Single);
    if ok {
        i := int(single)
        if i > 9 {
            // split: n -> [floor(n/2), ceil(n/2)]
            return Pair{Single(i/2), Single((i+1)/2)};
        }
    } else {
        pair := n.(Pair);
        fst := split(pair.fst);
        // break if fst was split (need to explode again - before splitting snd)
        if fst != pair.fst {
            pair.fst = fst;
            return pair;
        }
        snd := split(pair.snd);
        if snd != pair.snd { // break if snd was split (need to explode again)
            pair.snd = snd;
            return pair;
        }
    }
    return n; // no change
}
func reduce(n SingleFisher) SingleFisher {
    reducing := true;
    for reducing {
        // explode works through entire tree - doesn't need to wait for split
        exploding := true;
        for exploding {
            n2 := explode(n);
            exploding = n != n2;
            n = n2;
        }

        // we need to explode anew after every split
        n2 := split(n)
        reducing = n != n2; // we are done when nothing changes.
        n = n2;
    }
    return n;
}

func add(lhs, rhs SingleFisher) SingleFisher {
    return reduce(Pair{lhs, rhs});
}
func sum(lines []string) SingleFisher {
    var ret SingleFisher;
    for _, line := range lines {
        n := parse(line);
        if ret == nil {
            ret = n;
        } else {
            ret = add(ret, n);
        }
    }

    return ret;
}

func magnitude(n SingleFisher) int {
    switch t := n.(type) {
    case Pair:
        pair := n.(Pair)
        return 3*magnitude(pair.fst) + 2*magnitude(pair.snd);
    case Single:
        return int(n.(Single));
    default: 
        fmt.Printf("unexpected type %T", t)
    }
    return -1;
}

var part2_test_input = []string{
    `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
};
var part2_test_output = []string{
    `3993`,
};
func part2(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    result := 0;
    for i:=0; i < len(lines)-1; i++ {
        for j:=i+1; j < len(lines); j++ {
            magn := magnitude(sum( []string{lines[i], lines[j]} ));
            if result < magn {
                result = magn;
            }
        }
    }
    for i:=len(lines)-1; i >= 1; i-- {
        for j:=i-1; j >= 0; j-- {
            magn := magnitude(sum( []string{lines[i], lines[j]} ));
            if result < magn {
                result = magn;
            }
        }
    }
    return strconv.Itoa(result);
}
