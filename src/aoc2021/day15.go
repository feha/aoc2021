package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "sort"
    "os"
    "errors"
);

/**
  * Start - 16:10:20
  * p1 done - 16:57:51
  * p2 done - 17:10:25
  */

func main() {
    input, _ := utils.Get_input(2021, 15);
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
    `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
};
var part1_test_output = []string{
    `40`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    // grid := [][]Cell{};
    grid := map[Coord]Cell{};
    width, height := 0, 0;
    for y, line := range inputs {
        height = y+1;

        nums, _ := utils.StrToInt_array(strings.Split(line, ""));
        // row := []Cell{};
        for x, n := range nums {
            width = x+1;
            // row = append(row, Cell{cost: n});
            grid[Coord{x:x, y:y}] = Cell{cost: n, path_cost: -1};
        }
        // grid = append(grid, row);
    }

    start, end := Coord{x:0, y:0}, Coord{x:width-1, y:height-1};

    // first cell knows its path_cost
    cell_start := grid[start];
    cell_start.path_cost = 0; // "(the starting position is never entered, so its risk is not counted)." // = cell_start.cost;
    grid[start] = cell_start;

    candidates := []Coord{start};
    // candidates := map[Coord]bool{start:true};
    blacklist := map[Coord]bool{start:true};

    // candidates := map[Coord]bool{};
    // neighbours := get_neighbours(pos, width, height);
    // for _, p := range neighbours {
    //     candidates[p] = true;
    // }

    for len(candidates) > 0 {
        // var pos Coord;
        // // pop a value from the map
        // for p, _ := range candidates {
        //     delete(candidates, p);
        //     pos = p;
        //     break;
        // }

        candidates_redeclared, pos, _ := array_pop(candidates);
        candidates = candidates_redeclared; // ":=" operator hides in inner scope, but is required for pos

        cell := grid[pos];

        adjs := get_neighbours(pos, width, height);
        for _, adj := range adjs {
            cell_adj := grid[adj];
            cost := cell.path_cost + cell_adj.cost;
            if cell_adj.path_cost == -1 || cell_adj.path_cost > cost {
                cell_adj.path_cost = cost;
                cell_adj.path = pos;
                grid[adj] = cell_adj;
            }
            if _, ok := blacklist[adj]; !ok {
                candidates = append(candidates, adj);
                blacklist[adj] = true;
            }
        }

        // sort high -> low
        sort.Slice(candidates, func(i, j int) bool {
            return grid[candidates[i]].path_cost > grid[candidates[j]].path_cost
        })
    }

    cell_end := grid[end];

    path := []Cell{};
    path2 := []Coord{end};
    cur := cell_end;
    for cur != cell_start {
        path = append(path, cur);
        path2 = append(path2, cur.path);
        cur = grid[cur.path];
    }

    result := cell_end.path_cost;

    return strconv.Itoa(result);
}

type Coord struct {
    x, y int;
}
func (lhs Coord) Coord_add(rhs Coord) Coord {
    return Coord{x: lhs.x + rhs.x, y: lhs.y + rhs.y};
}

type Cell struct {
    cost, path_cost int;
    path Coord;
}
func (c Cell) String() string {
    return fmt.Sprintf("cost:%d, path_cost:%d, path:%v", c.cost, c.path_cost, c.path);
}

func get_neighbours(pos Coord, width int, height int) []Coord {
    neighbours := []Coord{};
    for x:=-1; x < 2; x++ {
        for y:=-1; y < 2; y++ {
            dir := Coord{x:x, y:y};
            pos2 := pos.Coord_add(dir);
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

var part2_test_input = []string{
    `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
};
var part2_test_output = []string{
    `315`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    tile_template := map[Coord]Cell{};
    template_width, template_height := 0, 0;
    for y, line := range inputs {
        template_height = y+1;

        nums, _ := utils.StrToInt_array(strings.Split(line, ""));
        for x, n := range nums {
            template_width = x+1;
            tile_template[Coord{x:x, y:y}] = Cell{cost: n};
        }
    }
    width, height := template_width*5, template_height*5;

    grid := map[Coord]Cell{};
    for x:=0; x < width; x++ {
        template_x := x % template_width;
        tile_x := x / template_width;
        for y:=0; y < height; y++ {
            template_y := y % template_height;
            tile_y := y / template_height;

            offset := tile_x + tile_y;

            template_pos := Coord{x:template_x, y:template_y};
            template_cell := tile_template[template_pos];

            pos := Coord{x:x, y:y};
            cost := template_cell.cost + offset;
            cost = ((cost-1) % 9)+1 // wrap 9 -> 1
            cell := Cell{cost: cost, path_cost: -1};
            grid[pos] = cell;
        }
    }

    start, end := Coord{x:0, y:0}, Coord{x:width-1, y:height-1};

    // first cell knows its path_cost
    cell_start := grid[start];
    cell_start.path_cost = 0; // "(the starting position is never entered, so its risk is not counted)." // = cell_start.cost;
    grid[start] = cell_start;

    candidates := []Coord{start};
    blacklist := map[Coord]bool{start:true};

    for len(candidates) > 0 {
        candidates_redeclared, pos, _ := array_pop(candidates);
        candidates = candidates_redeclared; // ":=" operator hides in inner scope, but is required for pos

        cell := grid[pos];

        adjs := get_neighbours(pos, width, height);
        for _, adj := range adjs {
            cell_adj := grid[adj];
            cost := cell.path_cost + cell_adj.cost;
            if cell_adj.path_cost == -1 || cell_adj.path_cost > cost {
                cell_adj.path_cost = cost;
                cell_adj.path = pos;
                grid[adj] = cell_adj;
            }
            if _, ok := blacklist[adj]; !ok {
                candidates = append(candidates, adj);
                blacklist[adj] = true;
            }
        }

        // sort high -> low
        sort.Slice(candidates, func(i, j int) bool {
            return grid[candidates[i]].path_cost > grid[candidates[j]].path_cost
        })
    }

    cell_end := grid[end];

    path := []Cell{};
    path2 := []Coord{end};
    cur := cell_end;
    for cur != cell_start {
        path = append(path, cur);
        path2 = append(path2, cur.path);
        cur = grid[cur.path];
    }

    result := cell_end.path_cost;

    return strconv.Itoa(result);
}
