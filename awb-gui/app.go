package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// App struct
type App struct {
	ctx context.Context
	rng *rand.Rand
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// CalculateCIN calculates the check digit for a given 7-digit AWB number
func (a *App) CalculateCIN(input string) (string, error) {
	if len(input) != 7 {
		return "", fmt.Errorf("devi inserire esattamente 7 cifre")
	}

	awbNum, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("l'input deve essere un numero")
	}

	cin := awbNum % 7
	return fmt.Sprintf("%d", cin), nil
}

// GenerateRandomAWB generates a random 7-digit AWB with its check digit
func (a *App) GenerateRandomAWB() string {
	awb := a.rng.Intn(10_000_000) // 0..9,999,999
	cin := awb % 7
	return fmt.Sprintf("%07d%d", awb, cin)
}
