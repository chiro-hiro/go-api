# GOAPI Authentication API

## GOAPI Standard of API

**Request:**

Almost request are using `POST` method with restricted [Cross-Origin Resource Sharing](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

```
POST /validation/isHexString HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
Content-Length: 12
Content-Type: application/x-www-form-urlencoded; charset=UTF-8

value=11ffaa
```

**Response:**

All response will have the same format, 

```golang
	type apiResonse struct {
		Success bool
		Message string
		Data    interface{}
	}
```

Where:

- **Success:** `boolean` value that's describe request is success or not.
    - `true`: No internal error.
    - `false`: Internal error found.
- **Message:** `string` value that's describe _reason of failure_ only available when **Success** is equal to `false`.
- **Data:**  `object` that's contain result of execution.

## Users package

### Validation

**URIs:**

There are serveral methods which were provided to validate input data, developer can request to server following URIs.

```
/validation/isHexString
/validation/isValidPassword
/validation/isValidUsername
/validation/isValidEmail 
/validation/isValidValue
```


**Response:**

_Success:_

```json
{"Success":true,"Message":"","Data":true}
```

_Fail:_

```json
{"Success":false,"Message":"Invalid hex string","Data":false}
```

**Example:**

```javascript
fetch("http://localhost:1988/validation/isHexString",
    {
        "credentials": "omit",
        "headers": {
            "content-type": "application/x-www-form-urlencoded; charset=UTF-8"
        },
        "referrer": "http://localhost:8080/",
        "referrerPolicy": "no-referrer-when-downgrade",
        "body": "value=1122ff",
        "method": "POST",
        "mode": "cors"
    }).then(function(response){
        try {
            response.json().then(function (resData) {
                console.log(resData);
            }, console.error);
        } catch (e) {
            console.error(e);
        }
    });
```
 
### User profile

#### Register

**Request:**

```
POST /user/register HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
Content-Length: 92
Content-Type: application/x-www-form-urlencoded; charset=UTF-8

username=chiro10144478&email=chiro5098682%40gmail.com.vn&password=23402425273131313031383231
```

- **username:** `string` register username
- **email:** `string` user's email
- **password:** `string` hex string of password

**Response:**

```json
{"Success":true,"Message":"","Data":true}
```

#### Login

**Request:**

```
POST /user/login HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
Content-Length: 58
Content-Type: application/x-www-form-urlencoded; charset=UTF-8

username=chiro10144478&password=23402425273131313031383231
```

- **username:** `string` register username
- **password:** `string` hex string of password

**Response:**

```json
{"Success":true,"Message":"","Data":"9d0c91df20904c90a57a7ca795de424ace02c65517f99b9f8f2c2d7cc30d3c95"}
```
- **Data:** Session ID string, all futher request need to contain this token


#### Get user's information

**Request:**

```
POST /user/getUser HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
X-Session-Id: 9d0c91df20904c90a57a7ca795de424ace02c65517f99b9f8f2c2d7cc30d3c95
```

**Response:**

```json
{
    "Success": true,
    "Message": "",
    "Data": {
        "ID": "26146486320dec386cd09d525903bb70",
        "Username": "chiro10144478",
        "Email": "chiro5098682@gmail.com.vn"
    }
}
```

#### Update profile's field

**Request:**

```
POST /user/updateProfile HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
Content-Length: 28
Content-Type: application/x-www-form-urlencoded; charset=UTF-8
X-Session-Id: 9d0c91df20904c90a57a7ca795de424ace02c65517f99b9f8f2c2d7cc30d3c95

field=first-name&value=Chiro
```

_Allowed field:_

```golang
	validMap := map[string]bool{
		"first-name":  true,
		"last-name":   true,
		"address":     true,
		"id-number":   true,
		"issued-date": true}
```

**Response:**

```json
{"Success":true,"Message":"","Data":true}
```

#### Get user's profile

**Request:**

```
POST /user/getProfile HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
X-Session-Id: 9d0c91df20904c90a57a7ca795de424ace02c65517f99b9f8f2c2d7cc30d3c95
```

**Response:**

```json
{
    "Success": true,
    "Message": "",
    "Data": {
        "address": "No where",
        "first-name": "Chiro",
        "last-name": "Hiro",
        "userID": "26146486320dec386cd09d525903bb70"
    }
}
```

#### Update password

**Request:**

```
POST /user/updatePassword HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
Content-Length: 87
Content-Type: application/x-www-form-urlencoded; charset=UTF-8
X-Session-Id: 9d0c91df20904c90a57a7ca795de424ace02c65517f99b9f8f2c2d7cc30d3c95

currentPassword=23402425273131313031383231&newPassword=23402425272222226439363437333933
```

**Response:**

```json
{"Success":true,"Message":"","Data":true}
```

#### Logout

**Request:**

```
POST /user/logout HTTP/1.1
Host: localhost:1988
Origin: http://localhost:8080
X-Session-Id: 9d0c91df20904c90a57a7ca795de424ace02c65517f99b9f8f2c2d7cc30d3c95
```

**Response:**

```json
{"Success":true,"Message":"","Data":true}
```