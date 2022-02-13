# Concurrency in Go

<img alt="cover image" src="https://github.com/golang-basics/concurrency/blob/master/cover.jpg?raw=true" width="800"/>

### Summary

<div> 
    <a href="https://youtu.be/_uQgGS_VIXM" target="_blank">
        <img src="https://img.shields.io/badge/Concurrency in Go %231 -- Introduction to Concurrency-0B0B0B?&style=for-the-badge&logo=youtube&logoColor=white" alt="youtube"/>
    </a>
    <a href="https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-1.tar.gz">
        <img src="https://img.shields.io/badge/Download Zip %231 📦-00C853?&style=for-the-badge&logoColor=white" alt="zip archive"/>
    </a>
    <br/>
    <a href="https://youtu.be/srb6fbioEY4" target="_blank">
        <img src="https://img.shields.io/badge/Concurrency in Go %232 -- WaitGroups Part 1-0B0B0B?&style=for-the-badge&logo=youtube&logoColor=white" alt="youtube"/>
    </a>
    <a href="https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-2.tar.gz">
        <img src="https://img.shields.io/badge/Download Zip %232 📦-00C853?&style=for-the-badge&logoColor=white" alt="zip archive"/>
    </a>
    <br/>
    <a href="https://youtu.be/zAMUKb6fCO0" target="_blank">
        <img src="https://img.shields.io/badge/Concurrency in Go %233 -- WaitGroups Part 2-0B0B0B?&style=for-the-badge&logo=youtube&logoColor=white" alt="youtube"/>
    </a>
    <a href="https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-3.tar.gz">
        <img src="https://img.shields.io/badge/Download Zip %233 📦-00C853?&style=for-the-badge&logoColor=white" alt="zip archive"/>
    </a>
    <br/>
    <a href="https://youtu.be/_QNcn7LAANY" target="_blank">
        <img src="https://img.shields.io/badge/Concurrency in Go %234 -- WaitGroups Part 3-0B0B0B?&style=for-the-badge&logo=youtube&logoColor=white" alt="youtube"/>
    </a>
    <a href="https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-4.tar.gz">
        <img src="https://img.shields.io/badge/Download Zip %234 📦-00C853?&style=for-the-badge&logoColor=white" alt="zip archive"/>
    </a>
    <br/>
    <a href="https://youtu.be/nVjAS0uEnVM" target="_blank">
        <img src="https://img.shields.io/badge/Concurrency in Go %235 -- Atomics Part 1-0B0B0B?&style=for-the-badge&logo=youtube&logoColor=white" alt="youtube"/>
    </a>
    <a href="https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-5.tar.gz">
        <img src="https://img.shields.io/badge/Download Zip %235 📦-00C853?&style=for-the-badge&logoColor=white" alt="zip archive"/>
    </a>
    <br/>
    <a href="https://youtu.be/lKds8lAzt6s" target="_blank">
        <img src="https://img.shields.io/badge/Concurrency in Go %236 -- Atomics Part 2-0B0B0B?&style=for-the-badge&logo=youtube&logoColor=white" alt="youtube"/>
    </a>
    <a href="https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-6.tar.gz">
        <img src="https://img.shields.io/badge/Download Zip %236 📦-00C853?&style=for-the-badge&logoColor=white" alt="zip archive"/>
    </a>
    <br/>
</div>

### Coding Examples

