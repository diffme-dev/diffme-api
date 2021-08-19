## Diffme API

WIP - this is an API to compute diffs between documents. 

It serves as a way to easily create audit logs for documents in your system, 
think version control for your documents. 

For example, if you had a document like the below you could easily answer questions like...

```typescript
const Item = {
    name: "Chaga",
    price: 10.50,
    description: "Superfood for your soul",
}
```

1. What changes were made to this document?
2. Who made these changes?
3. Who changed the "name" field on the document?

Even after you perform updates to it. So if the name changed from "Chaga" to "Chagachino", there would be a diff for that giving 
you a complete audit history of everything that happens.

## How does it work?

Diffme tracks document diffs over time by taking snapshots and then using jsondiff to compute patches between documents using the RFC6902 standard. 
It will compare documents fields, nested objects, arrays, etc... and track differences between json objects, and then automatically store 
and index those diffs for fast retrieval later.

TODO: implement a similar concept to git packfiles where snapshots are packaged with diffs to reduce size on disk.

### How can I use this?

In typescript you could use the below code to easily generate a difference with this api.

```typescript
    import axios from "axios"

    const client = axios.create({
        baseURL: "your diffme api url"
    })

    const document = { name: "Chaga", price: 10.50, id: "1"}

    // creates the first snapshot
    client.post("/v1/snapshots", {
        reference_id: document.id,
        data: document,
        editor: "andrew",
    })

    // edit the documents name
    document.name = "Chagachino"
    
    // this would create a diff saying the name was changed from "Chaga" to "Chagachino"
    // by olivia
    client.post("/v1/snapshots", {
        reference_id: document.id,
        data: document,
        editor: "olivia",
    }) 
```

And that is it! It is a simple rest API to easily get diffs of documents over time.

## Why would I want this?

Storing diffs of your documents has a lot of benefits including:

2. Easily searchable changes with elasticsearch.
3. Saves dev time debugging bc you can examine the diffs of documents over time.
4. Know who edited stuff when and where.

You could build this yourself, but it isn't differentiating your business AND this is free and open source so why not?

## API Documentation

TODO:

## Technology

- Go
- Redis (for pubsub)
- Protobufs (for pubsub data encoding)
- Elasticsearch (for searching diffs)
- Mongo (for storing diffs and snapshots of documents)

It exposes an easy to use REST api that you can post a document to, and it will track the diffs over time.



