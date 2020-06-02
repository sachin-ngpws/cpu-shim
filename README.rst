CPU USE TRACKER WITH FABRIC SHIM
================================

Low level fabric framework used to build smart contracts

Vendoring
---------

To make sure that we have all  the resources run

.. code:: bash

    GO111MODULE=on go mod vendor


Building
--------

To get the ``build`` command to use the vendor folder this will be done in the
docker container

.. code:: bash
  
    go build -mod vendor -o cpu

Env Variables
-------------

For the `Chaincode Server` we are using `env` Variables for that run
Add `CHAINCODE_SERVER_ADDRESS` and `CHAINCODE_ID`

.. code:: bash

    touch chaincode.env

Example
-------

.. code:: bash
    #Chaincode address
    CHAINCODE_SERVER_ADDRESS=cpu-shim:7053
    #ID
    CHAINCODE_ID=cpu-use:...

Docker
------

To build and run the chaincode

`BUILD`

.. code:: bash

    docker build -t ngp/cpu-shim .

`PUSH`

Push the docker image to a container registry

Kubernetes
----------

`CONFIG`

Config the kubectl to point to the cluster

`RUN`

..code:: bash
    
    kubectl create -f deployment.yaml

    kubectl create -f service.yaml


`UPDATE`

Install the chaincode with `Service External IP` and the update the deployment with `CHAINCODE_ID`

..code:: bash

    kubectl apply -f deployment.yaml


