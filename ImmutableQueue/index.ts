export class ImmutableQueue<T> {
    private items: T[]

    constructor(
        _items: T[]
    ) {
        this.items = [..._items]
        Object.freeze(this)
        Object.freeze(this.items)
    }

    isEmpty(): boolean {
        return this.items.length === 0
    }

    enQueue(item: T): ImmutableQueue<T> {
        return new ImmutableQueue([...this.items, item])
    }

    deQueue(): ImmutableQueue<T> {
        if (this.isEmpty()) return new ImmutableQueue([])

        const [_, ...rest]: T[] = [...this.items]
        return new ImmutableQueue(rest)
    }

    head(): T | null {
        if (this.isEmpty()) return null
        return this.items[0]
    }

    tail(): T | null {
        if (this.isEmpty()) return null
        return this.items[this.items.length - 1]
    }
}