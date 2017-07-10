package main


type inboundUserObj struct {
	Schemas []string `json:"schemas"`
	UserName string `json:"userName"`
	Name struct {
			Formatted string `json:"formatted"`
			GivenName string `json:"givenName"`
			FamilyName string `json:"familyName"`
		} `json:"name"`
	Emails []struct {
		Value string `json:"value"`
		Primary bool `json:"primary,omitempty"`
		Type string `json:"type"`
	} `json:"emails"`
	Password string `json:"password"`
	Active bool `json:"active"`
}

type outboundUserObj struct {
	Schemas []string `json:"schemas"`
	ID string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Active bool `json:"active"`
	Name struct {
			Formatted string `json:"formatted"`
			GivenName string `json:"givenName"`
			FamilyName string `json:"familyName"`
		} `json:"name"`
	Emails []struct {
		Value string `json:"value"`
		Primary bool `json:"primary,omitempty"`
		Type string `json:"type"`
	} `json:"emails"`
}

//This is what it looks like

/*
{
    "schemas": [
        "urn:scim:schemas:core:1.0",
        "urn:scim:schemas:extension:enterprise:1.0"
    ],
    "id": "104",
    "userName": "webuser@myokta.com",
    "password": "7S56g4dR",
    "active": true,
    "name": {
        "formatted": "web user",
        "givenName": "web",
        "familyName": "user"
    },
    "emails": [
        {
            "value": "webuser@myokta.com",
            "primary": true,
            "type": "primary"
        },
        {
            "value": "pmcdowell@okta.com",
            "type": "secondary"
        }
    ]
}
 */

type fullImportPayloadObj struct {
	TotalResults int `json:"totalResults"`
	Schemas []string `json:"schemas"`
	Resources []singleUserImportObj
}

type userNameStruct struct {
	Formatted string `json:"formatted"`
	GivenName string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type userEmailStruct struct { //Not Using emails
	Value string `json:"value"`
	Primary bool `json:"primary"`
	Type string `json:"type"`
}

type singleUserImportObj struct {
	Schemas []string `json:"schemas"`
	ID string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Active bool `json:"active"`
	Name userNameStruct `json:"name"`
	Emails [] userEmailStruct `json:"emails"`
	//Groups []struct { // No using Groups
	//	Value string `json:"value"`
	//	Display string `json:"display"`
	//} `json:"groups"`
	//UrnOktaOnpremApp11UserCustom struct { //Not using custom stuff
	//		IsAdmin bool `json:"isAdmin"`
	//		IsOkta bool `json:"isOkta"`
	//		DepartmentName string `json:"departmentName"`
	//	} `json:"urn:okta:onprem_app:1.1:user:custom"`
}
