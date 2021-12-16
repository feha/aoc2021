package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "sort"
    "os"
    "errors"
    "math"
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

    grid, width, height := generate_grid(inputs);
    // cell_end := djikstra(Coord{0, 0}, Coord{width-1, height-1}, grid, width, height);
    cell_end := astar(Coord{0, 0}, Coord{width-1, height-1}, grid, width, height);

    // grid, width, height := generate_grid_1D(inputs);
    // cell_end := astar_1D(as_1D(0, 0, width), as_1D(width-1, height-1, width), grid, width, height);

    result := cell_end.path_cost;

    return strconv.Itoa(result);
}

type Coord struct {
    x, y int;
}
func (lhs Coord) Coord_add(rhs Coord) Coord {
    return Coord{x: lhs.x + rhs.x, y: lhs.y + rhs.y};
}
func (p Coord) as_1D(width int) int {
    return p.x + p.y*width;
}
func as_1D(x int, y int, width int) int {
    return x + y*width;
}
func as_Coord(i int, width int) Coord {
    return Coord{i % width, i / width};
}

type Cell struct {
    cost, path_cost int;
    path Coord;
    heuristic int;
    pos Coord;
    path_1D, pos_1D int;
}
func (c Cell) String() string {
    return fmt.Sprintf("cost:%d, path_cost:%d, path:%v", c.cost, c.path_cost, c.path);
}

