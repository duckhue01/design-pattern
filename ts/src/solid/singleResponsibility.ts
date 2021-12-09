// the single responsibility principle state that a class should do one thing
// and therefore it should have only one reason to change

class Book {
  public name: string;
  public price: number;

  constructor(name: string, price: number) {
    this.name = name;
    this.price = price;
  }
}

// this wrong because there are three reason to change this class
class WrongInvoice {
  private total: number;

  constructor(book: Book, amount: number) {
    this.total = book.price * amount;
  }

  public getTotal() {
    return this.total;
  }

  public printInvoice() {
    console.log(this.total);
  }

  public saveInvoice() {
    console.log("invoice is saved");
  }
}

// this class will be changed when we want to alter the logic to calculate
// invoice
class Invoice {
  private total: number;

  constructor(book: Book, amount: number) {
    this.total = book.price * amount;
  }

  public getTotal() {
    return this.total;
  }
}
// same thing
class InvoicePrinter {
  private invoice: Invoice;

  constructor(invoice: Invoice) {
    this.invoice = invoice;
  }

  public printInvoice() {
    console.log(this.invoice);
  }
}
// same deal
class invoicePersistence {
  private invoice: Invoice;

  constructor(invoice: Invoice) {
    this.invoice = invoice;
  }
  public saveInvoice() {
    console.log("invoice is saved");
  }
}
