
var express = require('express');
var bodyParser = require('body-parser');
var app = express();
app.use(bodyParser.json());

const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', '..', 'basic-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

app.get('/api/queryAllEmployees', async function (req, res) {
 try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}` );

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');

        const result = await contract.evaluateTransaction('queryAllEmployees');
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        
res.status(200).json({response: result.toString()});
} catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        
res.status(500).json({error: error});        process.exit(1);
    }
}); 


app.get('/api/queryByID/:EmpID', async function (req, res) {
    try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');

        
        const result = await contract.evaluateTransaction('queryByID', req.params.EmpID);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        
res.status(200).json({response: result.toString()});
} catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        
res.status(500).json({error: error});        process.exit(1);
    }
});


app.post('/api/newEmployee/', async function (req, res) {
try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');

        
        
        await contract.submitTransaction('newEmployee', req.body.EmpID, req.body.Name, req.body.Phone, req.body.Email, req.body.Address, req.body.Designation);
        console.log('Transaction has been submitted');
        
res.send('Transaction has been submitted');
// Disconnect from the gateway.
        await gateway.disconnect();
} catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
})

 
app.put('/api/UpdateEmployeePhone/:EmpID', async function (req, res) {
    try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');
        

        
        
        await contract.submitTransaction('UpdateEmployeePhone', req.params.EmpID, req.body.Phone);
        console.log('Transaction has been submitted');
        
res.send('Transaction has been submitted');

        await gateway.disconnect();
} catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    } 
})


app.put('/api/DeleteEmployee/', async function (req, res) {
    try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');

        
        
        await contract.submitTransaction('DeleteEmployee', req.body.EmpID);
        console.log('Transaction has been evaluated');
        
res.send('Transaction has been evaluated');

        await gateway.disconnect();
} catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    } 
})


app.get('/api/HistoryOfEmployees/', async function (req, res) {
    try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');

        
       const result = await contract.evaluateTransaction('HistoryOfEmployees', req.body.EmpID);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        
res.status(200).json({response: result.toString()});
        
res.send('Transaction has been evaluated');

        await gateway.disconnect();
} catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    } 
})


app.get('/api/GetInfobyRange/', async function (req, res) {
    try {

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        const network = await gateway.getNetwork('channel1');

        const contract = network.getContract('Employeedb');

        
        
        const result = await contract.evaluateTransaction('GetInfobyRange', req.body.EmpID, req.body.EmpID1);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        
        //res.status(200).json({response: result.toString()});
        
        return res.send({status:200, message:'Transaction has been evaluated', resp: result.toString()});

        await gateway.disconnect();
} catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    } 
})

app.listen(8080, () => {
    console.log('Server started at port 8080');
});
