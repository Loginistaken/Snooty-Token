// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SnootyToken is ERC20, Ownable {
    uint256 public burnRate = 2;
    uint256 public maintenanceFeeRate = 0.002;
    uint256 public teamProfitRate = 1;

    uint256 public constant TOTAL_SUPPLY = 64000000 * 10 ** decimals();
    uint256 public constant OWNER_MINT = 1000000 * 10 ** decimals();

    address public teamAddress;
    address public medievalVault;
    address public userVault;

    string public puzzleAnswer = "defaultPuzzleAnswer"; // Placeholder for the puzzle answer
    mapping(address => bool) public solvedPuzzle;

    constructor(address _teamAddress, address _medievalVault, address _userVault) ERC20("Snooty Token", "SFT") {
        _mint(msg.sender, TOTAL_SUPPLY - OWNER_MINT);

        medievalVault = _medievalVault;
        _mint(medievalVault, OWNER_MINT);

        userVault = _userVault;
        _mint(userVault, TOTAL_SUPPLY - OWNER_MINT);

        teamAddress = _teamAddress;
    }

    function _transfer(address sender, address recipient, uint256 amount) internal override {
        uint256 burnAmount = (amount * burnRate) / 100;
        uint256 maintenanceFee = (amount * maintenanceFeeRate) / 10000;
        uint256 teamProfit = (burnAmount * teamProfitRate) / 10000;

        uint256 totalFee = burnAmount + maintenanceFee + teamProfit;

        require(amount > totalFee, "Transfer amount exceeds fee");

        _burn(sender, burnAmount);
        _transfer(sender, owner(), maintenanceFee);
        _transfer(sender, teamAddress, teamProfit);

        super._transfer(sender, recipient, amount - totalFee);
    }
Hardhat Configuration: hardhat.config.js

require('@nomiclabs/hardhat-ethers');
require('dotenv').config();

module.exports = {
  solidity: "0.8.0",
  networks: {
    sepolia: {
      url: process.env.INFURA_URL,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
      chainId: 11155111,
    },
  },
};

Hardhat Deployment Script: scripts/deploy.js

const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const teamAddress = "0xYourTeamAddressHere";
    const medievalVault = "0xMedievalVaultAddressHere";
    const userVault = "0xUserVaultAddressHere";

    const [deployer] = await hre.ethers.getSigners();

    console.log("Deploying contracts with the account:", deployer.address);

    const SnootyToken = await hre.ethers.getContractFactory("SnootyToken");
    const snootyToken = await SnootyToken.deploy(teamAddress, medievalVault, userVault);

    console.log("Snooty Token deployed to:", snootyToken.address);

    const contractAddress = snootyToken.address;
    fs.appendFileSync('.env', `CONTRACT_ADDRESS=${contractAddress}\n`);

    process.exit(0);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

Python API Integration: app.py

import os
from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
from web3 import Web3
import requests
from bs4 import BeautifulSoup

load_dotenv()

app = FastAPI()

w3 = Web3(Web3.HTTPProvider(os.getenv("INFURA_URL")))
contract_address = os.getenv("CONTRACT_ADDRESS")
private_key = os.getenv("PRIVATE_KEY")
account = w3.eth.account.privateKeyToAccount(private_key)

contract_abi = [ /* Add your ABI here */ ]
contract = w3.eth.contract(address=contract_address, abi=contract_abi)

def scrape_vault_address(url: str):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        vault_address = soup.find('div', {'class': 'vault-address'}).text
        return vault_address
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while scraping the data.")

@app.post("/create-vault/")
async def create_vault_and_mint_tokens():
    try:
        vault_address = scrape_vault_address('https://example.com/vault-address')

        user_address = account.address
        nonce = w3.eth.getTransactionCount(user_address)

        txn = contract.functions._mint(vault_address, 1000000 * 10**18).buildTransaction({
            'chainId': 11155111,  # Sepolia test network
            'gas': 2000000,
            'gasPrice': w3.toWei('10', 'gwei'),
            'nonce': nonce,
        })

        signed_txn = w3.eth.account.signTransaction(txn, private_key)
        txn_hash = w3.eth.sendRawTransaction(signed_txn.rawTransaction)
        receipt = w3.eth.waitForTransactionReceipt(txn_hash)

        return JSONResponse(status_code=200, content={"message": "Vault created, tokens minted!", "tx_hash": txn_hash.hex()})
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/scrape-data/{url}")
async def scrape_data(url: str):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        links = [a['href'] for a in soup.find_all('a', href=True)]
        return {"links": links}
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while scraping data.")

