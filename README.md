Eventwrite
==========


About
-----

_Eventwrite_ is an application which recieves events and writes them to back-end
storage for analysis, either with real-time stream analysis or bulk analysis.


Architecture
------------

_Eventwrite_ receives events in one of a few ways:

    * Event streams over websocket.¹
    * Event streams over MQTT¹
    * Single-shot event requests

Events are then written in batches to Amazon Kinesis, which then:

    * Stores all events in S3 in ORC format for use with AWS Athena
    * Passes events to Kinesis Analytics for use with real-time analytics

¹Not currently implemented, planned for a future version
