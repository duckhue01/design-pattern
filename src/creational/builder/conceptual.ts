class Product1 {
  public parts: string[] = [];

  public getProduct() {
    console.log(`Product parts: ${this.parts.join(" - ")}`);
  }
}

interface Builder {
  createPartA(): void;
  createPartB(): void;
  createPartC(): void;
}

class ConcreteBuilder implements Builder {
  private product: Product1;

  constructor() {
    this.product = new Product1();
  }

  reset(): void {
    this.product = new Product1();
  }

  public createPartA(): void {
    this.product.parts.push("part A");
  }
  public createPartB(): void {
    this.product.parts.push("part B");
  }

  public createPartC(): void {
    this.product.parts.push("part C");
  }

  public getProduct(): Product1 {
    const result = this.product;
    this.reset();
    return result;
  }
}

class Director {
  private builder: Builder | undefined;

  public setBuilder(builder: Builder) {
    this.builder = builder;
  }

  public buildMinimalViableProduct() {
    if (this.builder) {
      this.builder.createPartA();
    }
  }

  public buildFullFeaturedProduct() {
    if (this.builder) {
      this.builder.createPartA();
      this.builder.createPartB();
      this.builder.createPartC();
    }
  }
}
function clientCode (director: Director) {
  const builder = new ConcreteBuilder();
  director.setBuilder(builder);

  console.log('Standard basic product:');
  director.buildMinimalViableProduct();
  builder.getProduct().getProduct();

  console.log('Standard full featured product:');
  director.buildFullFeaturedProduct();
  builder.getProduct().getProduct();

  // Remember, the Builder pattern can be used without a Director class.
  console.log('Custom product:');
  builder.createPartA();
  builder.createPartA();
  builder.getProduct().getProduct();

}

clientCode(new Director());
