# Single Source of Truth

![The Owl, the Chihuahua, and the Panda at the milk depot](../assets/coffee_fable_trio.jpg)

## 🥛 The Store Milk Avalanche
Every morning **we must always have exactly 50 cartons of milk on the shelves for the morning rush.**

🦉 The Owl: Count inventory and Calculate the needed 
> *"8/1 Currently: 30 milks, Needed: 20"*
Panda import 20 milks 
> *"8/2 Currently: 15 milks, Needed: 35"*
Panda import 35 milks 

## 💥 The Incident: The Missing Milk
> *"8/3 Currently: 17 milks, Needed: 33"*
The Chihuahua, the newcomer, feel thirsty and take one milk from the store.
Panda import 33 milks 

At 9:00 AM, the store opened. The morning audit failed. An emergency post-mortem was called.

> *"It is not my problem!!!, My calculation is correct! Says the Owl!"*
> *"It is not my problem!!!, I always watch the needed entries! Says the Panda!"*
> *"I don't know... Says the Chihuahua"*
---

## 🪵 What is the Problem?

We don't need "needed" value to be write down on the whiteboard, it can be  calculated from the "currently" value.


```go
const TargetCap = 50

// ❌ The Anti-Pattern: Redundant Field NeedMilk
type MilkStore struct{
    CurMilk int
    NeededMilk int // ❌ Redundant, can be calculated from CurMilk
}

func (m *MilkStore) CalculateNeededMilk() {
    m.NeededMilk = TargetCap - m.CurMilk
}

func (m *MilkStore) NeededMilk() int {
    return m.NeededMilk
}

✅
type MilkStore struct{
    CurMilk int
}

func (m *MilkStore) NeededMilk() int {
    return TargetCap - m.CurMilk
}
```

## 🪵 The Real world example (User Suspend)

❌ The Bug-Prone Anti-Pattern
```go


type User struct {
    ID             int
    IsActive       bool       // ❌ Lying state: turns stale the moment time passes
    SuspendedUntil *time.Time
}
```
✅ The Single Source of Truth (SSoT) Solution
```go
type User struct {
    ID             int
    SuspendedUntil *time.Time // 🛡️ Single Source of Truth
}

// IsActive derives the status dynamically based on current time.
// It can NEVER become stale because time is evaluated at read time.
func (u *User) IsActive(now time.Time) bool {
    if u.SuspendedUntil == nil {
        return true
    }
    return now.After(*u.SuspendedUntil)
}
```