// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SnootyToken is ERC20, Ownable {
    uint256 public burnRate = 2;
    uint256 public maintenanceFeeRate = 0.002;
    uint256 public teamProfitRate = 1;

   int256 public constant TOTAL_SUPPLY = 64000000 * 10 ** 18;
    uint256 public constant OWNER_MINT = 1000000 * 10 ** 18;

    address public teamAddress;
    address public medievalVault;
    address public userVault;

    string public puzzleAnswer = "defaultPuzzleAnswer"; // Placeholder for the puzzle answer
    mapping(address => bool) public solvedPuzzle;

    constructor(address _teamAddress, address _medievalVault, address _userVault) ERC20("Snooty Token", "SNFT") {
        _mint(msg.sender, TOTAL_SUPPLY - OWNER_MINT);

        medievalVault = _medievalVault;
        _mint(medievalVault, OWNER_MINT);

        userVault = _userVault;
        _mint(userVault, TOTAL_SUPPLY - OWNER_MINT);

        teamAddress = _teamAddress;
    }
#!/bin/bash

# SPDX-License-Identifier: GPL-3.0
# Ensure dependencies are installed
echo "Checking dependencies..."
npm install

# Load environment variables
source .env

# Define variables
TEAM_ADDRESS="0xYourTeamAddressHere"
MEDIEVAL_VAULT="0xMedievalVaultAddressHere"
USER_VAULT="0xUserVaultAddressHere"

# Check if INFURA URL and private key are set in .env
if [[ -z "$INFURA_URL" || -z "$PRIVATE_KEY" ]]; then
  echo "Error: INFURA_URL or PRIVATE_KEY is not set in .env file."
  exit 1
fi

# Deploy the contract via Hardhat
echo "Deploying the contract..."
npx hardhat run --network sepolia scripts/deploy.js

# Store contract address in .env after deployment
DEPLOYED_CONTRACT_ADDRESS=$(npx hardhat run --network sepolia scripts/deploy.js | grep -oP "(?<=Snooty Token deployed to: )(0x[a-fA-F0-9]{40})")
echo "CONTRACT_ADDRESS=$DEPLOYED_CONTRACT_ADDRESS" >> .env

# Run post-deployment PowerShell script
echo "Running post-deployment PowerShell script..."
powershell.exe -File ./scripts/postDeploy.ps1

# Exit the script
echo "Deployment completed successfully. Exiting..."
exit 0
const hre = require("hardhat");
const fs = require('fs');

