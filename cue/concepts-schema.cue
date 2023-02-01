// https://cuelang.org/play/?id=JLRtLfwYsLl#cue@export@json

#UUID: string & =~"^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
#DATE: string &=~"^[0-9]{4}-[0-9]{2}-[0-9]{2}$"

// FIGI is a 12 character, alphanumeric, randomly generated ID
#FIGICODE: string &=~"^[0-9A-Z]{12}$"

#Concept: {
    uuid: #UUID
    type: string & =~"^[a-zA-Z]+$"
    prefLabel: string & =~"^[a-zA-Z 0-9.-]+$"
    authority: string & =~"^[a-zA-Z]+$"
    authorityValue: string & =~"^[a-zA-Z0-9=-]+$"
    
    aliases?: [...string]
    broaderUUIDs?: [...string]
    relatedUUIDs?: [...string]
    supersededByUUIDs?: [...string]
    
    scopeNote?: string
    isDeprecated?: bool
}

#Person: {#Concept, {
        salutation?: string
        birthYear?: int
        scopeNote?: string
        imageUrl?: string
        descriptionXML?: string
        emailAddress?: string
        twitterHandle?: string
        facebookPage?: string
    }
}

#MembershipRole: #Concept
#Genre: #Concept

#Membership: {#Concept, {
        organisationUUID: #UUID
        personUUID: #UUID
        inceptionDate: #DATE
        membershipRoles: [...{
                membershipRoleUUID: #UUID
                inceptionDate: #DATE
            }
        ]
    }
}

#FinancialInstrument: {#Concept, {
        figiCode: #FIGICODE
        issuedBy: #UUID
    }
}

#IndustryIdentifier: {#Concept, {
        industryIdentifier: string | int // TODO: more complete constraint
    }
}

#Organisation: {#Concept, {
        countryCode: string
        countryOfIncorporation: string
        countryOfIncorporationUUID: #UUID
        countryOfOperations: string
        countryOfOperationsUUID: #UUID
        countryOfRisk: string
        countryOfRiskUUID: #UUID
        
        legalName: string
        leiCode: string
        parentOrganisation: #UUID
        postalCode: string
        properName: string
        shortName: string
        yearFounded: int
        
        naicsIndustryClassifications: [...{
            uuid: #UUID
            rank: int
        }]
        
        formerNames?: [string]
    }
}

#Location: {#Concept, {
        iso31661?: string
        partOfUUIDs?: [...#UUID]
        geonamesFeatureCode?: string
    }
}

#Topic: {#Concept, {
        descriptionXML: string
    }
}

#Brand: {#Concept, {
        descriptionXML: string
        strapline: string
        parentUUIDs: [...#UUID]
        hasFocusUUIDs: [...#UUID]
        impliedByUUIDs: [...#UUID]
    }
}
