package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ServiceUnreachable = errors.New("service unreachable")

func main() {

	circuit := Circuit(func(ctx context.Context) (string, error) {
		if time.Now().Second()%2 == 1 {
			return "", errors.New("ошибка: нечетная секунда")
		}
		return "успех: четная секунда", nil
	})
	circuitBreaker := Breaker(circuit, 5)
	for {
		res, err := circuitBreaker(context.TODO())
		if ServiceUnreachable == err {
			fmt.Println(err)
			break
		} else if err != nil {
			fmt.Println(err, "программа продолжает работать...")
			continue
		} else {
			fmt.Println(res)
			break
		}
	}
	fmt.Println("программа завершила работу")
}

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex
	return func(ctx context.Context) (string, error) {
		m.RLock() // Установить блокировку чтения

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d) //  (2 seconds) * (2^d)
			if !time.Now().After(shouldRetryAt) {
				return "", ServiceUnreachable
			}
		}

		m.RUnlock()

		response, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}
		consecutiveFailures = 0
		return response, nil
	}
}
