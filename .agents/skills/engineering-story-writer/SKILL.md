---
name: engineering-story-writer
description: >-
  Use this skill whenever authoring, structuring, drafting, reviewing, or refactoring engineering stories, architectural fables, or technical blog posts in this repository. Enforces the signature living story formula (Everyday Living Fable -> Character Conflict with Owl, Chihuahua & Panda -> Technical Analogy & Go Code -> Proof/Gotchas -> Pragmatic Verdict).
---

# Engineering Fable & Story Pattern Guide

This skill defines the mandatory architectural story structure for this repository. Every story must break down abstract software concepts into engaging, intuitive narratives told through the eyes of **The Panda**, **The Owl**, and **The Chihuahua**, backed by real Go code and proof.

---

## 🎋 The Golden Rule: "Living Intuition Before Code Abstraction"

Never start a story with abstract theory, UML diagrams, or dry definitions and other buzzwords. Always begin with a **tangible, visual, physical real-world scenario** (e.g., pouring morning coffee, buying an off-road truck at a dealership, a restaurant prep kitchen, or a hardware multi-tool). 

Once the reader intuitively grasps the physical dilemma, smoothly bridge that intuition into software architecture and Go code.

---

## 🎭 The 3 Core Characters & Their Roles

Every story should feature the repository's signature trio:

1. **🦉 The Owl (Chief Architect & Systems Theorist)**
   * **Archetype:** The Ivory Tower Perfectionist.
   * **Behavior:** Imposes rigid abstractions, theoretical clean architecture rules, and enterprise patterns. Believes that if a system is elegant on paper, reality must bend to it.
   * **Visuals:** Gold wire-rimmed spectacles, multi-pocket utility vest, tablet with topology diagrams, blueprints.
   * **Catchphrase:** *"According to clean architecture paradigms, this coupling is unacceptable."*

2. **🐕 The Chihuahua (Feature Hacker & Speedrunner)**
   * **Archetype:** The 10x Panic Developer.
   * **Behavior:** 50% trembling panic, 50% unbridled rage. Bypasses reviews, skips tests, pushes to main, and gets trapped or panics when rigid abstractions block quick fixes during outages.
   * **Visuals:** Bulging bloodshot eyes, bared fangs in mid-bark, vibrating motion lines, flathead screwdriver held like a melee weapon.
   * **Catchphrase:** *"Ship it now, debug it live, refactor never!"*

3. **🐼 The Panda (On-Call SRE & Maintainer / Narrator "Me")**
   * **Archetype:** The Battle-Hardened Stoic.
   * **Behavior:** Weary, calm, and grounded in real-world simplicity. Has seen every outage, cleans up over-engineered messes, explains the pragmatic solution, and restores sanity.
   * **Visuals:** Permanent dark eye rings, icy fever cooling gel patch on forehead, heavy half-closed eyelids, steaming black coffee mug.
   * **Catchphrase:** *"I'm looking at the logs. Stop touching things."*

---

## 📐 Mandatory 4-Phase Story Structure

Every story must follow these four sequential phases:

### Phase 1: The Living Story / Everyday Fable
* Open with an everyday situation that anyone can visualize.
* Establish the physical problem without mentioning Go syntax or software jargon yet.
* Show the Owl introducing an over-complicated "pure" rule, creating friction or absurdity in everyday life.
* Show the Chihuahua hitting a wall or causing chaos due to that constraint.

### Phase 2: The Character Conflict & Discovery
* The Panda steps in with a calming, pragmatic realization.
* Break down how common sense solves the everyday problem (e.g., the car dealership gives you the whole truck; the highway toll booth only checks if you have wheels).

### Phase 3: The Technical Analogy & Concrete Go Code
* Directly translate the living fable into Go architecture:
  * **Producer perspective:** Show how constructors / libraries should be authored (`func NewTruck() *Truck`).
  * **Consumer perspective:** Show how downstream callers define and accept minimal interfaces (`func PayToll(d Drivable)`).
* Include clean, idiomatic Go code snippets demonstrating:
  * ❌ The Over-Engineered / Anti-Pattern approach.
  * ✅ The Idiomatic, Composable Go approach.
* Highlight Go language mechanics (e.g., implicit interfaces / structural typing, zero allocations, standard library alignment like `os.Open` returning `*os.File`).

### Phase 4: Hard Proof & The Pragmatic Verdict
* Provide concrete justification:
  * Reproducible Go benchmarks (`allocs/op`, `ns/op`), or
  * Compiler / type-safety proofs, or
  * Clear comparison tables.
* End with **The Panda’s Pragmatic Takeaways**: 3–4 punchy, memorable bullet points of engineering wisdom.

---

## ✍️ Checklist for Every New Story

Before publishing any story, verify:
- [ ] Does the story start with a relatable physical/living example before introducing code?
- [ ] Are all three characters (Owl, Chihuahua, Panda) integrated naturally with their distinct voices and catchphrases?
- [ ] Is the narrator written from the first-person perspective of the Panda?
- [ ] Does the technical section include concrete, runnable Go code examples?
- [ ] Is there hard proof (benchmarks, compiler mechanics, or standard library parallels)?
- [ ] Does it conclude with clear, pragmatic architectural advice?
