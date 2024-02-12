import pytest
import os
from pymongo import MongoClient
from testcontainers.mongodb import MongoDbContainer
from python.td_common.repository.model_repository import ModelRepository
from python.td_common.model.model import Model

class TestRepositoryAcceptance:
    @classmethod
    @pytest.fixture(scope="class", autouse=True)
    def mongo_container(cls):
        with MongoDbContainer("mongo:latest") as container:
            cls.container = container
            yield container

    @classmethod
    @pytest.fixture(scope="class", autouse=True)
    def mongo_client(cls, mongo_container):
        cls.client = MongoClient(mongo_container.get_connection_url())
        yield cls.client
        cls.client.close()

    @classmethod
    @pytest.fixture(scope="class")
    def repository_instance(cls, mongo_client):
        repository = ModelRepository(mongo_client['tracking-detector'])
        os.environ['MODELS_COLLECTION'] = 'models'

        yield repository

    def test_save_and_delete(self, repository_instance):
        # given
        model = {
            '_id': "id",
            'name': "name",
            'label': "label",
            'description': "description"
        }
        validModel = Model.from_json(model)
        # when
        repository_instance.save(validModel)
        # then
        assert repository_instance.find_by_id("id") == validModel
        # cleanup
        repository_instance.delete_by_id("id")
        assert repository_instance.find_by_id("id") is None
