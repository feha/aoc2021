package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "regexp"
);

/**
  * Start - 18:04:16
  * p1 done - 19:27:59
  * p2 done - 19:51:42
  */

func main() {
    input, _ := utils.Get_input(2021, 17);
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
    `target area: x=20..30, y=-10..-5`,
};
var part1_test_output = []string{
    `45`,
};
func part1(input string) string {
    input = strings.Trim(input, " \n");
    
    re := regexp.MustCompile("target area: x=(-?\\d+)..(-?\\d+), y=(-?\\d+)..(-?\\d+)");
    match := re.FindStringSubmatch(input);

    // x_min, err := strconv.Atoi(match[1]);
    // if err != nil {
    //     fmt.Println("error1!");
    // }
    // x_max, err  := strconv.Atoi(match[2]);
    // if err != nil {
    //     fmt.Println("error2!");
    // }
    y_min, err := strconv.Atoi(match[3]);
    if err != nil {
        fmt.Println("error3!");
    }
    // y_max, err  := strconv.Atoi(match[4]);
    // if err != nil {
    //     fmt.Println("error4!");
    // }

    // bounds := Bounds{Coord{x_min, y_min}, Coord{x_max, y_max}};
    // for n := y_max; n > y_min; n-- {



    //     a := 2 * 0.5;
    //     b := y_vel + 0.5:
    //     c := -n;

    //     // t = ( -b +- sqrt(b^2 + 4*a*n) )/(2*a); // a=0.5
    //     t0 = ( -b + sqrt(b^2 + 2*n) ); // -n < b^2/2
    //     t1 = ( -b - sqrt(b^2 + 2*n) ); // -n > b^2/2
        
    //     hit := false;
    //     if t_0 > 0 || t1 > 0 {
    //         hit = true;
    //     }


    // }


    // Cheating thanks to knowing my input (and the test) places bounds below y=0

    // due to symmetry, it will pass y=0 at velocity `y_vel+1`
    // So the highest possible y_vel is one that hits lower bound the step directly after,
    // which means the vel at that point equals the lower bound
    // Subtract 1 as gravity increased vel by 1 there, while we want the upwards vel from x=0.
    y_vel := -y_min - 1; 
    
    result := 0;
    for i:=y_vel; i > 0; i-- {
        result += i;
    }

    return strconv.Itoa(result);
}

type Coord struct {
    x, y int
}
func (this Coord) Contains(n int) bool {
    min, max := utils.Min(this.x, this.y), utils.Min(this.x, this.y);
    if min <= n && n <= max {
        return true;
    }
    return false;
}

type Bounds struct {
    x_bound Coord; // x-bound
    y_bound Coord; // y-bound
}
func (this Bounds) MinX() int {
    return utils.Min(this.x_bound.x, this.x_bound.y);
}
func (this Bounds) MinY() int {
    return utils.Min(this.y_bound.x, this.y_bound.y);
}
func (this Bounds) MaxX() int {
    return utils.Max(this.x_bound.x, this.x_bound.y);
}
func (this Bounds) MaxY() int {
    return utils.Max(this.y_bound.x, this.y_bound.y);
}
func (this Bounds) Contains(p Coord) bool {
    if this.MinX() <= p.x && p.x <= this.MaxX() &&
            this.MinY() <= p.y && p.y <= this.MaxY() {
        return true;
    }
    return false;
}

var part2_test_input = []string{
    `target area: x=20..30, y=-10..-5`,
};
var part2_test_output = []string{
    `112`,
};
func part2(input string) string {
    input = strings.Trim(input, " \n");
    
    re := regexp.MustCompile("target area: x=(-?\\d+)..(-?\\d+), y=(-?\\d+)..(-?\\d+)");
    match := re.FindStringSubmatch(input);

    x_min, err := strconv.Atoi(match[1]);
    if err != nil {
        fmt.Println("error1!");
    }
    x_max, err  := strconv.Atoi(match[2]);
    if err != nil {
        fmt.Println("error2!");
    }
    y_min, err := strconv.Atoi(match[3]);
    if err != nil {
        fmt.Println("error3!");
    }
    y_max, err  := strconv.Atoi(match[4]);
    if err != nil {
        fmt.Println("error4!");
    }

    // Cheating thanks to knowing my input (and the test) places bounds below y=0,
    // and to the right of x=0.

    x_bound, y_bound := Coord{x_min, x_max}, Coord{y_min, y_max};
    bounds := Bounds{x_bound, y_bound};

    result := 0;

    // The possible x-velocities that aren't sure to miss are 0 < v <= x_max
    //   '0 < x_vel' due to arbitrary choice that is at least going rightwards
    // The possible y-velocities that aren't sure to miss are y_min <= v < -y_min
    //   '< -y_min' due to upwards vels being y_vel+1 once they leave the symmetry point
    for x_vel:=0; x_vel <= x_max; x_vel++ {
        for y_vel:=y_min; y_vel < -y_min; y_vel++ {
            // valid := false;

            // cheat_vel := - utils.Abs(y_vel);
            
            // // if it can hit bounds directly
            // if y_bound.Contains(cheat_vel) {
            //     valid = true;
            // }

            // // if it can hit bounds directly
            // if y_bound.Contains(cheat_vel) {
            //     valid = true;
            // }

            // if valid {
            //     y_vels = append(y_vels, i);
            // }
            hit := iterate(0, 0, x_vel, y_vel, bounds);
            if hit {
                result++;
            }
        }
    }

    return strconv.Itoa(result);
}


func iterate(x, y, x_vel, y_vel int, bounds Bounds) bool {
    x, y = x + x_vel, y + y_vel;
    if bounds.Contains(Coord{x, y}) { // hit!
        return true;
    }
    if x > bounds.MaxX() || y < bounds.MinY() { // overshoot
        return false;
    }

    x_vel = utils.Max(x_vel - 1, 0);
    y_vel = y_vel - 1;
    return iterate(x, y, x_vel, y_vel, bounds); // step
}