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
    candidates := []IPacket{packet};
    for len(candidates)>0 {
        cur := candidates[len(candidates)-1];
        clone := []IPacket{};
        clone = append(clone, candidates[:len(candidates)-1]...);
        candidates = clone;

        // add subpackets
        candidates = append(candidates, cur.GetArgs()...);

        result += cur.GetVer();
    }

    return strconv.Itoa(result);
}

var extra_test_input = []string{
    `110100101111111000101000`,
    `00111000000000000110111101000101001010010001001000000000`,
    `11101110000000001101010000001100100000100011000001100000`,
};
var extra_test_output = []IPacket{
    &Literal{
        Packet: Packet{ ver:6, args: []IPacket{} },
        n: 2021,
    },
    &Operator{
        Packet{
            ver:1,
            args: []IPacket{
                &Literal{
                    Packet{ ver:6, args: []IPacket{} },
                    10,
                },
                &Literal{
                    Packet{ ver:2, args: []IPacket{} },
                    20,
                },
            },
        },
        &Lt{},
    },
    &Operator{
        Packet{
            ver:7,
            args: []IPacket{
                &Literal{
                    Packet{ ver:2, args: []IPacket{} },
                    1,
                },
                &Literal{
                    Packet{ ver:4, args: []IPacket{} },
                    2,
                },
                &Literal{
                    Packet{ ver:1, args: []IPacket{} },
                    3,
                },
            },
        },
        &Max{},
    },
};
func extra_test() {
    for i, input := range extra_test_input {
        output := extra_test_output[i];

        bit_arr, _ := utils.StrToInt_array(strings.Split(input, ""));
        packet, _ := parse_packet(bit_arr, 0);
        
        if !DeepEquals(packet, output) {
            fmt.Printf("ERROR!!! at test [%d]: '%s'\n%T == %T -> %t\n%v\n%v\n",
                    i,
                    input,
                    packet, output, fmt.Sprintf("%T",packet) != fmt.Sprintf("%T",output),
                    packet,
                    output);
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
func get_literal(bit_arr []int, i int, len int) (int, int, int) {
    prefix := bit_arr[i];
    i++;
    le := len-1; // we already retrieved prefix
    n := 0;
    for j:=i; j < i+le; j++ {
        n = n << 1 + bit_arr[j];
    }
    i+=le;

    switch prefix {
    case 0:
        return n, i, 1;
    case 1:
        n2, i, size := get_literal(bit_arr, i, len);
        shift := le * size;
        n = n << shift + n2;
        return n, i, size+1;
    default:
        fmt.Println("undefined");
    }
    return -1, -1, -1; // this should error hopefully
}
func get_literal2(bit_arr []int, i int, len int) (int, int) {
    n := 0;
    for  {
        prefix := bit_arr[i];
        for j:=1; j < len; j++ { // start j=1 to avoid prefix
            n = n << 1 + bit_arr[i+j];
        }
        i+= len;
        
        if prefix == 0 {
            break;
        }
    }
    return n, i;
}
func conv_Btoi(bit_arr []int, i int, len int) (int, int) {
    n := 0;
    for j:=0; j < len; j++ {
        n = n << 1 + bit_arr[i+j];
    }
    return n, i+len;
}

type IPacket interface {
    GetVer() int
    SetVer(int)
    GetArgs() []IPacket // consider whether this should be moved to IOperator
    SetArgs([]IPacket)
    GetValue() int
}
func DeepEquals(this IPacket, obj IPacket) bool { // can't make it a struct-method cause the struct is an interface ;(
    same_type := fmt.Sprintf("%T",this) == fmt.Sprintf("%T",obj);
    if !same_type {
        fmt.Println("==",fmt.Sprintf("%v  -  %v",this,obj));
        return false;
    }
    same_ver := this.GetVer() == obj.GetVer();
    if !same_ver {
        return false;
    }

    switch t := this.(type) {
    case *Literal:
        return this.(*Literal).n == obj.(*Literal).n;
    case *Operator:
        args, args2 := this.GetArgs(), obj.GetArgs();
        same_len := len(args) == len(args2);
        if !same_len {
            return false;
        }
        all := true;
        for i:=0; i < len(args); i++ {
            all = all && DeepEquals(args[i], args2[i]);
        }
        return all;
    default:
        fmt.Printf("unexpected type %T", t)
    }
    return false;
}

type IOperator interface {
    GetOperator() func(lhs, rhs int) int
};

type Packet struct {
    ver int;
    args []IPacket;
}
func (this *Packet) GetVer() int {
    return this.ver
}
func (this *Packet) SetVer(v int) {
    this.ver = v;
}
func (this *Packet) GetArgs() []IPacket {
   return this.args;
}
func (this *Packet) SetArgs(args []IPacket) {
    this.args = args;
}

// "embedding". Feels like wrapping/composition with some built-in type-checking and forwarding support for interfaces?
// Essentially how you make structs extend other structs in golang
type Operator struct {
    Packet;
    IOperator
};
func (this Operator) String() string {
    return fmt.Sprintf("%T{v:%d, args:%v }", this.IOperator, this.ver, this.args);
}
func (this *Operator) GetValue() int {
    f := this.IOperator.GetOperator();
    // Can't foldl without an initial value, which should be the first arg's value, and then fold over tail
    n := this.args[0].GetValue();
    for _, arg := range this.args[1:] {
        n = f(n, arg.GetValue());
    }
    return n;
}

type Literal struct {
    Packet;

    n int;
};
func (this Literal) String() string {
    return fmt.Sprintf("%T{v:%d, n:%v}", this, this.ver, this.n);
}
func (this *Literal) GetValue() int {
    return this.n;
}

// IOperator implementations
type Sum struct {};
type Product struct {};
type Min struct {};
type Max struct {};
type Gt struct {};
type Lt struct {};
type Eq struct {};
func (p *Sum) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { return lhs + rhs };
}
func (p *Product) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { return lhs * rhs };
}
func (p *Min) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { if lhs < rhs { return lhs } else { return rhs } };
}
func (p *Max) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { if lhs > rhs { return lhs } else { return rhs } };
}
func (p *Gt) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { if lhs > rhs { return 1 } else { return 0 } };
}
func (p *Lt) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { if lhs < rhs { return 1 } else { return 0 } };
}
func (p *Eq) GetOperator() func(lhs, rhs int) int {
    return func(lhs, rhs int) int { if lhs == rhs { return 1 } else { return 0 } };
}

