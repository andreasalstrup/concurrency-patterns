# Concurrency patterns

### Concurrency is not Parallelism
* Concurrency: Dealing with lots of things at once
    * About structure
* Parallelism: Doing lots of things at once 
    * About execution

>If getting the concurrency design right we dont have to worry about parallelism, as parallelism is then a free variable that we can decide on. - Rob Pike

>Dont think about running i parallel, think about how to break the problem down into independent components that you can separate and understand to compose to understand the whole problem together. - Rob Pike

* Concurrency enables parallelism.
* Concurrency makes parallelism easy.

`goroutines` starts a "light thread" what runds conceptually right away, that we dont have to wait for it to return. 

* Express ideas - Use threads as a clean way of expressing idears
* Locks are meant to protect invariants
* Channels are a asynchronous communications mechanismen (mutex and condition variables are often better to reason about)