# Peerdx, a Peer Diagnostics Tool

This tool allows you to run diagnostics for peers by means of analyzing your nodes' address books or directly connecting to the nodes themselves via RPC. In contrast to just retrieving the information via the /net_info endpoint, this program aims to present the data in a more readable way as well as make it possible to compare multiple nodes with each other in terms of peer diversity.

## How to Run Diagnostics on Address Books

In order to compare address books, put them into one directory and type:

```
$ peerdx addrbook --dir <your dir containing address books>
```

The output should look something like this:

```
2020/04/08 18:25:29 Looking for address book json files at /path/to/addrBooks
2020/04/08 18:25:29 Analyzing address books
2020/04/08 18:25:29 A total of 29 different addresses:
2020/04/08 18:25:29 c3469b7fcf414f26bc1e86a9abe019053587422d: addrbook1.json, addrbook2.json
2020/04/08 18:25:29 304ddd76ea750c7d2d72ac40a2525c37b10ad124: addrbook1.json
...
```

If you already know some of the nodes and want to give them a name, you can do so by providing a config file:

```
$ peerdx addrbook --dir <your dir containing address books> --config cfg.json
```
The config file should look like this:
```
{
    "known_ids": [
        {
            "name": "Awesome Node",
            "id": "c3469b7fcf414f26bc1e86a9abe019053587422d"
        }
    ]
}
```
And the output would look like this:
```
2020/04/08 18:25:29 Looking for address book json files at /path/to/addrBooks
2020/04/08 18:25:29 Analyzing address books
2020/04/08 18:25:29 A total of 29 different addresses:
2020/04/08 18:25:29 Awesome Node        : addrbook1.json, addrbook2.json
2020/04/08 18:25:29 304ddd76ea750c7d2d72ac40a2525c37b10ad124: addrbook1.json
...
```

## How to Run Diagnostics via RPC

### Comparing peers of nodes

If you want to compare peers of nodes without collecting all of their address books, you can check via RPC. You have to make the rpc endpoint accessible on the nodes you want to compare.
Then all you have to do is run:

```
$ peerdx rpc compare  <node_1 rpc address> <node_2 rpc address> <...>
```

You can pass in a config file with already known addresses as shown in the address book description.
The output will look exactly the same as when comparing address books.

### Show Peer Info of a Single Node

Additionally, you can take a more detailed look at the node's peers. Just type:

```
$ peerdx rpc list <node rpc address>
```

The result will lokk something like this:
```
OUT sebytza05 [d32432f9...]              : IP 164.64.104.66
                                           send active; receive active
                                           tendermint version 0.33.3

OUT prim-sentry [8aea2394...]            : IP 157.85.239.128
                                           send active; receive active
                                           tendermint version 0.33.3

OUT Nodeasy [3f931256...]                : IP 172.243.73.214
                                           send active; receive active
                                           tendermint version 0.33.3

OUT gunray [49d7656d...]                 : IP 65.222.131.87
                                           send active; receive active
                                           tendermint version 0.33.3
```