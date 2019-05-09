import { ImmutableQueue } from './ImmutableQueue'

let queue = new ImmutableQueue([0, 10, 100])

console.log("start with {0, 10, 100}")

console.log("queue.isEmpty() is: ", queue.isEmpty())
console.log("queue.head() is: ", queue.head())

console.log("queue.deQueue()")
queue = queue.deQueue()
console.log("queue.head() is: ", queue.head())

console.log("queue.deQueue()")
queue = queue.deQueue()
console.log("queue.head() is: ", queue.head())

console.log("queue.deQueue()")
queue = queue.deQueue()
console.log("queue.head() is: ", queue.head())
console.log("queue.isEmpty() is: ", queue.isEmpty())

console.log("queue.deQueue()")
queue = queue.deQueue()
console.log("queue.head() is: ", queue.head())

console.log("queue.enQueue(2)")
queue = queue.enQueue(2)
console.log("queue.tail() is: ", queue.tail())

console.log("queue.enQueue(22)")
queue = queue.enQueue(22)
console.log("queue.tail() is: ", queue.tail())


