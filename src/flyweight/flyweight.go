package main

import (
	"fmt"
	"sync"
)

// 享元对象
type Card struct {
	Name  string
	Color string
}

// 享元工厂（管理 Card 复用）
type CardFactory struct {
	cards map[string]*Card
	mu    sync.Mutex
}

func NewCardFactory() *CardFactory {
	return &CardFactory{cards: make(map[string]*Card)}
}

func (f *CardFactory) GetCard(name, color string) *Card {
	key := name + "_" + color
	f.mu.Lock()
	defer f.mu.Unlock()

	// 如果已存在该 Card，则返回共享对象
	if card, exists := f.cards[key]; exists {
		return card
	}

	// 否则创建新 Card 并存入池中
	card := &Card{Name: name, Color: color}
	f.cards[key] = card
	return card
}

// 业务结构体
type PokerGame struct {
	Cards map[int]*Card
}

// PokerGame 通过工厂获取共享的 Card
func NewPokerGame(factory *CardFactory) *PokerGame {
	return &PokerGame{
		Cards: map[int]*Card{
			1: factory.GetCard("A", "紅"),
			2: factory.GetCard("A", "黑"),
		},
	}
}

func main() {
	// 创建享元工厂
	factory := NewCardFactory()

	// 创建两个 PokerGame 实例
	game1 := NewPokerGame(factory)
	game2 := NewPokerGame(factory)

	// 由于 Card 是共享的，所以指针相等
	fmt.Println(game1.Cards[1] == game2.Cards[1]) // 输出 true
	fmt.Println(game1.Cards[2] == game2.Cards[2]) // 输出 true
}
