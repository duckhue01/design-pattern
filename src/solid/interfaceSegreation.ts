//interface segregation state many client-specific interface are better than
// one general purpose interface. client should not be forced to implement method
//  they do not need


// our interface should not much general
// it should more specific
interface WrongPersistence {
  createDatabase(): void;
  save(): void;
}

class WrongDatabase implements Persistence {
  createDatabase(): void {
    console.log("database is created");
  }
  save(): void {
    throw new Error("data is saved");
  }
}



class WrongFile implements Persistence {
  createDatabase(): void {
    console.log("this method do nothing");
  }
  save(): void {
    console.log("file is saved");
  }
}



interface Persistence {
  save(): void;
}


interface FilePersistence extends Persistence {

}
interface DatabasePersistence extends Persistence {
  createDatabase(): void;
}


class File implements FilePersistence {
  save(): void {
    console.log("data is saved")
  }

}

class Database implements DatabasePersistence {
  createDatabase(): void {
    console.log("database is created")
  }
  save(): void {
    console.log("data is saved")
  }

}


export {};
