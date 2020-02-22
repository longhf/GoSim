/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.16
 * Note:
 */

package common

import (
	"log"
	"time"
)

// CatchPanic calls a function and returns the error if function paniced
func CatchPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
		if err != nil {
			log.Printf("%s panic: %s", f, err)
		}
	}()

	f()
	return
}

// RunPanicless calls a function panic-freely
func RunPanicless(f func()) (panicless bool) {
	defer func() {
		err := recover()
		panicless = err == nil
		if err != nil {
			log.Printf("%s panic: %s", f, err)
		}
	}()

	f()
	return
}

// RepeatUntilPanicless runs the function repeatly until there is no panic
func RepeatUntilPanicless(f func()) {
	for !RunPanicless(f) {
	}
}


// RunForever runs the function forever...
func RunForever(restartInterval time.Duration, f func()error) {
	for {
		err := runForeverOnce(f)
		log.Printf("%s failed with error: %v, will restart after %s", f, err, restartInterval)
		time.Sleep(restartInterval)
	}
}

func runForeverOnce(f func()error) error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("serveTCPImpl: paniced with error %s", err)
		}
	}()

	return f()
}
