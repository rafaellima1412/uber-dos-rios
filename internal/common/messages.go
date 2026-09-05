package common

const (
	// -----------------------------
	// General error messages
	// -----------------------------
	ErrCreateTransaction = "error creating transaction: %w"
	ErrCommitTransaction = "error committing transaction: %w"

	// Ship error messages
	ErrCreateShip          = "error creating ship: %w"
	ErrFindShipByID        = "error finding ship by ID: %w"
	ErrShipNotFound        = "ship with ID %s not found: %w"
	ErrUpdateShip          = "error updating ship: %w"
	ErrDeleteShip          = "error deleting ship: %w"
	ErrGetShip             = "error retrieving ship: %w"
	ErrInvalidShipID       = "invalid ship ID: %w"
	ErrParseOrganizationID = "error parsing organization ID: %w"
	ErrParseShipID         = "error parsing ship ID: %w"
	// Ship configuration error messages
	ErrCreateShipConfig  = "error creating ship configuration: %w"
	ErrFindShipConfig    = "error finding ship configuration: %w"
	ErrUpdateShipConfig  = "error updating ship configuration: %w"
	ErrDeleteShipConfig  = "error deleting ship configuration: %w"
	ErrExecuteShipSelect = "error executing ship select: %w"
	ErrExecuteShipUpdate = "error executing ship update: %w"
	ErrExecuteSeatList   = "error executing seat list: %w"

	// -----------------------------
	// General log error messages
	// -----------------------------
	LogErrCommitTransaction = "failed to commit transaction:"

	// Ship log error messages
	LogErrCreateShip             = "error creating ship:"
	LogErrFindShip               = "error finding ship:"
	LogErrUpdateShip             = "error updating ship:"
	LogErrDeleteShip             = "error deleting ship:"
	LogErrGetShip                = "error retrieving ship:"
	LogShipNotFound              = "ship not found:"
	LogInvalidShipID             = "Invalid ship ID:"
	LogFailedToGetShip           = "failed to get Ship:"
	LogFailedToExecuteShipUpdate = "failed to execute ship update:"
	LogFailedToExecuteSeatList   = "failed to execute seat list:"
	// Ship configuration log error messages
	LogErrCreateShipConfig = "error creating ship configuration:"
	LogErrFindShipConfig   = "error finding ship configuration:"
	LogErrUpdateShipConfig = "error updating ship configuration:"
	LogErrDeleteShipConfig = "error deleting ship configuration:"
	LogShipConfigNotFound  = "ship configuration not found:"

	// -----------------------------
	// General info messages
	// -----------------------------
	InfoConnectedToDB = "connected to the database successfully!"

	// Ship info messages
	LogInfoCreateShip = "ship created successfully:"
	LogInfoGetShip    = "ship retrieved successfully:"
	LogInfoUpdateShip = "ship updated successfully:"
	LogInfoDeleteShip = "ship deleted successfully:"
	LogInfoCreateSeat = "seat created successfully:"

	// Ship configuration info messages
	LogInfoCreateShipConfig = "ship configuration created successfully:"
	LogInfoGetShipConfig    = "ship configuration retrieved successfully:"
	LogInfoUpdateShipConfig = "ship configuration updated successfully:"
	LogInfoDeleteShipConfig = "ship configuration deleted successfully:"
	LogFailedToDeleteShip   = "failed to delete ship:"
	// -----------------------------
	// General fatal messages
	// -----------------------------
	FailedToReadConfigFile      = "failed to read config file:"
	FailedToUnmarshalConfigFile = "failed to unmarshal config file:"
	FailedToCreateDBPool        = "failed to create database connection pool:"
	LogFailedToBeginTransaction = "failed to begin transaction:"

	//seat log
	LogFailedToExecuteSeatUpdate = "failed to execute seat update:"
	LogInfoUpdateSeat            = "seat updated successfully:"
	LogSeatNotFound              = "seat not found:"
	LogFailedToScanSeatRow       = "failed to scan seat row:"

	//seat error
	ErrExecuteSeatUpdate = "error executing seat update: %w"
	ErrCreateSeat        = "error creating seat: %w"
	ErrGetSeat           = "error retrieving seat: %w"
	ErrUpdateSeat        = "error updating seat: %w"
	ErrDeleteSeat        = "error deleting seat: %w"
	ErrReserveSeat       = "error reserving seat: %w"
	ErrExecuteSeatSelect = "error executing seat select: %w"
	ErrScanSeatRow       = "error scanning seat row: %w"
	// trip-config
	ErrArrivalBeforeDeparture = "arrival should be after departure"
	ErrInvalidRecurrence      = "invalid recurrence"
	ErrTripTooShort           = "trip interval too short"
	ErrTripTooLong            = "trip too long"
	ErrTimezoneMismatch       = "timezone mismatch"
	ErrTripTooShortSameDay    = "trip interval too short for same day"

)
