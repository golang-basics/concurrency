# Concurrency in Go

### OSX `sysctl`

```bash
# get the number of logical CPU cores
sysctl hw.logicalcpu

# get the number of physical CPU cores
sysctl hw.physicalcpu

# get the number of logical cores
sysctl hw.ncpu

# get the number of physical/logical cores
# also thread count meaning the total count of running threads in parallel
sysctl -a | grep machdep.cpu | grep count
```

### Resources

- [OSX - number of CPUs](https://github.com/golang/go/blob/master/src/runtime/os_darwin.go#L151)
- [Windows - number of CPUs](https://github.com/golang/go/blob/master/src/runtime/os_windows.go#L356)
- [Go Scheduler by rakyll](https://rakyll.org/scheduler/)
- [Go Scheduler by morsmachine](https://morsmachine.dk/go-scheduler)
- [Illustrated Tales of Go Runtime Scheduler](https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b)
- [Go Scheduler Implementation](https://github.com/golang/go/blob/master/src/runtime/proc.go)
- [Main Goroutine](https://github.com/golang/go/blob/master/src/runtime/proc.go#L144)
- [Go Scheduler Implementation](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3470)
- [G](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L403)
- [M](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L503)
- [P](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L596)
- [Go Scheduler Design Doc](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit)
- [Stack size](https://github.com/golang/go/blob/master/src/runtime/stack.go#L73)