func generate_grid(inputs []string) ([][]Cell, int, int) {
    width, height := len(inputs), len(inputs[0]);
    grid := make([][]Cell, height);
    for y, line := range inputs {
        grid[y] = make([]Cell, width);
        for x, c := range strings.Split(line, "") {
            n, _ := strconv.Atoi(c);
            grid[y][x] = Cell{cost: n, path_cost: -1};
        }
    }
    return grid, width, height;
}
func djikstra(start Coord, end Coord, grid [][]Cell, width int, height int) Cell {
    cell_start := grid[start.y][start.x];
    cell_start.path_cost = 0; // "(the starting position is never entered, so its risk is not counted)." // = cell_start.cost;
    grid[start.y][start.x] = cell_start;

    candidates := []Coord{start};
    blacklist := map[Coord]bool{start:true};
    for len(candidates) > 0 {
        candidates_redeclared, pos, _ := array_pop(candidates);
        candidates = candidates_redeclared; // ":=" operator hides in inner scope, but is required for variable pos

        cell := grid[pos.y][pos.x];
        cur_cost := cell.path_cost;

        adjs := get_neighbours(pos, width, height);
        for _, adj := range adjs {
            cell_adj := grid[adj.y][adj.x];
            cost := cur_cost + cell_adj.cost;
            if cell_adj.path_cost == -1 || cell_adj.path_cost > cost {
                cell_adj.path_cost = cost;
                cell_adj.path = pos;
                grid[adj.y][adj.x] = cell_adj;
            }
            if _, ok := blacklist[adj]; !ok {
                candidates = append(candidates, adj);
                blacklist[adj] = true;
            }
        }

        // sort high -> low
        sort.Slice(candidates, func(i, j int) bool {
            p1, p2 := candidates[i], candidates[j];
            return grid[p1.y][p1.x].path_cost > grid[p2.y][p2.x].path_cost
        })
    }
    return grid[end.y][end.x];


    // candidates := []Coord{start};
    // blacklist := map[Coord]bool{start:true};

    // for len(candidates) > 0 {
    //     candidates_redeclared, pos, _ := array_pop(candidates);
    //     candidates = candidates_redeclared; // ":=" operator hides in inner scope, but is required for pos

    //     cell := grid[pos];

    //     adjs := get_neighbours(pos, width, height);
    //     for _, adj := range adjs {
    //         cell_adj := grid[adj];
    //         cost := cell.path_cost + cell_adj.cost;
    //         if adj == end { // early termination!
    //             cell_adj.path_cost = cost;
    //             cell_adj.path = pos;
    //             grid[adj] = cell_adj;
    //             candidates = []Coord{}; // empty to escape loop
    //             break;
    //         }
    //         if cell_adj.path_cost == -1 || cell_adj.path_cost > cost {
    //             cell_adj.path_cost = cost;
    //             cell_adj.path = pos;
    //             grid[adj] = cell_adj;
    //         }
    //         if _, ok := blacklist[adj]; !ok {
    //             candidates = append(candidates, adj);
    //             blacklist[adj] = true;
    //         }
    //     }

    //     // sort high -> low
    //     sort.Slice(candidates, func(i, j int) bool {
    //         return grid[candidates[i]].path_cost > grid[candidates[j]].path_cost
    //     })
    // }
}
func astar(start Coord, end Coord, grid [][]Cell, width int, height int) Cell {
    cell_start := grid[start.y][start.x];
    cell_start.path_cost = 0; // "(the starting position is never entered, so its risk is not counted)." // = cell_start.cost;
    grid[start.y][start.x] = cell_start;
    
    candidates := map[Coord]Cell{start:grid[start.y][start.x]};
    seen := map[Coord]Cell{};

    for len(candidates) > 0 {
        best, cell_best := pop_best(candidates);
        cur_cost := cell_best.path_cost;

        adjs := get_neighbours(best, width, height);
        for _, adj := range adjs {
            cell_adj := grid[adj.y][adj.x];
            cell_adj.path_cost = cur_cost + cell_adj.cost;
            cell_adj.path = best;
            if adj == end {
                return cell_adj;
            }
            cell_adj.heuristic = cell_adj.path_cost + heuristic(adj, end);
            
            cell_prior, is_cand := candidates[adj];
            if !is_cand || cell_adj.heuristic < cell_prior.heuristic { // if not candidate OR cheaper than existing candidate
                cell_prior, is_seen := seen[adj];
                if !is_seen || cell_adj.heuristic < cell_prior.heuristic { // if not seen OR cheaper than prior seen
                    candidates[adj] = cell_adj;
                }
            }
        }

        seen[best] = cell_best;
    }

    return grid[end.y][end.x]; // Can't find end
}
func pop_best(candidates map[Coord]Cell) (Coord, Cell) {
    // TODO do better
    var arbitrary Coord;
    for k, _ := range candidates {
        arbitrary=k;
        break;
    }
    min := candidates[arbitrary].heuristic;
    k_min := arbitrary;
    for k, v := range candidates {
        if v.heuristic < min {
            min = v.heuristic;
            k_min= k;
        }
    }
    v_min := candidates[k_min];
    delete(candidates, k_min);
    return k_min, v_min;
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
func heuristic(pos Coord, end Coord) int {
    return int(math.Abs(float64(end.x - pos.x)) + math.Abs(float64(end.y - pos.y)));
}

func generate_grid_1D(inputs []string) ([]Cell, int, int) {
    width, height := len(inputs), len(inputs[0]);
    grid := make([]Cell, width*height);
    for y, line := range inputs {
        for x, c := range strings.Split(line, "") {
            n, _ := strconv.Atoi(c);
            i := as_1D(x, y, width);
            grid[i] = Cell{cost: n, pos_1D: i};
        }
    }
    return grid, width, height;
}
func astar_1D(start int, end int, grid []Cell, width int, height int) Cell {
    cell_start := grid[start];
    cell_start.path_cost = 0; // "(the starting position is never entered, so its risk is not counted)." // = cell_start.cost;
    grid[start] = cell_start;

    // candidates := []Cell{grid[start]};
    candidates := map[int]Cell{start:grid[start]};
    // seen := []int{};
    seen := map[int]Cell{};

    for len(candidates) > 0 {
        best, cell_best := pop_best_1D(candidates);
        cur_cost := cell_best.path_cost;

        adjs := get_neighbours_1D(best, width, height);
        for _, adj := range adjs {
            cell_adj := grid[adj];
            cell_adj.path_cost = cur_cost + cell_adj.cost;
            cell_adj.path_1D = best;
            if adj == end {
                return cell_adj;
            }
            cell_adj.heuristic = cell_adj.path_cost + heuristic_1D(adj, end, width);
            
            cell_prior, is_cand := candidates[adj];
            if !is_cand || cell_adj.heuristic < cell_prior.heuristic { // if not candidate OR cheaper than existing candidate
                cell_prior, is_seen := seen[adj];
                if !is_seen || cell_adj.heuristic < cell_prior.heuristic { // if not seen OR cheaper than prior seen
                    candidates[adj] = cell_adj;
                }
            }
        }

        seen[best] = cell_best;
    }

    return cell_start; // Can't find end
}

func pop_best_1D(candidates map[int]Cell) (int, Cell) {
    // TODO do better
    arbitrary := 0;
    for k, _ := range candidates {
        arbitrary=k;
        break;
    }
    min := candidates[arbitrary].heuristic;
    k_min := arbitrary;
    for k, v := range candidates {
        if v.heuristic < min {
            min = v.heuristic;
            k_min= k;
        }
    }
    v_min := candidates[k_min];
    delete(candidates, k_min);
    return k_min, v_min;
}
func get_neighbours_1D(i int, width int, height int) []int {
    neighbours := []int{};
    for x:=-1; x < 2; x++ {
        for y:=-1; y < 2; y++ {
            // ignore diagonals
            if x != 0 && y != 0 {
                continue;
            }
            j := i + x + y*width;
            if i == j {
                continue;
            }
            if 0 <= j && j < width*height { // inside bounds
                wrap := (i % width);
                wrap2 := (j % width);
                if math.Abs(float64(wrap-wrap2)) <= 1  { // didn't wrap to next/prior row
                    neighbours = append(neighbours, j);
                }
            }
        }
    }
    return neighbours;
}
func heuristic_1D(i int, j int, width int) int {
    pos := as_Coord(i, width);
    end := as_Coord(j, width);
    return heuristic(pos, end);
}
// ! can be optimized by also having a map[int]bool alongside candidates
func contains_1D(arr []Cell, e int) (bool, int) {
    for i, e2 := range arr {
        if e == e2.pos_1D {
            return true, i;
        }
    }
    return false, -1;
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

    wrap := 5;


    // // grid, width, height := generate_grid(inputs);
    // // grid, width, height = expand_grid(grid, width, height);
    // nums, width, height := parse(inputs);
    // grid, width, height := generate_grid2(nums, width, height, wrap);

    // // cell_end := djikstra(Coord{0, 0}, Coord{width-1, height-1}, grid, width, height);
    // cell_end := astar(Coord{0, 0}, Coord{width-1, height-1}, grid, width, height);
    

    // grid, width, height := generate_grid(inputs);
    // cell_end := astar_wrap(Coord{0, 0}, Coord{width*wrap-1, height*wrap-1}, grid, width, height, wrap);


    // grid, width, height := generate_grid_1D(inputs);
    // grid, width, height = expand_grid_1D(grid, width, height);
    // cell_end := astar_1D(as_1D(0, 0, width), as_1D(width-1, height-1, width), grid, width, height);


    nums, width, height := parse(inputs);
    grid := generate_graph(nums, width, height, wrap);
    // fmt.Println(grid);
    cell_end := astar_generic(grid[0][0], grid[width*wrap-1][height*wrap-1], PointerMap{});

    result := cell_end.GetPathCost();
    // result := cell_end.(Node).path_cost;

    // cur := cell_end.(*Node);
    // for cur != grid[0][0] {
    //     cur.blah = true;
    //     cur = cur.owner;
    // }
    // str := "";
    // for _, row := range grid {
    //     for _, n := range row {
    //         if n.blah {
    //             str+= strconv.Itoa(n.GetCost());
    //         } else {
    //             str+= "#";
    //         }
    //     }
    //     str+="\n";
    // }
    // fmt.Println(str)

    return strconv.Itoa(result);
}

func parse(inputs []string,) ([][]int, int, int) {
    width, height := len(inputs), len(inputs[0]);
    grid := make([][]int, height);
    for y, line := range inputs {
        grid[y] = make([]int, width);
        for x, c := range strings.Split(line, "") {
            n, _ := strconv.Atoi(c);
            grid[y][x] = n-1;
        }
    }
    return grid, width, height;
}
func generate_grid2(nums [][]int, width int, height int, wrap int) ([][]Cell, int, int) {
    full_width, full_height := width*wrap, height*wrap;
    grid := make([][]Cell, full_height);
    for y, row := range nums {
        for tile_y:=0; tile_y < wrap; tile_y++ {
            full_y := y + tile_y*height;
            grid[full_y] = make([]Cell, full_width);
            for tile_x:=0; tile_x < wrap; tile_x++ {
                for x, n := range row {
                    full_x := x + tile_x*width;
                    
                    offset := tile_y + tile_x;
                    cost := (n + offset) % 9 + 1; // wrap 9 -> 1
                    grid[full_y][full_x] = Cell{cost: cost};
                }
            }
        }
    }
    return grid, full_width, full_height;
}
func expand_grid(tile_template [][]Cell, template_width int, template_height int) ([][]Cell, int, int) {
    width, height := template_width*5, template_height*5;
    grid := make([][]Cell, height);

    for y:=0; y < height; y++ {
        grid[y] = make([]Cell, width);

        template_y := y % template_height;
        tile_y := y / template_height;
        for x:=0; x < width; x++ {
            template_x := x % template_width;
            tile_x := x / template_width;

            offset := tile_x + tile_y;

            template_pos := Coord{x:template_x, y:template_y};
            template_cell := tile_template[template_pos.y][template_pos.x];

            pos := Coord{x:x, y:y};
            cost := template_cell.cost + offset;
            cost = ((cost-1) % 9)+1 // wrap 9 -> 1
            cell := Cell{cost: cost, path_cost: -1};
            grid[pos.y][pos.x] = cell;
        }
    }

    return grid, width, height;
}
func expand_grid_1D(tile_template []Cell, template_width int, template_height int) ([]Cell, int, int) {
    width, height := template_width*5, template_height*5;
    grid := make([]Cell, width*height);

    for x:=0; x < width; x++ {
        template_x := x % template_width;
        tile_x := x / template_width;
        for y:=0; y < height; y++ {
            template_y := y % template_height;
            tile_y := y / template_height;

            offset := tile_x + tile_y;

            template_pos := template_x + template_y * template_width;
            template_cell := tile_template[template_pos];
            
            cost := template_cell.cost + offset;
            cost = ((cost-1) % 9)+1 // wrap 9 -> 1
            cell := Cell{cost: cost, path_cost: -1};

            pos := x + y*width;
            grid[pos] = cell;
        }
    }

    return grid, width, height;
}


func astar_wrap(start Coord, end Coord, grid [][]Cell, template_width int, template_height int, wrap int) Cell {
    width, height := template_width*wrap, template_height*wrap;

    cell_start := grid[start.y][start.x];
    cell_start.path_cost = 0; // "(the starting position is never entered, so its risk is not counted)." // = cell_start.cost;
    grid[start.y][start.x] = cell_start;
    
    candidates := map[Coord]Cell{start:grid[start.y][start.x]};
    seen := map[Coord]Cell{};

    for len(candidates) > 0 {
        best, cell_best := pop_best(candidates);
        cur_cost := cell_best.path_cost;

        adjs := get_neighbours(best, width, height);
        for _, adj := range adjs {
            wrapped_adj := Coord{x: adj.x % template_width, y: adj.y % template_height};
            cell_adj := grid[wrapped_adj.y][wrapped_adj.x];

            offset := adj.x / template_width + adj.y / template_height;
            cost := (cell_adj.cost + offset - 1) % 9 + 1;

            cell_adj.path_cost = cur_cost + cost;
            cell_adj.path = best;
            if adj == end {
                return cell_adj;
            }
            cell_adj.heuristic = cell_adj.path_cost + heuristic(adj, end);
            
            cell_prior, is_cand := candidates[adj];
            if !is_cand || cell_adj.heuristic < cell_prior.heuristic { // if not candidate OR cheaper than existing candidate
                cell_prior, is_seen := seen[adj];
                if !is_seen || cell_adj.heuristic < cell_prior.heuristic { // if not seen OR cheaper than prior seen
                    candidates[adj] = cell_adj;
                }
            }
        }

        seen[best] = cell_best;
    }

    return cell_start; // Can't find end
}



func generate_graph(nums [][]int, width, height int, wrap int) [][]*Node {
    full_width, full_height := width*wrap, height*wrap;
    grid := make([][]*Node, full_height);
    for y, row := range nums {
        for tile_y:=0; tile_y < wrap; tile_y++ {
            full_y := y + tile_y*height;
            grid[full_y] = make([]*Node, full_width);
            for tile_x:=0; tile_x < wrap; tile_x++ {
                for x, n := range row {
                    full_x := x + tile_x*width;
                    
                    offset := tile_y + tile_x;
                    cost := (n + offset) % 9 + 1; // wrap 9 -> 1 (subtraction is done in parse())
                    grid[full_y][full_x] = &Node{cost: cost, pos: Point{full_x, full_y}};
                }
            }
        }
    }
    link_graph(grid);

    return grid;
}
func link_graph(grid [][]*Node) {
    offsets := []Point{
        Point{x:-1, y: 0},
        Point{x: 1, y: 0},
        Point{x: 0, y:-1},
        Point{x: 0, y: 1},
    };
    for y, row := range grid {
        for x, n := range row {
            neighbours := []CollectionNode{};
            for _, p := range offsets {
                adj := Get(grid, x + p.x, y + p.y);
                if adj != nil {
                    neighbours = append(neighbours, adj);
                } else {
                    // fmt.Fprintf(os.Stderr, "Found a nil-pointer at %d,%d!\n", x + p.x, y + p.y);
                }
            }
            n.SetNeighbours(neighbours);
        }
    }
}
func Get(this [][]*Node, x int, y int) *Node {
    if 0 <= y && y < len(this) && 0 <= x && x < len(this[0]) {
        return this[y][x];
    }
    return nil;
}

func astar_generic(start CollectionNode, end CollectionNode, candidates SortedCollection) CollectionNode {
    start.SetCost(0); // "(the starting position is never entered, so its risk is not counted)."
    
    start_node := create_astarnode(start);
    start_node.is_cand = true;
    candidates.Add(start_node);

    closed := map[astar_node]bool{};

    for candidates.Len() > 0 {
        best_node := candidates.PopBest();
        cur_cost := best_node.node.GetPathCost();
        for _, adj := range best_node.node.GetNeighbours() {
            adj_node := create_astarnode(adj);
            cost := cur_cost + adj.GetCost();
            // heuristic := cost + heuristic(adj, end);

            change := false;
            if cost < adj.GetPathCost() {
                // if is_cand {
                // }
                delete(closed, adj_node);
                change = true;
            }
            is_cand := candidates.Contains(adj_node);
            _, is_closed := closed[adj_node];
            if change || (!is_cand && !is_closed) {
                adj.SetPathCost(cost);
                adj.SetOwner(best_node.node);
                adj.SetHeuristic(end);
                if !is_cand {
                    candidates.Add(adj_node);
                }
            }

            if adj == end {
                return adj;
            }


            
            // is_cand, prior_cost := candidates.Contains(adj);
            // if is_cand && prior_cost < heuristic {
            //     continue;
            // }
            // cell_prior, is_cand := candidates[adj];
            // if !is_cand || heuristic < cell_prior.heuristic { // if not candidate OR cheaper than existing candidate
            //     cell_prior, is_seen := seen[adj];
            //     if !is_seen || heuristic < cell_prior.heuristic { // if not seen OR cheaper than prior seen
            //         candidates[adj] = cell_adj;
            //     }
            // }
        }

        // closed[*(best.(*Node))] = true;
        closed[best_node] = true;
    }

    fmt.Println("ERROR")
    return nil; // Can't find end
}

type astar_node struct {
    node CollectionNode;
    // cost, heuristic int;
    // owner *astar_node;
    // pos Point;
    is_cand, close bool;
}
// type pool map[CollectionNode]astar_node;
func create_astarnode(n CollectionNode) astar_node {
    return astar_node{
        node: n,
    };
}

type CollectionNode interface {
    Compare(rhs CollectionNode) int
    GetNeighbours() []CollectionNode
    SetNeighbours([]CollectionNode)
    GetCost() int
    SetCost(i int) int
    GetPathCost() int
    SetPathCost(i int) int
    GetHeuristic() int
    SetHeuristic(target CollectionNode) int
    SetOwner(n CollectionNode)
}

type SortedCollection interface {
    Len() int

    PopBest() astar_node

    Add(e astar_node)
    
    Contains(e astar_node) bool
}

type Point struct {
    x, y int;
}
func (lhs Point) add(rhs Point) Point {
    return Point{x: lhs.x + rhs.x, y: lhs.y + rhs.y};
}
func (p Point) as_1D(width int) int {
    return p.x + p.y*width;
}

type Node struct {
    cost, pathcost, heuristic int;
    owner *Node;
    pos Point;
    neighbours *[]*Node;
    blah bool
}
func (this Node) String() string {
    return fmt.Sprintf("cost:%d, heuristic:%d, pos:%v",
            this.cost,
            this.heuristic,
            this.pos);
}

// type Node_Not_Pointer Node;
// func (lhs Node_Not_Pointer) Compare(rhs CollectionNode) int {
//     if lhs.heuristic < rhs.GetHeuristic() {
//         return -1;
//     } else {
//         return 1;
//     }
// }
// func (this Node_Not_Pointer) GetNeighbours() []CollectionNode {
//     neighbours := make([]CollectionNode, len(*this.neighbours));
//     for i, n := range *this.neighbours {
//         neighbours[i] = n;
//     }
//     return neighbours;
// }
// func (this Node_Not_Pointer) SetNeighbours(ns []CollectionNode) {
//     arr := make([]*Node, len(ns))
//     this.neighbours = &arr;
//     for i, n := range ns {
//         (*this.neighbours)[i] = n.(*Node);
//     }
// }
// func (this Node_Not_Pointer) GetCost() int {
//     return this.cost;
// }
// func (this Node_Not_Pointer) SetCost(i int) int {
//     this.cost = i;
//     return this.cost;
// }
// func (this Node_Not_Pointer) GetHeuristic() int {
//     return this.heuristic;
// }
// func (this Node_Not_Pointer) SetHeuristic(target CollectionNode) int {
//     lhs, rhs := this.pos, target.(Node_Not_Pointer).pos;
//     this.heuristic = this.GetCost() + int(math.Abs(float64(rhs.x - lhs.x)) + math.Abs(float64(rhs.y - lhs.y)));
//     return this.heuristic;
// }
// func (this Node_Not_Pointer) SetOwner(n CollectionNode) {
//     this.owner = n.(*Node);
// }

// func (lhs Node) Compare(rhs CollectionNode) int {
//     if lhs.heuristic < rhs.GetHeuristic() {
//         return -1;
//     } else {
//         return 1;
//     }
// }
// func (this Node) GetNeighbours() []CollectionNode {
//     neighbours := make([]CollectionNode, len(this.neighbours));
//     for i, n := range this.neighbours {
//         neighbours[i] = n;
//     }
//     return neighbours;
// }
// func (this Node) SetNeighbours(ns []CollectionNode) {
//     this.neighbours = make([]*Node, len(ns));
//     for i, n := range ns {
//         this.neighbours[i] = n.(*Node);
//     }
// }
// func (this Node) GetCost() int {
//     return this.cost;
// }
// func (this Node) SetCost(i int) int {
//     this.cost = i;
//     return this.cost;
// }
// func (this Node) GetHeuristic() int {
//     return this.heuristic;
// }
// func (this Node) SetHeuristic(target CollectionNode) int {
//     lhs, rhs := this.pos, target.(Node).pos;
//     this.heuristic = this.GetCost() + int(math.Abs(float64(rhs.x - lhs.x)) + math.Abs(float64(rhs.y - lhs.y)));
//     return this.heuristic;
// }
// func (this Node) SetOwner(n CollectionNode) {
//     this.owner = n.(*Node);
// }

func (lhs *Node) Compare(rhs CollectionNode) int {
    if lhs.heuristic < rhs.GetHeuristic() {
        return -1;
    } else {
        return 1;
    }
}
func (this *Node) GetNeighbours() []CollectionNode {
    neighbours := make([]CollectionNode, len(*this.neighbours));
    for i, n := range *this.neighbours {
        neighbours[i] = n;
    }
    return neighbours;
}
func (this *Node) SetNeighbours(ns []CollectionNode) {
    arr := make([]*Node, len(ns))
    this.neighbours = &arr;
    for i, n := range ns {
        (*this.neighbours)[i] = n.(*Node);
    }
}
func (this *Node) GetCost() int {
    return this.cost;
}
func (this *Node) SetCost(i int) int {
    this.cost = i;
    return this.cost;
}
func (this *Node) GetPathCost() int {
    return this.pathcost;
}
func (this *Node) SetPathCost(i int) int {
    this.pathcost = i;
    return this.pathcost;
}
func (this *Node) GetHeuristic() int {
    return this.heuristic;
}
func (this *Node) SetHeuristic(target CollectionNode) int {
    lhs, rhs := this.pos, target.(*Node).pos;
    this.heuristic = this.GetPathCost() + int(math.Abs(float64(rhs.x - lhs.x)) + math.Abs(float64(rhs.y - lhs.y)));
    return this.heuristic;
}
func (this *Node) SetOwner(n CollectionNode) {
    this.owner = n.(*Node);
}


type PointerMap map[astar_node]int;
func (this PointerMap) Len() int {
    return len(this);
}
func (this PointerMap) Contains(e astar_node) bool {
    _, ok := this[e];
    return ok;
}
func (this PointerMap) Add(e astar_node) {
    this[e] = e.node.GetPathCost();
}
func (this PointerMap) PopBest() astar_node {
    var arbitrary astar_node;
    for k, _ := range this {
        arbitrary = k;
        break;
    }

    min := arbitrary;
    for k, _ := range this {
        if min.node.Compare(k.node) > 0 { // v < min
            min = k;
        }
    }
    delete(this, min);
    return min;
}



// func astar_generic(start CollectionNode, end CollectionNode, candidates SortedCollection) CollectionNode {
//     start.SetCost(0); // "(the starting position is never entered, so its risk is not counted)."
    
//     candidates.Add(start);
//     closed := map[Point]bool{};

//     for candidates.Len() > 0 {
//         best := candidates.PopBest();
//         cur_cost := best.GetCost();
//         for _, adj := range best.GetNeighbours() {
//             cost := cur_cost + adj.GetCost();
//             // heuristic := cost + heuristic(adj, end);

//             is_cand := candidates.Contains(adj);
//             // _, is_closed := closed[*(adj.(*Node))];
//             _, is_closed := closed[adj.(*Node).pos];
//             if cost < adj.GetCost() {
//                 // if is_cand {
//                 // }
//                 delete(closed, best.(*Node).pos);
//             }
//             if !is_cand && !is_closed {
//                 adj.SetCost(cost);
//                 adj.SetOwner(best);
//                 adj.SetHeuristic(end);
//                 if !is_cand {
//                     candidates.Add(adj);
//                 }
//             }
//             if adj == end {
//                 return adj;
//             }


            
//             // is_cand, prior_cost := candidates.Contains(adj);
//             // if is_cand && prior_cost < heuristic {
//             //     continue;
//             // }
//             // cell_prior, is_cand := candidates[adj];
//             // if !is_cand || heuristic < cell_prior.heuristic { // if not candidate OR cheaper than existing candidate
//             //     cell_prior, is_seen := seen[adj];
//             //     if !is_seen || heuristic < cell_prior.heuristic { // if not seen OR cheaper than prior seen
//             //         candidates[adj] = cell_adj;
//             //     }
//             // }
//         }

//         // closed[*(best.(*Node))] = true;
//         closed[best.(*Node).pos] = true;
//     }

//     fmt.Println("ERROR")
//     return nil; // Can't find end
// }


// type CollectionNode interface {
//     Compare(rhs CollectionNode) int
//     GetNeighbours() []CollectionNode
//     SetNeighbours([]CollectionNode)
//     GetCost() int
//     SetCost(i int) int
//     GetHeuristic() int
//     SetHeuristic(target CollectionNode) int
//     SetOwner(n CollectionNode)
// }

// type SortedCollection interface {
//     Len() int

//     PopBest() CollectionNode

//     Add(e CollectionNode)
    
//     Contains(e CollectionNode) bool
// }

// type Point struct {
//     x, y int;
// }
// func (lhs Point) add(rhs Point) Point {
//     return Point{x: lhs.x + rhs.x, y: lhs.y + rhs.y};
// }
// func (p Point) as_1D(width int) int {
//     return p.x + p.y*width;
// }

// type Node struct {
//     cost, heuristic int;
//     owner *Node;
//     pos Point;
//     neighbours *[]*Node;
//     blah bool
// }
// func (this Node) String() string {
//     return fmt.Sprintf("cost:%d, heuristic:%d, pos:%v",
//             this.cost,
//             this.heuristic,
//             this.pos);
// }

// // type Node_Not_Pointer Node;
// // func (lhs Node_Not_Pointer) Compare(rhs CollectionNode) int {
// //     if lhs.heuristic < rhs.GetHeuristic() {
// //         return -1;
// //     } else {
// //         return 1;
// //     }
// // }
// // func (this Node_Not_Pointer) GetNeighbours() []CollectionNode {
// //     neighbours := make([]CollectionNode, len(*this.neighbours));
// //     for i, n := range *this.neighbours {
// //         neighbours[i] = n;
// //     }
// //     return neighbours;
// // }
// // func (this Node_Not_Pointer) SetNeighbours(ns []CollectionNode) {
// //     arr := make([]*Node, len(ns))
// //     this.neighbours = &arr;
// //     for i, n := range ns {
// //         (*this.neighbours)[i] = n.(*Node);
// //     }
// // }
// // func (this Node_Not_Pointer) GetCost() int {
// //     return this.cost;
// // }
// // func (this Node_Not_Pointer) SetCost(i int) int {
// //     this.cost = i;
// //     return this.cost;
// // }
// // func (this Node_Not_Pointer) GetHeuristic() int {
// //     return this.heuristic;
// // }
// // func (this Node_Not_Pointer) SetHeuristic(target CollectionNode) int {
// //     lhs, rhs := this.pos, target.(Node_Not_Pointer).pos;
// //     this.heuristic = this.GetCost() + int(math.Abs(float64(rhs.x - lhs.x)) + math.Abs(float64(rhs.y - lhs.y)));
// //     return this.heuristic;
// // }
// // func (this Node_Not_Pointer) SetOwner(n CollectionNode) {
// //     this.owner = n.(*Node);
// // }

// // func (lhs Node) Compare(rhs CollectionNode) int {
// //     if lhs.heuristic < rhs.GetHeuristic() {
// //         return -1;
// //     } else {
// //         return 1;
// //     }
// // }
// // func (this Node) GetNeighbours() []CollectionNode {
// //     neighbours := make([]CollectionNode, len(this.neighbours));
// //     for i, n := range this.neighbours {
// //         neighbours[i] = n;
// //     }
// //     return neighbours;
// // }
// // func (this Node) SetNeighbours(ns []CollectionNode) {
// //     this.neighbours = make([]*Node, len(ns));
// //     for i, n := range ns {
// //         this.neighbours[i] = n.(*Node);
// //     }
// // }
// // func (this Node) GetCost() int {
// //     return this.cost;
// // }
// // func (this Node) SetCost(i int) int {
// //     this.cost = i;
// //     return this.cost;
// // }
// // func (this Node) GetHeuristic() int {
// //     return this.heuristic;
// // }
// // func (this Node) SetHeuristic(target CollectionNode) int {
// //     lhs, rhs := this.pos, target.(Node).pos;
// //     this.heuristic = this.GetCost() + int(math.Abs(float64(rhs.x - lhs.x)) + math.Abs(float64(rhs.y - lhs.y)));
// //     return this.heuristic;
// // }
// // func (this Node) SetOwner(n CollectionNode) {
// //     this.owner = n.(*Node);
// // }

// func (lhs *Node) Compare(rhs CollectionNode) int {
//     if lhs.heuristic < rhs.GetHeuristic() {
//         return -1;
//     } else {
//         return 1;
//     }
// }
// func (this *Node) GetNeighbours() []CollectionNode {
//     neighbours := make([]CollectionNode, len(*this.neighbours));
//     for i, n := range *this.neighbours {
//         neighbours[i] = n;
//     }
//     return neighbours;
// }
// func (this *Node) SetNeighbours(ns []CollectionNode) {
//     arr := make([]*Node, len(ns))
//     this.neighbours = &arr;
//     for i, n := range ns {
//         (*this.neighbours)[i] = n.(*Node);
//     }
// }
// func (this *Node) GetCost() int {
//     return this.cost;
// }
// func (this *Node) SetCost(i int) int {
//     this.cost = i;
//     return this.cost;
// }
// func (this *Node) GetHeuristic() int {
//     return this.heuristic;
// }
// func (this *Node) SetHeuristic(target CollectionNode) int {
//     lhs, rhs := this.pos, target.(*Node).pos;
//     this.heuristic = this.GetCost() + int(math.Abs(float64(rhs.x - lhs.x)) + math.Abs(float64(rhs.y - lhs.y)));
//     return this.heuristic;
// }
// func (this *Node) SetOwner(n CollectionNode) {
//     this.owner = n.(*Node);
// }


// type PointerMap map[CollectionNode]int;
// func (this PointerMap) Len() int {
//     return len(this);
// }
// func (this PointerMap) Contains(e CollectionNode) bool {
//     _, ok := this[e];
//     return ok;
// }
// func (this PointerMap) Add(e CollectionNode) {
//     this[e] = e.GetCost();
// }
// func (this PointerMap) PopBest() CollectionNode {
//     var arbitrary CollectionNode;
//     for k, _ := range this {
//         arbitrary = k;
//         break;
//     }

//     min := arbitrary;
//     for k, _ := range this {
//         if min.Compare(k) > 0 { // v < min
//             min = k;
//         }
//     }
//     delete(this, min);
//     return min;
// }

// TODO implement binary heap