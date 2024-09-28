# Adapter Desgin Pattern in Go

The **Adapter** design pattern is a structural design pattern that allows objects with incompatible interfaces to work together. The Adapter acts as a wrapper around an existing class, converting its interface into one that a client expects. This pattern is particularly useful when integrating third-party libraries, legacy code, or when class interfaces are incompatible.

## Overview

The Adapter pattern involves:

1. **Target Interface** : This is the interface that the client expects and interacts with.
2. **Adaptee** : An existing class that needs to be adapted to the target interface.
3. **Adapter** : A wrapper class that implements the target interface and translates the client's requests to the adaptee's methods.

## Benefits of the Adapter Pattern

* **Decoupling** : The client code (`main` function) does not need to know the specifics of the different socket types. It only interacts with the adapter through the `Charger` interface.
* **Reusability** : Allows integration of new types of sockets without modifying the existing code, as long as they implement the `Socket` interface.
* **Flexibility** : Easily adapt incompatible interfaces to work together by introducing new adapters.

## When to Use the Adapter Pattern

* When you need to integrate with an existing class that has an incompatible interface.
* When you want to create reusable classes that can work with unrelated classes or those that have incompatible interfaces.
* When you need to use several existing subclasses, but it's impractical to adapt their interface by subclassing every one.
