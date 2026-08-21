# Single Source of Truth

## 🥛 The Store Milk Avalanche

The Owl, the Chihuahua, and the Panda are working together in a milk store, The Owl is in charge of the inventory, each morning, the Owl checks the stock and writes down the number of milk needed(total: 50) on the whiteboard for the day. and the Panda is in charge of importing milk from the warehouse.

> *"Each day: 50 milks"*
> *"8/1 Currently: 30 milks, Needed: 20"*
Panda import 20 milks 
> *"8/2 Currently: 15 milks, Needed: 35"*
Panda import 35 milks 


> *"8/3 Currently: 17 milks, Needed: 33"*
The Chihuahua, the newcomer, feel thirsty and take one milk from the store.
Panda import 33 milks 

Total: 49 milks

> *"It is not my problem!!!, My calculation is correct! Says the Owl!"*
> *"It is not my problem!!!, I always watch the needed entries! Says the Panda!"*
> *"I don't know... Says the Chihuahua"*
---

## 🪵 What is the Problem?

We don't need "needed" value to be write down on the whiteboard, it can be  calculated from the "currently" value.

❌
```go
type MilkStore struct{
    CurMilk int
    NeededMilk int
}

func (m *MilkStore) CalculateNeededMilk() {
    m.NeededMilk = 50 - m.CurMilk
}

func (m *MilkStore) NeededMilk() int {
    return m.NeededMilk
}

```

✅
```go
type MilkStore struct{
    CurMilk int
}

func (m *MilkStore) NeededMilk() int {
    return 50 - m.CurMilk
}

```