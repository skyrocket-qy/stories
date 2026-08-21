# Single Source of Truth

![The Owl, the Chihuahua, and the Panda at the milk depot](../assets/coffee_fable_trio.jpg)

## 🥛 The Store Milk Avalanche

Every morning, **we must always have exactly 50 cartons of milk on the shelves for the morning rush.**

🦉 **The Owl:** Counts inventory and calculates the needed restock on the whiteboard.

> *"8/1 Currently: 30 milks, Needed: 20"*  
> 🐼 Panda imports 20 milks. Total: 50.

> *"8/2 Currently: 15 milks, Needed: 35"*  
> 🐼 Panda imports 35 milks. Total: 50.

---

## 💥 The Incident: The Missing Milk

> *"8/3 Currently: 17 milks, Needed: 33"*  
> 🐕 The Chihuahua, the newcomer, feels thirsty and takes one milk from the store.  
> 🐼 Panda imports 33 milks.  

At 9:00 AM, the store opened. Total: **49 milks**. The morning audit failed, and an emergency post-mortem was called:

> 🦉 *"It is not my problem! My calculation is correct!"* — The Owl  
> 🐼 *"It is not my problem! I always watch the needed entries!"* — The Panda  
> 🐕 *"I don't know..."* — The Chihuahua  

---

## 🪵 What is the Problem?

We don't need the `needed` value written on the whiteboard—it can simply be calculated from the `currently` value on the fly.

### ❌ The Anti-Pattern: Redundant Field
```go
const TargetCap = 50

type MilkStore struct {
    CurMilk    int
    NeededMilk int // ❌ Redundant: can be calculated from CurMilk
}

func (m *MilkStore) CalculateNeededMilk() {
    m.NeededMilk = TargetCap - m.CurMilk
}

func (m *MilkStore) NeededMilk() int {
    return m.NeededMilk
}
```

### ✅ The Flexible Solution
```go
const TargetCap = 50

type MilkStore struct {
    CurMilk int
}

func (m *MilkStore) NeededMilk() int {
    return TargetCap - m.CurMilk
}
```

---

## 🪵 Real-World Example: User Suspension

### ❌ The Bug-Prone Anti-Pattern
```go
type User struct {
    ID             int
    IsActive       bool       // ❌ Lying state: turns stale the moment time passes
    SuspendedUntil *time.Time
}
```

### ✅ The Single Source of Truth (SSoT) Solution
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