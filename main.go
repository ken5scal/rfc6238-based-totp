package main

import "time"

const timeStep = 30
const initialCountTimeStep = 0
const secretEx = "CUBU Z3LG DQET IKD7"

// https://tools.ietf.org/html/rfc6238
// The keys MAY be stored in a tamper-resistant device and SHOULD be
// protected against unauthorized access and usage.

func main() {
	totp(secretEx, initialCountTimeStep, timeStep)
}

// golang timestamp support int64 so that post-2038 timestamp will be supported
func totp(k string, t0, x int64) int64{
	return hotp(k, deriveNumberOfCurrentTimeStep(t0, x))
}

func hotp(k string, something int64) int64{

}

func deriveNumberOfCurrentTimeStep(t0, x int64) int64  {
	return (time.Now().Unix() - t0) / x
}