Transition Between PowerShell and Terminal:

# To execute the deployment and integration steps with terminal-to-PowerShell transitions, follow the commands below:
# In PowerShell, start the following:

# 1. Ensure Hardhat and dependencies are installed
npm install --save-dev hardhat

# 2. Run Hardhat deployment on Sepolia network
npx hardhat run scripts/deploy.js --network sepolia

# 3. Switch to terminal for further operations
exit

# 4. In terminal, run the FastAPI server
uvicorn app:app --reload
    // Puzzle-solving function
    function solvePuzzle(string memory _answer) external {
        require(!solvedPuzzle[msg.sender], "You have already solved the puzzle");
        require(keccak256(abi.encodePacked(_answer)) == keccak256(abi.encodePacked(puzzleAnswer)), "Incorrect answer");

        solvedPuzzle[msg.sender] = true;

        // Minting a small token amount after solving the puzzle
        _mint(msg.sender, 0.000000000000031415 * 10 ** decimals()); // Mint the updated amount after solving the puzzle
    }

    function updateBurnRate(uint256 newBurnRate) external onlyOwner {
        require(newBurnRate <= 50, "Burn rate cannot exceed 50%");
        burnRate = newBurnRate;
    }

    function updateMaintenanceFeeRate(uint256 newMaintenanceFeeRate) external onlyOwner {
        require(newMaintenanceFeeRate <= 100, "Maintenance fee rate cannot exceed 100%");
        maintenanceFeeRate = newMaintenanceFeeRate;
    }

    function updateTeamAddress(address newTeamAddress) external onlyOwner {
        teamAddress = newTeamAddress;
    }

    function burnRateAcknowledgment() external pure returns (string memory) {
        return "Burn rate can be updated up to 50%. Higher rates can be risky, and it is a best practice to keep the burn rate lower.";
    }
}
const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const teamAddress = "0xYourTeamAddressHere";
    const medievalVault = "0xMedievalVaultAddressHere";
    const userVault = "0xUserVaultAddressHere";

    const [deployer] = await hre.ethers.getSigners();

    console.log("Deploying contracts with the account:", deployer.address);

    const SnootyToken = await hre.ethers.getContractFactory("SnootyToken");
    const snootyToken = await SnootyToken.deploy(teamAddress, medievalVault, userVault);

    console.log("Snooty Token deployed to:", snootyToken.address);

    const contractAddress = snootyToken.address;
    fs.appendFileSync('.env', `CONTRACT_ADDRESS=${contractAddress}\n`);

    process.exit(0);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const teamAddress = "0xYourTeamAddressHere"; // Replace with your actual team address
    const medievalVault = "0xMedievalVaultAddressHere"; // Replace with your actual medieval vault address
    const userVault = "0xUserVaultAddressHere"; // Replace with your actual user vault address

    const [deployer] = await hre.ethers.getSigners();

    console.log("Deploying contracts with the account:", deployer.address);

    const SnootyToken = await hre.ethers.getContractFactory("SnootyToken");
    const snootyToken = await SnootyToken.deploy(teamAddress, medievalVault, userVault);

    console.log("Snooty Token deployed to:", snootyToken.address);

    // Save contract address to .env
    const contractAddress = snootyToken.address;
    fs.appendFileSync('.env', `CONTRACT_ADDRESS=${contractAddress}\n`);

    process.exit(0);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
require('@nomiclabs/hardhat-ethers');
require('dotenv').config();

module.exports = {
  solidity: "0.8.0",
  networks: {
    rinkeby: {
      url: process.env.INFURA_URL,
      accounts: [`0x${process.env.PRIVATE_KEY}`]
    }
  }
};

Python Backend (FastAPI with Google Cloud Lambda & Secret Management)

import os
from dotenv import load_dotenv
from cryptography.fernet import Fernet
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
from web3 import Web3
import requests
from bs4 import BeautifulSoup
from google.cloud import secretmanager
from google.cloud import storage

load_dotenv()
import os
from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
from web3 import Web3
import requests
from bs4 import BeautifulSoup

INFURA_URL=https://rinkeby.infura.io/v3/YOUR_INFURA_PROJECT_ID
PRIVATE_KEY=your_private_key_here
CONTRACT_ADDRESS=0xYourContractAddressHere
LAMBDA_PASSWORD=your_lambda_password
LAMBDA_USERNAME=your_lambda_username
ALLOWED_DOMAIN=your_allowed_domain.com

# Load environment variables from .env file
load_dotenv()

