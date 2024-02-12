from pymongo.database import Database

from python.td_common.repository.base_repository import BaseRepository
from python.td_common.model.dataset import Dataset
import python.td_common.config as config

class DatasetRepository(BaseRepository):
    def __init__(self, db: Database):
        super().__init__(db[config.env_dataset_collection()])

    def _doc_to_entity(self, doc):
        return Dataset.from_json(doc) if doc else None
    