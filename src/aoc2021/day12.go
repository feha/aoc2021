package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 19:47:14
  * p1 done - 20:56:05
  * p2 done - 22:23:06
  */

func main() {
    input, _ := utils.Get_input(2021, 12);
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
    `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
    `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
    `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
};
var part1_test_output = []string{
    `10`,
// start,A,b,A,c,A,end
// start,A,b,A,end
// start,A,b,end
// start,A,c,A,b,A,end
// start,A,c,A,b,end
// start,A,c,A,end
// start,A,end
// start,b,A,c,A,end
// start,b,A,end
// start,b,end`,
    `19`,
// start,HN,dc,HN,end
// start,HN,dc,HN,kj,HN,end
// start,HN,dc,end
// start,HN,dc,kj,HN,end
// start,HN,end
// start,HN,kj,HN,dc,HN,end
// start,HN,kj,HN,dc,end
// start,HN,kj,HN,end
// start,HN,kj,dc,HN,end
// start,HN,kj,dc,end
// start,dc,HN,end
// start,dc,HN,kj,HN,end
// start,dc,end
// start,dc,kj,HN,end
// start,kj,HN,dc,HN,end
// start,kj,HN,dc,end
// start,kj,HN,end
// start,kj,dc,HN,end
// start,kj,dc,end`,
    `226`,
// `,
};
func part1(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    nodes := create_nodes(lines);
    // fmt.Println("created:",nodes);
    // nodes = link_nodes(nodes);
    // fmt.Println("linked:",nodes);

    // visit start only once (aka only start there)
    start := nodes[start_node];
    delete(nodes, start_node);
    // paths := generate_paths(start, []Node{}, map[string]bool{}, nodes);
    paths := generate_paths(start, []*Node2{}, map[string]bool{}, nodes);

    result := len(paths);

    return strconv.Itoa(result);
}
var start_node string = "start";
var end_node string = "end";

type Node struct {
    name string;
    adjs_strs []string;
    adjs []Node;
    big bool;
}
type Node2 struct {
    name string;
    adjs_strs []string;
    adjs []*Node2;
    big bool;
}
type Node3 struct {
    name string;
    adjs_strs *[]string;
    adjs *[]*Node3;
    big bool;
}
// func (n Node) String() string {
//     adjs := "[";
//     for _, adj := range n.adjs {
//         adjs += adj.name + ",";
//     }
//     adjs += "]";
//     return fmt.Sprintf("{name:%s, adjs:%v, adjs_strs:%v}", n.name, adjs, n.adjs_strs)
// }
func (n Node) String() string {
    return fmt.Sprintf("%s, ", n.name)
}
func (n *Node2) String() string {
    return fmt.Sprintf("%s", n.name)
}
func format_pretty_node2(paths [][]*Node2) string {
    str := "[\n";
    for _, path := range paths {
        str += fmt.Sprintf("%v",path);
        str += "\n";
    }
    str += "]";
    return str;
}

// func create_nodes(lines []string) map[string]Node {
//     nodes := map[string]Node{};
//     for _, line := range lines {
//         kv := strings.Split(line, "-");
//         lhs := kv[0];
//         rhs := kv[1];
//         lhs_big := strings.ToLower(lhs) != lhs;
//         rhs_big := strings.ToLower(rhs) != rhs;
//         // create lhs node (if doesn't exist)
//         lhs_node, ok := nodes[lhs];
//         if !ok {
//             nodes[lhs] = Node{name: lhs, big: lhs_big};
//             lhs_node = nodes[lhs];
//         }
//         // create rhs node (if doesn't exist)
//         rhs_node, ok := nodes[rhs];
//         if !ok {
//             nodes[rhs] = Node{name: rhs, big: rhs_big};
//             rhs_node = nodes[rhs];
//         }
//         // bidirectional & only occur once (opposite dir is never written)
//         lhs_node.adjs_strs = append(lhs_node.adjs_strs, rhs);
//         rhs_node.adjs_strs = append(rhs_node.adjs_strs, lhs);
//         // Why the hell is this needed? did adjs_strs get cloned somehow (pass-by-value), and I dont append the src?
//         nodes[lhs] = lhs_node;
//         nodes[rhs] = rhs_node;
//     }
//     return nodes;
// }
func create_nodes(lines []string) map[string]*Node2 {
    nodes := map[string]*Node2{};
    for _, line := range lines {
        kv := strings.Split(line, "-");
        lhs := kv[0];
        rhs := kv[1];
        lhs_big := strings.ToLower(lhs) != lhs;
        rhs_big := strings.ToLower(rhs) != rhs;
        // create lhs node (if doesn't exist)
        lhs_node, ok := nodes[lhs];
        if !ok {
            lhs_node = &Node2{name: lhs, big: lhs_big};
            nodes[lhs] = lhs_node;
        }
        // create rhs node (if doesn't exist)
        rhs_node, ok := nodes[rhs];
        if !ok {
            rhs_node = &Node2{name: rhs, big: rhs_big};
            nodes[rhs] = rhs_node;
        }
        // bidirectional & only occur once (opposite dir is never given)
        lhs_node.adjs_strs = append(lhs_node.adjs_strs, rhs);
        rhs_node.adjs_strs = append(rhs_node.adjs_strs, lhs);
        // // Why the hell is this needed? did adjs_strs get cloned somehow (pass-by-value), and I dont append the src?
        // nodes[lhs] = lhs_node;
        // nodes[rhs] = rhs_node;
    }
    return nodes;
}
// func link_nodes(nodes map[string]Node) map[string]Node {
//     for _, node := range nodes {
//         for _, name := range node.adjs_strs {
//             adj := nodes[name];
//             node.adjs = append(node.adjs, adj);
//             // adj.adjs = append(adj.adjs, node);
//             // nodes[name] = adj;
//         }
//         node.adjs_strs = []string{};
//         nodes[node.name] = node;
//     }
//     return nodes;
// }

