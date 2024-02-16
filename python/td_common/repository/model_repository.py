from pymongo.database import Database
from python.td_common.repository.base_repository import BaseRepository
from python.td_common.model.model import Model
import python.td_common.config.env as config

class ModelRepository(BaseRepository):
    def __init__(self, db: Database):
        super().__init__(db[config.env_model_collection()])

    def _doc_to_entity(self, doc):
        return Model.from_json(doc) if doc else None