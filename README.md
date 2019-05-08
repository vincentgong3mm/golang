# golang
vincent golang 

## ToDo List
- complete) Log Method
- --)Play Card Action  Method
- --)Action Interface

## Action Card Interface
### Actioner interface 
- Draw()
- AddBuy()
- AddAction()
- DoSpecialAction()

### test call
- (r \*Player) PlayCard(c interface{})
- r.hands += c.hands <- draw n card  ex) Smithy is draw 3 card from deck.
- r.actions += c.actions
- r.coins = += c.coins
- c.DoSpecialAction()

### Thief Card's SpecialAction Sample
- gain Gold from Supply
- Reveal the top card from theirs deck, if the card is silver or gold then trash it -> Other player 

### Patrol Card's SpecialAction Sample
- draw 3 card from deck
- reveal 4 card from deck, if it is victory card then gain to hands, else put it onto deck.

### Upgrade Card's SpecialAction Sample
- trahs a card from hand, cost +1 it's cost gain a card 


## SmartyPants

SmartyPants converts ASCII punctuation characters into "smart" typographic punctuation HTML entities. For example:

|                |ASCII                          |HTML                         |
|----------------|-------------------------------|-----------------------------|
|Single backticks|`'Isn't this fun?'`            |'Isn't this fun?'            |
|Quotes          |`"Isn't this fun?"`            |"Isn't this fun?"            |
|Dashes          |`-- is en-dash, --- is em-dash`|-- is en-dash, --- is em-dash|
