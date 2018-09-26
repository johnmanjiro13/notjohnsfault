package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/johnmanjiro13/notjohnsfault/game/card"
	"github.com/johnmanjiro13/notjohnsfault/game/deck"
	"github.com/johnmanjiro13/notjohnsfault/game/discard"
	"github.com/johnmanjiro13/notjohnsfault/game/downcard"
	"github.com/johnmanjiro13/notjohnsfault/game/field"
	"github.com/johnmanjiro13/notjohnsfault/game/milestone"
	"github.com/johnmanjiro13/notjohnsfault/game/player"
	"github.com/johnmanjiro13/notjohnsfault/util"
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
		"Player1",
		playDeck,
		playDowncard,
	)
	p2 := player.NewPlayer(
		"Player2",
		playDeck,
		playDowncard,
	)
	p3 := player.NewPlayer(
		"Player3",
		playDeck,
		playDowncard,
	)
	p4 := player.NewPlayer(
		"Player4",
		playDeck,
		playDowncard,
	)
	// フィールドに山札、伏せ札、捨て札を追加
	playField := field.NewField(playDeck, playDowncard, playDiscard, *p1, *p2, *p3, *p4)

	event := "reset"
	isFirst := true
GAME_ROOP:
	for i := 1; ; i++ {
		currentPlayer := &playField.CurrentPlayer
		lastPlayer := &playField.LastPlayer
		switch event {
		// ゲーム状態の初期化
		case "reset":
			playMilestone.ResetCurrentPoint()
			playField.ResetYellowCards()
			event = "standby"
		// スタンバイフェイズ
		case "standby":
			fmt.Printf("GM：現在のプレイヤーは%sです。", currentPlayer.ID)
			fmt.Println("申告しますか？　監査しますか？ 1:申告 2:監査")
			inputSelect := nextLine()
			if isFirst && (inputSelect == "2") {
				fmt.Println("初回は監査できません。")
				continue
			}
			if inputSelect == "1" {
				fmt.Println("GM：申告します。")
				event = "draw"
			} else if inputSelect == "2" {
				fmt.Println("GM：監査！")
				event = "judge"
			}
			isFirst = false
		// ドローアクション
		case "draw":
			drawedCard := currentPlayer.Draw()
			fmt.Printf("GM：カードをドローしました。%v\n", drawedCard)
			currentPlayer.ToDowncard(drawedCard)
			if playField.GetDeckLength() == 0 {
				playField.DiscardToDeck()
			}
			event = "report"
		// 申告アクション
		case "report":
			if !currentPlayer.Warned {
				fmt.Printf("GM：イエローカードを使用しますか？ 現在の進捗：%d 1:使用 2:不使用\n", playMilestone.GetCurrentPoint())
				if nextLine() == "1" {
					currentPlayer.UseYellowCard()
					playMilestone.SetWhiteValid()
					fmt.Println("戒告！")
				}
			}
			fmt.Printf("GM：進捗を申告してください。現在の進捗：%d\n", playMilestone.GetCurrentPoint())
			event = "milestone"
		// 数字選択アクション
		case "milestone":
			inputNum, _ := strconv.Atoi(nextLine())
			if inputNum <= playMilestone.GetCurrentPoint() {
				fmt.Println("GM：現在の進捗より大きな数字を申告してください。")
				continue
			}
			err := playMilestone.SetCurrentPoint(inputNum)
			if err != nil {
				fmt.Println("GM：無効な数字です。もう一度申告してください。")
				continue
			}
			fmt.Printf("GM：進捗を報告しました。%d\n", playMilestone.GetCurrentPoint())
			// 進捗が30を超えていたら最終判定へ
			if playMilestone.GetCurrentPoint() >= 30 {
				event = "judge"
				continue
			}
			playMilestone.RemoveWhiteValid()
			// プレイヤーの状態移動
			util.TransitState(playField)
			event = "standby"
		// 監査アクション
		case "judge":
			sumProgress := playField.ComputeSumProgress()
			fmt.Printf("GM：合計進捗：%d\n", sumProgress)
			// 申告数より小さければレッドカード（最終申告は30）
			if sumProgress < playMilestone.GetCurrentPoint() {
				if playMilestone.GetCurrentPoint() != 30 {
					fmt.Println("GM：監査成功！")
				} else if playMilestone.GetCurrentPoint() == 30 {
					fmt.Println("GM：目標未達！")
				}

				if lastPlayer.Suspended {
					fmt.Printf("GM：%sの敗北！他全員の勝利です！\n", lastPlayer.ID)
					fmt.Println("GM：もう一度プレイしますか？ 1：はい 2：いいえ")
					if nextLine() == "1" {
						playField.ResetAllCards()
						event = "reset"
						isFirst = true
						continue
					}
					break GAME_ROOP
				}
				fmt.Printf("GM：%sにレッドカードが付与されます。\n", lastPlayer.ID)
				lastPlayer.SetSuspend()
			}
			// 申告数より大きければ監査した側にレッドカード
			if sumProgress >= playMilestone.GetCurrentPoint() {
				fmt.Println("GM：監査失敗！")
				if currentPlayer.Suspended {
					fmt.Printf("GM：%sの敗北！他全員の勝利です！\n", currentPlayer.ID)
					fmt.Println("GM：もう一度プレイしますか？ 1：はい 2：いいえ")
					if nextLine() == "1" {
						playField.ResetAllCards()
						event = "reset"
						isFirst = true
						continue
					}
					break GAME_ROOP
				}
				fmt.Printf("GM：%sにレッドカードが付与されます。\n", currentPlayer.ID)
				currentPlayer.SetSuspend()
			}
			// 30より大きければゲーム終了
			if sumProgress >= 30 {
				fmt.Println("GM：目標達成！")
				fmt.Printf("GM：%sの勝利！他全員の敗北です！\n", currentPlayer.ID)
				fmt.Println("GM：もう一度プレイしますか？ 1：はい 2：いいえ")
				if nextLine() == "1" {
					playField.ResetAllCards()
					event = "reset"
					isFirst = true
					continue
				}
				break GAME_ROOP
			}
			playField.DowncardToDiscard()
			// プレイヤーの状態移動
			util.TransitState(playField)
			isFirst = true
			event = "reset"
		}
	}
}
