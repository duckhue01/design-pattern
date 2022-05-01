### Abstract factory - a factory of factories
the Abstract Factory design pattern is a new layer of grouping to archive a bigger(and more complex) composite object, which is used through its interface. The idea behind grouping objects in families and grouping families is to have big factories that can be interchangeable and can grow more easily. In the early stages of development, it is also easier to work with factories and abstract factory than to wait until all concrete implementations are done to start your code. Also, you won't write an Abstract Factory from beginning unless you know that your object's inventory for a particular field is going to be very large and it could be easily grouped into families

#### the objectives
1. provide a new layer of encapsulation for Factory methods that return a common interface for all factories
2. group common factory into a super Factory