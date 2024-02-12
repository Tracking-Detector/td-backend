from python.td_common.model.model import Model
import unittest

class TestModel(unittest.TestCase):
    def test_model_serial(self):
        # given
        model = Model("id", "name", "description", [1, 2, 3], {"key": "value"})
        # when
        model_json = model.to_json()
        # then
        self.assertEqual(model.ID, model_json['_id'])
        self.assertEqual(model.Name, model_json['name'])
        self.assertEqual(model.Description, model_json['description'])
        self.assertEqual(model.Dims, model_json['dims'])
        self.assertEqual(model.TensorflowLayers, model_json['tfLayers'])

    def test_mode_deserial(self):
        # given
        model_json = {
            '_id': "id",
            'name': "name",
            'description': "description",
            'dims': [1, 2, 3],
            'tfLayers': {"key": "value"}
        }
        # when
        model = Model.from_json(model_json)
        # then
        self.assertEqual(model_json['_id'], model.ID)
        self.assertEqual(model_json['name'], model.Name)
        self.assertEqual(model_json['description'], model.Description)
        self.assertEqual(model_json['dims'], model.Dims)
        self.assertEqual(model_json['tfLayers'], model.TensorflowLayers)

    