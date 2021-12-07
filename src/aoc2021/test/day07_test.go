package test;

import (
    "aoc/libs/utils"
    "aoc/src/aoc2021/test"
    "strings"
    "testing"
);

const separator string = ",";
var input, _ = utils.Get_input(2021, 07);
var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator+"\n"), separator));
var nums, _ = utils.StrToInt_array(inputs);

// func Benchmark1(b *testing.B) {
//     for n := 0; n < b.N; n++ {
//         test.Calculate_fuels1(nums)
//     }
// }
// func Benchmark2(b *testing.B) {
//     for n := 0; n < b.N; n++ {
//         test.Calculate_fuels2(nums)
//     }
// }
// func Benchmark_oldest(b *testing.B) {
//     for n := 0; n < b.N; n++ {
//         test.Calculate_fuels_old(nums)
//     }
// }
// func Benchmark_TriNum_Normalized_old(b *testing.B) {
//     for n := 0; n < b.N; n++ {
//         test.Calculate_fuels_Trinum_Normalized_old(nums)
//     }
// }
func Benchmark_TriNum_Normalized(b *testing.B) {
    for n := 0; n < b.N; n++ {
        test.Calculate_fuels_Trinum_Normalized(nums)
    }
}
// func Benchmark_Trinum_Normalized_Hist(b *testing.B) {
//     for n := 0; n < b.N; n++ {
//         test.Calculate_fuels_Trinum_Normalized_Hist(nums)
//     }
// }
func Benchmark_Avg(b *testing.B) {
    for n := 0; n < b.N; n++ {
        test.Calculate_fuels_Avg(nums)
    }
}
func Benchmark_Avg_Hist(b *testing.B) {
    for n := 0; n < b.N; n++ {
        test.Calculate_fuels_Avg_Hist(nums)
    }
}