- [Introduction to Concurrency](https://github.com/golang-basics/concurrency/tree/master/intro)
- [Go Routines](https://github.com/golang-basics/concurrency/tree/master/go-routines)
- [Channels](https://github.com/golang-basics/concurrency/tree/master/channels)
- [Select](https://github.com/golang-basics/concurrency/tree/master/select)
- [Concurrency Patterns](https://github.com/golang-basics/concurrency/tree/master/patterns)
- [Atomic(s) - sync/atomic](https://github.com/golang-basics/concurrency/tree/master/atomics)
- [WaitGroup(s) - sync.WaitGroup](https://github.com/golang-basics/concurrency/tree/master/waitgroups)
- [Mutexes - sync.Mutex](https://github.com/golang-basics/concurrency/tree/master/mutexes)
- [Pool - sync.Pool](https://github.com/golang-basics/concurrency/tree/master/pool)
- [Map - sync.Map](https://github.com/golang-basics/concurrency/tree/master/mutexes/syncmap)
- [Cond - sync.Cond](https://github.com/golang-basics/concurrency/tree/master/mutexes/cond)
- [Once - sync.Once](https://github.com/golang-basics/concurrency/tree/master/mutexes/once)
- [Race Conditions](https://github.com/golang-basics/concurrency/tree/master/go-routines/race-condition)
- [Threads](https://github.com/golang-basics/concurrency/tree/master/threads)
- [GOMAXPROCS](https://github.com/golang-basics/concurrency/tree/master/gomaxprocs)
- [Testing](https://github.com/golang-basics/concurrency/tree/master/testing)
- [Profiling](https://github.com/golang-basics/concurrency/tree/master/profiling)
- [HTTP REST API (Error Handling)](https://github.com/golang-basics/concurrency/blob/master/patterns/error-handling/http/main.go)
- [AWS S3 Bucket Clone](https://github.com/golang-basics/concurrency/tree/master/s3)

### Presentation Notes

- [Concurrency in Go #1 - Introduction to Concurrency](https://github.com/golang-basics/concurrency/raw/master/presentations/1_introduction-to-concurrency)
- [Concurrency in Go #2, #3, #4 - WaitGroups](https://github.com/golang-basics/concurrency/raw/master/presentations/2_3_4_waitgroups)
- [Concurrency in Go #5, #6 - Atomic(s)](https://github.com/golang-basics/concurrency/raw/master/presentations/5_6_atomics)
- [Concurrency in Go #7 - Mutexes](https://github.com/golang-basics/concurrency/raw/master/presentations/7_mutexes)

### Go routines

A go routines can block for one of these reasons:

- Sending/Receiving on channel
- Network or I/O
- Blocking System Call
- Timers
- Mutexes

Here's the full list of Go routines statuses:

- Gidle,            // 0
- Grunnable,        // 1 runnable and on a run queue
- Grunning,         // 2 running
- Gsyscall,         // 3 performing a syscall
- Gwaiting,         // 4 waiting for the runtime
- Gmoribund_unused, // 5 currently unused, but hardcoded in gdb scripts
- Gdead,            // 6 goroutine is dead
- Genqueue,         // 7 only the Gscanenqueue is used
- Gcopystack,       // 8 in this state when newstack is moving the stack

Feel free to check the rest of the statuses in the [runtime](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L34) source code

#### Fairness

- Infinite loop — preemption (~10ms time slice)
- Local Run queue — preemption (~10ms time slice)
- Global run queue starvation is avoided by checking the global run queue for every 61 scheduler tick
- Network Poller Starvation Background thread poll network occasionally if not polled by the main worker thread

### Channels

Here are couple of simple rules to make sure channels are used correctly

- Before writing to a channel, make sure someone else is reading from it (deadlock)
- Before reading from a channel, make sure someone else is writing to it (deadlock)
- When ranging over a channel, ALWAYS make sure the producer closes the channel eventually (deadlock)
- Writing to a closed channel will result in a runtime panic
- Reading from a closed channel won't have any effects
- A channel close, is considered a write operation

#### Poll Order

- Local Run queue
- Global Run queue
- Network Poller
- Work Stealing

#### Deadlocks

The Coffman Conditions are known as the techniques/conditions to help detect, prevent and correct deadlocks.
The Coffman Conditions are as follows:

- `Mutual Exclusion`

A concurrent process holds exclusive rights to a resource,
at any one time.

- `Wait for Condition`

A concurrent process must simultaneously hold a resource
and be waiting for an additional resource.

- `No Preemption`

A resource held by a concurrent process can only be released
by that process

- `Circular Wait`

A concurrent process (P1) must be waiting on a chain of other
concurrent processes (P2), which are in turn waiting on it (P1)

### Go Scheduler

Primarily the Go scheduler has the opportunity to get triggered on these events:

- The use of the keyword go
- Garbage collection
- System calls (i.e. open file, read file, e.t.c.)
- Synchronization and Orchestration (channel read/write)

#### P, M, G

G - goroutine
M - worker thread, or machine
P - processor, a resource that is required to execute Go code.
    M must have an associated P to execute Go code

Once the syscall exists Go tries to apply one of the rules:

- try to acquire the exact same P, and resume the execution
- try to acquire a P in the idle list and resume the execution
- put the goroutine in the global queue and put the associated M back to the idle list

Goroutines do not go in the global queue only when the local queue is full;
it is also pushed in it when Go inject a list of goroutines to the scheduler,
e.g. from the network poller or goroutines asleep during the garbage collection

### Spinning Threads

### Net Poller

### SysMon

`sysmon` is smart enough to not consume resources when there is nothing to do.
Its cycle time is dynamic and depends on the current activity of the running program.
The initial pace is set at 20 nanoseconds, meaning the thread is constantly looking to help.
Then, after some cycles, if the thread is not doing anything, the sleep between two cycles
will double until it reaches 10ms.
If your application does not have many system calls or long-running goroutines,
the thread should back off to a 10ms delay most of its time, giving
a very light overhead to your application.

For the implementation details checkout [`sysmon` source code](https://github.com/golang/go/blob/35ea62468bf7e3a79011c3ad713e847daa9a45a2/src/runtime/proc.go#L4245)

The thread is also able to detect when it should not run. Here are two cases:

- The garbage collector is going to run. sysmon will resume when the garbage collector ends.
- All the threads are idle, nothing is running.

### Work Stealing

Here's how Go makes sure to equally distribute & balance work
and make use of computer resources as efficient as possible:

- pull work from the local queue
- pull work from the global queue
- pull work from network poller
- steal work from the other P’s local queues

### Tracing

```bash
GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
```

### Concurrency Patterns

#### Pipelines

In general terms a pipeline is a mechanism for inter-process communication using message passing,
where the output of a pipeline is the input for the next pipeline.

Suppose that assembling one car requires three tasks that take 20, 10, and 15 minutes, respectively.
Then, if all three tasks were performed by a single station, the factory would output one car every 45 minutes.
By using a pipeline of three stations, the factory would output the first car in 45 minutes,
and then a new one every 20 minutes.

### Resources

- [OSX - number of CPUs](https://github.com/golang/go/blob/master/src/runtime/os_darwin.go#L151)
- [Windows - number of CPUs](https://github.com/golang/go/blob/master/src/runtime/os_windows.go#L356)
- [OSX - osinit](https://github.com/golang/go/blob/master/src/runtime/os_darwin.go#L128)
- [Windows - osinit](https://github.com/golang/go/blob/master/src/runtime/os_windows.go#L545)
- [Linux - osinit](https://github.com/golang/go/blob/master/src/runtime/os_linux.go#L301)
- [Go Scheduler by rakyll](https://rakyll.org/scheduler/)
- [Go Scheduler by morsmachine](https://morsmachine.dk/go-scheduler)
- [Illustrated Tales of Go Runtime Scheduler](https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b)
- [Go Scheduler Implementation](https://github.com/golang/go/blob/master/src/runtime/proc.go)
- [Main Goroutine](https://github.com/golang/go/blob/master/src/runtime/proc.go#L144)
- [Go Scheduler Implementation](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3470)
- [G - Source Code](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L403)
- [M - Source Code](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L511)
- [P - Source Code](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L604)
- [GRQ Check - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3343)
- [Force Preempt Duration for G - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5435)
- [Go Scheduler Design Doc](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit)
- [Scheduling in Go (Part 2) - Ardan Labs](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
- [Scheduling in Go (Part 3 - Concurrency) - Ardan Labs](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)
- [Scheduler Tracing in Go - Ardan Labs](https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html)
- [The Scheduler Saga](https://about.sourcegraph.com/go/gophercon-2018-the-scheduler-saga/#:~:text=The%20scheduler%20needs%20to%20exist,it%20multiplexes%20goroutines%20onto%20threads.)
- [Stack size](https://github.com/golang/go/blob/master/src/runtime/stack.go#L73)
- [Golang Net Poller Source Code](https://github.com/golang/go/blob/master/src/runtime/netpoll.go)
- [Golang Net Poller](https://morsmachine.dk/netpoller)
- [Preemptive vs Non-Preemptive Scheduling](https://www.guru99.com/preemptive-vs-non-preemptive-scheduling.html#:~:text=In%20Preemptive%20Scheduling%2C%20the%20CPU,Schedulign%20no%20switching%20takes%20place.)
- [Preemptive vs Cooperative](https://www.geeksforgeeks.org/difference-between-preemptive-and-cooperative-multitasking/#:~:text=Preemptive%20multitasking%20is%20a%20task,running%20process%20to%20another%20process.)
- [Go 1.14 Release Notes - Runtime](https://golang.org/doc/go1.14#runtime)
- [Go Asynchronous Preemption](https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c)
- [Go Routine & Preemption](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7)
- [Go Routine, OS Thread and CPU Management](https://medium.com/a-journey-with-go/go-goroutine-os-thread-and-cpu-management-2f5a5eaf518a)
- [Go Work Stealing - Go Scheduler](https://medium.com/a-journey-with-go/go-work-stealing-in-go-scheduler-d439231be64d)
- [Guarded Command Language](https://en.wikipedia.org/wiki/Guarded_Command_Language)
- [System Monitor - sysmon](https://en.wikipedia.org/wiki/System_monitor)
- [Go SysMon Runtime Monitoring](https://medium.com/@blanchon.vincent/go-sysmon-runtime-monitoring-cff9395060b5)
- [How does Go routine stack size evolve? - Medium](https://medium.com/a-journey-with-go/go-how-does-the-goroutine-stack-size-evolve-447fc02085e5)
- [SysMon - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5273)
- [Garbage Collector Period - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5268)
- [Preemption - suspend](https://github.com/golang/go/blob/master/src/runtime/preempt.go#L105)
- [Preemption - resume](https://github.com/golang/go/blob/master/src/runtime/preempt.go#L105)
- [Preemption - asyncPreempt](https://github.com/golang/go/blob/master/src/runtime/preempt.go#L302)
- [Preemption - preemptPark](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3563)
- [Scheduling - schedule](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3289)
- [CSP](https://levelup.gitconnected.com/communicating-sequential-processes-csp-for-go-developer-in-a-nutshell-866795eb879d)
- [Visualising Concurrency in Go](https://divan.dev/posts/go_concurrency_visualize/)
- [NUMA Deep Dive](https://frankdenneman.nl/2016/07/07/numa-deep-dive-part-1-uma-numa/)
- [How to Reduce Lock Contention with Atomic Package](https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549)
- [Go Mutex and Starvation](https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50)
- [Linux Kernel Mutex Lock Hand Off](https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/commit/?id=9d659ae14b545c4296e812c70493bfdc999b5c1c)
- [Sync Map - Medium](https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c)
- [Pipelines](https://blog.golang.org/pipelines)
- [Nil Channels](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
- [Context Cancellation](https://www.sohamkamani.com/golang/2018-06-17-golang-using-context-cancellation/)
- [Concurrency in Go - O Reilly](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html)
- [Pipeline - Wiki](https://en.wikipedia.org/wiki/Pipeline_(computing))
- [Pipeline - Unix](https://en.wikipedia.org/wiki/Pipeline_(Unix))
- [Mutex of Channel - Golang Docs](https://github.com/golang/go/wiki/MutexOrChannel)
- [Understand the Design of Sync Pool - Medium](https://medium.com/a-journey-with-go/go-understand-the-design-of-sync-pool-2dde3024e277)
- [clearpools - sync.Pool, GC - Go source code](https://github.com/golang/go/blob/master/src/runtime/mgc.go#L1547)
- [poolCleanup - sync.Pool - Go source code](https://github.com/golang/go/blob/master/src/sync/pool.go#L233)
- [Using Sync Pool](https://developer20.com/using-sync-pool/)
- [Concurrency in Go eBook - Amazon](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers-ebook/dp/B0742NH2SG)
