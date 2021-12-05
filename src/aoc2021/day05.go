package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 20:43:09
  * p1 done - 21:28:02
  * p2 done - 21:47:28
  */

func main() {
    input, _ := utils.Get_input(2021, 05);
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
    `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
};
var part1_test_output = []string{
    `5`,
};
func part1(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    segments := parse(lines);

    width, height := get_dimensions(segments);

    // prepare field
    field := make([][]int, height);
    for y := range field {
        field[y] = make([]int, width);
    }

    // ignore diagonals
    pass := make([]Segment, 0);
    for _, segment := range segments {
        start := segment.start;
        end := segment.end;
        if start.x != end.x && start.y != end.y {
            continue;
        }
        pass = append(pass, segment);
    }
    segments = pass;

    for _, segment := range segments {
        add_segment(field, segment);
    }

    count := Count(field, 1);

    return strconv.Itoa(count);
}

var part2_test_input = []string{
    `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
};
var part2_test_output = []string{
    `12`,
};
func part2(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    segments := parse(lines);

    width, height := get_dimensions(segments);

    // prepare field
    field := make([][]int, height);
    for y := range field {
        field[y] = make([]int, width);
    }

    for _, segment := range segments {
        add_segment(field, segment);
    }

    count := Count(field, 1);

    return strconv.Itoa(count);
}

func parse(lines []string) []Segment {
    segments := make([]Segment, 0);
    for _, line := range lines {
        coords := strings.Split(line, " -> ");
        start, err := utils.StrToInt_array(strings.Split(coords[0], ","));
        if err != nil {
            fmt.Printf("error1 = %s \n", err);
        }
        end, err := utils.StrToInt_array(strings.Split(coords[1], ","));
        if err != nil {
            fmt.Printf("error2 = %s \n", err);
        }
        segments = append(segments, Segment{
            Coord{x: start[0], y: start[1]},
            Coord{x: end[0], y: end[1]},
        });
    }
    return segments;
}

func get_dimensions(segments []Segment) (int, int) {
    width, height := 0, 0;
    for _, segment := range segments {
        start := segment.start;
        end := segment.end;
        if start.x > width {
            width = start.x;
        }
        if start.y > height {
            height = start.y;
        }
        if end.x > width {
            width = end.x;
        }
        if end.y > height {
            height = end.y;
        }
    }
    return width+1, height+1;
}

func add_segment(field [][]int, segment Segment) {
    start := segment.start;
    end := segment.end;
    steps := utils.Abs(end.x - start.x);
    if start.y != end.y {
        steps = utils.Abs(end.y - start.y);
    }
    dir_x := utils.Sign(end.x - start.x);
    dir_y := utils.Sign(end.y - start.y);
    x:=start.x;
    y:=start.y;
    for i:=0; i <= steps; i++ {
        field[y][x]++;
        x+=dir_x;
        y+=dir_y;
    }
}

func Count(field [][]int, threshold int) int {
    count := 0;
    for y:=0; y < len(field); y++ {
        for x:=0; x < len(field[y]); x++ {
            if field[y][x] > threshold {
                count++;
            }
        }
    }
    return count;
}


func print_pretty_int(arr [][]int) string {
    str := "";

    for _, row := range arr {
        str += strings.Join(utils.Itoa_array(row), "");
        str += "\n";
    }

    fmt.Println(str);
    return str;
}

type Coord struct {
    x int;
    y int;
}
type Segment struct {
    start Coord;
    end Coord;
}