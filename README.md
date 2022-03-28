# Golang

![Golang Image](golang.png)

# Hash Tables

* Hash table is one of the most important complex data structures in computer science. Data is stored in hash table in an associative manner. 

* In a hash table, each data value has its own **unique** index which provides very fast access to data if we know the index of the desired data.

## Hashing

* Hashing is a technique to convert a range of key values into a range of indexes of an array. We're going to use modulo operator to get a range of key values.

## Hash Function

* Hash function is a function which provides necessary **hash codes** for data which is going to be stored in hash table.

![Hashing](https://upload.wikimedia.org/wikipedia/commons/7/7d/Hash_table_3_1_1_0_1_0_0_SP.svg)

The picture is taken from [Hash table - Wikipedia](https://en.wikipedia.org/wiki/Hash_table) . 

## Compression function

* We can use **compression function**s to map hash codes to an integer in ***[0, N-1]*** where N is the size of the hash table. The result of this operation is an index to the hash table.

* There are some methods for compression functions to use.

### 1. The division method

The index is calculated as follows, i mod N, where N is the size of the hash table. N is recommended to be a prime number. Using a prime number does not guarantee a collision-free hashing, however, decreases the chance to 1/N (e.g., hash codes 2*37+3 and 5*37+3 would collide.).

### 2. The MAD method

MAD is an abbreviation for multiply-add-and-divide. With this method, a hash code i is mapped to ***[(ai + b) mod p]*** mod N where N is the size of the hash table, p is a prime number larger than N, a and b are integers chosen from [0, p-1] randomly and a > 0 .

## Collisions

* Collisions occur when there are two distinct keys (***k1*** and ***k2*** ) whose hash functions create the same value (***h(k1) = h(k2)***).

* There are some methods to fixed these kind of problems. One of them is **open addressing(probing)** and the other one is **seperate chaining**.

* In this implementation, we use **seperate chaining**. However, **probing** is also very important topic to deal with collisions. 

### Seperate chaining

* A hash table is comprised of buckets, and in each bucket there can be a small map instance implemented with a list (an unsorted map).

* For each key-value (k,v) pair in a bucket, h(k) would be the same.

![Seperate Chaining](https://upload.wikimedia.org/wikipedia/commons/d/d0/Hash_table_5_0_1_1_1_1_1_LL.svg)

* Generally linked-lists would be very nice choise to store data in buckets.

#### Load Factor

* The expected size of a bucket is ***n/N***, where n is the number of ***(k,v)*** pairs and N is the size of hash table. 

* With a good hash function, core map operations run in ***O(⌈n/N⌉)***.

* ***<u> λ = n/N</u>*** is called the load factor.

* If ***λ*** is **O(1)**, the core map operations run in O(1) expected time.

## Implementation

* In the implementation of hash table, linked-list structure is used for buckets. 

* The implementation can be seen clearly in **HashTables.go**.

```go
    myHashTable := Init()
    list := []string{
        "england",
        "france",
        "germany",
        "italy",
        "netherlands",
        "greece",
        "turkey",
    }

    for _, v := range list {
        myHashTable.Insert(v)
    }

    myHashTable.Delete("italy")
    myHashTable.Delete("germany")
    myHashTable.Delete("england")

    fmt.Println(myHashTable.Search("italy"))
    fmt.Println(myHashTable.Search("france"))
    fmt.Println(myHashTable.Search("england"))
    fmt.Println(myHashTable.Search("turkey"))
```

```[console]
$ go run HashTables.go
false
true
false
true
```

## Licence

MIT
