abstract class Component {
  public add(com: Component): void {}

  public remove(com: Component): void {}

  public isComposite(): boolean {
    return false;
  }

  public abstract operation(): string;
}

class Leaf extends Component {
  public operation(): string {
    return " left ";
  }
}

class Composite extends Component {
  protected children: Component[] = [];

  public add(com: Component): void {
    this.children.push(com);
  }

  public remove(com: Component): void {
    const componentIndex = this.children.indexOf(com);
    this.children.splice(componentIndex, 1);
  }

  public isComposite(): boolean {
    return true;
  }

  public operation(): string {
    const results = [];
    for (const child of this.children) {
      results.push(child.operation());
    }

    return `Branch(${results.join("+")})`;
  }
}

const client = (com: Component): void => {
  console.log(com.operation());
};
const composite = new Composite();


composite.add(new Leaf());
composite.add(new Leaf());
const composite1 = new Composite();
composite1.add(new Leaf());
composite1.add(new Leaf())
composite.add(composite1)

client(composite)
export {};
