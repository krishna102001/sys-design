# BASE62 Hash
- base62 hash basically convert the number into the string using 62 alphabet and numbers excluding the any special character like +,-,_,/
- to convert the long url we will have the long url in the database and we will use snowflake id to generate the unique short base62 string
- base62 is used when you have id in number form if you have mix of alphabet and number and character then use base64 encoding or some crypto hash.