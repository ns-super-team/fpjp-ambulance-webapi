const mongoHost = process.env.AMBULANCE_API_MONGODB_HOST;
const mongoPort = process.env.AMBULANCE_API_MONGODB_PORT;

const mongoUser = process.env.AMBULANCE_API_MONGODB_USERNAME;
const mongoPassword = process.env.AMBULANCE_API_MONGODB_PASSWORD;

const database = process.env.AMBULANCE_API_MONGODB_DATABASE;
const collection = process.env.AMBULANCE_API_MONGODB_COLLECTION;
const departmentsCollection = "departments";
const roomsCollection = "rooms";

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// Try to connect to MongoDB until it is available
let connection;
while (true) {
    try {
        connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        break;
    } catch (exception) {
        print(`Cannot connect to MongoDB: ${exception}`);
        print(`Will retry after ${retrySeconds} seconds`);
        sleep(retrySeconds * 1000);
    }
}

// If database and collections exist, exit with success - already initialized
const databases = connection.getDBNames();
if (databases.includes(database)) {
    const dbInstance = connection.getDB(database);
    const collections = dbInstance.getCollectionNames();
    if (collections.includes(departmentsCollection) && collections.includes(roomsCollection)) {
        print(`Collections '${departmentsCollection}' and '${roomsCollection}' already exist in database '${database}'`);
        process.exit(0);
    }
}

// Initialize: create database and collections
const db = connection.getDB(database);

if (!db.getCollectionNames().includes(departmentsCollection)) {
    db.createCollection(departmentsCollection);
    db[departmentsCollection].createIndex({ "id": 1 });
}

if (!db.getCollectionNames().includes(roomsCollection)) {
    db.createCollection(roomsCollection);
    db[roomsCollection].createIndex({ "id": 1 });
}

// Insert sample data into the departments collection
let result = db[departmentsCollection].insertMany([
  {"id": "1", "name": "Pediatrické oddelenie"},
  {"id": "2", "name": "Chirurgia"},
  {"id": "3", "name": "Alergológia"},
  {"id": "4", "name": "Ortopédia"},
  {"id": "5", "name": "Neurológia"},
]);

if (result.writeError) {
  console.error(result);
  print(`Error when writing the data: ${result.errmsg}`);
}

// Insert sample data into the rooms collection
result = db[roomsCollection].insertMany([
  {"id": "1", "department_id": "1", "name": "Miestnosť 1.1"},
  {"id": "2", "department_id": "1", "name": "Miestnosť 1.2"},
  {"id": "3", "department_id": "2", "name": "Miestnosť 2.1"},
  {"id": "4", "department_id": "2", "name": "Miestnosť 2.2"},
  {"id": "5", "department_id": "2", "name": "Miestnosť 2.3"},
  {"id": "6", "department_id": "3", "name": "Miestnosť 3.1"},
  {"id": "7", "department_id": "4", "name": "Miestnosť 4.1"},
  {"id": "8", "department_id": "5", "name": "Miestnosť 5.1"},
  {"id": "9", "department_id": "5", "name": "Miestnosť 5.2"},
]);

if (result.writeError) {
  console.error(result);
  print(`Error when writing the data: ${result.errmsg}`);
}

// Exit with success
process.exit(0);