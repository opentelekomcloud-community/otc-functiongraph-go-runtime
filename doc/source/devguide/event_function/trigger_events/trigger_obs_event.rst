OBS Event Source
================

For details, see `Using an OBS Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_obs_trigger.html>`__.

Hints
-----

.. note::

   - Triggers on OBS can only be used for FunctionGraphs in the main Project,
     **not** in Sub-Projects!
   - For each OBS bucket, only one FunctionGraph can be triggered (no multiple
     FunctionGraphs listening on same bucket)

OBS example event
-----------------

.. code-block:: json

   {
     "Records": [
       {
         "eventVersion": "2.0",
         "eventSource": "aws:s3",
         "awsRegion": "eu-de",
         "eventTime": "2024-12-02T09:49:37.939Z",
         "eventName": "ObjectCreated:Put",
         "userIdentity": { "principalId": "00a4a44afa9f43ceb520c82166407d58" },
         "requestParameters": { "sourceIPAddress": "172.27.48.44" },
         "responseElements": {
           "x-amz-request-id": "4e788bb2744a49a1ab04148b47bf51dc",
           "x-amz-id-2": "s7A0rE7V5QL/mQEYrVUdcFPrBcb1wrieHbjrAs4Gy9J09f2uLOC9deJSIhmXKGET"
         },
         "s3": {
           "s3SchemaVersion": "1.0",
           "configurationId": "8ba3f66e-aa19-459d-8f4d-337998c50dc9",
           "bucket": {
             "name": "fg-sample-obs-go-runtime-sdk",
             "ownerIdentity": {
               "PrincipalId": "65b49ad6e14149e0b01abde940685697"
             },
             "arn": "arn:aws:s3:::fg-sample-obs-go-runtime-sdk"
           },
           "object": {
             "key": "test2.txt",
             "eTag": "5d41402abc4b2a76b9719d911017c592",
             "size": 5,
             "versionId": "null",
             "sequencer": "0000000019386C6871FB059ED0000000"
           }
         }
       }
     ]
   }

Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 20 15 30 35

   * - Parameter
     - Type
     - Example Value
     - Description
   * - eventVersion
     - String
     - 2.0
     - Event version
   * - eventSource
     - String
     - aws:s3
     - Event source
   * - awsRegion
     - String
     - eu-de
     - AWS region
   * - eventTime
     - String
     - 2024-12-02T09:49:37.939Z
     - Time when an event occurs
   * - eventName
     - String
     - ObjectCreated:Put
     - See below
   * - userIdentity
     - String
     - { "principalId": "00a4a44afa9f43ceb520c82166407d58" }
     - User identity information
   * - requestParameters
     - String
     - { "sourceIPAddress": "172.27.48.44" }
     - Request parameters
   * - responseElements
     - String
     - { "x-amz-request-id": "...", "x-amz-id-2": "..." }
     - Response elements
   * - s3
     - Object
     -
     - See below

Possible values for eventName
-----------------------------

.. list-table::
   :header-rows: 1
   :widths: 40 60

   * - EventName
     - Description
   * - | ObjectCreated:Put
       | ObjectCreated:Post
       | ObjectCreated:Copy
     - Operations such as PUT, POST, and COPY can create an object. With these
       event types, you can enable notifications when an object is created
       using a specific API operation.
   * - ObjectCreated:CompleteMultipartUpload
     - ObjectCreated:CompleteMultipartUpload includes objects that are created
       using UploadPartCopy for Copy operations.
   * - | ObjectRemoved:Delete
       | ObjectRemoved:DeleteMarkerCreated
     - By using the ObjectRemoved event types, you can enable notification when
       an object or a batch of objects is removed from a bucket. You can
       request notification when an object is deleted or a versioned object is
       permanently deleted by using the s3:ObjectRemoved:Delete event type.
       Alternatively, you can request notification when a delete marker is
       created for a versioned object using
       s3:ObjectRemoved:DeleteMarkerCreated. These event notifications don't
       alert you for automatic deletes from lifecycle configurations or from
       failed operations.

Parameter s3
------------

.. list-table::
   :header-rows: 1
   :widths: 25 15 30 30

   * - Parameter
     - Type
     - Example Value
     - Description
   * - s3SchemaVersion
     - String
     - 1.0
     - S3 schema version
   * - configurationId
     - String
     - 8ba3f66e-aa19-459d-8f4d-337998c50dc9
     - Configuration ID
   * - bucket
     - Object
     -
     - See below

Parameter bucket
----------------

.. list-table::
   :header-rows: 1
   :widths: 25 15 30 30

   * - Parameter
     - Type
     - Example Value
     - Description
   * - name
     - String
     - fg-sample-obs-go-runtime-sdk
     - Name of the bucket
   * - ownerIdentity
     - String
     - { "PrincipalId": "65b49ad6e14149e0b01abde940685697" }
     - Owner identity information
   * - arn
     - String
     - arn:aws:s3:::fg-sample-obs-go-runtime-sdk
     - Amazon Resource Name

Parameter object
----------------

.. list-table::
   :header-rows: 1
   :widths: 25 15 30 30

   * - Parameter
     - Type
     - Example Value
     - Description
   * - key
     - String
     - test2.txt
     - The name that has been assigned to an object.
   * - eTag
     - String
     - 5d41402abc4b2a76b9719d911017c592
     - The entity tag is a hash of the object.
   * - size
     - Long
     - 5
     - Size in bytes of the object
   * - versionId
     - String
     - "null"
     - Version ID
   * - sequencer
     - String
     - 0000000019386C6871FB059ED0000000
     - Sequencer

Example
-------

.. literalinclude:: /../../samples-events/obss3/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/obss3/handler.go>`
