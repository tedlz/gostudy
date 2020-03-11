package files_test

import (
	cake "gostudy/08、Goroutines和Channels/files"

	"testing"
	"time"
)

var defaults = cake.Shop{
	Verbose:      testing.Verbose(),
	Cakes:        20,
	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
	// 基线：一位烘焙师，一位上糖师，一位雕花师
	// 每步仅 10ms，无缓冲区
	cakeshop := defaults
	cakeshop.Work(b.N) // 224ms
}

func BenchmarkBuffers(b *testing.B) {
	// 添加缓冲区无影响
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 224ms
}

func BenchmarkVariable(b *testing.B) {
	// 为每个步骤的速率增加可变性
	// 由于 channel 延迟而增加了总时间
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) // 259ms
}

func BenchmarkVariableBuffers(b *testing.B) {
	// 添加 channel 缓冲区减少
	// 变异性造成的延误
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 244ms
}

func BenchmarkSlowIcing(b *testing.B) {
	// 使中间阶段变慢
	// 直接添加到关键路径
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) // 1.032s
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// 添加更多的上糖师，降低上糖的成本
	// 遵循阿姆达尔定律的顺序组件
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) // 288ms
}
