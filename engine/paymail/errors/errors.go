package pmerrors

import "github.com/bitcoin-sv/spv-wallet/models"

// ErrPaymailHostResponseError is when the paymail host is responding with errors.
var ErrPaymailHostResponseError = models.SPVError{Message: "paymail host is responding with error", StatusCode: 500, Code: "error-paymail-host-error"}

// ErrPaymailHostNotSupportingP2P is when the paymail host is not supporting P2P capabilities.
var ErrPaymailHostNotSupportingP2P = models.SPVError{Message: "paymail host is not supporting P2P capabilities", StatusCode: 400, Code: "error-paymail-host-not-supporting-p2p"}

// ErrPaymailHostInvalidResponse is when the paymail host is responding with invalid response.
var ErrPaymailHostInvalidResponse = models.SPVError{Message: "paymail host invalid response", StatusCode: 500, Code: "error-paymail-host-invalid-response"}

// ErrPaymailDBFailed is when the paymail database operation failed.
var ErrPaymailDBFailed = models.SPVError{Message: "paymail database operation failed", StatusCode: 500, Code: "error-paymail-db-failed"}

// ErrPaymailNotFound is when the paymail is not found.
var ErrPaymailNotFound = models.SPVError{Message: "paymail not found", StatusCode: 404, Code: "error-paymail-not-found"}

// ErrPaymailPKI is when the paymail PKI operation failed.
var ErrPaymailPKI = models.SPVError{Message: "paymail PKI operation failed", StatusCode: 500, Code: "error-paymail-pki"}

// ErrPaymentDestination is when the payment destination operation failed.
var ErrPaymentDestination = models.SPVError{Message: "payment destination operation failed", StatusCode: 500, Code: "error-payment-destination"}

// ErrAddressSave is when the address save operation failed.
var ErrAddressSave = models.SPVError{Message: "address save operation failed", StatusCode: 500, Code: "error-address-save"}

// ErrPaymailMerkleRootVerificationFailed is when merkle root verification could not be processed
var ErrPaymailMerkleRootVerificationFailed = models.SPVError{Message: "merkle root verification could not be processed", StatusCode: 500, Code: "error-paymail-merkle-root-verification-failed"}

// ErrPaymailInvalidMerkleRoots is when merkle roots verification by BHS returns status: INVALID
var ErrPaymailInvalidMerkleRoots = models.SPVError{Message: "invalid merkle roots", StatusCode: 400, Code: "error-paymail-invalid-merkle-roots"}

// ErrParseIncomingTransaction is when the incoming hex transaction could not be parsed
var ErrParseIncomingTransaction = models.SPVError{Message: "incoming hex transaction could not be parsed", StatusCode: 400, Code: "error-parse-incoming-hex-transaction"}

// ErrRecordTransaction is when the transaction could not be recorded
var ErrRecordTransaction = models.SPVError{Message: "transaction could not be recorded", StatusCode: 500, Code: "error-record-transaction"}
