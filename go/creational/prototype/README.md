### Prototype design pattern
the aim of the prototype pattern is to have an object or a set of objects that is already created at compilation time, but which you can clone as many time as you want at runtime. This is useful, for example, as a default template for a user who has just registered with your webpage or default pricing plan in some service. The key difference between this and a builder pattern is that objects are cloned for the uer instead of building them at runtime. you can also build a cache-like solution, storing information using a prototype.

#### Objective
- maintain a set of objects that will be cloned to create new instances
- provide a default value of some type to start working on top of it
- free CPU of complex object initialization to take more memory resource

