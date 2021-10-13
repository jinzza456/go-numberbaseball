package main

import (
	"fmt"
	"math/rand" // 랜덤 숫자를 위한 패키지 호출
	"time"      //항상 변하는 seed 값을 주기위해 time을 호출
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 랜덤 숫자를 위해 seed값 설정
	numbers := MakeNumbers()         // 무작위 숫자 3개를 만든다.

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

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false // 겹치는지 안겹치는지에 대한 플래그 변수 초기값
			for j := 0; j < i; j++ {
				if rst[j] == n { //숫자가 겹친다면 다시 뽑는다.
					duplicated = true //겹치면 true
					break
				}
			}
			if !duplicated { //겹치지 않는다면 for문을 빠져 나간다.
				rst[i] = n
				break
			}
		}
	}

	fmt.Println(rst)
	return rst
} // 0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환

func InputNumbers() [3]int {
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")
		var no int
		_, err := fmt.Scanf("%d\n", &no) // scanf 를 통해 값을 입력 받음 \n 를 하여 엔터값까지 입력 받아야함
		// _ 사용하지 않는 변수 , err 에러를 받음
		if err != nil { // 에러가 null이 아닌 경우: 에러가 났을 경우
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}
		success := true
		idx := 0 // rst에 넣을 변수
		for no > 0 {
			n := no % 10 // 숫자를 1자리씩 뽑아냄 123 % 10 = 3
			no = no / 10 // 남은 120 / 10 을 하면 12가 나옴 12가 다시가서 반복

			duplicated := false // 겹치는지 안겹치는지에 대한 플래그 변수 초기값
			for j := 0; j < idx; j++ {
				if rst[j] == n { //숫자가 겹친다면 다시 뽑는다.
					duplicated = true //겹치면 true
					break
				}
			}
			if duplicated { //겹치지 않는다면 for문을 빠져 나간다.
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("3개 보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n // rst[0] 부터 들어가기 때문에 순서가 바껴서 입력됨
			idx++
		}
		if success && idx < 3 { //success가 ture고 3보다 idx가 작을때
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success { // 숫자가 겹치면 다시 for 문으로 돌아감
			continue
		}
		break // 숫자를 받았으면 for문을 빠져나온다.
	}
	rst[0], rst[2] = rst[2], rst[0] // 순서를 똑바로 돌려줌
	fmt.Println(rst)
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
