# **GIS Dynamo tables**

## **Dynamo DB Tables:**
---

- generic-ingest-point-configuration
- generic-ingest-point-lookup
- generic-ingest-point-polling-status
- generic-ingest-point-service-account

---
## 1. generic-ingest-point-configuration

### For ingest configuration

### **POST Ingest Configuration**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|configuration|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

### **PATCH Ingest Configuration**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|configuration|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

### **DELETE Ingest Configuration**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|configuration|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

As POST, PATCH, and DELETE are write operations, they can only be implemented using the primary table.
If GSI or LSI are used to access the database, this won't function.

### **GET All Ingest Configurations**

|type (hashKey)| cid | partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|configuration|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

### **GET Ingest Configuration by configuration Id**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|configuration|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|


### For Ingest Instance

### **POST Ingest Instance**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|instance \| 37ff194e-be52-4deb-931a-9d142e40ca56|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

### **PATCH Ingest Instance**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|instance \| 37ff194e-be52-4deb-931a-9d142e40ca56|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

### **DELETE Ingest Instance**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|instance \| 37ff194e-be52-4deb-931a-9d142e40ca56|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|

As POST, PATCH, and DELETE are write operations, they can only be implemented using the primary table.
If GSI or LSI are used to access the database, this won't function.

### **GET All Ingest Instance**

|partitionId (hashKey)| cid (sortKey)| type| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|instance \| 37ff194e-be52-4deb-931a-9d142e40ca56| | schemaData|

### **GET Ingest Instance by Instance Id**

|cid (hashKey)| type (sortKey)| partitionId| .....| schemaDefination|
|:-------------|:-------------|:-----------|:-----|:----------------|
|72609d2a-70a4-4078-9cc5-f42d5cec2ab2|instance \| 37ff194e-be52-4deb-931a-9d142e40ca56|partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| | schemaData|


---
## 2. generic-ingest-point-lookup

### **POST Single Point in Lookup Table**

|hashKey| sortKey| instanceId|
|:------|:-------|:----------|
|point-7ceb9a68-6ef9-4584-a631-954b13cde40c| part-partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| |

### **POST Points in Lookup Table**

|hashKey| sortKey| instanceId|
|:------|:-------|:----------|
|point-7ceb9a68-6ef9-4584-a631-954b13cde40c| part-partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| |

Batch insert is used to add multiple points in lookup table simultaneously, this is used in event replay.

### **PATCH Point in Lookup Table**

|hashKey| sortKey| instanceId|
|:------|:-------|:----------|
|point-7ceb9a68-6ef9-4584-a631-954b13cde40c| part-partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| |

### **DELETE Point from Lookup Table**

|hashKey| sortKey| instanceId|
|:------|:-------|:----------|
|point-7ceb9a68-6ef9-4584-a631-954b13cde40c| part-partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| |

There are two ways to use delete query on lookup table
- using hashKey only
- using both hashKey and sortKey

### **GET Point from Lookup Table**

|hashKey| sortKey| instanceId|
|:------|:-------|:----------|
|point-7ceb9a68-6ef9-4584-a631-954b13cde40c| part-partition.id-10d729c8-1905-41aa-93d1-48017c8e4503| |

There are two ways to get point data from lookup table
- using hashKey only
- using both hashKey and sortKey

---
## 3. generic-ingest-point-polling-status

### **POST Polling Status**

|instanceId (hashKey)| message| pointId| pollingStatus| timestamp|
|:-------------------|:-------|:-------|:-------------|:---------|
|daf6fcc3-3fc4-47fb-852a-c82217a3312c| Point- Accepted| 3cd6e214-4271-48ea-8e7c-7c17c069c7f9| 202| 2023-01-24T23:47:04Z|

### **PATCH Polling Status**

|instanceId (hashKey)| message| pointId| pollingStatus| timestamp|
|:-------------------|:-------|:-------|:-------------|:---------|
|daf6fcc3-3fc4-47fb-852a-c82217a3312c| Point- Accepted| 3cd6e214-4271-48ea-8e7c-7c17c069c7f9| 202| 2023-01-24T23:47:04Z|

### **DELETE Polling Status**

|instanceId (hashKey)| message| pointId| pollingStatus| timestamp|
|:-------------------|:-------|:-------|:-------------|:---------|
|daf6fcc3-3fc4-47fb-852a-c82217a3312c| Point- Accepted| 3cd6e214-4271-48ea-8e7c-7c17c069c7f9| 202| 2023-01-24T23:47:04Z|

### **GET Polling Status**

|instanceId (hashKey)| message| pointId| pollingStatus| timestamp|
|:-------------------|:-------|:-------|:-------------|:---------|
|daf6fcc3-3fc4-47fb-852a-c82217a3312c| Point- Accepted| 3cd6e214-4271-48ea-8e7c-7c17c069c7f9| 202| 2023-01-24T23:47:04Z|

---
## 4. generic-ingest-point-service-account

### **POST Service Account**

|partitionId (hashKey)| id (sortKey)| serviceAccountId| name| ....| userDetails|
|:--------------------|:------------|:----------------|:----|:----|:-----------|
|partition.id-87362731-c1b8-4a96-9660-08f2e56fc03a| e52f1939-5801-4a30-b23e-ef808147f357 \| 6111eb58-88ee-4a65-86b2-474445d97173| 6111eb58-88ee-4a65-86b2-474445d97173| Service Account| |{"username":{"S":"testUser"}}|

### **PATCH Service Account**

|partitionId (hashKey)| id (sortKey)| serviceAccountId| name| ....| userDetails|
|:--------------------|:------------|:----------------|:----|:----|:-----------|
|partition.id-87362731-c1b8-4a96-9660-08f2e56fc03a| e52f1939-5801-4a30-b23e-ef808147f357 \| 6111eb58-88ee-4a65-86b2-474445d97173| 6111eb58-88ee-4a65-86b2-474445d97173| Service Account| |{"username":{"S":"testUser"}}|

### **DELETE Service Account**

|partitionId (hashKey)| id (sortKey)| serviceAccountId| name| ....| userDetails|
|:--------------------|:------------|:----------------|:----|:----|:-----------|
|partition.id-87362731-c1b8-4a96-9660-08f2e56fc03a| e52f1939-5801-4a30-b23e-ef808147f357 \| 6111eb58-88ee-4a65-86b2-474445d97173| 6111eb58-88ee-4a65-86b2-474445d97173| Service Account| |{"username":{"S":"testUser"}}|

### **Get Service Account by Service Account id **

|partitionId (hashKey)| serviceAccountId (sortKey)| id| name| ....| userDetails|
|:--------------------|:------------|:----------------|:----|:----|:-----------|
|partition.id-87362731-c1b8-4a96-9660-08f2e56fc03a|  6111eb58-88ee-4a65-86b2-474445d97173| e52f1939-5801-4a30-b23e-ef808147f357 \| 6111eb58-88ee-4a65-86b2-474445d97173| Service Account| |{"username":{"S":"testUser"}}|

### **GET Service Accounts by Configuration Id**

|partitionId (hashKey)| id (sortKey)| serviceAccountId| name| ....| userDetails|
|:--------------------|:------------|:----------------|:----|:----|:-----------|
|partition.id-87362731-c1b8-4a96-9660-08f2e56fc03a| e52f1939-5801-4a30-b23e-ef808147f357 \| 6111eb58-88ee-4a65-86b2-474445d97173| 6111eb58-88ee-4a65-86b2-474445d97173| Service Account| |{"username":{"S":"testUser"}}|
