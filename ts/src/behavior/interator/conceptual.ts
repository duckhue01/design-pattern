

interface Iterator<Type> {

  next(): Type;
  current(): Type;
  key(): number;
  valid(): boolean;

  rewind(): void;

}

interface Aggregator {
  getIteration(): Iterator<string>;
}






export { }