# COMP429PROJ02
Spring 2022 COMP429 PROJ02

ciRDT

For this project we want to implement a simple transport layer protcol: cirdt. Our transport layer will provide reliable transport on top of UDP. 
While it is not necessarily efficient, it still in fact provides the same guarentees that TCP does for the application layer.

Documentation
RDT was covered in detail in week 10 and is also covered well in our textbook. 
cirdt varies slightly in that the sequence numbers are ever increasing (not alternating bit).

Transmit data reliably by implementing a (slightly modified) RDT version 3.0 as covered in class.

Implement a sender and a receiver.

Uni-directional data transmission.

NOTE: Sequence numbers are not byte based, they are block based.

When "packetizing" the data, use 255 byte blocks. (Completed for you)

Acknowledgement numbers are cumulative.

Flags (completed for you)

Use bit flag 0 for a packet carrying DATA.

Use bit flag 1 for an ACK packets.

Use bit flag 10 for FIN packets.

Only need to support a single "connection"

Over localhost on port 9001.

Most of this is completed for you.

Implement the TODO sections.

You can test your reliablity by programming the receiver to randomly drop packets.

RDT Header
+-------------+-------------+--------------------+--------------------+
|   1 byte    |   1 byte    |      2 bytes       |      2 bytes       |
+-------------+-------------+--------------------+--------------------+
| data length |    flags    |       seq no.      |      ack no.       |
+-------------+-------------+--------------------+--------------------+

How to run program
So we have two sides where sender would act as the server and the receiver would be
the user that is trying to get access.

Since one of the text files has something we want to deliver to the other text file, we have to make sure that the
receiver is well-connected and if so pass on the written text in the receiver's text.

There were a few complications in terms of having the file not be able to read and also the way it was sent and
receiving the message. It uses the sequence number and the expected number to keep the connection alive. 

After several tries, it appears that both receiver and sender have been succefully executed. 