func clone_add(set map[string]bool, e string) map[string]bool {
    new_set := map[string]bool{};
    for k := range set {
        new_set[k]=true;
    }
    new_set[e]=true;
    return new_set;
}

func generate_paths(cur_node *Node2, path []*Node2, visited map[string]bool, nodes map[string]*Node2) [][]*Node2 {
    clone := []*Node2{};
    clone = append(clone, path...);
    path = append(clone, cur_node);
    // path = append(path, cur_node);
    
    if cur_node.name == end_node {
        return [][]*Node2{path};
    }

    visited = clone_add(visited, cur_node.name);

    paths := [][]*Node2{};
    for _, adj_name := range cur_node.adjs_strs {
        adj, ok := nodes[adj_name];
        if !ok {
            continue;
        }
        if adj.big || !visited[adj_name] {
            sub_paths := generate_paths(adj, path, visited, nodes);
            paths = append(paths, sub_paths...);
        }
    }
    return paths;
}
// func generate_paths(cur_node Node, path []Node, visited map[string]bool, nodes map[string]Node) [][]Node {
//     if cur_node.name == end_node {
//         return [][]Node{[]Node{cur_node}};
//     }

//     visited = clone_add(visited, cur_node.name);
//     path = append(path, cur_node);

//     paths := [][]Node{};
//     for _, adj_name := range cur_node.adjs_strs {
//         adj := nodes[adj_name];
//         if !adj.big && visited[adj_name] {
//             // visited small cave already
//             continue;
//         }
//         sub_paths := generate_paths(adj, path, visited, nodes);
//         paths = append(paths, sub_paths...);
//     }
//     return paths;
// }

var part2_test_input = []string{
    `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
    `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
    `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
};
var part2_test_output = []string{
    `36`,
    `103`,
    `3509`,
};
func part2(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    nodes := create_nodes2(lines);

    // visit start only once (aka only start there)
    start := nodes[start_node];
    delete(nodes, start_node);
    paths := generate_paths2(start, []Node{}, map[string]bool{}, false, nodes);

    result := len(paths);

    return strconv.Itoa(result);
}

func create_nodes2(lines []string) map[string]Node {
    nodes := map[string]Node{};
    for _, line := range lines {
        kv := strings.Split(line, "-");
        lhs := kv[0];
        rhs := kv[1];
        lhs_big := strings.ToLower(lhs) != lhs;
        rhs_big := strings.ToLower(rhs) != rhs;
        // create lhs node (if doesn't exist)
        lhs_node, ok := nodes[lhs];
        if !ok {
            nodes[lhs] = Node{name: lhs, big: lhs_big};
            lhs_node = nodes[lhs];
        }
        // create rhs node (if doesn't exist)
        rhs_node, ok := nodes[rhs];
        if !ok {
            nodes[rhs] = Node{name: rhs, big: rhs_big};
            rhs_node = nodes[rhs];
        }
        // bidirectional & only occur once (opposite dir is never written)
        lhs_node.adjs_strs = append(lhs_node.adjs_strs, rhs);
        rhs_node.adjs_strs = append(rhs_node.adjs_strs, lhs);
        // Why the hell is this needed? did ajds_strs get cloned somehow (pass-by-value), and I dont append the src?
        nodes[lhs] = lhs_node;
        nodes[rhs] = rhs_node;
    }
    return nodes;
}

func clone_add2(hist map[string]int, e string) map[string]int {
    new_hist := map[string]int{};
    for k, v := range hist {
        new_hist[k]=v;
    }
    new_hist[e]++;
    return new_hist;
}
func generate_paths2(cur_node Node, path []Node, visited map[string]bool, exception_cave bool, nodes map[string]Node) [][]Node {
    clone := []Node{};
    clone = append(clone, path...);
    path = append(clone, cur_node);
    // path = append(path, cur_node);

    // visit end only once, and end path there
    // aka terminate when we arrive at the end
    if cur_node.name == end_node {
        return [][]Node{path};
    }

    visited = clone_add(visited, cur_node.name);

    paths := [][]Node{};
    for _, adj_name := range cur_node.adjs_strs {
        // fmt.Println("adj_name=",adj_name, "path=",path);
        adj, ok := nodes[adj_name];
        if !ok {
            continue;
        }
        // visit small caves once at most, except for a single exception
        blah := exception_cave;
        if !adj.big && visited[adj_name] {
            if exception_cave {
                continue;
            } else {
                blah = true
            }
        }
        sub_paths := generate_paths2(adj, path, visited, blah, nodes);
        paths = append(paths, sub_paths...);
    }
    return paths;
}
func format_pretty(arrs [][]Node) string {
    str := "[\n";

    for _, row := range arrs {
        str += fmt.Sprintf("%v",row);
        str += "\n";
    }
    str += "]";

    return str;
}