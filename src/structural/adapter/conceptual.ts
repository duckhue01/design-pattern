class Target {
  public request() {
    return "this is content that program can understand";
  }
}

class Adaptee {
  public AnotherRequest() {
    return ".eetpadA eht fo roivaheb laicepS";
  }
}

class Adapter extends Target {
  public adaptee: Adaptee;

  constructor(adaptee: Adaptee) {
    super();
    this.adaptee = adaptee;
  }

  public request() {
    return `(adapted): ${this.adaptee
      .AnotherRequest()
      .split("")
      .reverse()
      .join("")}`;
  }
}

const client = (target: Target) => {
  console.log(target.request());
};


client(new Target())


// this will cause an error because client code would'nt work with adaptee class
// client(new Adaptee())

// this kind of like wrapper of adaptee so that it can work with client code
client(new Adapter(new Adaptee()))


export {};
