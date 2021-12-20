package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "regexp"
);

/**
  * Start - 14:39:14
  * p1 done - 18:28:30
  * p2 done - 18:59:35
  */

func main() {
    input, _ := utils.Get_input(2021, 19);
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
    `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`,
};
var part1_test_output = []string{
    ``,
    // `79`,
};
func part1(input string) string {
    scanners := parse(input);

    generate_rotations();
    // From AoC: we are guaranteed it's an overlap if _12_ beacons match
    overlap := 12;

    lhs := scanners[0];
    candidates := append([]*Scanner{}, scanners[1:]...);
    for len(candidates) > 0 {
        // pop safely
        rhs := candidates[0];
        candidates = append([]*Scanner{}, candidates[1:]...);

        overlap, unique := has_overlap(lhs, rhs, overlap);
        if overlap {
            // fmt.Printf("%d overlaps: %d, %d\n", len(lhs.beacons), len(lhs.beacons)+len(unique));
            for _, p := range unique {
                any := false;
                for _, p2 := range lhs.beacons {
                    if p == p2 {
                        any = true;
                    }
                }
                if any {
                    fmt.Println("BIG ERROR!");
                }
                // lhs.beacons = append(lhs.beacons, p);
                // lhs.beacons[p] = true;
            }
            lhs.beacons = append(lhs.beacons, unique...);
        } else {
            candidates = append(candidates, rhs);
        }
        fmt.Println(len(candidates));
    }

    result := len(lhs.beacons);
    return strconv.Itoa(result);
}
// func part1(input string) string {
//     return "";
// }
// func part1(input string) string {
//     scanners := parse(input);

//     generate_rotations();
//     // From AoC: we are guaranteed it's an overlap if _12_ beacons match
//     overlap := 12;

//     pop := scanners[0];
//     beacons := []Coord{};
//     lhs := map[Coord]*Scanner{Coord{}:pop};


//     candidates := append([]*Scanner{}, scanners[1:]...);
//     for len(candidates) > 0 {
//         // pop safely
//         rhs := candidates[0];
//         candidates = append([]*Scanner{}, candidates[1:]...);

//         overlap, unique := has_overlap(lhs, rhs, overlap);
//         if overlap {
//             // fmt.Printf("%d overlaps: %d, %d\n", len(lhs.beacons), len(lhs.beacons)+len(unique));
//             for _, p := range unique {
//                 any := false;
//                 for _, p2 := range lhs.beacons {
//                     if p == p2 {
//                         any = true;
//                     }
//                 }
//                 if any {
//                     fmt.Println("BIG ERROR!");
//                 }
//                 // lhs.beacons = append(lhs.beacons, p);
//                 // lhs.beacons[p] = true;
//             }
//             lhs.beacons = append(lhs.beacons, unique...);
//         } else {
//             candidates = append(candidates, rhs);
//         }
//         fmt.Println(len(candidates));
//     }

//     result := len(lhs.beacons);
//     return strconv.Itoa(result);
// }

type Coord struct {
    x, y, z int
}
func (this Coord) Add(rhs Coord) Coord {
    return Coord{this.x + rhs.x, this.y + rhs.y, this.z + rhs.z};
}
func (this Coord) Sub(rhs Coord) Coord {
    return Coord{this.x - rhs.x, this.y - rhs.y, this.z - rhs.z};
}

type Scanner struct {
    beacons []Coord
    // beacons map[Coord]Beacon
    // beacons map[Coord]bool
    // orientations *[]map[Coord]bool
    orientations [][]Coord
}
// type Beacon struct {
//     pos Coord
// }

func parse(input string) []*Scanner {
    inputs := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));

    re := regexp.MustCompile("(-?\\d+),(-?\\d+),(-?\\d+)");
    scanners := []*Scanner{};
    for _, input := range inputs {
        lines := strings.Split(input, "\n");
        scanner := Scanner{};
        for _, line := range lines[1:] {
            match := re.FindStringSubmatch(line);
            x, _ := strconv.Atoi(match[1]);
            y, _ := strconv.Atoi(match[2]);
            z, _ := strconv.Atoi(match[3]);
            p := Coord{x, y, z};
            scanner.beacons = append(scanner.beacons, p);
            // scanner.beacons[p] = true;
        }
        scanners = append(scanners, &scanner);
    }

    return scanners;
}

var rots []Coord;
func generate_rotations() {
    // all rotations
    all_rots := []Coord{};
    for z:=0; z<4; z++ {
        for y:=0; y<4; y++ {
            for x:=0; x<4; x++ {
                all_rots = append(all_rots, Coord{x,y,z});
            }
        }
    }

    test := Coord{1,2,3};
    exists := map[Coord]bool{};
    rots = []Coord{};
    for _, rot := range all_rots {
        rot_test := rotate(test, rot);
        if !exists[rot_test] {
            rots = append(rots, rot);
        }
        exists[rot_test]=true;
    }
    fmt.Println("generate_rotations - ", len(rots), rots);
}

