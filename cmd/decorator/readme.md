# Decorator Design Patterns in Go

The **Decorator** design pattern is a structural pattern that allows you to dynamically add behavior to an object at runtime without modifying its original code. It provides a flexible alternative to subclassing for extending functionalities. In Go, which doesn't support inheritance, the Decorator pattern is particularly useful for adding functionalities to objects by embedding one struct into another.

### **When to Use the Decorator Pattern**

* When you need to add behavior to objects dynamically at runtime.
* When extending functionalities using subclassing would result in an explosion of classes.
* When you want to add responsibilities to objects transparently without modifying their original code.

### **Components of the Decorator Pattern**

1. **Component Interface** : Defines the common interface for both concrete components and decorators.
2. **Concrete Component** : The core object to which additional behavior can be added.
3. **Decorator** : Implements the component interface and has a reference to a component object.
4. **Concrete Decorators** : Extend the functionalities of the component by adding additional behaviors.
