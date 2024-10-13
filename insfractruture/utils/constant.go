package utils

const AWS_BUCKET_NAME string = "bucket-all-is-safe-school"
const AWS_REGION string = "us-east-1"
const AWS_URL_S3 string = "https://%s.s3.amazonaws.com/%s"
const UPLOADS string = "./uploads/"
const UPLOADS_FILE string = "./uploads/%s"

//pagianation

const (
	PAGE        string = "page"
	LIMIT       int    = 5
	OFFSET             = "offset"
	TOTAL_COUNT string = "total_count"
	PAGE_COUNT  string = "page_count"
	BEGIN       string = "begin"
)

const (
	ID                    = "id"
	MESSAGE               = "message"
	STATUS                = "status"
	DATA                  = "data"
	CREATED               = "was successfully created"
	UPDATED               = "was updated successfully"
	REMOVED               = "was successfully removed"
	ERROR_QUERY           = "error query, please try again later"
	ERROR_CREATE          = "error creating"
	ERROR_PARSING_BODY    = "error parsing body"
	ERROR_UPDATE          = "error updating"
	ERROR_DELETE          = "error deleting"
	ERROR_REQUIRED_FIELDS = " is required."
	EMPTY                 = ""
)
const (
	DB_DIFF_ID        = "id <>?"
	DB_EQUAL_ID       = "id=?"
	DB_EQUAL_NAME     = "name =?"
	DB_ORDER_DESC     = "id desc"
	DB_EQUAL_EMAIL_ID = "email=?"
)

const (
	ID_NO_EXIST         = "The id not exists"
	NAME_ALREADY_EXIST  = "Name already exists"
	EMAIL_ALREADY_EXIST = "Email already exists"
)
const (
	NAME             = "name"
	ADDRESS          = "address"
	PHONE            = "phone"
	ZIP_CODE         = "zip_code"
	EMAIL            = "email"
	ACTIVE           = "active"
	SCHOOL_ID        = "school_id"
	FILE      string = "file"
)
const (
	NAME_IS_FIELD_REQUIRED    = "The field name is required."
	ADDRESS_IS_FIELD_REQUIRED = "The field address is required."
	PHONE_IS_FIELD_REQUIRED   = "The field phone is required."
	EMAIL_IS_FIELD_REQUIRED   = "The field email is required."
)

const (
	NAME_IS_REQUIRED     = "The name is required."
	ADDRESS_IS_REQUIRED  = "The address is required."
	PHONE_IS_REQUIRED    = "The phone is required."
	ZIP_CODE_IS_REQUIRED = "The zip code is required."
	EMAIL_IS_REQUIRED    = "The email is required."
	EMAIL_IS_INVALID     = "The email is invalid."
)
const STATE_NOT_FOUND string = "State not found"
