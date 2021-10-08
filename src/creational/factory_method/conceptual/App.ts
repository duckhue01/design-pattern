

/**
 * Factory class have an abstract create method that return Product
 * subclass of Factory will implement create method
 */
abstract class Factory {
  abstract create(): Product;

  /**
   * this method is some core business logic of application it will change depend on
   * the Product object
   */
  public print(): void {
    console.log(this.create().getProduct());
  }
}

class ConcreteFactory1 extends Factory {
  public create(): Product {
    return new ConcreteProduct1();
  }
}

class ConcreteFactory2 extends Factory {
  public create(): Product {
    return new ConcreteProduct2();
  }
}

interface Product {
  getProduct(): string;
}

class ConcreteProduct1 implements Product {
  public getProduct(): string {
    return "this is concrete product 1";
  }
}

class ConcreteProduct2 implements Product {
  public getProduct(): string {
    return "this is concrete product 2";
  }
}


//  this is when we use factory method
// in this case we just add product class and concrete factory
function createProduct(factory: Factory): void {
  factory.print();
}
createProduct(new ConcreteFactory1());
createProduct(new ConcreteFactory2());


// this is normal case
function normal() {
  
  const condition  = 1;

  
  if(condition == 1) {
    const product1 = new ConcreteProduct1();
    
  } else if(condition == 2) {

    const product2 = new ConcreteProduct2();
  }
  // if there is another product we need to add product class and change the code base
  // again and again
  
}





export {};
