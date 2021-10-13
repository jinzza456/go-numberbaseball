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
