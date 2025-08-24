# Three-Body Problem Simulation

This project is a real-time interactive simulation of the classical **three-body problem** in physics. It demonstrates chaotic orbital dynamics of three point masses under Newton’s law of gravitation.

Users can tweak initial conditions, pause/resume/reset the simulation.

---

## Demo

![](demo.mov)

---

## Details

* **Language:** Go (Golang)
* **Graphics Library:** [Ebiten](https://ebitengine.org/) v2
* **Numerical Method:** Explicit Euler integration (basic), extendable to Runge-Kutta or other ODE solvers
* **User Controls:**

  * Arrow keys: Adjust positions/velocities of bodies
  * P/O: Pause/Resume simulation
  * R: Reset simulation
  * Esc: Quit

---

## Physics

* **Newton’s Law of Gravitation:**
  $F = G \frac{m_1 m_2}{r^2} \hat{r}$

* **Newton’s Second Law of Motion:**
  $F = m a$ → $a = \frac{F}{m}$

* Each body experiences the gravitational pull of the other two, resulting in a system of coupled second-order differential equations.

* The three-body system is **chaotic**, meaning that small changes in initial conditions can lead to vastly different outcomes.

---

## Equations in action

* **Equations of Motion:**

  * Position update:
    $\mathbf{x}_{t+dt} = \mathbf{x}_t + \mathbf{v}_t \, dt$
  * Velocity update:
    $\mathbf{v}_{t+dt} = \mathbf{v}_t + \mathbf{a}_t \, dt$

* **Numerical Integration:**

  * Current version uses the **Euler method** (simple, but may accumulate error).
  * Can be upgraded to **RK4 (Runge-Kutta 4th order)** for better stability.

* **Complexity:**

  * For N bodies, the force computation is $O(N^2)$.
  * Here N = 3, so updates are efficient in real-time.

---

## How to Run

### Prerequisites

* [Go](https://go.dev/) 1.20+
* Ebiten v2 (`go get github.com/hajimehoshi/ebiten/v2`)

### Run Locally

```bash
git clone https://github.com/omgupta1608/three-body-simulation.git
cd three-body-simulation
go run main.go
```

### Build

```bash
go build -o threebody main.go
./threebody
```
