# **Panda NIP-05 Identifier System Database Schemas**

This document outlines the MongoDB schemas for a system designed to sell and resolve NIP-05 identifiers.

---

## **1. Users Collection**

Stores user information and credentials.

### **Collection Name**: `Users`  
### **Type**: User Management  

| **Field**        | **Type**    | **Constraints**                              | **Description**                                      |
|-------------------|-------------|-----------------------------------------------|------------------------------------------------------|
| `_id`            | ObjectId    | Unique, Indexed                               | Unique identifier for the user.                     |
| `npub`           | string      | Required, Unique, Indexed                     | Nostr public key (e.g., npub...).                   |
| `passwordHash`   | string      | Required                                      | Hashed password.                                    |
| `createdAt`      | ISODate     | Required                                      | Account creation timestamp.                         |
| `updatedAt`      | ISODate     | Required                                      | Last updated timestamp.                             |

---

## **2. Domains Collection**

Defines the available domains for identifiers and their pricing.

### **Collection Name**: `Domains`  
### **Type**: Domain Management  

| **Field**              | **Type**    | **Constraints**                              | **Description**                                      |
|-------------------------|-------------|-----------------------------------------------|------------------------------------------------------|
| `_id`                  | ObjectId    | Unique, Indexed                               | Unique identifier for the domain.                   |
| `domain`               | string      | Required, Unique, Indexed                     | Domain name (e.g., "example.com").                  |
| `basePricePerIdentifier` | number     | Required                                      | Cost per identifier (in sats).                      |
| `defaultTTL`           | number      | Optional                                      | Default Time-to-Live for JSON cache (in seconds).   |
| `status`               | string      | Required, Indexed                             | Status of the domain (e.g., "active", "inactive").  |
| `createdAt`            | ISODate     | Required                                      | Creation timestamp.                                 |
| `updatedAt`            | ISODate     | Required                                      | Last updated timestamp.                             |

---

## **3. Identifiers Collection**

Represents individual NIP-05 identifiers.

### **Collection Name**: `Identifiers`  
### **Type**: Identifier Management  

| **Field**        | **Type**    | **Constraints**                              | **Description**                                      |
|-------------------|-------------|-----------------------------------------------|------------------------------------------------------|
| `_id`            | ObjectId    | Unique, Indexed                               | Unique identifier for the NIP-05 identifier.         |
| `name`           | string      | Required                                      | Identifier name (e.g., "alice").                    |
| `domainId`       | ObjectId    | Required, Indexed                             | Reference to the `Domains` collection.              |
| `fullIdentifier` | string      | Required, Unique, Indexed                     | Full identifier (e.g., "alice@example.com").         |
| `userId`         | ObjectId    | Required, Indexed                             | Reference to the `Users` collection.                |
| `expiresAt`      | ISODate     | Required                                      | Expiration date of the identifier.                  |
| `status`         | string      | Required, Indexed                             | Status of the identifier (e.g., "active", "inactive"). |
| `createdAt`      | ISODate     | Required                                      | Creation timestamp.                                 |
| `updatedAt`      | ISODate     | Required                                      | Last updated timestamp.                             |

---

## **4. Transactions Collection**

Tracks all financial transactions, including payments for identifiers or other services.

### **Collection Name**: `Transactions`  
### **Type**: Financial Management  

| **Field**           | **Type**    | **Constraints**                              | **Description**                                      |
|----------------------|-------------|-----------------------------------------------|------------------------------------------------------|
| `_id`              | ObjectId    | Unique, Indexed                               | Unique identifier for the transaction.              |
| `transactionId`    | string      | Required, Unique, Indexed                     | Unique ID for the transaction (e.g., UUID).         |
| `userId`           | ObjectId    | Required, Indexed                             | Reference to the `Users` collection.                |
| `domainId`         | ObjectId    | Optional, Indexed                             | Reference to the `Domains` collection.              |
| `identifierId`     | ObjectId    | Optional, Indexed                             | Reference to the `Identifiers` collection.          |
| `amount`           | number      | Required                                      | Amount paid (e.g., in sats or fiat).                |
| `currency`         | string      | Required                                      | Currency type (e.g., "BTC", "USD").                 |
| `type`             | string      | Required, Indexed                             | Transaction type (e.g., "purchase", "renewal").     |
| `paymentProcessor` | string      | Optional                                      | Payment processor used (e.g., "Stripe", "Bitcoin"). |
| `paymentDetails`   | object      | Optional                                      | Additional details from the payment processor.      |
| `paymentDetails.referenceId` | string | Optional, Indexed                          | Payment processor's reference ID.                   |
| `paymentDetails.method` | string | Optional                                      | Payment method (e.g., "credit_card").               |
| `paymentDetails.confirmedAt` | ISODate | Optional                                  | When the payment was confirmed (if applicable).     |
| `paymentStatus`    | string      | Required, Indexed                             | Payment status (e.g., "completed", "pending").      |
| `status`           | string      | Required, Indexed                             | Overall transaction status (e.g., "active").        |
| `createdAt`        | ISODate     | Required                                      | Transaction creation timestamp.                     |
| `updatedAt`        | ISODate     | Required                                      | Last updated timestamp.                             |

---

## **5. ResolveRecords Collection**

Stores resolution records for identifiers, similar to DNS records.

### **Collection Name**: `ResolveRecords`  
### **Type**: Resolution Management  

| **Field**        | **Type**    | **Constraints**                              | **Description**                                      |
|-------------------|-------------|-----------------------------------------------|------------------------------------------------------|
| `_id`            | ObjectId    | Unique, Indexed                               | Unique identifier for the resolution record.         |
| `identifierId`   | ObjectId    | Required, Indexed                             | Reference to the `Identifiers` collection.           |
| `type`           | string      | Required, Indexed                             | Record type (e.g., "NOSTR", "CNAME", "URL", "TXT").  |
| `value`          | string      | Required                                      | Record value (e.g., public key, alias, URL).         |
| `priority`       | number      | Optional                                      | Priority for sorting multiple records.              |
| `createdAt`      | ISODate     | Required                                      | Creation timestamp.                                 |
| `updatedAt`      | ISODate     | Required                                      | Last updated timestamp.                             |
