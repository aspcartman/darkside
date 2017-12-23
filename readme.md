# DarkSide

[![buddy pipeline](https://buddy.aspc.me/aspc/darkside/pipelines/pipeline/6/badge.svg?token=39b71296da891e3a6a485732a532b589eac0d22fa92f1a74b48afcd7c0ddfd87 "buddy pipeline")](https://buddy.aspc.me/aspc/darkside/pipelines/pipeline/6)

The darkside of the golang. WIP

## Note
### Is it safe?
As safe as a lightsaber. In hands of a Jadi it is a powerful tool, but in a hand of a youngling...

You need to know following things before using this package wisely:
1. System Development fundamentals
2. C, asm, goasm
3. GoLang runtime & gc (it's required that you read sources and understand how all of the involved functions and mechanics work)

The author does not take any responsibility for cut off hands and lost money.

Having that said, most features are safe to use if you follow the rules.

### State
The project is in WiP. As for now it uses two publicly available libraries ([fore-export](https://github.com/alangpierce/go-forceexport) and [monkey](https://github.com/bouk/monkey)), which will be replaced later as they are a bit ambiguous and do not provide a futureset required.

## Features
### Panic SafeNet
`SetUnrecoveredPanicHandler` allows you to set a handler for an unrecovered panic. You are not allowed to call `recover()` in there, instead you are provided with a `Panic` structure, from which you can get all the info you need.
Accessing the stacktrace is the same as usual.
### Rules
1. Do not recover
2. Do not grow the stack (call only relatively small functions, or pass the work to another goroutine and wait)
3. Set it only once in the init() of the main package or an init of an imported in the main package package.
### What the worst thing could happen if I break the rules?
You *don't* crash; instead loop indefinitely catching the "growstack on system stack" panic. You will have to kill the process manually. More then that, other goroutines will continue to work.


