interface ProductA {
  getProductA(): void;
}

interface ProductB {
  getProductB(): void;
}

class concreteProductALevel1 implements ProductA {
  public getProductA() {
    console.log("product A level 1 is created");
  }
}

class concreteProductALevel2 implements ProductA {
  public getProductA() {
    console.log("product A level 2 is created");
  }
}

class concreteProductBLevel1 implements ProductB {
  public getProductB() {
    console.log("product B level 1 is created");
  }
}

class concreteProductBLevel2 implements ProductB {
  public getProductB() {
    console.log("product B level 2 is created");
  }
}

interface Factory {
  createProductA(): ProductA;
  createProductB(): ProductB;
}

class FactoryLevel1 implements Factory {
  public createProductA(): ProductA {
    return new concreteProductALevel1();
  }

  public createProductB(): ProductB {
    return new concreteProductBLevel1();
  }
}

class FactoryLevel2 implements Factory {
  createProductA(): ProductA {
    return new concreteProductALevel2();
  }
  createProductB(): ProductB {
    return new concreteProductBLevel2();
  }
}

const client = (factory: Factory) => {

  factory.createProductA().getProductA();
  factory.createProductB().getProductB();

};


client(new FactoryLevel1())