func create_orientations(beacons []Coord) [][]Coord {
    orientations := [][]Coord{};
// func create_orientations(beacons map[Coord]bool) []map[Coord]bool {
//     orientations := []map[Coord]bool{beacons};

    // all rotations
    for _, rot := range rots {
        orientation := []Coord{};
        for _, p := range beacons {
        // orientation := map[Coord]bool{};
        // for p, _ := range beacons {
            orientation = append(orientation, rotate(p, rot));
            // orientation[rotate(p, rot)] = true;
        }
        orientations = append(orientations, orientation)
    }
    return orientations;
    // // filter duplicates (4*4*4 != 24)
    // ret := [][]Coord{};
    // // ret := []map[Coord]bool{};
    // for i:=0; i<len(orientations)-1; i++ {
    //     any_duplicate := false;
    //     for j:=i+1; j<len(orientations); j++ {
    //         diff := symmetric_diff(orientations[i], orientations[j])
    //         if len(diff) == 0 {
    //             fmt.Println("identical orientations!! i=",i," j=",j);
    //             any_duplicate = true;
    //             break;
    //         }
    //         // all := len(orientations[i]) == len(orientations[j]);
    //         // if all {
    //         //     for k, p := range orientations[i] {
    //         //         all = all && p == orientations[j][k];
    //         //     }
    //         //     if all {
    //         //         fmt.Println("identical orientations!! i=",i," j=",j);
    //         //         any_duplicate = true;
    //         //     }
    //         // }
    //     }
    //     if !any_duplicate {
    //         ret = append(ret, orientations[i]);
    //     }
    // }
    // return ret;
}
func rotate(p Coord, axis Coord) Coord { // axis = {0-3,0-3,0-3}
    if (axis.x == 0 && axis.y == 0 && axis.z == 0) ||
            (axis.x == 2 && axis.y == 2 && axis.z == 2) {
        // fmt.Println("warning! no rotation", axis);
        return p;
    }
    if axis.x < 0 || axis.y < 0 || axis.z < 0 {
        fmt.Println("error! rotating wrong direction");
        return p;
    }
    if axis.x > 3 || axis.y > 3 || axis.z > 3 {
        fmt.Println("error! rotating more than 1 lap");
        return p;
    }
    
    old_p := p;
    // x-axis
    for i:=0; i < axis.x; i++ {
        p = Coord{p.x, -p.z, p.y};
    }
    // y-axis
    for i:=0; i < axis.y; i++ {
        p = Coord{p.z, p.y, -p.x};
    }
    // z-axis
    for i:=0; i < axis.z; i++ {
        p = Coord{p.y, -p.x, p.z};
    }
    if old_p == p {
        fmt.Println("error! nothing changed", old_p, p, axis);
        return p;
    }
    return p;
}
func symmetric_diff(lhs []Coord, rhs []Coord) []Coord {
    unique := []Coord{};
    unique = append(unique, diff(lhs, rhs)...);
    unique = append(unique, diff(rhs, lhs)...);
    return unique;
}
func diff(lhs []Coord, rhs []Coord) []Coord {
    unique := []Coord{};
    // lhs - rhs
    for _, p1 := range lhs {
        not_contains := true;
        for _, p2 := range rhs {
            not_contains = not_contains && p1 != p2
        }
        if not_contains {
            unique = append(unique, p1);
        }
    }
    return unique;
}

