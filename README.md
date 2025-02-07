# Unexpected Defer Behavior with Panic in Go

This repository demonstrates a subtle issue in Go's handling of `defer` statements when a `panic` occurs.  The example showcases a situation where the deferred function might not be executed as expected.

## The Problem

Go's `defer` keyword is used to schedule a function call to be executed at the end of the surrounding function. However, if a `panic` occurs before the end of the function, the `defer` function's execution becomes less predictable.

This example specifically highlights that `defer` statements are not guaranteed to execute after a `panic` occurs. The `defer` function is skipped in this specific situation due to the panic.