# EmployeeDB
Employee Database using hyperledger

A basic Hyperledger-Fabric App that  can:

Query about an existing Employee credentials.
Change/Update their credentials and
Add a new employee to the DataBase.

PROJECT FLOW
============

1.Generation of crypto/certificate using cryptogen
2.Generation of Configuration Transaction using configtxgen
3.Bring up the nodes based on what is defined in docker-compose file
4.Use CLI to setup the First Network
5.Use CLI to install and instantiate the chaincode
6.Use CLI to invoke the chaincode
7.Tear down all the setup
