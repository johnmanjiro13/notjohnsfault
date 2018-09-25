package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/johnmanjiro13/notjohnsfault/player/card"
	"github.com/johnmanjiro13/notjohnsfault/player/deck"
	"github.com/johnmanjiro13/notjohnsfault/player/discard"
	"github.com/johnmanjiro13/notjohnsfault/player/downcard"
	"github.com/johnmanjiro13/notjohnsfault/player/field"
	"github.com/johnmanjiro13/notjohnsfault/player/milestone"
	"github.com/johnmanjiro13/notjohnsfault/player/player"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	// 各カード要素の初期化
	cards := []card.ICard{}
	for i := 0; i < 32; i++ {
		switch {
		case (i >= 0 && i < 2):
			newCard := card.Card{Number: 0}
			cards = append(cards, newCard)
		case (i >= 2 && i < 6):
			newCard := card.Card{Number: 1}
			cards = append(cards, newCard)
		case (i >= 6 && i < 11):
			newCard := card.Card{Number: 2}
			cards = append(cards, newCard)
		case (i >= 11 && i < 16):
			newCard := card.Card{Number: 3}
			cards = append(cards, newCard)
		case (i >= 16 && i < 21):
			newCard := card.Card{Number: 4}
			cards = append(cards, newCard)
		case (i >= 21 && i < 26):
			newCard := card.Card{Number: 5}
			cards = append(cards, newCard)
		case (i >= 26 && i < 31):
			newCard := card.Card{Number: 6}
			cards = append(cards, newCard)
		}
	}
	playDeck := deck.NewDeck(cards)
	playDowncard := downcard.NewDowncard()
	playDiscard := discard.NewDiscard()

	// マイルストーン
	playMilestone := milestone.NewMilestone()

	// プレイヤーの追加
	p1 := player.NewPlayer(
		"p1",
		playDeck,
		playDowncard,
	)
	p2 := player.NewPlayer(
		"p2",
		playDeck,
		playDowncard,
	)
	p3 := player.NewPlayer(
		"p3",
		playDeck,
		playDowncard,
	)
	p4 := player.NewPlayer(
		"p4",
		playDeck,
		playDowncard,
	)
	// フィールドに山札、伏せ札、捨て札を追加
	playField := field.NewField(playDeck, playDowncard, playDiscard, *p1, *p2, *p3, *p4)

	event := "standby"
	for i := 1; ; i++ {
		currentPlayer := &playField.CurrentPlayer
		switch event {
		// スタンバイフェイズ
		case "standby":
			fmt.Printf("現在のプレイヤーは%sです。", currentPlayer.ID)
			fmt.Println("申告しますか？　監査しますか？ 1:申告 2:監査")
			inputSelect := nextLine()
			if (i == 1) && (inputSelect == "2") {
				fmt.Println("初回は監査できません。")
				continue
			}
			if inputSelect == "1" {
				fmt.Println("申告します。")
				event = "draw"
			}
		// ドローアクション
		case "draw":
			drawedCard := currentPlayer.Draw()
			fmt.Printf("カードをドローしました。%v\n", drawedCard)
			event = "report"
		// 申告アクション
		case "report":
			if !currentPlayer.Warned {
				fmt.Println("イエローカードを使用しますか？ 1:使用 2:不使用")
				if nextLine() == "1" {
					currentPlayer.UseYellowCard()
					playMilestone.SetWhiteValid()
				}
			}
			fmt.Printf("進捗を申告してください。現在の進捗：%d\n", playMilestone.GetCurrentPoint())
			event = "milestone"
		// 数字選択アクション
		case "milestone":
			inputNum, _ := strconv.Atoi(nextLine())
			if inputNum <= playMilestone.GetCurrentPoint() {
				fmt.Println("現在の進捗より大きな数字を申告してください。")
				continue
			}
			err := playMilestone.SetCurrentPoint(inputNum)
			if err != nil {
				fmt.Println("無効な数字です。もう一度申告してください。")
				continue
			}
			fmt.Printf("進捗を報告しました。%d\n", playMilestone.GetCurrentPoint())
			playMilestone.RemoveWhiteValid()
			// プレイヤーの状態移動
			tmpPlayer := *currentPlayer
			playField.SetCurrentPlayer(playField.NextPlayer)
			playField.SetNextPlayer(playField.OppPlayer)
			playField.SetOppPlayer(playField.LastPlayer)
			playField.SetLastPlayer(tmpPlayer)
			event = "standby"
		}
	}
}
