interface Cloneable {
  clone(): Cloneable;
}

class Prototype implements Cloneable {
  public primitive: string | undefined;
  public object: object | undefined;
  public backReference: BackReference | undefined;

  // constructor(primitive: string, object: object, backReference: BackReference) {
  //   this.primitive = primitive;
  //   this.object = object;
  //   this.backReference = backReference;
  // }

  clone(): Prototype {
    const clone: Prototype = Object.assign({}, this);
    clone.object = {
      ...this.object,
    };
    clone.backReference = {
      reference: { ...this },
    };

    return clone;
  }
}

class BackReference {
  public reference;

  constructor(reference: Prototype) {
    this.reference = reference;
  }
}

(() => {
  const prototype = new Prototype();
  prototype.primitive = "duke";
  prototype.object = {
    a: 100,
    b: 1000,
  };
  prototype.backReference = new BackReference(prototype);

  const prototype2 = prototype.clone();
})();
