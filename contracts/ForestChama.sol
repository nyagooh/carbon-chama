// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ForestChama {
    // Structure to store credit information
    struct CreditInfo {
        string creditSerial;
        string projectName;
        string registry;
        string ownershipName;
        uint256 timestamp;
    }
    
    // Mapping to store credits by address
    mapping(address => CreditInfo[]) public userCredits;
    
    // Event emitted when a credit is saved
    event CreditSaved(
        address indexed user,
        string creditSerial,
        string projectName,
        string registry,
        string ownershipName,
        uint256 timestamp
    );
    
    // Function to save a new credit
    function saveCredit(
        string memory creditSerial,
        string memory projectName,
        string memory registry,
        string memory ownershipName
    ) public {
        // Create a new credit info
        CreditInfo memory newCredit = CreditInfo({
            creditSerial: creditSerial,
            projectName: projectName,
            registry: registry,
            ownershipName: ownershipName,
            timestamp: block.timestamp
        });
        
        // Add to user's credits
        userCredits[msg.sender].push(newCredit);
        
        // Emit event
        emit CreditSaved(
            msg.sender,
            creditSerial,
            projectName,
            registry,
            ownershipName,
            block.timestamp
        );
    }
    
    // Function to get the number of credits for a user
    function getUserCreditCount(address user) public view returns (uint256) {
        return userCredits[user].length;
    }
    
    // Function to get a specific credit by index
    function getUserCredit(address user, uint256 index) public view returns (
        string memory creditSerial,
        string memory projectName,
        string memory registry,
        string memory ownershipName,
        uint256 timestamp
    ) {
        require(index < userCredits[user].length, "Credit index out of bounds");
        CreditInfo memory credit = userCredits[user][index];
        return (
            credit.creditSerial,
            credit.projectName,
            credit.registry,
            credit.ownershipName,
            credit.timestamp
        );
    }
}