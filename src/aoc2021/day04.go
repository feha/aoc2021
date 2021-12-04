package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "regexp"
    "strconv"
);

/**
  * Start - 15:25:43
  * p1 done - 16:46:31
  * p2 done - 16:51:16
  */

func main() {
    input, _ := utils.Get_input(2021, 04);
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
    `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
};
var part1_test_output = []string{
    `4512`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    nums,  err := utils.StrToInt_array(strings.Split(inputs[0], ","));
    if err != nil {
        fmt.Printf("error1 = %s \n", err);
    }
    boards_str := inputs[1:];

    re := regexp.MustCompile("\\s+")
    boards := make([][][]int, 0);
    boards_marks := make([][][]bool, 0);
    for i, str := range boards_str {
        lines := utils.Trim_array(strings.Split(str, "\n"));
        height := len(lines);

        boards = append(boards, make([][]int, height));
        boards_marks = append(boards_marks, make([][]bool, height));
        for y, line := range lines {
            nums, err := utils.StrToInt_array(re.Split(line, -1));
            if err != nil {
                fmt.Printf("error2 = %s \n", err);
            }
            boards[i][y] = nums;
            boards_marks[i][y] = make([]bool, len(nums));
        }
    }

    for _, n := range nums {
        for i, board := range boards {
            board_marks := boards_marks[i];
            changed, _ , _ := markValue(n, board, board_marks);
            if changed {
                won := boardWin(board_marks);
                if won {
                    score := n * boardScore(board, board_marks);
                    return strconv.Itoa(score);
                }
            }
        }
        // fmt.Println("n=", n);
        // print_pretty_int(boards);
        // print_pretty_bool(boards_marks);
    }
    // ...

    // return "";
    return strconv.Itoa(-1);
}

var part2_test_input = []string{
    `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
};
var part2_test_output = []string{
    `1924`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    nums,  err := utils.StrToInt_array(strings.Split(inputs[0], ","));
    if err != nil {
        fmt.Printf("error1 = %s \n", err);
    }
    boards_str := inputs[1:];

    re := regexp.MustCompile("\\s+")
    boards := make([][][]int, 0);
    boards_marks := make([][][]bool, 0);
    for i, str := range boards_str {
        lines := utils.Trim_array(strings.Split(str, "\n"));
        height := len(lines);

        boards = append(boards, make([][]int, height));
        boards_marks = append(boards_marks, make([][]bool, height));
        for y, line := range lines {
            nums, err := utils.StrToInt_array(re.Split(line, -1));
            if err != nil {
                fmt.Printf("error2 = %s \n", err);
            }
            boards[i][y] = nums;
            boards_marks[i][y] = make([]bool, len(nums));
        }
    }

    score := -1;
    blacklist := make([]bool, len(boards));
    for _, n := range nums {
        for i, board := range boards {
            if blacklist[i] {
                continue;
            }
            
            board_marks := boards_marks[i];
            changed, _ , _ := markValue(n, board, board_marks);
            if changed {
                won := boardWin(board_marks);
                if won {
                    score = n * boardScore(board, board_marks);
                    blacklist[i] = true;
                }
            }
        }
        // fmt.Println("n=", n);
        // print_pretty_int(boards);
        // print_pretty_bool(boards_marks);
    }
    // ...

    // return "";
    return strconv.Itoa(score);
}

func print_pretty_int(arrs [][][]int) string {
    str := "";

    r1 := arrs[0];
    for y, _ := range r1 {
        for _, arr := range arrs {
            str += strings.Join(utils.Itoa_array(arr[y]), ",") + " | ";
        }
        str += "\n";
    }

    fmt.Println(str);
    return str;
}
func print_pretty_bool(arrs [][][]bool) string {
    str := "";

    r1 := arrs[0];
    for y, _ := range r1 {
        for _, arr := range arrs {
            str += strings.Join(utils.BoolToStr_array(arr[y]), ",") + " | ";
        }
        str += "\n";
    }

    fmt.Println(str);
    return str;
}

func boardScore(board [][]int, board_marks [][]bool) int {
    sum := 0;
    for y, row := range board {
        for x, cell := range row {
            if board_marks[y][x] == false {
                sum += cell;
            }
        }
    }
    return sum;
}

func markValue(val int, board [][]int, board_marks [][]bool) (bool, [][]int, [][]bool) {
    change := false;
    for y, row := range board {
        for x, cell := range row {
            if cell == val {
                board_marks[y][x] = true;
                change = true;
            }
        }
    }
    return change, board, board_marks;
}

func boardWin(board_marks [][]bool) bool {
    return rowWin(board_marks) || colWin(board_marks);
}
func rowWin(board_marks [][]bool) bool {
    win := false;
    for _, row := range board_marks {
        win_row := true;
        for _, cell := range row {
            win_row = win_row && cell;
        }
        win = win || win_row;
    }
    return win;
}
func colWin(board_marks [][]bool) bool {
    win := false;
    for x, _ := range board_marks[0] {
        win_col := true;
        for y, _ := range board_marks {
            win_col = win_col && board_marks[y][x];
        }
        win = win || win_col;
    }
    return win;
}