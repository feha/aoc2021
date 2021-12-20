package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 20:09:47
  * p1 done - 23:43:26
  * p2 done - 00:12:07
  */

func main() {
    input, _ := utils.Get_input(2021, 20);
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
    // `#.#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#...
    `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`,
};
var part1_test_output = []string{
    `35`,
};
func part1(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    alg_str := strings.Join(strings.Split(inputs[0], "\n"), ""); // ensure newlines inside it doesn't break things by removing them
    lines := strings.Split(inputs[1], "\n");

    alg := make([]int, 512);
    for i, c := range strings.Split(alg_str, "") {
        n := 0;
        if c == "#" {
            n = 1;
        }
        alg[i] = n;
    }

    grid := [][]int{};
    for _, line := range lines {
        row := []int{};
        for _, c := range strings.Split(line, "") {
            n := 0;
            if c == "#" {
                n = 1;
            }
            row = append(row, n);
        }
        grid = append(grid, row);
    }

    window_size := 3;
    steps := 2;

    output := prepare_output(grid, steps);

    grid = output.Grid(window_size, alg);
    result := 0;
    for _, row := range grid {
        for _, n := range row {
            if n == 1 {
                result++;
            }
        }
    }
    return strconv.Itoa(result);
}

type Coord struct {
    x, y int
}

type Grid interface {
    Get(int, int, int, []int) int
    Read_Window(int, int, int, []int) int
    Width() int
    Height() int
    Grid(int, []int) [][]int
}

func Read_Window(this Grid, x, y, window_size int, alg []int) int {
    half := window_size / 2;
    n := 0;
    for i:=-half; i <= half; i++ {
        for j:=-half; j <= half; j++ {
            n = n << 1 + this.Get(x+j, y+i, window_size, alg);
        }
    }
    if n > 512-1 {
        fmt.Println("BIG ERROR!! n=",n, half);
    }
    return n;
}

type Inner struct {
    grid [][]int
    // cache map[Coord]int
}
func (this Inner) Width() int {
    return len(this.grid[0]);
}
func (this Inner) Height() int {
    return len(this.grid);
}
func (this Inner) Get(x, y, _ int, _ []int) int {
    // fmt.Printf("Inner Get %d, %d ", x, y)
    if 0 <= y && y < this.Height() &&
            0 <= x && x < this.Width() {
        // fmt.Println(this.grid[y][x])
        return this.grid[y][x];
    }
    // fmt.Println("-")
    return 0;
}
func (this Inner) Read_Window(x, y, window_size int, alg []int) int {
    // fmt.Println("Inner Read_Window", x, y)
    return Read_Window(this, x, y, window_size, alg);
    
    // k := Coord{x,y};
    // n, cached := this.cache[k];
    // if !cached {
    //     half := window_size / 2;
    //     n = 0;
    //     for i:=-half; i <= half; i++ {
    //         for j:=-half; j <= half; j++ {
    //             n = n << 1 + this.Get(x+j, y+i, window_size, alg);
    //         }
    //     }
    //     if n > 512-1 {
    //         fmt.Println("BIG ERROR!! n=",n, half);
    //     }

    //     this.cache[k] = n;
    // }
    // return n;
}
func (this Inner) Grid(_ int, _ []int) [][]int {
    return this.grid;
}

type Padded struct {
    inner Grid
    cache map[Coord]int
}
func (this Padded) Width() int {
    return this.inner.Width()+2;
}
func (this Padded) Height() int {
    return this.inner.Height()+2;
}
func (this Padded) Get(x, y, window_size int, alg []int) int {
    k := Coord{x,y};
    n, cached := this.cache[k];
    if !cached {
        n = alg[this.inner.Read_Window(x-1,y-1, window_size, alg)];
        this.cache[k] = n;
    }
    return n;
    return alg[this.inner.Read_Window(x-1,y-1, window_size, alg)];
}
func (this Padded) Read_Window(x, y, window_size int, alg []int) int {
    // fmt.Println("padded Read_Window", x, y)
    return Read_Window(this, x, y, window_size, alg);

    // k := Coord{x,y};
    // n, cached := this.cache[k];
    // if !cached {
    //     half := window_size / 2;
    //     n = 0;
    //     for i:=-half; i <= half; i++ {
    //         for j:=-half; j <= half; j++ {
    //             n = n << 1 + this.Get(x+j, y+i, window_size, alg);
    //         }
    //     }
    //     if n > 512-1 {
    //         fmt.Println("BIG ERROR!! n=",n, half);
    //     }

    //     this.cache[k] = n;
    // }
    // return n;
}
func (this Padded) Grid(window_size int, alg []int) [][]int {
    grid := [][]int{};
    for y:=0; y<this.Height(); y++ {
        row := []int{};
        for x:=0; x<this.Width(); x++ {
            row = append(row, this.Get(x,y, window_size, alg));
        }
        grid = append(grid, row);
    }
    return grid;
}

