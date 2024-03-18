pragma solidity ^0.8;

interface IComplianceBridge {
    function addVerificationDetails(
        address userAddress,
        uint32 verificationType,
        uint32 issuanceTimestamp,
        uint32 expirationTimestamp,
        bytes memory proofData
    ) external;

    function hasVerification(
        address userAddress,
        uint32 verificationType,
        uint32 expirationTimestamp,
        address[] memory allowedIssuers
    ) external returns (bool);
}

contract ComplianceProxy {
    uint32 constant public VERIFICATION_TYPE = 2;

    function markUserAsVerified(address userAddress) public {
        bytes memory proofData = new bytes(0);

        IComplianceBridge(address(1028)).addVerificationDetails(
            userAddress,
            VERIFICATION_TYPE,
            uint32(block.timestamp % 2**32),
            0,
            proofData
        );
    }
}
