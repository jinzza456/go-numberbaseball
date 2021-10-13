package main

import (
	"fmt"
)

func main() {
	numbers := MakeNumbers() // 무작위 숫자 3개를 만든다.

	cnt := 0

	for {
		cnt++
		inputNumbers := InputNumbers() // 사용자의 입력을 받는다.

		result := CompareNumbers(numbers, inputNumbers) // 결과를 비교한다.

		PrintResult(result)

		if IsGameEnd(result) {
			break
		} // 3S 일경우 게임 끝}
	}
	fmt.Printf("%d번 만에 맞췄습니다.", cnt)
	//몇번만에 맞췄는지 출력
}

func MakeNumbers() [3]int {
	var rst [3]int
	return rst
} // 0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환

func InputNumbers() [3]int {
	var rst [3]int
	return rst
} // 키보드로 부터 0~9 사이의 겹치지 않는 숮 3개를 입력받아 반환

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	return true
} // 두개의 숫자 3개를 비교해서 결과를 반환

func PrintResult(result bool) {
	fmt.Println(result)
} // 스트라이크 볼 을 판단하는 결과를 반환

func IsGameEnd(result bool) bool {
	return result
} // 비교 결과가 3스트라이크 인지 확인
