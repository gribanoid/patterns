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
		time.Sleep(time.Millisecond * 150) // имитируем работу
		return "успех: четная секунда", nil
	})
	wrapped := Breaker(circuit, 3)
	for {
		res, err := wrapped(context.TODO())
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
			shouldRetryAt := lastAttempt.Add(time.Microsecond * 1 << d) //  (2 seconds) * (2^d)
			if !time.Now().After(shouldRetryAt) {
				return "", ServiceUnreachable
			}
			fmt.Println("Даем программе еще один шанс")
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
