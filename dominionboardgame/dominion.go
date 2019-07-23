package main

import (
	"bufio"
	"fmt"
	"os"

	dom "github.com/vincentgong3mm/golang/dominionboardgame/dominion"
)

func waitExit() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')

	fmt.Println(in)

}

func CreateGameManAndSetSuppy() *dom.GameMan {
	gman := dom.CreateNewGameMan()
	gman.CreateAllCard()

	gman.RegistCardToSuppy(dom.SetFirstGame, 2)

	return gman
}

func CreateTwoPlayer(g *dom.GameMan) {
	g.CreateNewPlayer("jong")
	g.CreateNewPlayer("seong")
}

func main() {
	fmt.Println("TestPlay")
	g := dom.CreateNewGameMan()
	g.CreateAllCard()
	g.RegistCardToSuppy(dom.SetFirstGame, 1)

	p := g.CreateNewPlayer("gong")
	p.DrawCard(5)

	p.GainCardGM(dom.Artisan)

	for true {
		switch p.GetStatus() {
		case dom.PlayerAction:
			p.Action(g)
			// ??? action 처리를 메시지로 던져서 여기 지나가고 또 action 진행함.
			// -> action 이 완료된 후에 다른 action 가능하게 해야함.

			fmt.Println("test---------action 끝나기를 기다리는 중.....")
			p.Wait()
			fmt.Println("test---------action 완료 .....")
		case dom.PlayerBuy:
			p.Buy(g)
		case dom.PlayerCleanUp:
			//p.CleanUp()
		}

		// player 상태 메시지 보여주고 입력 받아서 진행
		// 1. 액션상태 표시 : 액션 진행할 카드 선택
		//		- 액션 없는 카드 사용 못하세 제한
		// 2. buy 상태 표시, 현재 coin개수 + treature  카드 표시
		// 3. clean up
	}
	//g.Wait()
}
