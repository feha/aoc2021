package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 16:27:46
  * p1 done - 18:29:19
  * p2 done - 19:59:51
  */

func main() {
    extra_test();
    input, _ := utils.Get_input(2021, 16);
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

const separator string = "";

var part1_test_input = []string{
    `8A004A801A8002F478`,
    `620080001611562C8802118E34`,
    `C0015000016115A2E0802F182340`,
    `A0016C880162017C3686B18A3D4780`,
};
var part1_test_output = []string{
    `16`,
    `12`,
    `23`,
    `31`,
};
func part1(input string) string {
    input = strings.Trim(input, " \n");
    
    bit_arr := parse_hex(input);

    packet, _ := parse_packet(bit_arr, 0);

    result := 0;
    candidates := []Packet{packet};
    for len(candidates)>0 {
        cur := candidates[len(candidates)-1];
        clone := []Packet{};
        clone = append(clone, candidates[:len(candidates)-1]...);
        candidates = clone;

        // add subpackets
        candidates = append(candidates, cur.args...);

        result += cur.ver;
    }

    return strconv.Itoa(result);
}

var extra_test_input = []string{
    `110100101111111000101000`,
    `00111000000000000110111101000101001010010001001000000000`,
    `11101110000000001101010000001100100000100011000001100000`,
};
var extra_test_output = []Packet{
    Packet{
        ver:6,
        op: 4,
        n: 2021,
        args: []Packet{},
    },
    Packet{
        ver:1,
        op: 6,
        n: -1,
        args: []Packet{
            Packet{
                ver:6,
                op: 4,
                n: 10,
                args: []Packet{},
            },
            Packet{
                ver:2,
                op: 4,
                n: 20,
                args: []Packet{},
            },
        },
    },
    Packet{
        ver:7,
        op: 3,
        n: -1,
        args: []Packet{
            Packet{
                ver:2,
                op: 4,
                n: 1,
                args: []Packet{},
            },
            Packet{
                ver:4,
                op: 4,
                n: 2,
                args: []Packet{},
            },
            Packet{
                ver:1,
                op: 4,
                n: 3,
                args: []Packet{},
            },
        },
    },
};
func extra_test() {
    for i, input := range extra_test_input {
        output := extra_test_output[i];

        bit_arr, _ := utils.StrToInt_array(strings.Split(input, ""));
        packet, _ := parse_packet(bit_arr, 0);
        
        candidates := []Packet{packet};
        candidates2 := []Packet{output};
        for len(candidates)>0 {
            cur := candidates[len(candidates)-1];
            clone := []Packet{};
            clone = append(clone, candidates[:len(candidates)-1]...);
            candidates = append(clone, cur.args...);

            cur_out := candidates2[len(candidates2)-1];
            clone_out := []Packet{};
            clone_out = append(clone_out, candidates2[:len(candidates2)-1]...);
            candidates2 = append(clone_out, cur_out.args...);

            if cur.ver != cur_out.ver ||
                cur.op != cur_out.op || 
                cur.n != cur_out.n || 
                len(cur.args) != len(cur_out.args) {
                fmt.Println("ERROR!!! at test", i, input, packet, output);
            }
        }
    }
}

func parse_hex(hex string) []int {
    bit_arr := []int{};
    for _, c := range strings.Split(hex, "") {
        n, err := strconv.ParseInt(c, 16, 64);
        if err != nil {
            fmt.Println("err:", err);
        }
        hex_as_bit := []int{};
        // each hex is 4 bits (cant just check `n > 0` as we want to preserve 'padded' zeroes)
        for i:=0; i < 4; i++ {
            hex_as_bit = append(hex_as_bit, int(n) & 1);
            n = n >> 1;
        }
        // reverse as we parsed wrong dir (LSB first for each hex)
        for i, j := 0, len(hex_as_bit)-1; i < j; i, j = i+1, j-1 {
            hex_as_bit[i], hex_as_bit[j] = hex_as_bit[j], hex_as_bit[i];
        }
        bit_arr = append(bit_arr, hex_as_bit...);
    }

    if len(bit_arr) != len(hex)*4 {
        fmt.Println("ERROR!", hex, len(bit_arr), len(hex)*4);
    }
    return bit_arr;
}

// len = 5
func get_literal(packet []int, i int, len int) (int, int, int) {
    prefix := packet[i];
    i++;
    le := len-1; // we already retrieved prefix
    n := 0;
    for j:=i; j < i+le; j++ {
        n = n << 1 + packet[j];
    }
    i+=le;

    switch prefix {
    case 0:
        return n, i, 1;
    case 1:
        n2, i, size := get_literal(packet, i, len);
        shift := le * size;
        n = n << shift + n2;
        return n, i, size+1;
    default:
        fmt.Println("undefined");
    }
    return -1, -1, -1; // this should error hopefully
}
func conv_Btoi(packet []int, i int, len int) (int, int) {
    n := 0;
    for j:=0; j < len; j++ {
        n = n << 1 + packet[i+j];
    }
    return n, i+len;
}
func get_subpackets(packet []int, i int) ([]Packet, int) {
    length_type_id := packet[i];
    i++;
    subpackets := []Packet{};
    switch length_type_id {
    case 0: // 'total length in bits' - a 15 bit int
        total_len, new_i := conv_Btoi(packet, i, 15);
        i = new_i;
        for i < new_i + total_len {
            var subpacket Packet;
            subpacket, i = parse_packet(packet, i);
            subpackets = append(subpackets, subpacket);
        }
    case 1: // 'number of sub-packets immediately contained' - a 11 bit int
        num, new_i := conv_Btoi(packet, i, 11);
        i = new_i;
        for n:=0; n < num; n++ {
            var subpacket Packet;
            subpacket, i = parse_packet(packet, i);
            subpackets = append(subpackets, subpacket);
        }
    default:
        fmt.Println("impossible length_type_id!");
    }
    return subpackets, i;
}

func parse_packet(packet []int, i int) (Packet, int) {
    ver, i := conv_Btoi(packet, i, 3);
    op, i := conv_Btoi(packet, i, 3);

    p := Packet{ver: ver, op: op};
    p.f = get_operation(p.op);
    p.n = -1; // impossible value, lets us foldl with sensible init values (or in practice, return rhs for the first iteration) - check get_operator()
    switch op {
    case 4: // literal
        n, new_i, _ := get_literal(packet, i, 5) // literal groups len=5
        i = new_i;
        p.n = n;
    default: // any operator
        subpackets, new_i := get_subpackets(packet, i);
        i = new_i;
        p.args = subpackets;
    }
    return p, i;
}

type Packet struct {
    ver int;
    op int;
    n int;
    args []Packet;
    f func(lhs, rhs int) int;
}
func (p Packet) String() string {
    return fmt.Sprintf("{ ver:%d op:%d n:%d args:%v }", p.ver, p.op, p.n, p.args);
}

var part2_test_input = []string{
    `C200B40A82`,
    `04005AC33890`,
    `880086C3E88112`,
    `CE00C43D881120`,
    `D8005AC2A8F0`,
    `F600BC2D8F`,
    `9C005AC2F8F0`,
    `9C0141080250320F1802104A08`,
};
var part2_test_output = []string{
    `3`,
    `54`,
    `7`,
    `9`,
    `1`,
    `0`,
    `0`,
    `1`,
};
func part2(input string) string {
    input = strings.Trim(input, " \n");
    
    bit_arr := parse_hex(input);

    packet, _ := parse_packet(bit_arr, 0);

    packet = execute_packet(packet);

    result := packet.n;

    return strconv.Itoa(result);
}

func execute_packet(p Packet) Packet {
    for _, p2 := range p.args {
        p2 = execute_packet(p2);
        p.n = p.f(p.n, p2.n);
    }
    return p
}

func get_operation(op int) func(lhs int, rhs int) int {
    switch op {
    case 0: // sum
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; return lhs + rhs };
    case 1: // product
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; return lhs * rhs };
    case 2: // min
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; if lhs < rhs { return lhs } else { return rhs } };
    case 3: // max
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; if lhs > rhs{ return lhs } else { return rhs } };
    case 4: // literal
        return func(lhs, rhs int) int { return lhs };
    case 5: // gt
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; if lhs > rhs { return 1 } else { return 0 } };
    case 6: // lt
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; if lhs < rhs { return 1 } else { return 0 } };
    case 7: // eq
        return func(lhs, rhs int) int { if lhs < 0 {return rhs}; if lhs == rhs { return 1 } else { return 0 } };
    default: // any operator
        fmt.Println("impossible operator!", op);
    }
    return nil;
}