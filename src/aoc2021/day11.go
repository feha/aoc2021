package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 18:58:54
  * p1 done - 19:41:49
  * p2 done - 19:45:46
  */

func main() {
    input, _ := utils.Get_input(2021, 11);
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
    `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
};
var part1_test_output = []string{
    `1656`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    grid := [][]int{};
    for _, line := range inputs {
        row, _ := utils.StrToInt_array(strings.Split(line,""));
        grid = append(grid, row);
    }
    
    result := 0;

    for i:=0; i < 100; i++ {
        result += step(grid);
    }

    return strconv.Itoa(result);
}

func flash(pos Coord, grid [][]int, flashed map[Coord]bool) {
    if grid[pos.y][pos.x] > 9 {
        flashed[pos]=true;
        grid[pos.y][pos.x] = 0;

        width, height := len(grid), len(grid[0]);

        adjs := get_neighbours(pos, width, height);
        for _, adj := range adjs {
            if _, ok := flashed[adj]; !ok {
                grid[adj.y][adj.x]++;
                flash(adj, grid, flashed);
            }
        }
    }
}
func step(grid [][]int) int {
    flashed := map[Coord]bool{};
    for y, row := range grid {
        for x, cell := range row {
            grid[y][x] = cell+1;
        }
    }
    for y, row := range grid {
        for x, _ := range row {
            flash(Coord{x:x, y:y}, grid, flashed);
        }
    }
    return len(flashed);
}

type Coord struct {
    x, y int;
}
func Coord_add(lhs Coord, rhs Coord) Coord {
    return Coord{x: lhs.x + rhs.x, y: lhs.y + rhs.y};
}
func Coord_sub(lhs Coord, rhs Coord) Coord {
    return Coord{x: lhs.x - rhs.x, y: lhs.y - rhs.y};
}
func get_neighbours(pos Coord, width int, height int) []Coord {
    neighbours := []Coord{};
    for x:=-1; x < 2; x++ {
        for y:=-1; y < 2; y++ {
            dir := Coord{x:x, y:y};
            pos2 := Coord_add(pos, dir);
            if pos == pos2 {
                continue;
            }
            // // ignore diagonals
            // if dir.x != 0 && dir.y != 0 {
            //     continue;
            // }
            if 0 <= pos2.x && pos2.x < width {
                if 0 <= pos2.y && pos2.y < height {
                    neighbours = append(neighbours, pos2);
                }
            }
        }
    }
    return neighbours;
}

func format_pretty_int(arrs [][]int) string {
    str := "";

    for _, row := range arrs {
            str += strings.Join(utils.Itoa_array(row), "");
        str += "\n";
    }

    return str;
}

var part2_test_input = []string{
    `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
};
var part2_test_output = []string{
    `195`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    grid := [][]int{};
    for _, line := range inputs {
        row, _ := utils.StrToInt_array(strings.Split(line,""));
        grid = append(grid, row);
    }
    
    result := 0;

    i:=0;
    for sum(grid) != 0 {
        result += step(grid);
        i++;
    }
    result = i;

    return strconv.Itoa(result);
}

func sum(grid [][]int) int {
    sum := 0;
    for _, row := range grid {
        for _, cell := range row {
            sum += cell;
        }
    }
    return sum;
}