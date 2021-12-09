interface Component {
  operation(): string;
}

class ConcreteComponent implements Component {
  public operation(): string {
    return "creteComponent"
  }
}


class Decorator implements Component {
  private com: Component;
  constructor(com: Component) {
    this.com = com;
  }
  public operation() : string {

    return this.com.operation();
  }


}

class ExtendedDecoratorA extends Decorator {
  public operation(): string {
    return `extended decorator a (${super.operation()})`
  }
}

class ExtendedDecoratorB extends Decorator {
  public operation(): string {
    return `extended decorator b (${super.operation()})`
  }
}


const clientCode = (com: Component) => {
  console.log(com.operation());
  
}


const simple = new ConcreteComponent();
console.log('Client: I\'ve got a simple component:');
clientCode(simple);
console.log('');

/**
 * ...as well as decorated ones.
 *
 * Note how decorators can wrap not only simple components but the other
 * decorators as well.
 */
const decorator1 = new ExtendedDecoratorA(simple);
const decorator2 = new ExtendedDecoratorB(decorator1);
console.log('Client: Now I\'ve got a decorated component:');
clientCode(decorator2);



export {};
