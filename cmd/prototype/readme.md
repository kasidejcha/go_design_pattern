# Prototype Design Pattern in Go

The **Prototype** design pattern is a creational pattern that allows you to create new objects by copying or cloning an existing object (the prototype). This pattern is particularly useful when the cost of creating a new object is expensive, or when you want to avoid the complexities of creating objects using a factory method.

## Overview

The Prototype pattern involves:

1. **Prototype Interface** : Defines a `Clone` method for creating a copy of an object.
2. **Concrete Prototype** : Implements the `Clone` method to return a copy of the object.
3. **Client** : Uses the prototype to clone new objects.

### When to Use the Prototype Pattern

* When creating new objects is resource-intensive, and you want to avoid creating them from scratch.
* When you need to create an object that is a copy of another object.
* When the object structure is complex and has many configuration options

### **Key Features of the Prototype Pattern**

* **Cloning** : The main feature of this pattern is creating new objects by copying an existing object. The `Clone()` method in each concrete prototype provides this capability.
* **Independence** : The cloned object is independent of the original. Changes to the clone do not affect the original, as shown in the example.

### **Benefits of the Prototype Pattern**

* **Avoids Costly Initialization** : Allows creating new objects by copying existing ones, which can be more efficient than creating objects from scratch.
* **Simplifies Object Creation** : Provides a simple way to create complex objects with many configuration options by copying an existing prototype.
* **Extensible** : You can easily extend the pattern by adding new prototypes that implement the `Clone()` method.
