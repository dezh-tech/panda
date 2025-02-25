# Panda Document Schema Diagram

### **Schema: Users**  
**Description**: Stores user credentials and metadata.  
```jsonc
{
  "_id": "ObjectId",
  "pubKey": "string",  // Nostr public key
  "createdAt": "ISODate",  // User creation timestamp
  "updatedAt": "ISODate"   // Last updated timestamp
}
```

---

### **Schema: Domains**  
**Description**: Defines domain settings, including pricing and TTL values.  
```jsonc
{
  "_id": "ObjectId",
  "domain": "string",  // Domain name (e.g., "example.com")
  "basePricePerIdentifier": "number",  // Base cost per identifier (in sats)
  "pricePerChar": "number",  // Cost per character (in sats)
  "defaultTTL": "number",  // Default Time-to-Live for JSON cache
  "createdAt": "ISODate",  // Domain creation timestamp
  "updatedAt": "ISODate"   // Last updated timestamp
}
```

---

### **Schema: Identifiers**  
**Description**: Tracks identifiers assigned to users within a domain.  
```jsonc
{
  "_id": "ObjectId",
  "name": "string",  // Identifier name (e.g., "alice")
  "domainId": "ObjectId",  // Reference to Domains table
  "fullIdentifier": "string",  // Full identifier (e.g., "alice@example.com")
  "userId": "ObjectId",  // Reference to Users table
  "expiresAt": "ISODate",  // Subscription expiration date
  "createdAt": "ISODate",  // Identifier creation timestamp
  "updatedAt": "ISODate"   // Last updated timestamp
}
```

---

### **Schema: Records**  
**Description**: Stores resolution records (e.g., TXT, NOSTR, CNAME) for identifiers.  
```jsonc
{
  "name": "Records",
  "_id": "ObjectId",
  "identifierId": "ObjectId",  // Reference to Identifiers table
  "type": "string",  // Record type (e.g., "NOSTR", "CNAME", "URL", "TXT")
  "value": "string",  // Record value (e.g., public key, alias, URL)
  "priority": "number",  // Sorting priority for multiple records
  "ttl": "number",  // Time-to-Live for the record in cache
  "createdAt": "ISODate",  // Record creation timestamp
  "updatedAt": "ISODate"   // Last updated timestamp
}
```

---

### **Schema: ReservedIdentifiers**  
**Description**: Tracks reserved identifiers that cannot be registered by users and their associated pricing.  
```jsonc
{
  "name": "ReservedIdentifiers",
  "_id": "ObjectId",
  "fullIdentifier": "string",  // Reserved identifier (e.g., "google@abc.com", "a@abc.com")
  "reservedBy": "ObjectId",  // Reference to Users table (optional)
  "reason": "string",  // Reason for reservation (e.g., "Trademark", "Premium")
  "price": "number",  // Price for the reserved identifier (in sats or specified currency)
  "createdAt": "ISODate",  // Reservation creation timestamp
  "updatedAt": "ISODate"   // Last updated timestamp
}
```

---

### **Schema: Logs**  
**Description**: Tracks various system events, including record changes, user logins, purchases, and more.  
```jsonc
{
  "name": "Logs",
  "_id": "ObjectId",
  "userId": "ObjectId",  // Reference to Users table (optional, if applicable to a user)
  "eventType": "string",  // Type of event (e.g., "record_change", "login", "purchase")
  "entityType": "string",  // Affected entity type (e.g., "Records", "Users", "Identifiers")
  "entityId": "ObjectId",  // ID of the affected entity (optional)
  "description": "string",  // Detailed description of the event
  "metadata": "object",  // Additional metadata (e.g., changes made, IP address, amount, etc.)
  "createdAt": "ISODate"  // Log creation timestamp
}
```