# Initialize FastAPI app
app = FastAPI()

# Web3 setup for interacting with the Ethereum blockchain
w3 = Web3(Web3.HTTPProvider(os.getenv("INFURA_URL")))  # Infura URL
contract_address = "0xYourContractAddressHere"  # Replace with your contract address
private_key = os.getenv("PRIVATE_KEY")  # Your private key
account = w3.eth.account.privateKeyToAccount(private_key)

# List of numbers (could be part of the puzzle)
numbers = [
    3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679,
    6.2831853071795864769252867665590057683943387987502116419498891846156328125724179972560696506842341358,
]
@app.post("/solve-puzzle/")
async def solve_puzzle(number: float):
    # Validate puzzle input (dynamic interaction)
    valid_numbers = [3.1415, 2.71828, 1.61803]  # Example valid numbers
    if number not in valid_numbers:
        raise HTTPException(status_code=400, detail="Incorrect number, try again.")

   # Get user address and build the transaction to interact with the smart contract
user_address = account.address
nonce = w3.eth.getTransactionCount(user_address)
txn = contract.functions.solvePuzzle().buildTransaction({
    'chainId': 11155111,  # Sepolia test network
    'gas': 2000000,
    'gasPrice': w3.toWei('10', 'gwei'),
    'nonce': nonce,
})

# Sign the transaction and send it
signed_txn = w3.eth.account.signTransaction(txn, private_key)
txn_hash = w3.eth.sendRawTransaction(signed_txn.rawTransaction)

# Wait for the transaction receipt
receipt = w3.eth.waitForTransactionReceipt(txn_hash)

return JSONResponse(status_code=200, content={"message": "Puzzle solved, tokens minted!", "tx_hash": txn_hash.hex()})

    })

    signed_txn = w3.eth.account.signTransaction(txn, private_key)
    txn_hash = w3.eth.sendRawTransaction(signed_txn.rawTransaction)
    receipt = w3.eth.waitForTransactionReceipt(txn_hash)

    return JSONResponse(status_code=200, content={"message": "Puzzle solved, tokens minted!", "tx_hash": txn_hash.hex()})
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Snooty Token</title>
</head>
<body>
    <h1>Snooty Token Info</h1>
    <div id="token-info">
        <p>Token Name: Snooty Token</p>
        <p>Symbol: SFT</p>
        <p>Balance: <span id="balance"></span></p>
    </div>

    <h2>Solve the Puzzle</h2>
    <input type="number" id="puzzle-input" placeholder="Enter Number">
    <button onclick="solvePuzzle()">Solve Puzzle</button>

    <script>
        async function solvePuzzle() {
            const number = document.getElementById('puzzle-input').value;
            const response = await fetch('/solve-puzzle/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ number: parseFloat(number) })
            });
            const data = await response.json();
            alert(data.message);
        }
    </script>
</body>
</html>

# Your contract ABI (replace with the actual ABI)
contract_abi = [ /* Add the ABI of your contract here */ ]

# Create contract instance
contract = w3.eth.contract(address=contract_address, abi=contract_abi)

# API to solve the puzzle by sending a correct number
@app.post("/solve-puzzle/")
async def solve_puzzle(number: float):
    if number not in numbers:
        raise HTTPException(status_code=400, detail="Incorrect number, try again.")

    # Get user address and build the transaction to interact with the smart contract
    user_address = account.address
    nonce = w3.eth.getTransactionCount(user_address)
    txn = contract.functions.solvePuzzle().buildTransaction({
        'chainId': 4,  # Rinkeby test network (use 1 for mainnet)
        'gas': 2000000,
        'gasPrice': w3.toWei('10', 'gwei'),
        'nonce': nonce,
    })

    # Sign the transaction and send it
    signed_txn = w3.eth.account.signTransaction(txn, private_key)
    txn_hash = w3.eth.sendRawTransaction(signed_txn.rawTransaction)

    # Wait for the transaction receipt
    receipt = w3.eth.waitForTransactionReceipt(txn_hash)

    return JSONResponse(status_code=200, content={"message": "Puzzle solved, tokens minted!", "tx_hash": txn_hash.hex()})

# API to scrape data from a given URL
@app.get("/scrape-data/{url}")
async def scrape_data(url: str):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        links = [a['href'] for a in soup.find_all('a', href=True)]
        return {"links": links}
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while scraping data.")

# API to check the token balance of an address
@app.get("/token_balance/{address}")
async def get_token_balance(address: str):
    try:
        balance = contract.functions.balanceOf(address).call()
        return {"balance": balance / 10**18}  # Convert from wei to ether/token decimals
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while fetching token balance.")


