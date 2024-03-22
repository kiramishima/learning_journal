# Liskov Substitution Principle with Golang

> When extending a class, remember that you should be able to pass objects of the subclass in place of objects of the parent class without breaking the client code.

Let’s go over this checklist in detail.

- **Parameter types in a method of a subclass should match or be more abstract than parameter types in the method of the superclass**. Sounds confusing? Let’s have an example.

    - Say there’s a class with a method that’s supposed to feed cats: `feed(Cat c)` . Client code always passes cat objects into this method.

    - **Good**: Say you created a subclass that overrode the method so that it can feed any animal (a superclass of cats): `feed(Animal c)` . Now if you pass an object of this subclass instead of an object of the superclass to the client code,
everything would still work fine. The method can feed all animals, so it can still feed any cat passed by the client.

    - **Bad**: You created another subclass and restricted the feeding method to only accept Bengal cats (a subclass of cats): `feed(BengalCat c)`. What will happen to the client code if you link it with an object like this instead of with the original class? Since the method can only feed a specific breed of cats, it won’t serve generic cats passed by the client, breaking all related functionality.

- **The return type in a method of a subclass should match or be a subtype of the return type in the method of the superclass**. As you can see, requirements for a return type are inverse to requirements for parameter types.

## Example
Let’s look at an example of a hierarchy of document classes that violates the substitution principle.

BEFORE

```go
type ReadOnlyDocument interface {
    save()
}

type Document struct {
    data any
    filename string
}
var _ ReadOnlyDocument = (*Document)(nil) // aplica la interface

func (doc *Document) open() {}
func (doc *Document) save() {}

type Projects struct {
    documents []Document
}
func (proj *Projects) openAll() {
    for _, doc := range proj.documents {
        doc.open()
    }
}
func (proj *Projects) saveAll() {
    for _, doc := range proj.documents {
        _, ok := doc.(ReadOnlyDocument) // Este es un ejemplo :P, puede que este mal mi reinterpretación
        if !ok {
            doc.save()
        }
    }
}
```

The save method in the ReadOnlyDocuments subclass throws an exception if someone tries to call it. The base method
doesn’t have this restriction. This means that the client code will break if we don’t check the document type before saving it

AFTER

```go
type WritableDocument interface {
    save()
}

type Document struct {
    data any
    filename string
    WritableDocument
}

func (doc *Document) open() {}

type Project struct {
    allDocs []Document
    writableDocs []WritableDocument
}
func (proj *Project) openAll() {
    for _, doc := range proj.allDocs {
        doc.open()
    }
}
func (proj *Projects) saveAll() {
    for _, doc := range proj.writableDocs {
        doc.save()
    }
}
```