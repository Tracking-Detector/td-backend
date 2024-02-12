from minio import Minio
class StorageService:
    def __init__(self, minio_client: Minio):
        self.minio_client = minio_client
    
    def get_object(self, bucket_name: str, object_name: str):
        return self.minio_client.get_object(bucket_name, object_name)
    
    def put_object(self, bucket_name: str, object_name: str, data, length):
        return self.minio_client.put_object(bucket_name, object_name, data, length)
    
    def remove_object(self, bucket_name: str, object_name: str):
        return self.minio_client.remove_object(bucket_name, object_name)
    
    def list_objects(self, bucket_name: str, prefix: str = None, recursive: bool = False):
        return self.minio_client.list_objects(bucket_name, prefix, recursive)
    
    def bucket_exists(self, bucket_name: str):
        return self.minio_client.bucket_exists(bucket_name)
    
    def make_bucket(self, bucket_name: str):
        return self.minio_client.make_bucket(bucket_name)
    
    def remove_bucket(self, bucket_name: str):
        return self.minio_client.remove_bucket(bucket_name)
    
    def list_buckets(self):
        return self.minio_client.list_buckets()
    