func parse_packet(bit_arr []int, i int) (IPacket, int) {
    ver, i := conv_Btoi(bit_arr, i, 3);
    op, i := conv_Btoi(bit_arr, i, 3);

    packet := Packet{};
    packet.SetVer(ver);

    var operator IOperator;
    switch op {
    case 0: // sum
        operator = &Sum{};
    case 1: // product
        operator = &Product{};
    case 2: // min
        operator = &Min{};
    case 3: // max
        operator = &Max{};
    case 4: // literal
        n, i := get_literal2(bit_arr, i, 5) // literal groups len=5
        return &Literal{packet, n}, i;
    case 5: // gt
        operator = &Gt{};
    case 6: // lt
        operator = &Lt{};
    case 7: // eq
        operator = &Eq{};
    default: // any operator
        fmt.Println("impossible operator!", op);
    }

    subpackets, i := get_subpackets(bit_arr, i);
    packet.SetArgs(subpackets);

    return &Operator{packet, operator}, i;
}

func get_subpackets(packet []int, i int) ([]IPacket, int) {
    length_type_id := packet[i];
    i++;
    subpackets := []IPacket{};
    switch length_type_id {
    case 0: // 'total length in bits' - a 15 bit int
        total_len, new_i := conv_Btoi(packet, i, 15);
        i = new_i;
        for i < new_i + total_len {
            var subpacket IPacket;
            subpacket, i = parse_packet(packet, i);
            subpackets = append(subpackets, subpacket);
        }
    case 1: // 'number of sub-packets immediately contained' - a 11 bit int
        num, new_i := conv_Btoi(packet, i, 11);
        i = new_i;
        for n:=0; n < num; n++ {
            var subpacket IPacket;
            subpacket, i = parse_packet(packet, i);
            subpackets = append(subpackets, subpacket);
        }
    default:
        fmt.Println("impossible length_type_id!");
    }
    return subpackets, i;
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

    result := packet.GetValue();

    return strconv.Itoa(result);
}
