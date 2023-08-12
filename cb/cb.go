package cb

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
)

func Cb_test() {

	callRemoteMultipleTimes(35)
}

func notImportant() {
	mb := circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{Name: "ttt"}))
	fmt.Println(mb)
}

func callRemoteMultipleTimes(n int) {
	var cb_settings gobreaker.Settings
	cb_settings.Name = "cb_test"
	cb_settings.MaxRequests = 3
	cb_settings.Interval = 30 * time.Second
	cb_settings.Timeout = 500 * time.Millisecond
	cb_settings.ReadyToTrip = func(counts gobreaker.Counts) (res bool) {
		defer func() {
			fmt.Printf("ReadyToTrip - %v ----- res:%v\n", counts, res)
		}()
		return counts.ConsecutiveFailures >= 5
	}
	cb_settings.OnStateChange = func(name string, from, to gobreaker.State) {
		fmt.Printf("------------------ OnStateChange - %s - state changed. %s -> %s\n", name, from, to)
	}
	cb_settings.IsSuccessful = func(err error) bool {
		return err == nil
	}

	cb := gobreaker.NewCircuitBreaker(cb_settings)

	for i := 0; i < n; i++ {
		err := callRemoteServiceWithCb(cb, i+1)
		fmt.Printf("callRemoteServiceWithCb(%d). err:%v  --  %v\n", i+1, err, cb.Counts())
		time.Sleep(100 * time.Millisecond)
	}
}

func callRemoteServiceWithCb(cb *gobreaker.CircuitBreaker, seq int) error {
	_, e := cb.Execute(func() (interface{}, error) {
		return remoteCallMock(seq)
	})
	// if e != nil {
	// 	fmt.Printf("cb.execute: %s\n", e)
	// }
	// fmt.Printf("method: %d\n", i)
	return e
}

func remoteCallMock(seq int) (int, error) {
	if seq > 10 && seq < 23 {
		return seq, errors.New("my func error")
	}
	return seq, nil
}
