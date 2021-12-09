package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "errors"
    "os"
    "sort"
);

/**
  * Start - 16:10:40
  * p1 done - 16:43:38
  * p2 done - 17:39:42
  */

func main() {
    input, _ := utils.Get_input(2021, 9);
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
    `2199943210
3987894921
9856789892
8767896789
9899965678`,
};
var part1_test_output = []string{
    `15`,
};
func part1(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    grid := parse_grid(lines);
    lows := get_lows(grid);

    result := 0;
    for _, pos := range lows {
        result += 1 + grid[pos.y][pos.x];
    }

    return strconv.Itoa(result);
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
            // ignore diagonals
            if dir.x != 0 && dir.y != 0 {
                continue;
            }
            if 0 <= pos2.x && pos2.x < width {
                if 0 <= pos2.y && pos2.y < height {
                    neighbours = append(neighbours, pos2);
                }
            }
        }
    }
    return neighbours;
}
func get_lows(grid [][]int) []Coord {
    height := len(grid);
    width := len(grid[0]);
    lows := []Coord{};
    for y, row := range grid {
        for x, cell := range row {
            pos := Coord{x:x, y:y};
            neighbours := get_neighbours(pos, width, height);
            low := true;
            for _, neighbour := range neighbours {
                cell_neighbour := grid[neighbour.y][neighbour.x];
                delta := cell_neighbour - cell;
                if delta <= 0 { // apparently low-points are allowed to have ones next to eachother
                    low  = false;
                    break;
                }
            }
            if low {
                lows = append(lows, Coord{x:x, y:y});
            }
        }
    }
    return lows;
}
func parse_grid(lines []string) [][]int {
    grid := [][]int{};
    for _, line := range lines {
        cells := strings.Split(line, "");
        nums, err := utils.StrToInt_array(cells);
        if err != nil {
            fmt.Printf("error = %s \n", err);
        }
        grid = append(grid, nums);
    }
    return grid;
}

var part2_test_input = []string{
    `2199943210
3987894921
9856789892
8767896789
9899965678`,
};
var part2_test_output = []string{
    `1134`,
};
func part2(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    grid := parse_grid(lines);
    lows := get_lows(grid);
    
    height := len(grid);
    width := len(grid[0]);

    basins := []Basin{};
    for _, pos := range lows {
        basin := Basin{pos};
        blacklist := map[Coord]bool{pos: true};
        candidates := get_neighbours(pos, width, height);
        set_add_collection(blacklist, candidates);

        for len(candidates) > 0 {
            candidates_redeclared, candidate, _ := array_pop(candidates);
            candidates = candidates_redeclared;
            cell := grid[candidate.y][candidate.x];

            // Assuming all basins are terminated by cells of magnitude 9
            // No basins whose contour is delimeted by the gradient contour
            // (aka basins with multiple lows, or at least that they are counted as part of eachothers basin)
            if cell == 9 {
                continue;
            }
            basin = append(basin, candidate);

            neighbours := get_neighbours(candidate, width, height);
            neighbours = set_difference(neighbours, blacklist);
            set_add_collection(blacklist, neighbours);
            candidates = append(candidates, neighbours...);
            // blacklist = set_union(blacklist, neighbours);
            // candidates = set_union(candidates, neighbours);
        }

        basins = append(basins, basin);
    }

    maxs := []int{};
    for _, basin := range basins {
        l := len(basin);
        if (len(maxs) < 3) {
            maxs = append(maxs, l);
            sort.Ints(maxs);
            continue;
        } else if maxs[0] < l {
            maxs[0] = l;
            sort.Ints(maxs);
        }
    }

    result := 1;
    for _, l := range maxs {
        result *= l;
    }

    return strconv.Itoa(result);
}

type Basin []Coord;

//! mutating!
func set_add_collection(set map[Coord]bool, arr []Coord) map[Coord]bool {
    for _, e := range arr {
        set[e]=true;
    }
    return set;
}

func set_difference(lhs []Coord, rhs map[Coord]bool) []Coord {
    result_set := []Coord{};
    for _, e := range lhs {
        if _, ok := rhs[e]; !ok {
            result_set = append(result_set, e);
        }
    }
    return result_set;
}
// func set_difference(lhs map[Coord]bool, rhs map[Coord]bool) map[Coord]bool {
//     result_set := map[Coord]bool{};
//     for e, _ := range lhs {
//         if _, ok := rhs[e]; !ok {
//             result_set[e] = true;
//         }
//     }
//     return result_set;
// }
// func set_union(lhs map[Coord]bool, rhs map[Coord]bool) map[Coord]bool {
//     result_set := map[Coord]bool{};
//     for e, _ := range lhs {
//         result_set[e] = true;
//     }
//     for e, _ := range rhs {
//         result_set[e] = true;
//     }
//     return result_set;
// }

func array_clone(arr []Coord) []Coord {
    ret := []Coord{};
    for _, e := range arr {
        ret = append(ret, e);
    }
    return ret
}
func array_pop(arr []Coord) ([]Coord, Coord, error) {
    if len(arr) == 0 {
        msg := fmt.Sprintf("WARNING - Tried to pop empty array\n");
        fmt.Fprintf(os.Stderr, msg);
        return arr, Coord{}, errors.New(msg);
    }
    e := arr[len(arr)-1];
    ret := []Coord{};
    ret = append(ret, arr[:len(arr)-1]...);
    return ret, e, nil;
}