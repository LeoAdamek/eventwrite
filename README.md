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

Deployment
----------

Eventwrite deploys natively to Amazon ECS, including AWS Fargate.
A CloudFormation stack template is provided which will create all required
resources and IAM roles to deploy the application on AWS Fargate.

The following cost-incurring resources will be created:

* Kinesis data stream
* S3 Bucket
* Kinesis Firehose delivery stream
* Application Load Balancer
* ECS service
* DynamoDB table

The following cost-free resources will also be created:

* Security groups
* IAM roles & polices

