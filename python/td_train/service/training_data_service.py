import gzip
import io
import numpy as np
import pandas as pd
from python.td_common.config import env
from python.td_common.minio.storage_service import StorageService
from python.td_common.model.model import Model

class TrainingDataService:
    def __init__(self, storage_service: StorageService):
        self.storage_service = storage_service

    def get_training_data(self, training_data_export: str, model: Model) -> tuple[np.ndarray]:
        resp = self.storage_service.get_object(env.env_export_bucket_name(), training_data_export + '.csv.gz')
        if resp.status != 200:
            raise Exception(f'Failed to get training data: {resp.status}')

        # Read the gzip content from the response
        gzip_content = resp.read()

        # Decompress the gzip content
        with gzip.GzipFile(fileobj=io.BytesIO(gzip_content), mode='rb') as gzip_file:
            df = pd.read_csv(gzip_file, header=None)

        if model.Dims[0]+1 != df.shape[1]:
            raise Exception(f'Model dimensions do not match training data dimensions: {model.Dims[0]} != {df.shape[0]}')

        X = df.iloc[:, 0:model.Dims[0]].values
        y = df.iloc[:, -1].values

        return X, y
        

        