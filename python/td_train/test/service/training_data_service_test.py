import unittest
from unittest.mock import Mock, patch
from python.td_common.minio.storage_service import StorageService
from python.td_common.model.model import Model
from python.td_train.service.training_data_service import TrainingDataService

class TrainingDataServiceTests(unittest.TestCase):
    def setUp(self):
        self.storage_service = Mock(spec=StorageService)
        self.training_data_service = TrainingDataService(self.storage_service)

    def test_get_training_data_success(self):
        # given
        s3location = 'training_data'
        model = Model(ID="test", Dims=[204, 1], Name="test", Description="test", TensorflowLayers={"key": "value"})

        mock_resp = Mock()
        mock_resp.status = 200
        with open('./python/td_train/test/service/testdata/expected_encoding.csv.gz', 'rb') as file:
            mock_resp.read.return_value = file.read()
        self.storage_service.get_object.return_value = mock_resp

        # when
        X, y = self.training_data_service.get_training_data(s3location, model)
        # then
        self.assertEqual(X.shape, (10, 204))
        self.assertEqual(y.shape, (10,))
        self.storage_service.get_object.assert_called_once_with(None, 'training_data.csv.gz')

    def test_get_training_data_error_dims(self):
        # given
        s3location = 'training_data'
        model = Model(ID="test", Dims=[203, 1], Name="test", Description="test", TensorflowLayers={"key": "value"})

        mock_resp = Mock()
        mock_resp.status = 200
        with open('./python/td_train/test/service/testdata/expected_encoding.csv.gz', 'rb') as file:
            mock_resp.read.return_value = file.read()
        self.storage_service.get_object.return_value = mock_resp

        # when
        with self.assertRaises(Exception) as context:
            self.training_data_service.get_training_data(s3location, model)
        # then
        self.assertEqual(str(context.exception), 'Model dimensions do not match training data dimensions: 203 != 10')
       



if __name__ == '__main__':
    unittest.main()