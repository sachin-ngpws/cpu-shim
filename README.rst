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

.. code:: bash

    touch chaincode.env

Docker
------

To build and run the chaincode

`BUILD`

.. code:: bash

    docker build -t ngp/cpu-shim .


`RUN`

..code:: bash

    docker run -it --rm --name cpu-shim --hostname cpu-shim --env-file chaincode.env --network=bridge ngp/cpu-shim

