package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "regexp"
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
    // nums, _ := utils.StrToInt_array(inputs);

    segments := make([]Segment, 0);

    // re := regexp.MustCompile("(\\d+,\\d+) -> (\\d+,\\d+)");
    re := regexp.MustCompile("() -> ()");
    for _, line := range lines {
        segment_str := re.Split(line, -1);
        start, err := utils.StrToInt_array(strings.Split(segment_str[0], ","));
        if err != nil {
            fmt.Printf("error1 = %s \n", err);
        }
        end, err := utils.StrToInt_array(strings.Split(segment_str[1], ","));
        if err != nil {
            fmt.Printf("error2 = %s \n", err);
        }
        segments = append(segments, Segment{Coord{x: start[0], y: start[1]}, Coord{x: end[0], y: end[1]}});
    }

    // needs dimensions for the field
    max_x, max_y := 0, 0;
    for _, segment := range segments {
        start := segment.start;
        end := segment.end;
        if start.x > max_x {
            max_x = start.x;
        }
        if start.y > max_y {
            max_y = start.y;
        }
        if end.x > max_x {
            max_x = end.x;
        }
        if end.y > max_y {
            max_y = end.y;
        }
    }

    // prepare field
    field := make([][]int, max_y+1);
    for y:=0; y < len(field); y++{
        field[y] = make([]int, max_x+1);
        // for x:=0; x < len(field[y]); x++ {

        // }
    }

    // ignore diagonals
    horizontal := make([]Segment, 0);
    vertical := make([]Segment, 0);
    for _, segment := range segments {
        start := segment.start;
        end := segment.end;
        if (start.x == end.x) {
            horizontal = append(horizontal, segment);
        }
        if (start.y == end.y) {
            vertical = append(vertical, segment);
        }
    }
    // segments = pass;

    // increment field
    for _, segment := range horizontal {
        start := segment.start;
        end := segment.end;
        if (start.y < end.y) {
            for y:=start.y; y <= end.y; y++ {
                field[y][start.x]++;
            }
        } else {
            for y:=end.y; y <= start.y; y++ {
                field[y][start.x]++;
            }
        }
    }
    for _, segment := range vertical {
        start := segment.start;
        end := segment.end;
        if (start.x < end.x) {
            for x:=start.x; x <= end.x; x++ {
                field[start.y][x]++;
            }
        } else {
            for x:=end.x; x <= start.x; x++ {
                field[start.y][x]++;
            }
        }
    }

    // print_pretty_int(field);

    count := 0;
    for y:=0; y < len(field); y++ {
        for x:=0; x < len(field[y]); x++ {
            if field[y][x] > 1 {
                count++;
            }
        }
    }

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
    // nums, _ := utils.StrToInt_array(inputs);

    segments := make([]Segment, 0);

    // re := regexp.MustCompile("(\\d+,\\d+) -> (\\d+,\\d+)");
    re := regexp.MustCompile("() -> ()");
    for _, line := range lines {
        segment_str := re.Split(line, -1);
        start, err := utils.StrToInt_array(strings.Split(segment_str[0], ","));
        if err != nil {
            fmt.Printf("error1 = %s \n", err);
        }
        end, err := utils.StrToInt_array(strings.Split(segment_str[1], ","));
        if err != nil {
            fmt.Printf("error2 = %s \n", err);
        }
        segments = append(segments, Segment{Coord{x: start[0], y: start[1]}, Coord{x: end[0], y: end[1]}});
    }

    // needs dimensions for the field
    max_x, max_y := 0, 0;
    for _, segment := range segments {
        start := segment.start;
        end := segment.end;
        if start.x > max_x {
            max_x = start.x;
        }
        if start.y > max_y {
            max_y = start.y;
        }
        if end.x > max_x {
            max_x = end.x;
        }
        if end.y > max_y {
            max_y = end.y;
        }
    }

    // prepare field
    field := make([][]int, max_y+1);
    for y:=0; y < len(field); y++{
        field[y] = make([]int, max_x+1);
        // for x:=0; x < len(field[y]); x++ {

        // }
    }

    // ignore diagonals
    // horizontal := make([]Segment, 0);
    // vertical := make([]Segment, 0);
    // diagonal := make([]Segment, 0);
    // for _, segment := range segments {
    //     start := segment.start;
    //     end := segment.end;
    //     if (start.x == end.x) {
    //         horizontal = append(horizontal, segment);
    //     } else if (start.y == end.y) {
    //         vertical = append(vertical, segment);
    //     } else {
    //         diagonal = append(diagonal, segment);
    //     }
    // }

    // increment field
    // for _, segment := range horizontal {
    //     start := segment.start;
    //     end := segment.end;
    //     if (start.y < end.y) {
    //         for y:=start.y; y <= end.y; y++ {
    //             field[y][start.x]++;
    //         }
    //     } else {
    //         for y:=end.y; y <= start.y; y++ {
    //             field[y][start.x]++;
    //         }
    //     }
    // }
    // for _, segment := range vertical {
    //     start := segment.start;
    //     end := segment.end;
    //     if (start.x < end.x) {
    //         for x:=start.x; x <= end.x; x++ {
    //             field[start.y][x]++;
    //         }
    //     } else {
    //         for x:=end.x; x <= start.x; x++ {
    //             field[start.y][x]++;
    //         }
    //     }
    // }
    // for _, segment := range diagonal {
    //     start := segment.start;
    //     end := segment.end;
    //     steps := Abs(end.x - start.x);
    //     dir_x := Sign(end.x - start.x);
    //     dir_y := Sign(end.y - start.y);
    //     x:=start.x;
    //     y:=start.y;
    //     for i:=0; x < steps; i++ {
    //         field[y][x]++;
    //         x+=dir_x;
    //         y+=dir_y;
    //     }
    // }
    for _, segment := range segments {
        start := segment.start;
        end := segment.end;
        steps := Abs(end.x - start.x);
        if (start.y != end.y) {
            steps = Abs(end.y - start.y);
        }
        dir_x := Sign(end.x - start.x);
        dir_y := Sign(end.y - start.y);
        x:=start.x;
        y:=start.y;
        for i:=0; i <= steps; i++ {
            field[y][x]++;
            x+=dir_x;
            y+=dir_y;
        }
    }

    // print_pretty_int(field);

    count := 0;
    for y:=0; y < len(field); y++ {
        for x:=0; x < len(field[y]); x++ {
            if field[y][x] > 1 {
                count++;
            }
        }
    }

    return strconv.Itoa(count);
}

func Sign(x int) int {
    if x == 0 {
        return 0;
    }
    return x / Abs(x);
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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