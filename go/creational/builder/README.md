### builder design pattern - reusing an algorithm to create many implementation of an interface

#### Objective
1. abstract complex creations so that object creation is separated from the object user
2. create an object step by step by filling its fields and creating the embedded objects
3. reuse the object creation algorithm between many object


#### Wrapping up
1. the builder design pattern helps us maintain an unpredictable number of products by using a common construction algorithms that is used by the director. The construction process is always abstracted from the user of the product.
2. having a defined construction pattern helps when a newcomer to our source code needs to add a new product to the pipeline. The BuildProcess interface specifies
3. try to avoid the Builder pattern when you are not completely sure that the algorithms is going to be more or less stable because any small change in this interface will affect all your builders and it could be awkward if you add a new method that some of your builders need and other builder do not