import { ImmutableQueue } from './index'
import * as assert from 'power-assert';


(function testIsEmpty() {
    const q1 = new ImmutableQueue([])
    assert.equal(q1.isEmpty(), true)

    const q2 = new ImmutableQueue([1])
    assert.equal(q2.isEmpty(), false)
})();

(function testHead() {
    const q1 = new ImmutableQueue([])
    assert.equal(q1.head(), null)

    const q2 = new ImmutableQueue([1])
    assert.equal(q2.head(), 1)
})();

(function testTail() {
    const q1 = new ImmutableQueue([])
    assert.equal(q1.tail(), null)

    const q2 = new ImmutableQueue([1])
    assert.equal(q2.tail(), 1)
})();

(function testEnQueue() {
    const q1 = new ImmutableQueue([])
    assert.equal(q1.tail(), null)

    const q2 = q1.enQueue(1)
    assert.equal(q1.tail(), null)
    assert.equal(q2.tail(), 1)
})();

(function testDeQueue() {
    const q1 = new ImmutableQueue([1, 2, 3])
    assert.equal(q1.head(), 1)

    const q2 = q1.deQueue()
    assert.equal(q1.head(), 1)
    assert.equal(q2.head(), 2)
})();

(function testItems() {
    const q1 = new ImmutableQueue([1, 2, 3])
    assert.throws(() => {
        // `q1.items` causes compile error
        q1['items'].push(100)
    })
})();