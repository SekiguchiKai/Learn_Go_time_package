package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 遡る日付をyyyy-mmで入力を受け付ける
	fmt.Println("いつまで遡りますか?\nyyyy-mmで入力してください")

	var input string

	// 適切な入力がなされない限りループ
	for {
		// 標準入力をスキャン
		fmt.Scan(&input)

		// 正規表現のチェック
		if b, err := regexp.MatchString(`^(\d{4})-(0[1-9]|1[0-2])$`, input); !b || err != nil {
			fmt.Println("入力いただいた文字列が不適切です。 yyyy-mmで入力してください")
		} else {
			// 適切な値が入力されたら、ループを抜ける
			break
		}

	}

	// 入力された文字列を"-"で分割
	s := strings.Split(input, "-")
	// 遡る西暦の限界値
	limitYear, err := strconv.Atoi(s[0])
	if err != nil {
		log.Printf("can not strconv.Atoi :%v \n", err)
	}

	// 遡る月の限界値
	limitMonth, err := strconv.Atoi(s[1])
	if err != nil {
		log.Printf("can not strconv.Atoi :%v \n", err)
	}

	ShowSpecificTermList(limitYear, limitMonth)
}

// 現在の時刻から指定した日付まで遡った月初と月末の日付の一覧を表示する
func ShowSpecificTermList(limitYear, limitMonth int) {
	// 現在の時刻を取得
	now := time.Now()
	// 基準となる日付を設定
	criterionDate := time.Date(now.Year(), now.Month(), 1, 12, 0, 0, 0, time.UTC)

	// 終了の条件となる日付を設定
	finDate := time.Date(limitYear, time.Month(limitMonth), 30, 15, 0, 0, 0, time.UTC)

	i := 0
	for {
		// 月初を設定
		beginningOfTheMonth := criterionDate.AddDate(0, -i, 0)
		// 月末を設定
		endOfTheMonth := criterionDate.AddDate(0, -i+1, -1)

		fmt.Printf("月初:%v", beginningOfTheMonth)
		fmt.Printf("月末:%v\n", endOfTheMonth)

		// 終了条件の日付以前の日にちまで遡ったら、breakしてループを抜ける
		if beginningOfTheMonth.Before(finDate) {
			fmt.Println("終了")
			break
		}

		// インクリメント
		i++
	}
}