// From AoC: scanners coordinate system has a random 90-degree rotation around any of the 3 axes (x, y, z)
func has_overlap(lhs, rhs *Scanner, overlap int) (bool, []Coord, Coord) {
    lhs_orientation := lhs.beacons; // we already rotate rhs, no need to rotate lhs too
    if len(rhs.orientations) == 0 {
        rhs.orientations = create_orientations(rhs.beacons);
    }
    rhs_orientations := rhs.orientations;
    
    for _, rhs_orientation := range rhs_orientations {
        overlap, unique, rhs_in_lhs := has_overlap_helper(lhs_orientation, rhs_orientation, overlap)
        if overlap {
            return true, unique, rhs_in_lhs;
        }
    }
    return false, []Coord{}, Coord{};
}
// func has_overlap_helper(lhs, rhs []Coord, overlap int) (bool, []Coord) {
// // func has_overlap_helper(lhs, rhs map[Coord]bool) (int, []Coord) {
//     for _, lhs_p := range lhs {
//     // for lhs_p, _ := range lhs {
//         // without := []Coord{};
//         // without = append(without, lhs[0:i]);
//         // without = append(without, lhs[i+1:]);
//         // // without := map[Coord]bool{};
//         // // for lhs_p2, _ := range lhs {
//         // //     if lhs_p != lhs_p2 {
//         // //         without[lhs_p2] = true;
//         // //     }
//         // // }
//         // lhs_local := to_local(lhs_p, without);
//         lhs_local := to_local(lhs_p, lhs);
//         for _, rhs_p := range rhs {
//         // for rhs_p, _ := range rhs {
//             // without := []Coord{};
//             // without = append(without, rhs[0:i]);
//             // without = append(without, rhs[i+1:]);
//             // // without := map[Coord]bool{};
//             // // for rhs_p2, _ := range rhs {
//             // //     if rhs_p != rhs_p2 {
//             // //         without[rhs_p2] = true;
//             // //     }
//             // // }
//             // rhs_local := to_local(rhs_p, without);
//             rhs_local := to_local(rhs_p, rhs); // local_ps = rhs.ps - rhs.to_local(origo)
//             count := 0;
//             group := []Coord{};
//             for _, p := range lhs_local {
//                 for _, p2 := range rhs_local {
//             // for lhs_p_local, _ := range lhs_local {
//             //     for rhs_p_local, _ := range rhs_local {
//                     if p == p2 {
//                         count++;
//                         // group = append(group, rhs_p.Add(p));
//                         group = append(group, p);
//                     }
//                 }
//             }
//             // arbitrary2 := group[0];
//             // rhs_diff := diff(rhs, group);
//             complement := diff(rhs_local, group); // only want to add those missing
//             unique := []Coord{};
//             for _, p := range complement {
//                 unique = append(unique, lhs_p.Add(p)); // lhs.ps = local_ps + lhs.to_local(origo)
//             }
//             if count >= overlap {
//                 return true, unique;
//             }
//         }
//     }
//     return false, []Coord{};
// }
func has_overlap_helper(lhs, rhs []Coord, overlap int) (bool, []Coord, Coord) {
    for _, lhs_p := range lhs {
        lhs_local := to_local(lhs_p, lhs);
        for _, rhs_p := range rhs {
            rhs_local := to_local(rhs_p, rhs); // local_ps = rhs.ps - rhs.origo
            count := 0;
            group := []Coord{};
            for _, p := range lhs_local {
                for _, p2 := range rhs_local {
                    if p == p2 {
                        count++;
                        group = append(group, p);
                    }
                }
            }
            complement := diff(rhs_local, group); // only want to add those missing
            unique := []Coord{};
            for _, p := range complement {
                unique = append(unique, lhs_p.Add(p)); // lhs.ps = local_ps + lhs.origo
            }
            if count >= overlap {
                sc_local_to_origo := Coord{0,0,0}.Sub(rhs_p);
                sc_in_lhs := lhs_p.Add(sc_local_to_origo)
                return true, unique, sc_in_lhs;
            }
        }
    }
    return false, []Coord{}, Coord{};
}
func to_local(p Coord, rhs []Coord) []Coord {
    local := []Coord{};
// func to_local(p Coord, rhs map[Coord]bool) map[Coord]bool {
//     local := map[Coord]bool{};
    for _, rhs_p := range rhs {
    // for rhs_p, _ := range rhs {
        local = append(local, rhs_p.Sub(p));
    }
    return local;
}

var part2_test_input = []string{
    `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`,
};
var part2_test_output = []string{
    `3621`,
};
func part2(input string) string {
    scanners := parse(input);

    generate_rotations();
    // From AoC: we are guaranteed it's an overlap if _12_ beacons match
    overlap := 12;
    
    lhs := scanners[0];
    candidates := append([]*Scanner{}, scanners[1:]...);
    scs := []Coord{Coord{}};
    for len(candidates) > 0 {
        // pop safely
        rhs := candidates[0];
        candidates = append([]*Scanner{}, candidates[1:]...);

        overlap, unique, sc_in_lhs := has_overlap(lhs, rhs, overlap);
        if overlap {
            scs = append(scs, sc_in_lhs);
            lhs.beacons = append(lhs.beacons, unique...);
        } else {
            candidates = append(candidates, rhs);
        }
        fmt.Println(len(candidates));
    }

    result := 0;
    for i:=0; i < len(scs)-1; i++ {
        for j:=i+1; j < len(scs); j++ {
            lhs := scs[i];
            rhs := scs[j];
            dist := lhs.Sub(rhs).Sum();
            if result < dist {
                result = dist;
            }
        }
    }

    // return strconv.Itoa(len(lhs.beacons));
    return strconv.Itoa(result);
}
func (this Coord) Sum() int {
    return utils.Abs(this.x) + utils.Abs(this.y) + utils.Abs(this.z);
}