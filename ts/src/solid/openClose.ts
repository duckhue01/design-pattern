// the open close principle require that classes should be
// open for extension and close to modification

class WrongPersistence {
  private data: string;

  constructor(data: string) {
    this.data = data;
  }

  saveToFile() {
    console.log("this data is saved to file");
  }

  // each time we want to add feature we have to change this class
  // we are taking the risk of creating potential bugs
  saveToDatabase() {
    console.log("this data is saved to database");
  }
}


// we create an common interface for each individual persistence class
// so that we can easily extend in the future
interface Persistence {
  save(): void;
}

class FilePersistence implements Persistence {
  private data: string;

  constructor(data: string) {
    this.data = data;
  }

  public save() {
    console.log("data is saved to file");
  }
}
// each time we want to create a new feature new just need to implement
// common interface without change the existing code, avoiding potential bug
class DatabasePersistence implements Persistence {
  private data: string;
  constructor(data: string) {
    this.data = data;
  }

  public save() {
    console.log("data is saved to database");
  }
}