async function main() {
    // Addresses for team, medieval vault, and user vault
    const teamAddress = "0xYourTeamAddressHere";
    const medievalVault = "0xMedievalVaultAddressHere";
    const userVault = "0xUserVaultAddressHere";

    const [deployer] = await hre.ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    const SnootyToken = await hre.ethers.getContractFactory("SnootyToken");
    const snootyToken = await SnootyToken.deploy(teamAddress, medievalVault, userVault);

    console.log("Snooty Token deployed to:", snootyToken.address);

    // Save contract address in .env
    fs.appendFileSync('.env', `CONTRACT_ADDRESS=${snootyToken.address}\n`);

    return snootyToken;
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
# Example Post Deployment PowerShell Script

Write-Host "Post-deployment script started..."

# Perform actions like notifying the team, or triggering other processes
Write-Host "Notifying team about the contract deployment..."
# (For example, send an email, API call, etc.)

# You could also add logic to interact with your smart contract
# For example, interacting with the contract via web3.js or ethers.js

Write-Host "Post-deployment actions completed."

# End of script
npm install -g slither
slither contracts/SnootyToken.sol

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

const hre = require("hardhat");
const fs = require('fs');
const { exec } = require("child_process");

async function main() {
    const teamAddress = "0xYourTeamAddressHere";
    const medievalVault = "0xMedievalVaultAddressHere";
    const userVault = "0xUserVaultAddressHere";

    const [deployer] = await hre.ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    const SnootyToken = await hre.ethers.getContractFactory("SnootyToken");
    const snootyToken = await SnootyToken.deploy(teamAddress, medievalVault, userVault);

    console.log("Snooty Token deployed to:", snootyToken.address);

    // Store contract address in .env
    fs.appendFileSync('.env', `CONTRACT_ADDRESS=${snootyToken.address}\n`);

    // Auto-exit & trigger PowerShell script
    exec("powershell.exe -File .\\scripts\\postDeploy.ps1", (err, stdout, stderr) => {
        if (err) console.error(`Error: ${err.message}`);
        if (stderr) console.error(`PowerShell Error: ${stderr}`);
        console.log(`PowerShell Output: ${stdout}`);
    });

    process.exit(0);
}

main().catch(error => {
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
        _mint(msg.sender, 0.314150000000000000 * 18 ** decimals()); // Mint the updated amount after solving the puzzle
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
    sepolia: {
      url: process.env.INFURA_URL, // Infura URL for Sepolia
      accounts: [`0x${process.env.PRIVATE_KEY}`], // Private key from .env
      chainId: 11155111, // Sepolia chain ID
    },
  },
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

INFURA_URL=https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID
PRIVATE_KEY=your_private_key_here
CONTRACT_ADDRESS=0xYourContractAddressHere
LAMBDA_PASSWORD=your_lambda_password
LAMBDA_USERNAME=your_lambda_username
ALLOWED_DOMAIN=your_allowed_domain.com

from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
from web3 import Web3
from hashlib import sha256
import math
import random
import os
import requests
from dotenv import load_dotenv
from cryptography.fernet import Fernet
from google.cloud import secretmanager
from bs4 import BeautifulSoup

# Load environment variables from .env file
load_dotenv()

# Initialize FastAPI app
app = FastAPI()

# Web3 setup for interacting with the Ethereum blockchain
w3 = Web3(Web3.HTTPProvider(os.getenv("INFURA_URL")))  # Infura URL
contract_address = os.getenv("CONTRACT_ADDRESS")  # Replace with your contract address
private_key = os.getenv("PRIVATE_KEY")  # Your private key
account = w3.eth.account.privateKeyToAccount(private_key)

# Constants for the puzzle (math and spacetime constants)
pi = math.pi
c = 3e8  # Speed of light in m/s
lunar_cycle = 29.5  # Lunar cycle days

# List of provided numbers
numbers = [
    3.141592653589793, 6.283185307179586, 9.424777960769379, 12.566370614359172, 
    15.707963267948966, 45.237963267948966, 48.379555921538759, 54.662741228718345, 
    60.944840265978473, 67.227045573157946, 73.509148386347530, 89.256074193173765, 
    96.538171964695060, 102.820320831942309, 107.934021437251987, 111.888212378964123,
    118.264972451276451, 124.549851247812578, 131.489572985192758, 135.278412785912758, 
    141.984712758912758, 148.621289758912758, 154.789612758912758, 167.589124785912758,
    173.758912758912758, 179.512758912758912, 184.789127589127589, 192.758912758912758, 
    199.751278591275891, 209.999999999999999
]

# Spatial dimensions (length, width, height) using multiples of pi
length = pi * 2  # Pi scaled by 2 for length
width = pi * 3   # Pi scaled by 3 for width
height = pi * 4  # Pi scaled by 4 for height

# Time as the lunar cycle
time = lunar_cycle  # Time as a sliding variable, the lunar cycle

# Introduce randomness by selecting random numbers from the provided list
random_numbers = random.sample(numbers, 5)  # Pick 5 random numbers from the list

# Equation combining all values (using random numbers for added complexity)
result = (length * width * height * random_numbers[0] * random_numbers[1] * random_numbers[2] * random_numbers[3]) / (time * c)

# Now, hash the result using SHA-256
result_str = str(result)  # Convert the result to a string
hash_object = sha256(result_str.encode())  # Hash the string result
hashed_result = hash_object.hexdigest()  # Get the hexadecimal hash

# Expected hash (This is a random example; you'd want to specify this for verification)
expected_hash = "d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2"

# Compare the computed hash with the expected hash
if hashed_result == expected_hash:
    print("Success: The hashes match!")
else:
    print("Error: The hashes do not match. Try again!")

# API route to solve the puzzle
@app.post("/solve-puzzle/")
async def solve_puzzle(number: float):
    if number not in numbers:
        raise HTTPException(status_code=400, detail="Incorrect number, try again.")

    # Get user address and build the transaction to interact with the smart contract
    user_address = account.address
    nonce = w3.eth.getTransactionCount(user_address)
    contract = w3.eth.contract(address=contract_address, abi=[ /* Add your ABI here */ ])
    
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

# API route to scrape data from a URL
@app.get("/scrape-data/{url}")
async def scrape_data(url: str):
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, 'html.parser')
        links = [a['href'] for a in soup.find_all('a', href=True)]
        return {"links": links}
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while scraping data.")

# API route to check token balance of an address
@app.get("/token_balance/{address}")
async def get_token_balance(address: str):
    try:
        balance = contract.functions.balanceOf(address).call()
        return {"balance": balance / 10**18}  # Convert from wei to ether/token decimals
    except Exception as e:
        raise HTTPException(status_code=500, detail="Error occurred while fetching token balance.")

# Google Cloud Secret Manager client for secure data handling
def get_secret(secret_name):
    client = secretmanager.SecretManagerServiceClient()
    project_id = os.getenv('GOOGLE_CLOUD_PROJECT')
    secret_version = f"projects/{project_id}/secrets/{secret_name}/versions/latest"
    response = client.access_secret_version(name=secret_version)
    secret_data = response.payload.data.decode('UTF-8')
    return secret_data

# Encryption setup using Fernet for private key storage
SECRET_KEY = os.getenv("SECRET_KEY")
cipher_suite = Fernet(SECRET_KEY)

def load_secure_env():
    encrypted_private_key = os.getenv("ENCRYPTED_PRIVATE_KEY")
    decrypted_private_key = cipher_suite.decrypt(encrypted_private_key.encode()).decode()
    return decrypted_private_key

# API route to create a profile and mint tokens
@app.post("/create-profile/")
async def create_profile():
    try:
        user_address = load_secure_env()  # Retrieve the private key securely
        vault_address = scrape_vault_address('https://example.com/vault-address')

        nonce = w3.eth.getTransactionCount(user_address)
        contract = w3.eth.contract(address=contract_address, abi=[ /* Add your ABI here */ ])

        txn = contract.functions._mint(vault_address, 1000000 * 10**18).buildTransaction({
            'chainId': 4,  # Rinkeby test network
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

# Encrypt and store environment variables securely
def encrypt_and_store_env_variables():
    private_key = os.getenv("PRIVATE_KEY")
    encrypted_key = cipher_suite.encrypt(private_key.encode()).decode()

    with open(".env", "a") as env_file:
        env_file.write(f"ENCRYPTED_PRIVATE_KEY={encrypted_key}\n")

# Main entry to run the FastAPI app
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)


    // 1. Deploy and Setup: Deploy the contract with team, medieval, and user vault addresses.
    // 2. Puzzle Interaction: Users can solve the puzzle to mint tokens.
    // 3. Token Distribution: Tokens are distributed based on burn, maintenance, and team profit rates.
    // 4. Burn and Fee Application: The burn fee and maintenance fee are applied on every transfer.
    // 5. Safe Burn Rate: The burn rate is limited to a maximum of 50% to ensure safe tokenomics.
    // 6. Puzzle Answer: The puzzle answer "42" can be changed as needed.
    // 7. Legal Acknowledgment: Acknowledgment for burn rate and its limits.
    <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Snooty Token Legal Agreement</title>
    <style>
        /* Modal Styles */
        .modal {
            display: none; 
            position: fixed; 
            z-index: 1; 
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            overflow: auto;
            background-color: rgb(0, 0, 0); 
            background-color: rgba(0, 0, 0, 0.4);
            padding-top: 60px;
        }

        .modal-content {
            background-color: #fefefe;
            margin: 5% auto;
            padding: 20px;
            border: 1px solid #888;
            width: 80%;
            max-width: 800px;
            overflow-y: auto;
            height: 80%;
        }

        .close {
            color: #aaa;
            float: right;
            font-size: 28px;
            font-weight: bold;
        }

        .close:hover,
        .close:focus {
            color: black;
            text-decoration: none;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <!-- Trigger Button -->
    <button id="showModalBtn">View Snooty Token Legal Agreement</button>

    <!-- Modal Structure -->
    <div id="legalModal" class="modal">
        <div class="modal-content">
            <span class="close">&times;</span>
            <h2>Snooty Token Legal Agreement</h2>
            <p><strong>Version 1.0</strong></p>
            <p>This Agreement ("Agreement") governs the rights, responsibilities, and legal obligations related to the use and ownership of the Snooty Token ("Token"). By interacting with the Token, users and owners agree to comply with the following terms and conditions.</p>

            <h3>1. Ownership and Control Rights</h3>
            <p><strong>1.1 Ownership Rights:</strong> The creator ("Owner") of the Token shall retain full ownership rights and control over the distribution, burn rate, and maintenance fees of the Token, as outlined in the contract. The Owner is responsible for ensuring the Token's integrity and governance, including the rights to mint additional tokens and adjust the burn rate.</p>
            <p><strong>1.2 Burn Rate Control:</strong> The Owner reserves the right to change the Token’s burn rate up to a maximum of 50%. This is a configurable value within the Token contract. The Owner acknowledges that setting the burn rate beyond 50% may be detrimental to the Token's value and sustainability. It is understood that best practice suggests keeping the burn rate significantly lower than the maximum threshold to maintain a healthy ecosystem.</p>
            <p><strong>1.3 Initial Minting and Vault Creation:</strong> The Owner is granted the right to mint the initial supply of the Token, totaling 64,000,000 tokens. Of this supply, 1,000,000 tokens will be reserved in a medieval vault for personal or discretionary use. The Owner also has the right to allocate specific tokens into other vaults, such as a user vault or team address, as per the code.</p>

            <h3>2. Cryptographic Puzzle and Minting Rewards</h3>
            <p><strong>2.1 Puzzle Mechanism:</strong> The Token contract includes a cryptographic puzzle-solving mechanism ("Puzzle"). To interact with the Token, users are required to solve a puzzle defined within the smart contract. Upon solving the puzzle, users are eligible to receive a mint reward. The mint reward is a small quantity of tokens minted and distributed to the user’s address. This reward is pre-defined in the contract and is issued upon the successful solution of the puzzle.</p>
            <p><strong>2.2 Puzzle Answer:</strong> The puzzle answer is initially set by the Owner and stored in the contract. The Owner may update or change the answer to reflect new challenges or rewards. The challenge presented to users is cryptographically secure and verified within the smart contract. Users who provide the correct solution to the puzzle will trigger the minting of tokens to their address.</p>
            <p><strong>2.3 Puzzle Answer Limitations:</strong> Users may only attempt the puzzle once. If a user successfully solves the puzzle, they will be recorded as having "solved" the puzzle, preventing further attempts. The Owner retains the right to modify the mechanics of the puzzle, the reward amount, or even remove the puzzle entirely.</p>

            <h3>3. Burn Rate Acknowledgment and Changes</h3>
            <p><strong>3.1 Burn Rate Modification:</strong> The burn rate of the Token may be modified by the Owner at any time, but it may never exceed 50% of the transaction value. This parameter is intended to control the deflationary effect of the Token. Any adjustment made will be documented and transparently communicated to all Token holders through the contract.</p>
            <p><strong>3.2 Best Practice for Burn Rate:</strong> While the Owner may alter the burn rate, it is strongly recommended that the burn rate remain significantly below 50% to prevent excessive deflation. The burn rate above 10% is discouraged for long-term sustainability.</p>

            <h3>4. API Scraping and Data Management</h3>
            <p><strong>4.1 Google Cloud Integration:</strong> The Token contract interacts with external systems, including API scraping and Web3-based Ethereum interactions, via Google Cloud. This is done to gather data necessary for specific Token operations, including scraping vault addresses, interacting with contract data, and executing blockchain transactions.</p>
            <p><strong>4.2 Web3 and Scraping Mechanism:</strong> The scraping process utilizes Web3 and is facilitated through the Google Cloud Lambda service. The information obtained via scraping is used for minting and updating vault addresses within the Token ecosystem. All actions are executed securely with user consent and according to the latest legal guidelines regarding data collection and privacy.</p>
            <p><strong>4.3 Legal Compliance of API Usage:</strong> The scraping and API integrations comply with all relevant data protection and privacy laws, including GDPR and CCPA (as applicable). Data is used only for the purposes outlined in this contract and will not be shared with third parties without the user’s explicit consent. The Owner ensures the use of encryption and secure communication methods for all data transmission.</p>

            <h3>5. Team Address and Profit Sharing</h3>
            <p><strong>5.1 Team Address Control:</strong> The Owner has the right to define and update the team address within the contract. This address is used to distribute profit shares, and the contract allows for the transfer of specific profits from transactions to the team’s address. The distribution is controlled by the Owner and may be adjusted for operational needs.</p>
            <p><strong>5.2 Profit Sharing Mechanism:</strong> A team profit rate is defined in the contract, which determines the portion of each transaction allocated to the team. This is set to 1% of each transaction. The Owner reserves the right to adjust this profit rate in the future, with the understanding that any changes must be clearly communicated to the community.</p>

            <h3>6. General Disclaimer and Legal Limitations</h3>
            <p><strong>6.1 No Liability:</strong> The Owner of the Token assumes no liability for any losses incurred by users as a result of interacting with the Token or participating in the cryptographic puzzle-solving process. Users are responsible for conducting their own due diligence and ensuring their participation aligns with all applicable local laws and regulations.</p>
            <p><strong>6.2 Risk Acknowledgment:</strong> By using the Token, users acknowledge that the Token’s ecosystem, cryptographic puzzles, burn rate changes, and external API scraping mechanisms involve inherent risks. The Owner makes no representations or warranties about the performance or future value of the Token.</p>
            <p><strong>6.3 Governing Law:</strong> This Agreement shall be governed by and construed in accordance with the laws of the jurisdiction in which the Owner resides, without regard to its conflict of law principles. Any disputes arising from this Agreement shall be resolved in the appropriate courts in the jurisdiction specified.</p>
            <p><strong>6.4 Updates and Modifications:</strong> The Owner reserves the right to update or modify the terms of this Agreement at any time. Any such changes will be communicated through the appropriate channels, including direct notifications to users, smart contract updates, and official documentation.</p>

            <h3>7. Licensing and Legal Considerations</h3>
            <p><strong>7.1 Licensing:</strong> The Snooty Token code is released under the GNU General Public License v3.0 (GPL-3.0). By using this token, users agree to adhere to the GPL-3.0 terms. However, the Owner offers the option to negotiate a proprietary license for use in commercial projects that may not comply with the open-source requirements of the GPL.</p>
            <p><strong>7.2 Ethereum Mining and Token Rights:</strong> This Agreement acknowledges the rights associated with mining Ethereum and ERC-20 tokens. While the smart contract operates within Ethereum's network, it cannot directly enforce real-world laws but will provide transparency regarding token rights and responsibilities. This includes maintaining proper burn rates, profit sharing, and governance as outlined in this agreement.</p>
            <p><strong>7.3 Compliance with Ethereum Network Laws:</strong> Users of this contract should ensure they are compliant with the relevant laws governing Ethereum mining, ERC-20 tokens, and other blockchain-related activities in their jurisdiction.</p>

        </div>
    </div>
    <script>
        // Get the modal
        var modal = document.getElementById("legalModal");
        // Get the button that opens the modal
        var btn = document.getElementById("showModalBtn");
        // Get the <span> element that closes the modal
        var span = document.getElementsByClassName("close")[0];
        // When the user clicks the button, open the modal
        btn.onclick = function() {
            modal.style.display = "block";
        }
        // When the user clicks on <span> (x), close the modal
        span.onclick = function() {
            modal.style.display = "none";
        }
        // When the user clicks anywhere outside of the modal, close it
        window.onclick = function(event) {
            if (event.target == modal) {
                modal.style.display = "none";
            }
        }
    </script>
</body>
</html>

# Setup for contract deployment and testing
echo "Setting up environment variables..."
source setup_environment.sh

echo "Deploying the contract..."
npx hardhat run --network sepolia scripts/deploy.js


echo "Running tests..."
npx hardhat test

echo "Cleaning up..."
./clean_up.sh
