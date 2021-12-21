package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "regexp"
);

/**
  * Start - 16:25:24
  * p1 done - 16:48:04
  * p2 done - 18:08:19
  */

func main() {
    input, _ := utils.Get_input(2021, 21);
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
    `Player 1 starting position: 4
Player 2 starting position: 8`,
};
var part1_test_output = []string{
    `739785`,
};
func part1(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    re := regexp.MustCompile("Player (\\d) starting position: (\\d)");
    players := make([]int, len(lines));
    for _, line := range lines {
        match := re.FindStringSubmatch(line);
        ply, _ := strconv.Atoi(match[1]);
        pos, _ := strconv.Atoi(match[2]);
        players[ply-1] = pos;
    }

    win_criteria := 1000;
    dice_size := 100;
    board_size := 10;

    scores := make([]int, len(players));
    max_score := 0;
    i := 0;
    dice := 1;
    for max_score < win_criteria {
        player := i % len(players);

        roll := dice*3 + 1 + 2; // three rolls
        dice += 3;
        dice = ((dice-1) % dice_size) + 1;

        n := ((players[player] + roll - 1) % board_size) + 1;
        score := scores[player] + n
        scores[player] = score;
        players[player] = n;
        if max_score < score {
            max_score = score;
        }

        i++;
    }

    min := max_score;
    for _, score := range scores {
        if min > score {
            min = score;
        }
    }

    result := i*3 * min;

    return strconv.Itoa(result);
}

var part2_test_input = []string{
    `Player 1 starting position: 4
Player 2 starting position: 8`,
};
var part2_test_output = []string{
    `444356092776315`,
};
func part2(input string) string {
    lines := utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    
    re := regexp.MustCompile("Player (\\d) starting position: (\\d)");
    players := make([]int, len(lines));
    for _, line := range lines {
        match := re.FindStringSubmatch(line);
        ply, _ := strconv.Atoi(match[1]);
        pos, _ := strconv.Atoi(match[2]);
        players[ply-1] = pos;
    }

    win_criteria := 21;
    dice_size := 3;
    num_rolls := 3; // unused
    board_size := 10;

    scores := make([]int, len(players));
    ply_wrap := wrap(0,len(players), 1);
    board_wrap := wrap(1,board_size, 11);
    dirac_dice := dirac_dice_rolls(num_rolls, rang(dice_size));
    fmt.Println(ply_wrap, board_wrap, dirac_dice)
    wins := iterate(0, 0, 0, win_criteria, players, scores, ply_wrap, board_wrap, dirac_dice);
    fmt.Println(wins)

    result := 0;
    for _, w := range wins {
        if result < w {
            result = w;
        }
    }

    return strconv.Itoa(result);
}

func iterate(ply, max_player, max_score, win_criteria int, players, scores, ply_wrap, board_wrap []int, dirac_dice map[int]int) []int {
    wins := make([]int, len(players));
    if max_score >= win_criteria { // base-case
        wins[max_player] = 1;
        return wins;
    }

    next_ply := ply_wrap[ply+1];
    pos := players[ply];
    score := scores[ply];
    for r, num := range dirac_dice {
        new_pos := board_wrap[pos + r];

        score := score + new_pos;

        // shadows the func args
        players := append([]int{}, players...); // clone the arrays
        scores := append([]int{}, scores...);
        max_ply := max_player; 
        max_score := max_score;

        players[ply] = new_pos;
        scores[ply] = score;
        if max_score < score {
            max_ply = ply;
            max_score = score;
        }

        sub_wins := iterate(next_ply, max_ply, max_score, win_criteria, players, scores, ply_wrap, board_wrap, dirac_dice);

        // add the 3 vectors together
        for i, w := range sub_wins {
            wins[i] += w * num;
        }
    }
    return wins;
}

func wrap(min, size, margin int) []int {
    arr := []int{};
    max := size + min;
    for i:=0; i < max+margin; i++ {
        if i < min { // useless case
            delta := min-i;
            arr = append(arr, max-delta);
        } else {
            arr = append(arr, min + (i-min) % size);
        }
    }
    return arr;
}

func dirac_dice_rolls(num_rolls int, rang []int) map[int]int {
    rolls := map[int]int{};
    for _, r1 := range rang {
        for _, r2 := range rang {
            for _, r3 := range rang {
                r := r1+r2+r3;
                rolls[r]++;
            }
        }
    }
    return rolls;
}

func rang(size int) []int {
    r := make([]int, size);
    for i:=0; i < size; i++ {
        r[i] = i+1;
    }
    return r;
}