import pytest
import os
from pymongo import MongoClient
from testcontainers.mongodb import MongoDbContainer
from python.td_common.repository.dataset_repository import DatasetRepository
from python.td_common.model.dataset import Dataset

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
        repository = DatasetRepository(mongo_client['tracking-detector'])
        os.environ['DATASET_COLLECTION'] = 'datasets'

        yield repository

    def test_save_and_delete(self, repository_instance):
        # given
        dataset = {
            '_id': "id",
            'name': "name",
            'label': "label",
            'description': "description"
        }
        validDataset = Dataset.from_json(dataset)
        # when
        repository_instance.save(validDataset)
        # then
        assert repository_instance.find_by_id("id") == dataset
        # cleanup
        repository_instance.delete_by_id("id")
        assert repository_instance.find_by_id("id") is None
