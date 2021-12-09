
// abstraction maintains a reference to an object of Implementation
// hierarchy and delegates all of the real work to this object
class Abstraction {
  protected impl: Implementation;

  constructor(imp: Implementation) {
    this.impl = imp;
  }

  public operation() {
    return `Abstraction: ${this.impl.implOperation()}`;
  }
}

class ExtendedAbstraction extends Abstraction {
  public operation() {
    return `Extended Abstraction: ${this.impl.implOperation()}`;
  }
}



//  Implementation provides only primitive operations, while the abstraction
// defines higher level of operation based on those primitive
interface Implementation {
  implOperation(): string;
}
class ConcreteImplementationA implements Implementation {
  implOperation() {
    return "this is concrete implementation a";
  }
}

class ConcreteImplementationB implements Implementation {
  implOperation() {
    return "this is concrete implementation b";
  }
}



// client code should only depend on Abstraction class
// so that client can support any abstraction implementation
function client(ab: Abstraction) {
  console.log(ab.operation())
}

client(new Abstraction(new ConcreteImplementationA()));

client(new ExtendedAbstraction(new ConcreteImplementationB()))

export {}