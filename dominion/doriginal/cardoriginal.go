package doriginal

type Actioner interface {
	Draw()
	AddBuy()
	AddAction()
	//DoSpecailACtion()
}

type Thief struct {
	*Card
}

/*
type Smithy struct {
	*Card
}
*/
