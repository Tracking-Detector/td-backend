import pytest
import os
from pymongo import MongoClient
from testcontainers.mongodb import MongoDbContainer
from python.td_common.repository.training_run_repository import TrainingRunRepository
from python.td_common.model.training_run import TrainingRun

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
        repository = TrainingRunRepository(mongo_client['tracking-detector'])
        os.environ['TRAINING_RUNS_COLLECTION'] = 'training-runs'

        yield repository

    def test_save_and_delete(self, repository_instance):
        # given
        training_run = {
            '_id': "id",
            'name': "name",
            'label': "label",
            'description': "description"
        }
        validTrainingRun = TrainingRun.from_json(training_run)
        # when
        repository_instance.save(validTrainingRun)
        # then
        assert repository_instance.find_by_id("id") == validTrainingRun
        # cleanup
        repository_instance.delete_by_id("id")
        assert repository_instance.find_by_id("id") is None