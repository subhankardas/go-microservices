package core

const (
	TRANSACTION_ID = "TRANSACTION_ID"
	ID             = "id"
)

// Error message ID constants
const (
	ENV_ERROR    = "ENV_ERROR"
	CONFIG_ERROR = "CONFIG_ERROR"
	SERVER_ERROR = "SERVER_ERROR"
	DB_ERROR     = "DB_ERROR"
)

// Error messages
const (
	UNABLE_TO_LOAD_ENV              = "unable to load environment"
	UNABLE_TO_LOAD_CONFIG_FILE      = "unable to load config file"
	UNABLE_TO_CREATE_CIPHER         = "unable to create cipher from key"
	UNABLE_TO_READ_CONFIG_FILE      = "unable to read config properties"
	UNABLE_TO_RUN_SERVER            = "unable to run server"
	INVALID_REQUEST_DATA            = "invalid request data"
	UNABLE_TO_CONNECT_DB            = "unable to connect to database"
	UNABLE_TO_READ_ALL_MENU_FROM_DB = "unable to read list of all menu from database"
	UNABLE_TO_ADD_MENU_TO_DB        = "unable to add menu details to database"
)
