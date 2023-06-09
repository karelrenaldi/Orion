package constants

// Model Name.
const (
	UserModel                         = "Model-User"
	AccidentModel                     = "Model-Accident"
	CertificateModel                  = "Model-Certificate"
	MaintenanceModel                  = "Model-Maintenance"
	CertificateTransactionModel       = "Model-Certificate-Transaction"
	CertificateTransactionDetailModel = "Model-Certificate-Transaction-Detail"
)

// Transaction Status.
const (
	InProgress         = "in-progress"
	ApprovedByNewOwner = "approved-by-new-owner"
	Success            = "success"
	Canceled           = "canceled"
)

// Org identifier.
const (
	OrionOrg       = "Org1MSP"
	GovermentOrg   = "Org2MSP"
	MaintenanceOrg = "Org3MSP"
)

// Private Data Collection.
const (
	CollectionTransaction       = "collectionTransaction"
	CollectionTransactionDetail = "collectionTransactionDetails"
)
