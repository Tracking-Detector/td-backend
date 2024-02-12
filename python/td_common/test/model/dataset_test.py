from python.td_common.model.dataset import Dataset
import unittest

class TestDataset(unittest.TestCase):
    def test_dataset_serial(self):
        # given
        dataset = Dataset("id", "name", "label", "description")
        # when
        dataset_json = dataset.to_json()
        # then
        self.assertEqual(dataset.ID, dataset_json['_id'])
        self.assertEqual(dataset.Name, dataset_json['name'])
        self.assertEqual(dataset.Label, dataset_json['label'])
        self.assertEqual(dataset.Description, dataset_json['description'])

    def test_dataset_deserial(self):
        # given
        dataset_json = {
            '_id': "id",
            'name': "name",
            'label': "label",
            'description': "description"
        }
        # when
        dataset = Dataset.from_json(dataset_json)
        # then
        self.assertEqual(dataset_json['_id'], dataset.ID)
        self.assertEqual(dataset_json['name'], dataset.Name)
        self.assertEqual(dataset_json['label'], dataset.Label)
        self.assertEqual(dataset_json['description'], dataset.Description)