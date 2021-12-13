package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "regexp"
);

/**
  * Start - 17:33:43
  * p1 done - 18:20:57
  * p2 done - 18:26:01
  */

func main() {
    input, _ := utils.Get_input(2021, 13);
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

const separator string = "\n\n";

var part1_test_input = []string{
    `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
};
var part1_test_output = []string{
    `17`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, " \n"), separator));
    lines := strings.Split(inputs[0], "\n");
    instrs_str := strings.Split(inputs[1], "\n");

    // hor := []uint{};
    // ver := []uint{};
    // for y, line := range lines {
    //     pos_arr := utils.StrToInt_array(strings.split(line, ","));
    //     pos := Coord{x:pos_arr[0], y:pos_arr[1]};
    //     hor[pos.y] = hor[pos.y] || 1 << pos.x;
    //     ver[pos.x] = ver[pos.x] || 1 << pos.y;
    // }

    dots := map[Coord]bool{};
    // dots := []Coord{};
    for _, line := range lines {
        pos_arr, _ := utils.StrToInt_array(strings.Split(line, ","));
        pos := Coord{x:pos_arr[0], y:pos_arr[1]};
        dots[pos]=true;
    }

    re := regexp.MustCompile("fold along ([xy])=(\\d+)");
    // instrs := []Instr{};
    for _, instr := range instrs_str[0:1] {
    // instr := instrs_str[0];
        match := re.FindStringSubmatch(instr);
        hor := match[1] == "x";
        mirror, _ := strconv.Atoi(match[2]);
        // instrs = append(instrs, Instr{hor: hor, mirror:mirror});

        new_dots := map[Coord]bool{};
        // fmt.Println(pretty_format(dots));
        for pos, _ := range dots {
            offset := 0;
            new_pos := Coord{x:pos.x, y:pos.y};
            if hor {
                offset = pos.x - mirror;
                new_pos.x = mirror - offset;
            } else {
                offset = pos.y - mirror;
                new_pos.y = mirror - offset;
            }
            if offset == 0 {
                continue; // AoC didn't define this case, but seems to imply removing them
            } else if offset < 0 {
                new_dots[pos]=true;
            } else {
                new_dots[new_pos]=true;
            }
        }
        dots = new_dots;
    }

    // fmt.Println(pretty_format(dots));
    return strconv.Itoa(len(dots));
}

type Coord struct {
    x, y int;
}
type Instr struct {
    hor bool;
    mirror int;
}

func Transpose(grid []uint) []uint {
    transposed := make([]uint, 64); //uints are 64 bits long
    for i, n := range grid {
        j := 0;
        for n != 0 {
            cell := n & 1;
            n = n >> 1;
            transposed[j] = transposed[j] | cell << i;
            j++;
        }
    }
    return transposed;
}

func pretty_format(dots map[Coord]bool) string {
    width, height := 0, 0;
    for pos, _ := range dots {
        if width < pos.x {
            width = pos.x;
        }
        if height < pos.y {
            height = pos.y;
        }
    }

    grid := make([][]bool, height+1);
    for i, _ := range grid {
        grid[i] = make([]bool, width+1);
    }
    for pos, _ := range dots {
        grid[pos.y][pos.x] = true;
    }
    str := "";
    for _, row := range grid {
        for _, b := range row {
            if b {
                str+="#";
            } else {
                str+=".";
            }
        }
        str+="\n";
    }
    return str;
}

var part2_test_input = []string{
    `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
};
var part2_test_output = []string{
    `
#####
#...#
#...#
#...#
#####`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, " \n"), separator));
    lines := strings.Split(inputs[0], "\n");
    instrs_str := strings.Split(inputs[1], "\n");

    // hor := []uint{};
    // ver := []uint{};
    // for y, line := range lines {
    //     pos_arr := utils.StrToInt_array(strings.split(line, ","));
    //     pos := Coord{x:pos_arr[0], y:pos_arr[1]};
    //     hor[pos.y] = hor[pos.y] || 1 << pos.x;
    //     ver[pos.x] = ver[pos.x] || 1 << pos.y;
    // }

    dots := map[Coord]bool{};
    // dots := []Coord{};
    for _, line := range lines {
        pos_arr, _ := utils.StrToInt_array(strings.Split(line, ","));
        pos := Coord{x:pos_arr[0], y:pos_arr[1]};
        dots[pos]=true;
    }

    re := regexp.MustCompile("fold along ([xy])=(\\d+)");
    // instrs := []Instr{};
    for _, instr := range instrs_str {
    // instr := instrs_str[0];
        match := re.FindStringSubmatch(instr);
        hor := match[1] == "x";
        mirror, _ := strconv.Atoi(match[2]);
        // instrs = append(instrs, Instr{hor: hor, mirror:mirror});

        new_dots := map[Coord]bool{};
        // fmt.Println(pretty_format(dots));
        for pos, _ := range dots {
            offset := 0;
            new_pos := Coord{x:pos.x, y:pos.y};
            if hor {
                offset = pos.x - mirror;
                new_pos.x = mirror - offset;
            } else {
                offset = pos.y - mirror;
                new_pos.y = mirror - offset;
            }
            if offset == 0 {
                continue; // AoC didn't define this case, but seems to imply removing them
            } else if offset < 0 {
                new_dots[pos]=true;
            } else {
                new_dots[new_pos]=true;
            }
        }
        dots = new_dots;
        // fmt.Println(pretty_format(dots));
    }

    return "\n"+strings.Trim(pretty_format(dots),"\n");
}
