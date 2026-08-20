# Retrun structs, not interfaces

## The Owl's Factory

DiDong! The Panda walks into the Owl's factory.
> *"I need something to put my laptop and coffee"*

The Owl builds a sturdy table with flat top and says:
> *"Here you go! This is a Table interface. You can use it to put your laptop and coffee!"*

The Chihuahua Rushes In
> *"Hey! I need a table to hold pizzas and beers, today is my birthday! I want to throw a party!"*

The Chihuahua spots the table sitting at the corner and say:
> *"That table seems perfect!, give me the same table!"*

The Owl takes out the blueprints:
> *"No, it was designed for laptop and coffee. Not for pizzas and beers."*

The Chihuahua’s Rage
> *"Idiot! It is a table, what difference does it make, I want the same table, NOW!"*

## What is the problem?

A table with the flat top can be used for many things, it is no need to restrict its usage!

```Go
// The panda want
type WorkTable interface{
    PutCoffee()
    PutLaptop()
}

func StartWorking(w WorkTable) {
    w.PutLaptop()
    w.PutCoffee()
}

// The Chihuahua want
type PartyTable interface{
    PutPizzas()
    PutBeers()
}

func HostParty(p PartyTable) {
    p.PutPizzas()
    p.PutBeers()
}

// The Owl thinking(WRONG!!!)
func NewWorkTable() WorkTable{}
func NewPartyTable() PartyTable{}
    

// But the reality is 
type Table struct{}
func NewTable() Table{}
func (t Table) PutCoffee(){}
func (t Table) PutLaptop(){}
func (t Table) PutPizzas(){}
func (t Table) PutBeers(){}

func main() {
    table := NewTable()

    StartWorking(table) // ✅ Works!
    HostParty(table)    // ✅ Also works!
}


```
> *""*