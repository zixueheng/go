package main

import (
	"fmt"
	"math"
)

// 二维矢量模拟玩家移动
// 在游戏中，一般使用二维矢量保存玩家的位置，使用矢量运算可以计算出玩家移动的位置，本例子中，首先实现二维矢量对象，接着构造玩家对象，最后使用矢量对象和玩家对象共同模拟玩家移动的过程。

// 1) 实现二维矢量结构

// Vec2 二维矢量拥有两个方向的信息
type Vec2 struct {
	X, Y float32
}

// Add 加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// Sub 减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// Scale 乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// DistanceTo 距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// Normalize 插值
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}

// 2) 实现玩家对象
// 玩家对象负责存储玩家的当前位置、目标位置和速度，使用 MoveTo() 方法为玩家设定移动的目标，使用 Update() 方法更新玩家位置，在 Update() 方法中，通过一系列的矢量计算获得玩家移动后的新位置。

// Player 玩家对象
type Player struct {
	currPos   Vec2    // 当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// MoveTo 移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// Pos 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// IsArrived 是否到达
func (p *Player) IsArrived() bool {
	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// Update 逻辑更新
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// 移动完成后，更新当前位置
		p.currPos = newPos
	}
}

// NewPlayer 创建新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}

func main() {
	// 将 Player 实例化后，设定玩家移动的最终目标点，之后开始进行移动的过程，这是一个不断更新位置的循环过程，每次检测玩家是否靠近目标点附近，
	// 如果还没有到达，则不断地更新位置，让玩家朝着目标点不停的修改当前位置

	// 实例化玩家对象，并设速度为0.5
	p := NewPlayer(0.5)
	// 让玩家移动到3,1点
	p.MoveTo(Vec2{3, 1})
	// 如果没有到达就一直循环
	for !p.IsArrived() {
		// 更新玩家位置
		p.Update()
		// 打印每次移动后的玩家位置
		fmt.Println(p.Pos())
	}
}
