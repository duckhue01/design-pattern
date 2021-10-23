class Facade {
  protected sub1: Subsystem1;

  protected sub2: Subsystem2;

  constructor(sub1: Subsystem1 , sub2: Subsystem2) {
    this.sub1 = sub1 || new Subsystem1();
    this.sub2 = sub2 || new Subsystem2();
  }

  public operation(): string {
    let result = "Facade initializes subsystems:\n";
    result += this.sub1.operation1();
    result += this.sub2.operation1();
    result += "Facade orders subsystems to perform the action:\n";
    result += this.sub2.operation2();
    result += this.sub1.operation2();

    return result;
  }
}

class Subsystem1 {
  public operation1(): string {
    return "do operation 1";
  }

  public operation2(): string {
    return "do operation 2";
  }
}

class Subsystem2 {
  public operation1(): string {
    return "do operation 1";
  }

  public operation2(): string {
    return "do operation 2";
  }
}

function client (facade: Facade) {
  console.log(facade.operation())
}

client(new Facade(new Subsystem1(), new Subsystem2()))



export {}