func prepare_output(input [][]int, steps int) Grid {
    // output := Grid(Inner{input, map[Coord]int{}});
    output := Grid(Inner{input});
    for i:=0; i < steps; i++ {
        output = Padded{output, map[Coord]int{}};
    }
    return output;
    // // assume final-output is not an infinite field of white pixels
    // return pad_grid(input, steps);
}

func pad_grid(grid [][]int, steps int) [][]int {
    // pad in all directions
    pad := 2*steps;
    width, height := len(grid[0]) + pad, len(grid)+ pad;
    padded := make([][]int, height);
    for i, _ := range padded {
        padded[i] = make([]int, width);
    }
    // fill padded grid with the data from grid
    for y:=steps; y < height-steps; y++ {
        for x:=steps; x < width-steps; x++ {
            padded[y][x] = grid[y][x];
        }
    }
    
    return padded;
}

// func lazy_eval(output, input [][]int, window_size, steps int) {
//     for y, row := range output {
//         for x, _ := range row {
//             n := alg[read_window(x, y, sdasd, window_size)];
//             output[y][x] = n;
//         }
//     }
// }

func format_pretty(grid [][]int) string {
    str := "";
    for _, row := range grid {
        for _, n := range row {
            if n == 1 {
                str+="#";
            } else {
                str+=".";
            }
        }
        str+="\n";
    }
    return str;
}
func format_sidebyside(grids [][][]int) string {
    last := grids[len(grids)-1];
    width, height := len(last[len(last)-1]), len(last);
    str := "";
    for y:=0; y < height; y++ {
        for _, grid := range grids[:2] {
            for x:=0; x < width; x++ {
                if y < len(grid) && x < len(grid[0]) {
                    n := grid[y][x];
                    if n == 1 {
                        str+="#";
                    } else {
                        str+=".";
                    }
                } else {
                    str+=" ";
                }
            }
            str+= " | "
        }
        str+="\n";
    }
    return str;
}



func read_cell(x, y int, grid [][]int) int {
    if 0 <= y && y < len(grid) &&
            0 <= x && x < len(grid[0]) {
        return grid[y][x];
    }
    return 0;
}
func read_window(x, y int, grid [][]int, window_size int) int {
    n := 0;
    str := "";
    for i:=0; i < window_size; i++ {
        for j:=0; j < window_size; j++ {
            x_offset, y_offset := x + j-1, y + i-1;
            cell := read_cell(x_offset, y_offset, grid);
            n = n << 1 + cell;
            str += strconv.Itoa(cell);
        }
        str+="\n"
    }
    if n > 512-1 {
        fmt.Println("BIG ERROR!!");
    }
    // fmt.Println(str, x, y, n)
    return n;
}

func enhance(grid [][]int, alg []int, window_size int) [][]int {
    width, height := len(grid[0]), len(grid);
    new_grid := [][]int{};

    // [-1 -> width+1, -1 -> height+1] - pad grid by 1 in all directions
    for y:=-1; y < height+1; y++ {
        new_grid = append(new_grid, []int{});
        for x:=-1; x < width+1; x++ {
            n := read_window(x, y, grid, window_size);
            new_grid[y+1] = append(new_grid[y+1], alg[n]);
        }
    }

    return new_grid;
}

var part2_test_input = []string{
    `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`,
};
var part2_test_output = []string{
    `3351`,
};
func part2(input string) string {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    alg_str := strings.Join(strings.Split(inputs[0], "\n"), ""); // ensure newlines inside it doesn't break things by removing them
    lines := strings.Split(inputs[1], "\n");

    window_size := 3;
    steps := 50;

    alg := make([]int, 512);
    for i, c := range strings.Split(alg_str, "") {
        n := 0;
        if c == "#" {
            n = 1;
        }
        alg[i] = n;
    }

    grid := [][]int{};
    for _, line := range lines {
        row := []int{};
        for _, c := range strings.Split(line, "") {
            n := 0;
            if c == "#" {
                n = 1;
            }
            row = append(row, n);
        }
        grid = append(grid, row);
    }

    output := prepare_output(grid, steps);
    
    grid = output.Grid(window_size, alg);
    result := 0;
    for _, row := range grid {
        for _, n := range row {
            if n == 1 {
                result++;
            }
        }
    }
    return strconv.Itoa(result);
}