# Initialize FastAPI
app = FastAPI()

# Encryption setup using Fernet
SECRET_KEY = os.getenv("SECRET_KEY")
cipher_suite = Fernet(SECRET_KEY)

w3 = Web3(Web3.HTTPProvider(os.getenv("INFURA_URL")))
contract_address = os.getenv("CONTRACT_ADDRESS")
private_key = os.getenv("PRIVATE_KEY")
account = w3.eth.account.privateKeyToAccount(private_key)

# Google Cloud Secret Manager client
def get_secret(secret_name):
    client = secretmanager.SecretManagerServiceClient()
    project_id = os.getenv('GOOGLE_CLOUD_PROJECT')
    secret_version = f"projects/{project_id}/secrets/{secret_name}/versions/latest"
    response = client.access_secret_version(name=secret_version)
    secret_data = response.payload.data.decode('UTF-8')
    return secret_data

# Decrypt environment variables securely
def load_secure_env():
    encrypted_private_key = os.getenv("ENCRYPTED_PRIVATE_KEY")
    decrypted_private_key = cipher_suite.decrypt(encrypted_private_key.encode()).decode()
    return decrypted_private_key

# Scrape vault address from a webpage
def scrape_vault_address(url: str):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        vault_address = soup.find('div', {'class': 'vault-address'}).text
        return vault_address
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while scraping the data.")

# API route to create a profile and mint tokens
@app.post("/create-profile/")
async def create_profile():
    try:
        user_address = load_secure_env()  # Retrieve the private key securely
        vault_address = scrape_vault_address('https://example.com/vault-address')

        nonce = w3.eth.getTransactionCount(user_address)
        txn = contract.functions._mint(vault_address, 1000000 * 10**18).buildTransaction({
            'chainId': 4,
            'gas': 2000000,
            'gasPrice': w3.toWei('10', 'gwei'),
            'nonce': nonce,
        })

        signed_txn = w3.eth.account.signTransaction(txn, private_key)
        txn_hash = w3.eth.sendRawTransaction(signed_txn.rawTransaction)
        receipt = w3.eth.waitForTransactionReceipt(txn_hash)

        return JSONResponse(status_code=200, content={"message": "Profile created successfully!", "tx_hash": txn_hash.hex()})
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# API to scrape data from a URL
@app.get("/scrape-data/{url}")
async def scrape_data(url: str):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        links = [a['href'] for a in soup.find_all('a', href=True)]
        return {"links": links}
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while scraping data.")

# Encrypt and store environment variables
def encrypt_and_store_env_variables():
    private_key = os.getenv("PRIVATE_KEY")
    encrypted_key = cipher_suite.encrypt(private_key.encode()).decode()

    with open(".env", "a") as env_file:
        env_file.write(f"ENCRYPTED_PRIVATE_KEY={encrypted_key}\n")

Google Cloud Setup for Secure Management

    Google Cloud KMS: Use Google Cloud Key Management to store the encryption key securely.

    Secret Manager: Store the private key, encrypted tokens, and other sensitive information in Google Cloud Secret Manager.

Shell Commands for Setup

# Setting up environment variables
export INFURA_API_KEY="your_infura_api_key"
export QUICKNODE_API_KEY="your_quicknode_api_key"
export CONTRACT_ADDRESS="your_contract_address"
export ETH_PRIVATE_KEY="your_private_key"
export GOOGLE_CLOUD_PROJECT="your_project_id"
export SECRET_KEY="your_secret_key_for_encryption"
// Steps Implementation

    // 1. Deploy and Setup: Deploy the contract with team, medieval, and user vault addresses.
    // 2. Puzzle Interaction: Users can solve the puzzle to mint tokens.
    // 3. Token Distribution: Tokens are distributed based on burn, maintenance, and team profit rates.
    // 4. Burn and Fee Application: The burn fee and maintenance fee are applied on every transfer.
    // 5. Safe Burn Rate: The burn rate is limited to a maximum of 50% to ensure safe tokenomics.
    // 6. Puzzle Answer: The puzzle answer "42" can be changed as needed.
    // 7. Legal Acknowledgment: Acknowledgment for burn rate and its limits.
# Setup for contract deployment and testing
echo "Setting up environment variables..."
source setup_environment.sh

echo "Deploying the contract..."
npx hardhat run --network rinkeby scripts/deploy.js

echo "Running tests..."
npx hardhat test

echo "Cleaning up..."
./clean_up.sh
