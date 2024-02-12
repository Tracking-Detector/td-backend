from pymongo import Collection
from dataclasses import asdict

class BaseRepository:
    def __init__(self, collection: Collection):
        self.collection = collection

    def save(self, entity):
        if '_id' in entity:
            self.collection.replace_one({'_id': entity['_id']}, entity, upsert=True)
        else:
            result = self.collection.insert_one(asdict(entity))
            entity._id = str(result.inserted_id)
        return entity

    def save_all(self, entities):
        result = self.collection.insert_many([asdict(entity) for entity in entities])
        for idx, entity in enumerate(entities):
            entity._id = str(result.inserted_ids[idx])
        return entities

    def find_all(self):
        return [self._doc_to_entity(doc) for doc in self.collection.find()]

    def find_by_id(self, id):
        doc = self.collection.find_one({'_id': id})
        return self._doc_to_entity(doc) if doc else None

    def delete_by_id(self, id):
        result = self.collection.delete_one({'_id': id})
        return result.deleted_count == 1

    def delete_all(self):
        result = self.collection.delete_many({})
        return result.deleted_count

    def count(self):
        return self.collection.count_documents({})

    def _doc_to_entity(self, doc):
        raise NotImplementedError("Must be implemented in subclasses")