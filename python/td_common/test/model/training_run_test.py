from python.td_common.model.training_run import TrainingRun
import unittest

class TestTrainingRun(unittest.TestCase):

    def test_training_run_serial(self):
        # given
        training_run = TrainingRun("id", "modelId", "name", "dataSet", "time", 0.1, 0.2, {"key": "value"}, 32, 10)
        # when
        training_run_json = training_run.to_json()
        # then
        self.assertEqual(training_run.ID, training_run_json['_id'])
        self.assertEqual(training_run.ModelId, training_run_json['modelId'])
        self.assertEqual(training_run.Name, training_run_json['name'])
        self.assertEqual(training_run.DataSet, training_run_json['dataSet'])
        self.assertEqual(training_run.Time, training_run_json['time'])
        self.assertEqual(training_run.F1Train, training_run_json['f1Train'])
        self.assertEqual(training_run.F1Test, training_run_json['f1Test'])
        self.assertEqual(training_run.TrainingHistory, training_run_json['trainingHistory'])
        self.assertEqual(training_run.BatchSize, training_run_json['batchSize'])
        self.assertEqual(training_run.Epochs, training_run_json['epochs'])

    def test_training_run_deserial(self):
        # given
        training_run_json = {
            '_id': "id",
            'modelId': "modelId",
            'name': "name",
            'dataSet': "dataSet",
            'time': "time",
            'f1Train': 0.1,
            'f1Test': 0.2,
            'trainingHistory': {"key": "value"},
            'batchSize': 32,
            'epochs': 10
        }
        # when
        training_run = TrainingRun.from_json(training_run_json)
        # then
        self.assertEqual(training_run_json['_id'], training_run.ID)
        self.assertEqual(training_run_json['modelId'], training_run.ModelId)
        self.assertEqual(training_run_json['name'], training_run.Name)
        self.assertEqual(training_run_json['dataSet'], training_run.DataSet)
        self.assertEqual(training_run_json['time'], training_run.Time)
        self.assertEqual(training_run_json['f1Train'], training_run.F1Train)
        self.assertEqual(training_run_json['f1Test'], training_run.F1Test)
        self.assertEqual(training_run_json['trainingHistory'], training_run.TrainingHistory)
        self.assertEqual(training_run_json['batchSize'], training_run.BatchSize)
        self.assertEqual(training_run_json['epochs'], training_run.Epochs)

