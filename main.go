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
GAME_ROOP:
	for i := 1; ; i++ {
		currentPlayer := &playField.CurrentPlayer
		lastPlayer := &playField.LastPlayer
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
			} else if inputSelect == "2" {
				fmt.Println("監査します。")
				event = "audit"
			}
		// ドローアクション
		case "draw":
			drawedCard := currentPlayer.Draw()
			fmt.Printf("カードをドローしました。%v\n", drawedCard)
			currentPlayer.ToDowncard(drawedCard)
			event = "report"
		// 申告アクション
		case "report":
			if !currentPlayer.Warned {
				fmt.Printf("イエローカードを使用しますか？ 現在の進捗：%d 1:使用 2:不使用\n", playMilestone.GetCurrentPoint())
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
			// 進捗が30を超えていたら最終判定へ
			if playMilestone.GetCurrentPoint() >= 30 {
				event = "judge"
				continue
			}
			playMilestone.RemoveWhiteValid()
			// プレイヤーの状態移動
			tmpPlayer := *currentPlayer
			playField.SetCurrentPlayer(playField.NextPlayer)
			playField.SetNextPlayer(playField.OppPlayer)
			playField.SetOppPlayer(*lastPlayer)
			playField.SetLastPlayer(tmpPlayer)
			event = "standby"
		// 監査アクション
		case "audit":
			sumProgress := playField.ComputeSumProgress()
			fmt.Printf("合計進捗：%d\n", sumProgress)
			// 申告数より小さければレッドカード
			if sumProgress < playMilestone.GetCurrentPoint() {
				fmt.Println("監査成功！")
				if lastPlayer.Suspended {
					fmt.Printf("%sの敗北！他全員の勝利です！\n", lastPlayer.ID)
					break GAME_ROOP
				}
				fmt.Printf("%sにレッドカードが付与されます。\n", lastPlayer.ID)
				lastPlayer.SetSuspend()
			}
			// プレイヤーの状態移動
			tmpPlayer := *currentPlayer
			playField.SetCurrentPlayer(playField.NextPlayer)
			playField.SetNextPlayer(playField.OppPlayer)
			playField.SetOppPlayer(*lastPlayer)
			playField.SetLastPlayer(tmpPlayer)
			event = "standby"
		}

	}
}
