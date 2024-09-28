# Flyweight Design Patterns in Go

The **Flyweight** design pattern is a structural pattern that aims to minimize memory usage by sharing as much data as possible between similar objects. It is particularly useful when a program needs to create a large number of objects that share common properties.

In the Flyweight pattern, objects are divided into:

1. **Intrinsic State** : The state that is shared between multiple objects (e.g., a common shape, color, or texture).
2. **Extrinsic State** : The state that varies between objects and is provided externally (e.g., position, size).

### **When to Use the Flyweight Pattern**

* When you have a large number of similar objects that consume a lot of memory.
* When objects can share common intrinsic states to reduce memory footprint.
* When you need to create a large number of objects, but most of their states are repetitive.

### **Advantages of the Flyweight Pattern**

* **Memory Efficiency** : Reduces memory usage by sharing common intrinsic states between objects.
* **Performance** : Can significantly improve the performance of applications that need to create a large number of similar objects.
* **Decoupling** : Separates the intrinsic state from the extrinsic state, making objects more lightweight and flexible.

### **When to Avoid the Flyweight Pattern**

* When objects have very few shared states, and most states are unique, making sharing impractical.
* When object creation is not a performance bottleneck or does not lead to excessive memory consumption.

### **Summary**

The **Flyweight** design pattern in Go allows you to efficiently share common states (intrinsic states) between objects and manage unique states (extrinsic states) externally. In this example, we used a `CircleFactory` to create and reuse circle objects based on their color, reducing memory usage when drawing circles with different positions and sizes.

The Flyweight pattern is particularly useful for applications that need to handle a large number of similar objects, such as rendering graphics, managing UI components, or implementing caching mechanisms.
