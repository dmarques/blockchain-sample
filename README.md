# Blockchain sample API
A simple Blockchain example concept implemented in an API made with Golang



**Important**: When you run the application, a temporary json file will be created in the folder ```/tmp/chain.json```

#### Checking the chain

In the chain there will be only the genesis block at the beginning.

GET: ```/chain```

**Response**
```json
[
    {
        "index": 0,
        "date_created": "2020-07-05T01:32:46.563182+09:00",
        "data": "genesis_block",
        "previous_hash": "",
        "hash": "bdb51048cd695778ce0cb1825345689515120148cff22b49d6428e77e5072270"
    }
]

```

#### Create and Add a new block to the chain

POST: ```/block```

**Body**
```json
{
    "data": "Transaction Data"
}
```


If you check the chain again:

GET: ```/chain```

**Response**
```json
[
    {
        "index": 0,
        "date_created": "2020-07-05T01:32:46.563182+09:00",
        "data": "genesis_block",
        "previous_hash": "",
        "hash": "bdb51048cd695778ce0cb1825345689515120148cff22b49d6428e77e5072270"
    },
    {
        "index": 1,
        "date_created": "2020-07-05T01:32:55.7546+09:00",
        "data": "Transaction Data",
        "previous_hash": "bdb51048cd695778ce0cb1825345689515120148cff22b49d6428e77e5072270",
        "hash": "fc5079cbeab3a8091208749652ac093565e710a540aeea18d8a34ed0df756d6b"
    }
]
```
