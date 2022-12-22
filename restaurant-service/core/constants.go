package core

const (
	TRANSACTION_ID = "TRANSACTION_ID"
	ID             = "id"
)

// Error message ID constants
const (
	ENV_LOAD_ERROR      = "ENV_LOAD_ERROR"
	CONFIG_LOAD_ERROR   = "CONFIG_LOAD_ERROR"
	DB_CONNECTION_ERROR = "DB_CONNECTION_ERROR"
	DB_MIGRATION_ERROR  = "DB_MIGRATION_ERROR"
)

// Error messages
const (
	UNABLE_TO_LOAD_ENV              = "unable to load environment"
	UNABLE_TO_LOAD_CONFIG_FILE      = "unable to load config file"
	UNABLE_TO_READ_CONFIG_FILE      = "unable to read config properties"
	INVALID_REQUEST_DATA            = "invalid request data"
	UNABLE_TO_CONNECT_DB            = "unable to connect to database"
	UNABLE_TO_READ_ALL_MENU_FROM_DB = "unable to read list of all menu from database"
	UNABLE_TO_ADD_MENU_TO_DB        = "unable to add menu details to database"
)
