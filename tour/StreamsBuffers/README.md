Intro to streams and buffers :

Data streams are very important to process a large amount of data in smaller chunks at a time. In this article, we are going to make ourselves familiar with some of the basic APIs provided by Go to handle streams and data buffers.

A data stream as the name suggests is a stream of data. We can visualize this as the stream of water flowing from one end to another. Each droplet of water is the packet of data flowing in the stream and we can resize this data packet as per our needs.

If we have a massive file to deal with, we can process the file’s data one packet at a time instead of loading the entire file in the memory. Also, we can write to a file or data store one packet at a time.

Buffer, on the other hand, is a temporary region of volatile memory (RAM) where data can be stored before consuming it. We can read from or write to a data buffer using streams.

Reading from a Data Source
A data-source is a data container (in memory or file storage) that can stream data. This container implements basic interfaces provided by the io package and exposes some methods to read data from.

Writing to a Data Store
A data-store is a data container (in memory or file storage) that can store an incoming stream of data. This container implements basic interfaces provided by the io package and exposes some methods to write data to.

References : 
- https://medium.com/rungo/introduction-to-streams-and-buffers-d148c0cda0ad
