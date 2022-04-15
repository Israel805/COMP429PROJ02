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